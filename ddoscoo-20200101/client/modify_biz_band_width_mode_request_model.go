// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBizBandWidthModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyBizBandWidthModeRequest
	GetInstanceId() *string
	SetMode(v string) *ModifyBizBandWidthModeRequest
	GetMode() *string
}

type ModifyBizBandWidthModeRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s ModifyBizBandWidthModeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBizBandWidthModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyBizBandWidthModeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyBizBandWidthModeRequest) GetMode() *string {
	return s.Mode
}

func (s *ModifyBizBandWidthModeRequest) SetInstanceId(v string) *ModifyBizBandWidthModeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyBizBandWidthModeRequest) SetMode(v string) *ModifyBizBandWidthModeRequest {
	s.Mode = &v
	return s
}

func (s *ModifyBizBandWidthModeRequest) Validate() error {
	return dara.Validate(s)
}
