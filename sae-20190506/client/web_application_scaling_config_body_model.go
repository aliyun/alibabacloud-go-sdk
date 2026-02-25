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
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The request was successful.
	//
	// 	- **3xx**: The request was redirected.
	//
	// 	- **4xx**: The request failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *WebScalingConfig `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message. Valid values:
	//
	// 	- If the request was successful, a success message is returned.
	//
	// 	- If the request failed, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
