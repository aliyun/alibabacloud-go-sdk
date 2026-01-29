// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateCostUnitResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AllocateCostUnitResourceResponseBody
	GetCode() *string
	SetData(v *AllocateCostUnitResourceResponseBodyData) *AllocateCostUnitResourceResponseBody
	GetData() *AllocateCostUnitResourceResponseBodyData
	SetMessage(v string) *AllocateCostUnitResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *AllocateCostUnitResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AllocateCostUnitResourceResponseBody
	GetSuccess() *bool
}

type AllocateCostUnitResourceResponseBody struct {
	// The status code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *AllocateCostUnitResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// Successful！
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 845C629F-47A7-4F46-A470-ED5047C4C250
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AllocateCostUnitResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateCostUnitResourceResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateCostUnitResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *AllocateCostUnitResourceResponseBody) GetData() *AllocateCostUnitResourceResponseBodyData {
	return s.Data
}

func (s *AllocateCostUnitResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AllocateCostUnitResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateCostUnitResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AllocateCostUnitResourceResponseBody) SetCode(v string) *AllocateCostUnitResourceResponseBody {
	s.Code = &v
	return s
}

func (s *AllocateCostUnitResourceResponseBody) SetData(v *AllocateCostUnitResourceResponseBodyData) *AllocateCostUnitResourceResponseBody {
	s.Data = v
	return s
}

func (s *AllocateCostUnitResourceResponseBody) SetMessage(v string) *AllocateCostUnitResourceResponseBody {
	s.Message = &v
	return s
}

func (s *AllocateCostUnitResourceResponseBody) SetRequestId(v string) *AllocateCostUnitResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateCostUnitResourceResponseBody) SetSuccess(v bool) *AllocateCostUnitResourceResponseBody {
	s.Success = &v
	return s
}

func (s *AllocateCostUnitResourceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AllocateCostUnitResourceResponseBodyData struct {
	// Indicates whether resources are allocated to the specified cost center. Valid values:
	//
	// 	- true: The resources are allocated to the specified cost center.
	//
	// 	- false: The resources fail to be allocated to the specified cost center.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the destination cost center.
	//
	// example:
	//
	// 32857346527
	ToUnitId *int64 `json:"ToUnitId,omitempty" xml:"ToUnitId,omitempty"`
	// The user ID of the owner of the destination cost center.
	//
	// example:
	//
	// 34857693874
	ToUnitUserId *int64 `json:"ToUnitUserId,omitempty" xml:"ToUnitUserId,omitempty"`
}

func (s AllocateCostUnitResourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AllocateCostUnitResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AllocateCostUnitResourceResponseBodyData) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *AllocateCostUnitResourceResponseBodyData) GetToUnitId() *int64 {
	return s.ToUnitId
}

func (s *AllocateCostUnitResourceResponseBodyData) GetToUnitUserId() *int64 {
	return s.ToUnitUserId
}

func (s *AllocateCostUnitResourceResponseBodyData) SetIsSuccess(v bool) *AllocateCostUnitResourceResponseBodyData {
	s.IsSuccess = &v
	return s
}

func (s *AllocateCostUnitResourceResponseBodyData) SetToUnitId(v int64) *AllocateCostUnitResourceResponseBodyData {
	s.ToUnitId = &v
	return s
}

func (s *AllocateCostUnitResourceResponseBodyData) SetToUnitUserId(v int64) *AllocateCostUnitResourceResponseBodyData {
	s.ToUnitUserId = &v
	return s
}

func (s *AllocateCostUnitResourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
