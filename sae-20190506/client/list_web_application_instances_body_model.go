// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebApplicationInstancesBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListWebApplicationInstancesBody
	GetCode() *int32
	SetData(v *ListWebInstancesOutput) *ListWebApplicationInstancesBody
	GetData() *ListWebInstancesOutput
	SetMessage(v string) *ListWebApplicationInstancesBody
	GetMessage() *string
	SetRequestId(v string) *ListWebApplicationInstancesBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWebApplicationInstancesBody
	GetSuccess() *bool
}

type ListWebApplicationInstancesBody struct {
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
	Data *ListWebInstancesOutput `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s ListWebApplicationInstancesBody) String() string {
	return dara.Prettify(s)
}

func (s ListWebApplicationInstancesBody) GoString() string {
	return s.String()
}

func (s *ListWebApplicationInstancesBody) GetCode() *int32 {
	return s.Code
}

func (s *ListWebApplicationInstancesBody) GetData() *ListWebInstancesOutput {
	return s.Data
}

func (s *ListWebApplicationInstancesBody) GetMessage() *string {
	return s.Message
}

func (s *ListWebApplicationInstancesBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWebApplicationInstancesBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWebApplicationInstancesBody) SetCode(v int32) *ListWebApplicationInstancesBody {
	s.Code = &v
	return s
}

func (s *ListWebApplicationInstancesBody) SetData(v *ListWebInstancesOutput) *ListWebApplicationInstancesBody {
	s.Data = v
	return s
}

func (s *ListWebApplicationInstancesBody) SetMessage(v string) *ListWebApplicationInstancesBody {
	s.Message = &v
	return s
}

func (s *ListWebApplicationInstancesBody) SetRequestId(v string) *ListWebApplicationInstancesBody {
	s.RequestId = &v
	return s
}

func (s *ListWebApplicationInstancesBody) SetSuccess(v bool) *ListWebApplicationInstancesBody {
	s.Success = &v
	return s
}

func (s *ListWebApplicationInstancesBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
