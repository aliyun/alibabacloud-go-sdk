// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticBizBandWidthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetElasticBizBandwidth(v int32) *ModifyElasticBizBandWidthRequest
	GetElasticBizBandwidth() *int32
	SetInstanceId(v string) *ModifyElasticBizBandWidthRequest
	GetInstanceId() *string
	SetMode(v string) *ModifyElasticBizBandWidthRequest
	GetMode() *string
}

type ModifyElasticBizBandWidthRequest struct {
	// This parameter is required.
	ElasticBizBandwidth *int32 `json:"ElasticBizBandwidth,omitempty" xml:"ElasticBizBandwidth,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Mode       *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s ModifyElasticBizBandWidthRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticBizBandWidthRequest) GoString() string {
	return s.String()
}

func (s *ModifyElasticBizBandWidthRequest) GetElasticBizBandwidth() *int32 {
	return s.ElasticBizBandwidth
}

func (s *ModifyElasticBizBandWidthRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyElasticBizBandWidthRequest) GetMode() *string {
	return s.Mode
}

func (s *ModifyElasticBizBandWidthRequest) SetElasticBizBandwidth(v int32) *ModifyElasticBizBandWidthRequest {
	s.ElasticBizBandwidth = &v
	return s
}

func (s *ModifyElasticBizBandWidthRequest) SetInstanceId(v string) *ModifyElasticBizBandWidthRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyElasticBizBandWidthRequest) SetMode(v string) *ModifyElasticBizBandWidthRequest {
	s.Mode = &v
	return s
}

func (s *ModifyElasticBizBandWidthRequest) Validate() error {
	return dara.Validate(s)
}
