// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVodDomainStagingConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainConfigList(v []*SetVodDomainStagingConfigResponseBodyDomainConfigList) *SetVodDomainStagingConfigResponseBody
	GetDomainConfigList() []*SetVodDomainStagingConfigResponseBodyDomainConfigList
	SetRequestId(v string) *SetVodDomainStagingConfigResponseBody
	GetRequestId() *string
}

type SetVodDomainStagingConfigResponseBody struct {
	DomainConfigList []*SetVodDomainStagingConfigResponseBodyDomainConfigList `json:"DomainConfigList,omitempty" xml:"DomainConfigList,omitempty" type:"Repeated"`
	RequestId        *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetVodDomainStagingConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetVodDomainStagingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetVodDomainStagingConfigResponseBody) GetDomainConfigList() []*SetVodDomainStagingConfigResponseBodyDomainConfigList {
	return s.DomainConfigList
}

func (s *SetVodDomainStagingConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetVodDomainStagingConfigResponseBody) SetDomainConfigList(v []*SetVodDomainStagingConfigResponseBodyDomainConfigList) *SetVodDomainStagingConfigResponseBody {
	s.DomainConfigList = v
	return s
}

func (s *SetVodDomainStagingConfigResponseBody) SetRequestId(v string) *SetVodDomainStagingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetVodDomainStagingConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type SetVodDomainStagingConfigResponseBodyDomainConfigList struct {
	ConfigId     *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
}

func (s SetVodDomainStagingConfigResponseBodyDomainConfigList) String() string {
	return dara.Prettify(s)
}

func (s SetVodDomainStagingConfigResponseBodyDomainConfigList) GoString() string {
	return s.String()
}

func (s *SetVodDomainStagingConfigResponseBodyDomainConfigList) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *SetVodDomainStagingConfigResponseBodyDomainConfigList) GetDomainName() *string {
	return s.DomainName
}

func (s *SetVodDomainStagingConfigResponseBodyDomainConfigList) GetFunctionName() *string {
	return s.FunctionName
}

func (s *SetVodDomainStagingConfigResponseBodyDomainConfigList) SetConfigId(v int64) *SetVodDomainStagingConfigResponseBodyDomainConfigList {
	s.ConfigId = &v
	return s
}

func (s *SetVodDomainStagingConfigResponseBodyDomainConfigList) SetDomainName(v string) *SetVodDomainStagingConfigResponseBodyDomainConfigList {
	s.DomainName = &v
	return s
}

func (s *SetVodDomainStagingConfigResponseBodyDomainConfigList) SetFunctionName(v string) *SetVodDomainStagingConfigResponseBodyDomainConfigList {
	s.FunctionName = &v
	return s
}

func (s *SetVodDomainStagingConfigResponseBodyDomainConfigList) Validate() error {
	return dara.Validate(s)
}
