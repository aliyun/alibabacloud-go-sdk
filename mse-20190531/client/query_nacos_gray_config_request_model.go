// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryNacosGrayConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *QueryNacosGrayConfigRequest
	GetAcceptLanguage() *string
	SetDataId(v string) *QueryNacosGrayConfigRequest
	GetDataId() *string
	SetGrayName(v string) *QueryNacosGrayConfigRequest
	GetGrayName() *string
	SetGroup(v string) *QueryNacosGrayConfigRequest
	GetGroup() *string
	SetInstanceId(v string) *QueryNacosGrayConfigRequest
	GetInstanceId() *string
	SetNamespaceId(v string) *QueryNacosGrayConfigRequest
	GetNamespaceId() *string
	SetRegionId(v string) *QueryNacosGrayConfigRequest
	GetRegionId() *string
	SetRequestPars(v string) *QueryNacosGrayConfigRequest
	GetRequestPars() *string
}

type QueryNacosGrayConfigRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// demo-develop
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// example:
	//
	// test
	GrayName *string `json:"GrayName,omitempty" xml:"GrayName,omitempty"`
	// example:
	//
	// DEFAULT_GROUP
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mse-cn-st21ri2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 6bafdde3-4fa7-4d67-b3da-a607ab87f7e4
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s QueryNacosGrayConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryNacosGrayConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryNacosGrayConfigRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *QueryNacosGrayConfigRequest) GetDataId() *string {
	return s.DataId
}

func (s *QueryNacosGrayConfigRequest) GetGrayName() *string {
	return s.GrayName
}

func (s *QueryNacosGrayConfigRequest) GetGroup() *string {
	return s.Group
}

func (s *QueryNacosGrayConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryNacosGrayConfigRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *QueryNacosGrayConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryNacosGrayConfigRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *QueryNacosGrayConfigRequest) SetAcceptLanguage(v string) *QueryNacosGrayConfigRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *QueryNacosGrayConfigRequest) SetDataId(v string) *QueryNacosGrayConfigRequest {
	s.DataId = &v
	return s
}

func (s *QueryNacosGrayConfigRequest) SetGrayName(v string) *QueryNacosGrayConfigRequest {
	s.GrayName = &v
	return s
}

func (s *QueryNacosGrayConfigRequest) SetGroup(v string) *QueryNacosGrayConfigRequest {
	s.Group = &v
	return s
}

func (s *QueryNacosGrayConfigRequest) SetInstanceId(v string) *QueryNacosGrayConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryNacosGrayConfigRequest) SetNamespaceId(v string) *QueryNacosGrayConfigRequest {
	s.NamespaceId = &v
	return s
}

func (s *QueryNacosGrayConfigRequest) SetRegionId(v string) *QueryNacosGrayConfigRequest {
	s.RegionId = &v
	return s
}

func (s *QueryNacosGrayConfigRequest) SetRequestPars(v string) *QueryNacosGrayConfigRequest {
	s.RequestPars = &v
	return s
}

func (s *QueryNacosGrayConfigRequest) Validate() error {
	return dara.Validate(s)
}
