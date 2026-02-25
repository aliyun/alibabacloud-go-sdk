// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebApplicationResourceStaticsBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *WebApplicationResourceStaticsBody
	GetCode() *int32
	SetData(v *DescribeWebAppStaticsOutput) *WebApplicationResourceStaticsBody
	GetData() *DescribeWebAppStaticsOutput
	SetMessage(v string) *WebApplicationResourceStaticsBody
	GetMessage() *string
	SetRequestId(v string) *WebApplicationResourceStaticsBody
	GetRequestId() *string
	SetSuccess(v bool) *WebApplicationResourceStaticsBody
	GetSuccess() *bool
}

type WebApplicationResourceStaticsBody struct {
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
	// The data returned for the operation.
	Data *DescribeWebAppStaticsOutput `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned for the operation. Valid values:
	//
	// 	- success is returned when the request succeeds.
	//
	// 	- An error code is returned if the request failed.
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
	// Indicates whether the call was successful. Valid values:
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

func (s WebApplicationResourceStaticsBody) String() string {
	return dara.Prettify(s)
}

func (s WebApplicationResourceStaticsBody) GoString() string {
	return s.String()
}

func (s *WebApplicationResourceStaticsBody) GetCode() *int32 {
	return s.Code
}

func (s *WebApplicationResourceStaticsBody) GetData() *DescribeWebAppStaticsOutput {
	return s.Data
}

func (s *WebApplicationResourceStaticsBody) GetMessage() *string {
	return s.Message
}

func (s *WebApplicationResourceStaticsBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WebApplicationResourceStaticsBody) GetSuccess() *bool {
	return s.Success
}

func (s *WebApplicationResourceStaticsBody) SetCode(v int32) *WebApplicationResourceStaticsBody {
	s.Code = &v
	return s
}

func (s *WebApplicationResourceStaticsBody) SetData(v *DescribeWebAppStaticsOutput) *WebApplicationResourceStaticsBody {
	s.Data = v
	return s
}

func (s *WebApplicationResourceStaticsBody) SetMessage(v string) *WebApplicationResourceStaticsBody {
	s.Message = &v
	return s
}

func (s *WebApplicationResourceStaticsBody) SetRequestId(v string) *WebApplicationResourceStaticsBody {
	s.RequestId = &v
	return s
}

func (s *WebApplicationResourceStaticsBody) SetSuccess(v bool) *WebApplicationResourceStaticsBody {
	s.Success = &v
	return s
}

func (s *WebApplicationResourceStaticsBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
