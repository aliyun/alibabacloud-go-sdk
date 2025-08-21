// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpgradeApplicationRequest
	GetAppId() *string
	SetTemplate(v string) *UpgradeApplicationRequest
	GetTemplate() *string
	SetTimeout(v int32) *UpgradeApplicationRequest
	GetTimeout() *int32
}

type UpgradeApplicationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2bac6f4-75dc-455e-8389-2dc8e47526d3
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The information template for phased update. The value must be a JSON string and contain the following information:
	//
	// 	- Version range that you want to update
	//
	// 	- Configuration information of the target version
	//
	// 	- Canary release policy for resources
	//
	// 	- Intelligent upgrade policy that contains information such as the time window and resource usage limit
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"fromAppVersions\\":{        \\"operator\\":\\"In\\",        \\"values\\":[            \\"v1\\",            \\"v2\\"        ]    },    \\"toAppVersion\\":\\"v3\\",    \\"descrip\\":\\"xxx\\",    \\"workload\\":[        {            \\"name\\":\\"nginx\\",            \\"podSpec\\":{            }        }    ],    \\"upgradeStrategy\\":{        \\"name\\":\\"ScheduleToISP\\",        \\"parameters\\":{            \\"operator\\":\\"In\\",            \\"values\\":[                \\"telecom\\"            ]        }    },    \\"autoUpgradeStrategy\\":{        \\"name\\":\\"AdjustToPodUsage\\",        \\"checkInterval\\":600,        \\"startTime\\":\\"2021-02-19 00:00:00\\",        \\"startHourPoint\\":\\"0\\",        \\"endHourPoint\\":\\"8\\",        \\"endTime\\":\\"2021-02-19 08:00:00\\",        \\"level\\":\\"RegionId\\",        \\"rules\\":[            {                \\"regionCodes\\":[                    \\"cn-wuxi-telecom_unicom_cmcc\\",                    \\"cn-shijiazhuang-telecom_unicom_cmcc\\"                ],                \\"usageRatioLimit\\":{                    \\"maxPodUsageRatio\\":50                },                \\"maxUpgradingRatio\\":50            },            {                \\"regionCodes\\":[                    \\"cn-wuhan-telecom_unicom_cmcc\\"                ],                \\"usageRatioLimit\\":{                    \\"maxPodUsageRatio\\":30                },                \\"maxUpgradingRatio\\":20            },            {                \\"regionCodes\\":[                    \\"All\\"                ],                \\"usageRatioLimit\\":{                    \\"maxPodUsageRatio\\":20                },                \\"maxUpgradingRatio\\":50,                \\"maxUpgradingCount\\":2            }        ]    }}
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The timeout period for asynchronous upgrade. Unit: seconds. Default value: 300.
	//
	// example:
	//
	// 1800
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpgradeApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeApplicationRequest) GoString() string {
	return s.String()
}

func (s *UpgradeApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpgradeApplicationRequest) GetTemplate() *string {
	return s.Template
}

func (s *UpgradeApplicationRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpgradeApplicationRequest) SetAppId(v string) *UpgradeApplicationRequest {
	s.AppId = &v
	return s
}

func (s *UpgradeApplicationRequest) SetTemplate(v string) *UpgradeApplicationRequest {
	s.Template = &v
	return s
}

func (s *UpgradeApplicationRequest) SetTimeout(v int32) *UpgradeApplicationRequest {
	s.Timeout = &v
	return s
}

func (s *UpgradeApplicationRequest) Validate() error {
	return dara.Validate(s)
}
