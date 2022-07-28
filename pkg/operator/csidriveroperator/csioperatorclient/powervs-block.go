package csioperatorclient

import (
	"os"
	"strings"

	configv1 "github.com/openshift/api/config/v1"
)

const (
	POWERVSBlockCSIDriverName          = "powervs.csi.ibm.com"
	envPOWERVSBlockDriverOperatorImage = "POWERVS_BLOCK_CSI_DRIVER_OPERATOR_IMAGE"
	envPOWERVSBlockDriverImage         = "POWERVS_BLOCK_CSI_DRIVER_IMAGE"
)

func GetPOWERVSBlockCSIOperatorConfig() CSIOperatorConfig {
	pairs := []string{
		"${OPERATOR_IMAGE}", os.Getenv(envPOWERVSBlockDriverOperatorImage),
		"${DRIVER_IMAGE}", os.Getenv(envPOWERVSBlockDriverImage),
	}

	return CSIOperatorConfig{
		CSIDriverName:   POWERVSBlockCSIDriverName,
		ConditionPrefix: "POWERVSBlock",
		Platform:        configv1.NonePlatformType,
		StaticAssets: []string{
			"csidriveroperators/powervs-block/01_sa.yaml",
			"csidriveroperators/powervs-block/02_role.yaml",
			"csidriveroperators/powervs-block/03_rolebinding.yaml",
			"csidriveroperators/powervs-block/04_clusterrole.yaml",
			"csidriveroperators/powervs-block/05_clusterrolebinding.yaml",
		},
		CRAsset:         "csidriveroperators/powervs-block/07_cr.yaml",
		DeploymentAsset: "csidriveroperators/powervs-block/06_deployment.yaml",
		ImageReplacer:   strings.NewReplacer(pairs...),
		AllowDisabled:   false,
	}
}
