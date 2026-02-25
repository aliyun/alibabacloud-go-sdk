// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebApplicationTrafficConfigBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *WebApplicationTrafficConfigBody
	GetCode() *int32
	SetData(v *WebTrafficConfig) *WebApplicationTrafficConfigBody
	GetData() *WebTrafficConfig
	SetMessage(v string) *WebApplicationTrafficConfigBody
	GetMessage() *string
	SetRequestId(v string) *WebApplicationTrafficConfigBody
	GetRequestId() *string
	SetSuccess(v bool) *WebApplicationTrafficConfigBody
	GetSuccess() *bool
}

type WebApplicationTrafficConfigBody struct {
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
	Data *WebTrafficConfig `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s WebApplicationTrafficConfigBody) String() string {
	return dara.Prettify(s)
}

func (s WebApplicationTrafficConfigBody) GoString() string {
	return s.String()
}

func (s *WebApplicationTrafficConfigBody) GetCode() *int32 {
	return s.Code
}

func (s *WebApplicationTrafficConfigBody) GetData() *WebTrafficConfig {
	return s.Data
}

func (s *WebApplicationTrafficConfigBody) GetMessage() *string {
	return s.Message
}

func (s *WebApplicationTrafficConfigBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WebApplicationTrafficConfigBody) GetSuccess() *bool {
	return s.Success
}

func (s *WebApplicationTrafficConfigBody) SetCode(v int32) *WebApplicationTrafficConfigBody {
	s.Code = &v
	return s
}

func (s *WebApplicationTrafficConfigBody) SetData(v *WebTrafficConfig) *WebApplicationTrafficConfigBody {
	s.Data = v
	return s
}

func (s *WebApplicationTrafficConfigBody) SetMessage(v string) *WebApplicationTrafficConfigBody {
	s.Message = &v
	return s
}

func (s *WebApplicationTrafficConfigBody) SetRequestId(v string) *WebApplicationTrafficConfigBody {
	s.RequestId = &v
	return s
}

func (s *WebApplicationTrafficConfigBody) SetSuccess(v bool) *WebApplicationTrafficConfigBody {
	s.Success = &v
	return s
}

func (s *WebApplicationTrafficConfigBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
