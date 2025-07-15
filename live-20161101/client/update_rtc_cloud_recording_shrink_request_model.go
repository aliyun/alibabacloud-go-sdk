// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRtcCloudRecordingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMixLayoutParamsShrink(v string) *UpdateRtcCloudRecordingShrinkRequest
	GetMixLayoutParamsShrink() *string
	SetSubscribeParamsShrink(v string) *UpdateRtcCloudRecordingShrinkRequest
	GetSubscribeParamsShrink() *string
	SetTaskId(v string) *UpdateRtcCloudRecordingShrinkRequest
	GetTaskId() *string
}

type UpdateRtcCloudRecordingShrinkRequest struct {
	MixLayoutParamsShrink *string `json:"MixLayoutParams,omitempty" xml:"MixLayoutParams,omitempty"`
	// This parameter is required.
	SubscribeParamsShrink *string `json:"SubscribeParams,omitempty" xml:"SubscribeParams,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ******73-8501-****-8ac1-72295a******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateRtcCloudRecordingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtcCloudRecordingShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRtcCloudRecordingShrinkRequest) GetMixLayoutParamsShrink() *string {
	return s.MixLayoutParamsShrink
}

func (s *UpdateRtcCloudRecordingShrinkRequest) GetSubscribeParamsShrink() *string {
	return s.SubscribeParamsShrink
}

func (s *UpdateRtcCloudRecordingShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateRtcCloudRecordingShrinkRequest) SetMixLayoutParamsShrink(v string) *UpdateRtcCloudRecordingShrinkRequest {
	s.MixLayoutParamsShrink = &v
	return s
}

func (s *UpdateRtcCloudRecordingShrinkRequest) SetSubscribeParamsShrink(v string) *UpdateRtcCloudRecordingShrinkRequest {
	s.SubscribeParamsShrink = &v
	return s
}

func (s *UpdateRtcCloudRecordingShrinkRequest) SetTaskId(v string) *UpdateRtcCloudRecordingShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateRtcCloudRecordingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
