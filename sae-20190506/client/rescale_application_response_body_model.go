// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRescaleApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RescaleApplicationResponseBody
	GetCode() *string
	SetData(v *RescaleApplicationResponseBodyData) *RescaleApplicationResponseBody
	GetData() *RescaleApplicationResponseBodyData
	SetErrorCode(v string) *RescaleApplicationResponseBody
	GetErrorCode() *string
	SetMessage(v string) *RescaleApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *RescaleApplicationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RescaleApplicationResponseBody
	GetSuccess() *bool
}

type RescaleApplicationResponseBody struct {
	// The HTTP status code. Take note of the following rules:
	//
	// 	- **2xx**: The call was successful.
	//
	// 	- **3xx**: The call was redirected.
	//
	// 	- **4xx**: The call failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response.
	Data *RescaleApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed. Take note of the following rules:
	//
	// 	- The **ErrorCode*	- parameter is not returned if the request succeeds.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the "**Error codes**" section of this topic.
	//
	// example:
	//
	// Null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned for the operation.
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
	// Indicates whether the application is successfully scaled. Take note of the following rules:
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

func (s RescaleApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RescaleApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *RescaleApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *RescaleApplicationResponseBody) GetData() *RescaleApplicationResponseBodyData {
	return s.Data
}

func (s *RescaleApplicationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RescaleApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RescaleApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RescaleApplicationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RescaleApplicationResponseBody) SetCode(v string) *RescaleApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *RescaleApplicationResponseBody) SetData(v *RescaleApplicationResponseBodyData) *RescaleApplicationResponseBody {
	s.Data = v
	return s
}

func (s *RescaleApplicationResponseBody) SetErrorCode(v string) *RescaleApplicationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RescaleApplicationResponseBody) SetMessage(v string) *RescaleApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *RescaleApplicationResponseBody) SetRequestId(v string) *RescaleApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RescaleApplicationResponseBody) SetSuccess(v bool) *RescaleApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *RescaleApplicationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RescaleApplicationResponseBodyData struct {
	// The ID of the change order.
	//
	// example:
	//
	// 0e09829f-4914-4612-bc88-6f52fd81****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s RescaleApplicationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RescaleApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *RescaleApplicationResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *RescaleApplicationResponseBodyData) SetChangeOrderId(v string) *RescaleApplicationResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *RescaleApplicationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
