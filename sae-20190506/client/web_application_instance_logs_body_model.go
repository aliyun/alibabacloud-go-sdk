// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebApplicationInstanceLogsBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *WebApplicationInstanceLogsBody
	GetCode() *int32
	SetData(v *DescribeInstanceLogsOutput) *WebApplicationInstanceLogsBody
	GetData() *DescribeInstanceLogsOutput
	SetMessage(v string) *WebApplicationInstanceLogsBody
	GetMessage() *string
	SetRequestId(v string) *WebApplicationInstanceLogsBody
	GetRequestId() *string
	SetSuccess(v bool) *WebApplicationInstanceLogsBody
	GetSuccess() *bool
}

type WebApplicationInstanceLogsBody struct {
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
	Data *DescribeInstanceLogsOutput `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s WebApplicationInstanceLogsBody) String() string {
	return dara.Prettify(s)
}

func (s WebApplicationInstanceLogsBody) GoString() string {
	return s.String()
}

func (s *WebApplicationInstanceLogsBody) GetCode() *int32 {
	return s.Code
}

func (s *WebApplicationInstanceLogsBody) GetData() *DescribeInstanceLogsOutput {
	return s.Data
}

func (s *WebApplicationInstanceLogsBody) GetMessage() *string {
	return s.Message
}

func (s *WebApplicationInstanceLogsBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WebApplicationInstanceLogsBody) GetSuccess() *bool {
	return s.Success
}

func (s *WebApplicationInstanceLogsBody) SetCode(v int32) *WebApplicationInstanceLogsBody {
	s.Code = &v
	return s
}

func (s *WebApplicationInstanceLogsBody) SetData(v *DescribeInstanceLogsOutput) *WebApplicationInstanceLogsBody {
	s.Data = v
	return s
}

func (s *WebApplicationInstanceLogsBody) SetMessage(v string) *WebApplicationInstanceLogsBody {
	s.Message = &v
	return s
}

func (s *WebApplicationInstanceLogsBody) SetRequestId(v string) *WebApplicationInstanceLogsBody {
	s.RequestId = &v
	return s
}

func (s *WebApplicationInstanceLogsBody) SetSuccess(v bool) *WebApplicationInstanceLogsBody {
	s.Success = &v
	return s
}

func (s *WebApplicationInstanceLogsBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
