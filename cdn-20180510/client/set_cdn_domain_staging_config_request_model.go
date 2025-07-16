// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCdnDomainStagingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SetCdnDomainStagingConfigRequest
	GetDomainName() *string
	SetFunctions(v string) *SetCdnDomainStagingConfigRequest
	GetFunctions() *string
}

type SetCdnDomainStagingConfigRequest struct {
	// The accelerated domain name. You can specify only one domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The features that you want to configure. Format:
	//
	// > 	- **functionName**: The name of the feature. Separate multiple values with commas (,). For more information, see [A list of features](https://help.aliyun.com/document_detail/388460.html).
	//
	// >	- **argName**: The feature parameters for **functionName**.
	//
	// >	- **argValue**: The parameter values set for **functionName**.
	//
	//         [
	//
	//          {
	//
	//            "functionArgs": [
	//
	//             {
	//
	//              "argName": "Parameter A",
	//
	//              "argValue": "Value of Parameter A"
	//
	//             },
	//
	//           {
	//
	//             "argName": "Parameter B",
	//
	//             "argValue": "Value of Parameter B"
	//
	//              }
	//
	//          ],
	//
	//          "functionName": "Feature name"
	//
	//             }
	//
	//         ]
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"functionArgs":[{"argName":"enable","argValue":"on"},{"argName":"pri","argValue":"1"},{"argName":"rule","argValue":"xxx"}],"functionName":"edge_function"}]
	Functions *string `json:"Functions,omitempty" xml:"Functions,omitempty"`
}

func (s SetCdnDomainStagingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetCdnDomainStagingConfigRequest) GoString() string {
	return s.String()
}

func (s *SetCdnDomainStagingConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetCdnDomainStagingConfigRequest) GetFunctions() *string {
	return s.Functions
}

func (s *SetCdnDomainStagingConfigRequest) SetDomainName(v string) *SetCdnDomainStagingConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetCdnDomainStagingConfigRequest) SetFunctions(v string) *SetCdnDomainStagingConfigRequest {
	s.Functions = &v
	return s
}

func (s *SetCdnDomainStagingConfigRequest) Validate() error {
	return dara.Validate(s)
}
