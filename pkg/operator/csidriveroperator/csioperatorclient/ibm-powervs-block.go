package csioperatorclient

import (
	"os"
	"strings"

	configv1 "github.com/openshift/api/config/v1"
)

const (
	IBMPOWERVSBlockCSIDriverName          = "powervs.csi.ibm.com"
	envIBMPOWERVSBlockDriverOperatorImage = "IBM_POWERVS_BLOCK_CSI_DRIVER_OPERATOR_IMAGE"
	envIBMPOWERVSBlockDriverImage         = "IBM_POWERVS_BLOCK_CSI_DRIVER_IMAGE"
)

func GetIBMPOWERVSBlockCSIOperatorConfig() CSIOperatorConfig {
	pairs := []string{
		"${OPERATOR_IMAGE}", os.Getenv(envIBMPOWERVSBlockDriverOperatorImage),
		"${DRIVER_IMAGE}", os.Getenv(envIBMPOWERVSBlockDriverImage),
	}

	return CSIOperatorConfig{
		CSIDriverName:   IBMPOWERVSBlockCSIDriverName,
		ConditionPrefix: "IBMPOWERVSBlock",
		Platform:        configv1.NonePlatformType,
		StaticAssets: []string{
			"csidriveroperators/ibm-powervs-block/01_sa.yaml",
			"csidriveroperators/ibm-powervs-block/02_role.yaml",
			"csidriveroperators/ibm-powervs-block/03_rolebinding.yaml",
			"csidriveroperators/ibm-powervs-block/04_clusterrole.yaml",
			"csidriveroperators/ibm-powervs-block/05_clusterrolebinding.yaml",
		},
		CRAsset:         "csidriveroperators/ibm-powervs-block/07_cr.yaml",
		DeploymentAsset: "csidriveroperators/ibm-powervs-block/06_deployment.yaml",
		ImageReplacer:   strings.NewReplacer(pairs...),
		AllowDisabled:   false,
	}
}
