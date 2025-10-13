// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebApplicationScalingConfigBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *WebApplicationScalingConfigBody
	GetCode() *int32
	SetData(v *WebScalingConfig) *WebApplicationScalingConfigBody
	GetData() *WebScalingConfig
	SetMessage(v string) *WebApplicationScalingConfigBody
	GetMessage() *string
	SetRequestId(v string) *WebApplicationScalingConfigBody
	GetRequestId() *string
	SetSuccess(v bool) *WebApplicationScalingConfigBody
	GetSuccess() *bool
}

type WebApplicationScalingConfigBody struct {
	Code      *int32            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *WebScalingConfig `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WebApplicationScalingConfigBody) String() string {
	return dara.Prettify(s)
}

func (s WebApplicationScalingConfigBody) GoString() string {
	return s.String()
}

func (s *WebApplicationScalingConfigBody) GetCode() *int32 {
	return s.Code
}

func (s *WebApplicationScalingConfigBody) GetData() *WebScalingConfig {
	return s.Data
}

func (s *WebApplicationScalingConfigBody) GetMessage() *string {
	return s.Message
}

func (s *WebApplicationScalingConfigBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WebApplicationScalingConfigBody) GetSuccess() *bool {
	return s.Success
}

func (s *WebApplicationScalingConfigBody) SetCode(v int32) *WebApplicationScalingConfigBody {
	s.Code = &v
	return s
}

func (s *WebApplicationScalingConfigBody) SetData(v *WebScalingConfig) *WebApplicationScalingConfigBody {
	s.Data = v
	return s
}

func (s *WebApplicationScalingConfigBody) SetMessage(v string) *WebApplicationScalingConfigBody {
	s.Message = &v
	return s
}

func (s *WebApplicationScalingConfigBody) SetRequestId(v string) *WebApplicationScalingConfigBody {
	s.RequestId = &v
	return s
}

func (s *WebApplicationScalingConfigBody) SetSuccess(v bool) *WebApplicationScalingConfigBody {
	s.Success = &v
	return s
}

func (s *WebApplicationScalingConfigBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
