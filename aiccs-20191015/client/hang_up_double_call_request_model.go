// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHangUpDoubleCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcid(v string) *HangUpDoubleCallRequest
	GetAcid() *string
	SetInstanceId(v string) *HangUpDoubleCallRequest
	GetInstanceId() *string
}

type HangUpDoubleCallRequest struct {
	// Session ID.
	//
	// example:
	//
	// 68255155****
	Acid *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	// Artificial Intelligence Cloud Call Service instance ID. You can obtain it from the Artificial Intelligence Cloud Call Service console.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s HangUpDoubleCallRequest) String() string {
	return dara.Prettify(s)
}

func (s HangUpDoubleCallRequest) GoString() string {
	return s.String()
}

func (s *HangUpDoubleCallRequest) GetAcid() *string {
	return s.Acid
}

func (s *HangUpDoubleCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *HangUpDoubleCallRequest) SetAcid(v string) *HangUpDoubleCallRequest {
	s.Acid = &v
	return s
}

func (s *HangUpDoubleCallRequest) SetInstanceId(v string) *HangUpDoubleCallRequest {
	s.InstanceId = &v
	return s
}

func (s *HangUpDoubleCallRequest) Validate() error {
	return dara.Validate(s)
}
