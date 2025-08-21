// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstanceOperateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *InstanceOperateResponse
	GetCode() *int64
	SetInstanceId(v string) *InstanceOperateResponse
	GetInstanceId() *string
	SetMessage(v string) *InstanceOperateResponse
	GetMessage() *string
}

type InstanceOperateResponse struct {
	Code       *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s InstanceOperateResponse) String() string {
	return dara.Prettify(s)
}

func (s InstanceOperateResponse) GoString() string {
	return s.String()
}

func (s *InstanceOperateResponse) GetCode() *int64 {
	return s.Code
}

func (s *InstanceOperateResponse) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InstanceOperateResponse) GetMessage() *string {
	return s.Message
}

func (s *InstanceOperateResponse) SetCode(v int64) *InstanceOperateResponse {
	s.Code = &v
	return s
}

func (s *InstanceOperateResponse) SetInstanceId(v string) *InstanceOperateResponse {
	s.InstanceId = &v
	return s
}

func (s *InstanceOperateResponse) SetMessage(v string) *InstanceOperateResponse {
	s.Message = &v
	return s
}

func (s *InstanceOperateResponse) Validate() error {
	return dara.Validate(s)
}
