// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMOQuotaAlertThresholdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApikeyShrink(v string) *UpdateMOQuotaAlertThresholdShrinkRequest
	GetApikeyShrink() *string
	SetInstanceId(v string) *UpdateMOQuotaAlertThresholdShrinkRequest
	GetInstanceId() *string
}

type UpdateMOQuotaAlertThresholdShrinkRequest struct {
	// This parameter is required.
	ApikeyShrink *string `json:"Apikey,omitempty" xml:"Apikey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rds_copilot***_public_cn-*********6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateMOQuotaAlertThresholdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMOQuotaAlertThresholdShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMOQuotaAlertThresholdShrinkRequest) GetApikeyShrink() *string {
	return s.ApikeyShrink
}

func (s *UpdateMOQuotaAlertThresholdShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateMOQuotaAlertThresholdShrinkRequest) SetApikeyShrink(v string) *UpdateMOQuotaAlertThresholdShrinkRequest {
	s.ApikeyShrink = &v
	return s
}

func (s *UpdateMOQuotaAlertThresholdShrinkRequest) SetInstanceId(v string) *UpdateMOQuotaAlertThresholdShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateMOQuotaAlertThresholdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
