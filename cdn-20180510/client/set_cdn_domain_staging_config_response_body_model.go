// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCdnDomainStagingConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainConfigList(v []*SetCdnDomainStagingConfigResponseBodyDomainConfigList) *SetCdnDomainStagingConfigResponseBody
	GetDomainConfigList() []*SetCdnDomainStagingConfigResponseBodyDomainConfigList
	SetRequestId(v string) *SetCdnDomainStagingConfigResponseBody
	GetRequestId() *string
}

type SetCdnDomainStagingConfigResponseBody struct {
	// The list of domain configurations.
	DomainConfigList []*SetCdnDomainStagingConfigResponseBodyDomainConfigList `json:"DomainConfigList,omitempty" xml:"DomainConfigList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetCdnDomainStagingConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetCdnDomainStagingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetCdnDomainStagingConfigResponseBody) GetDomainConfigList() []*SetCdnDomainStagingConfigResponseBodyDomainConfigList {
	return s.DomainConfigList
}

func (s *SetCdnDomainStagingConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetCdnDomainStagingConfigResponseBody) SetDomainConfigList(v []*SetCdnDomainStagingConfigResponseBodyDomainConfigList) *SetCdnDomainStagingConfigResponseBody {
	s.DomainConfigList = v
	return s
}

func (s *SetCdnDomainStagingConfigResponseBody) SetRequestId(v string) *SetCdnDomainStagingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetCdnDomainStagingConfigResponseBody) Validate() error {
	if s.DomainConfigList != nil {
		for _, item := range s.DomainConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SetCdnDomainStagingConfigResponseBodyDomainConfigList struct {
	// The ID of the configuration.
	//
	// example:
	//
	// 1234567
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The domain name.
	//
	// example:
	//
	// www.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The name of the feature.
	//
	// example:
	//
	// set_resp_header
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
}

func (s SetCdnDomainStagingConfigResponseBodyDomainConfigList) String() string {
	return dara.Prettify(s)
}

func (s SetCdnDomainStagingConfigResponseBodyDomainConfigList) GoString() string {
	return s.String()
}

func (s *SetCdnDomainStagingConfigResponseBodyDomainConfigList) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *SetCdnDomainStagingConfigResponseBodyDomainConfigList) GetDomainName() *string {
	return s.DomainName
}

func (s *SetCdnDomainStagingConfigResponseBodyDomainConfigList) GetFunctionName() *string {
	return s.FunctionName
}

func (s *SetCdnDomainStagingConfigResponseBodyDomainConfigList) SetConfigId(v int64) *SetCdnDomainStagingConfigResponseBodyDomainConfigList {
	s.ConfigId = &v
	return s
}

func (s *SetCdnDomainStagingConfigResponseBodyDomainConfigList) SetDomainName(v string) *SetCdnDomainStagingConfigResponseBodyDomainConfigList {
	s.DomainName = &v
	return s
}

func (s *SetCdnDomainStagingConfigResponseBodyDomainConfigList) SetFunctionName(v string) *SetCdnDomainStagingConfigResponseBodyDomainConfigList {
	s.FunctionName = &v
	return s
}

func (s *SetCdnDomainStagingConfigResponseBodyDomainConfigList) Validate() error {
	return dara.Validate(s)
}
