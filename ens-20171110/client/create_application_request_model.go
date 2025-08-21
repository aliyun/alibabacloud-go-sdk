// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplate(v string) *CreateApplicationRequest
	GetTemplate() *string
	SetTimeout(v int32) *CreateApplicationRequest
	GetTimeout() *int32
}

type CreateApplicationRequest struct {
	// The edge application template. The value must be a JSON string that contains the following information:
	//
	// 	- Basic information such as the name of the application
	//
	// 	- Information such as resource specifications and network security configurations
	//
	// 	- Service specifications
	//
	// 	- Required resources
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"appMetaData\\":{        \\"appName\\":\\"nginx\\",        \\"clusterName\\":\\"poc\\",        \\"appType\\":\\"Common\\",        \\"description\\":\\"test\\"    },    \\"resourceAttribute\\":{        \\"resourceType\\":\\"\\",        \\"instanceSpec\\":\\"ens.sn1.tiny\\",        \\"systemDiskSize\\":20,        \\"dataDiskSize\\":0,        \\"bandwithOut\\":10,        \\"areaLevel\\":\\"National\\",        \\"netSecurityStrategy\\":null,        \\"initConfig\\":null    },    \\"resourceSelector\\":[        {            \\"count\\":1        }    ],    \\"workload\\":[        {            \\"podCount\\":1,            \\"serviceConfig\\":null,            \\"name\\":\\"nginx\\",            \\"podSpec\\":{                \\"containers\\":[                    {                        \\"name\\":\\"android\\",                        \\"image\\":\\"edge-registry.alicdn.com/test/nginx\\"                    }                ]            },            \\"count\\":1        }    ]}
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The timeout period for asynchronous processing. Unit: seconds. Default value: 1800.
	//
	// example:
	//
	// 1800
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequest) GetTemplate() *string {
	return s.Template
}

func (s *CreateApplicationRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateApplicationRequest) SetTemplate(v string) *CreateApplicationRequest {
	s.Template = &v
	return s
}

func (s *CreateApplicationRequest) SetTimeout(v int32) *CreateApplicationRequest {
	s.Timeout = &v
	return s
}

func (s *CreateApplicationRequest) Validate() error {
	return dara.Validate(s)
}
