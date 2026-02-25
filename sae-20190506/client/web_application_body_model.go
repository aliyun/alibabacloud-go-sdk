// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebApplicationBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *WebApplicationBody
	GetCode() *int32
	SetData(v *WebApplication) *WebApplicationBody
	GetData() *WebApplication
	SetMessage(v string) *WebApplicationBody
	GetMessage() *string
	SetRequestId(v string) *WebApplicationBody
	GetRequestId() *string
	SetSuccess(v bool) *WebApplicationBody
	GetSuccess() *bool
}

type WebApplicationBody struct {
	// The HTTP status code. Value description:
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
	Data *WebApplication `json:"Data,omitempty" xml:"Data,omitempty"`
	// Additional information about the call result. Value description:
	//
	// 	- If the request is successful, a success message is returned.
	//
	// 	- If the request fails, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WebApplicationBody) String() string {
	return dara.Prettify(s)
}

func (s WebApplicationBody) GoString() string {
	return s.String()
}

func (s *WebApplicationBody) GetCode() *int32 {
	return s.Code
}

func (s *WebApplicationBody) GetData() *WebApplication {
	return s.Data
}

func (s *WebApplicationBody) GetMessage() *string {
	return s.Message
}

func (s *WebApplicationBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WebApplicationBody) GetSuccess() *bool {
	return s.Success
}

func (s *WebApplicationBody) SetCode(v int32) *WebApplicationBody {
	s.Code = &v
	return s
}

func (s *WebApplicationBody) SetData(v *WebApplication) *WebApplicationBody {
	s.Data = v
	return s
}

func (s *WebApplicationBody) SetMessage(v string) *WebApplicationBody {
	s.Message = &v
	return s
}

func (s *WebApplicationBody) SetRequestId(v string) *WebApplicationBody {
	s.RequestId = &v
	return s
}

func (s *WebApplicationBody) SetSuccess(v bool) *WebApplicationBody {
	s.Success = &v
	return s
}

func (s *WebApplicationBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
