// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticBandWidthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetElasticBandwidth(v int32) *ModifyElasticBandWidthRequest
	GetElasticBandwidth() *int32
	SetInstanceId(v string) *ModifyElasticBandWidthRequest
	GetInstanceId() *string
	SetSourceIp(v string) *ModifyElasticBandWidthRequest
	GetSourceIp() *string
}

type ModifyElasticBandWidthRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30
	ElasticBandwidth *int32 `json:"ElasticBandwidth,omitempty" xml:"ElasticBandwidth,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ModifyElasticBandWidthRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticBandWidthRequest) GoString() string {
	return s.String()
}

func (s *ModifyElasticBandWidthRequest) GetElasticBandwidth() *int32 {
	return s.ElasticBandwidth
}

func (s *ModifyElasticBandWidthRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyElasticBandWidthRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyElasticBandWidthRequest) SetElasticBandwidth(v int32) *ModifyElasticBandWidthRequest {
	s.ElasticBandwidth = &v
	return s
}

func (s *ModifyElasticBandWidthRequest) SetInstanceId(v string) *ModifyElasticBandWidthRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyElasticBandWidthRequest) SetSourceIp(v string) *ModifyElasticBandWidthRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyElasticBandWidthRequest) Validate() error {
	return dara.Validate(s)
}
