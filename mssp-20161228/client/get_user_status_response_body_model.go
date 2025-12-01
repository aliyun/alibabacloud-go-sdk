// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUserStatusResponseBody
	GetCode() *string
	SetData(v *GetUserStatusResponseBodyData) *GetUserStatusResponseBody
	GetData() *GetUserStatusResponseBodyData
	SetHttpStatusCode(v int32) *GetUserStatusResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetUserStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUserStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUserStatusResponseBody
	GetSuccess() *bool
}

type GetUserStatusResponseBody struct {
	// Interface response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data *GetUserStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message of the returned result.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// D8DBD769-613E-5E6B-A9FD-B622375B152D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUserStatusResponseBody) GetData() *GetUserStatusResponseBodyData {
	return s.Data
}

func (s *GetUserStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetUserStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUserStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserStatusResponseBody) SetCode(v string) *GetUserStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserStatusResponseBody) SetData(v *GetUserStatusResponseBodyData) *GetUserStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetUserStatusResponseBody) SetHttpStatusCode(v int32) *GetUserStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUserStatusResponseBody) SetMessage(v string) *GetUserStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserStatusResponseBody) SetRequestId(v string) *GetUserStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserStatusResponseBody) SetSuccess(v bool) *GetUserStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserStatusResponseBodyData struct {
	// Customer type.
	//
	// example:
	//
	// official
	CustomerType *string `json:"CustomerType,omitempty" xml:"CustomerType,omitempty"`
	// End date.
	//
	// example:
	//
	// 2023-09-28 00:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// 726cec3c-4887-4354-8c21-c0ad12e10fc2
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Start date.
	//
	// example:
	//
	// 2023-09-20 00:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Status.
	//
	// example:
	//
	// FirstLogin
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Version.
	//
	// example:
	//
	// mdrjichu
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetUserStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUserStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserStatusResponseBodyData) GetCustomerType() *string {
	return s.CustomerType
}

func (s *GetUserStatusResponseBodyData) GetEndDate() *string {
	return s.EndDate
}

func (s *GetUserStatusResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetUserStatusResponseBodyData) GetStartDate() *string {
	return s.StartDate
}

func (s *GetUserStatusResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetUserStatusResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *GetUserStatusResponseBodyData) SetCustomerType(v string) *GetUserStatusResponseBodyData {
	s.CustomerType = &v
	return s
}

func (s *GetUserStatusResponseBodyData) SetEndDate(v string) *GetUserStatusResponseBodyData {
	s.EndDate = &v
	return s
}

func (s *GetUserStatusResponseBodyData) SetInstanceId(v string) *GetUserStatusResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetUserStatusResponseBodyData) SetStartDate(v string) *GetUserStatusResponseBodyData {
	s.StartDate = &v
	return s
}

func (s *GetUserStatusResponseBodyData) SetStatus(v string) *GetUserStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetUserStatusResponseBodyData) SetVersion(v string) *GetUserStatusResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetUserStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
