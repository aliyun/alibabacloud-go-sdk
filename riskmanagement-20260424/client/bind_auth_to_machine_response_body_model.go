// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAuthToMachineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BindAuthToMachineResponseBody
	GetCode() *string
	SetData(v *BindAuthToMachineResponseBodyData) *BindAuthToMachineResponseBody
	GetData() *BindAuthToMachineResponseBodyData
	SetMessage(v string) *BindAuthToMachineResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindAuthToMachineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BindAuthToMachineResponseBody
	GetSuccess() *bool
}

type BindAuthToMachineResponseBody struct {
	// example:
	//
	// OK
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *BindAuthToMachineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// BEE90F8C-EDC2-5394-953B-D07A121612B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindAuthToMachineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindAuthToMachineResponseBody) GoString() string {
	return s.String()
}

func (s *BindAuthToMachineResponseBody) GetCode() *string {
	return s.Code
}

func (s *BindAuthToMachineResponseBody) GetData() *BindAuthToMachineResponseBodyData {
	return s.Data
}

func (s *BindAuthToMachineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindAuthToMachineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindAuthToMachineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BindAuthToMachineResponseBody) SetCode(v string) *BindAuthToMachineResponseBody {
	s.Code = &v
	return s
}

func (s *BindAuthToMachineResponseBody) SetData(v *BindAuthToMachineResponseBodyData) *BindAuthToMachineResponseBody {
	s.Data = v
	return s
}

func (s *BindAuthToMachineResponseBody) SetMessage(v string) *BindAuthToMachineResponseBody {
	s.Message = &v
	return s
}

func (s *BindAuthToMachineResponseBody) SetRequestId(v string) *BindAuthToMachineResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAuthToMachineResponseBody) SetSuccess(v bool) *BindAuthToMachineResponseBody {
	s.Success = &v
	return s
}

func (s *BindAuthToMachineResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindAuthToMachineResponseBodyData struct {
	Body *BindAuthToMachineResponseBodyDataBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
}

func (s BindAuthToMachineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BindAuthToMachineResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindAuthToMachineResponseBodyData) GetBody() *BindAuthToMachineResponseBodyDataBody {
	return s.Body
}

func (s *BindAuthToMachineResponseBodyData) SetBody(v *BindAuthToMachineResponseBodyDataBody) *BindAuthToMachineResponseBodyData {
	s.Body = v
	return s
}

func (s *BindAuthToMachineResponseBodyData) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindAuthToMachineResponseBodyDataBody struct {
	// example:
	//
	// 5
	BindCount *int32 `json:"BindCount,omitempty" xml:"BindCount,omitempty"`
	// example:
	//
	// 2
	InsufficientCoreCount *int32 `json:"InsufficientCoreCount,omitempty" xml:"InsufficientCoreCount,omitempty"`
	// example:
	//
	// 1
	InsufficientEcsCount *int32 `json:"InsufficientEcsCount,omitempty" xml:"InsufficientEcsCount,omitempty"`
	// example:
	//
	// F799C1E4-D4C6-5964-A6D1-4BA9CCF105F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	ResultCode *int32 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// 4
	UnBindCount *int32 `json:"UnBindCount,omitempty" xml:"UnBindCount,omitempty"`
}

func (s BindAuthToMachineResponseBodyDataBody) String() string {
	return dara.Prettify(s)
}

func (s BindAuthToMachineResponseBodyDataBody) GoString() string {
	return s.String()
}

func (s *BindAuthToMachineResponseBodyDataBody) GetBindCount() *int32 {
	return s.BindCount
}

func (s *BindAuthToMachineResponseBodyDataBody) GetInsufficientCoreCount() *int32 {
	return s.InsufficientCoreCount
}

func (s *BindAuthToMachineResponseBodyDataBody) GetInsufficientEcsCount() *int32 {
	return s.InsufficientEcsCount
}

func (s *BindAuthToMachineResponseBodyDataBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindAuthToMachineResponseBodyDataBody) GetResultCode() *int32 {
	return s.ResultCode
}

func (s *BindAuthToMachineResponseBodyDataBody) GetUnBindCount() *int32 {
	return s.UnBindCount
}

func (s *BindAuthToMachineResponseBodyDataBody) SetBindCount(v int32) *BindAuthToMachineResponseBodyDataBody {
	s.BindCount = &v
	return s
}

func (s *BindAuthToMachineResponseBodyDataBody) SetInsufficientCoreCount(v int32) *BindAuthToMachineResponseBodyDataBody {
	s.InsufficientCoreCount = &v
	return s
}

func (s *BindAuthToMachineResponseBodyDataBody) SetInsufficientEcsCount(v int32) *BindAuthToMachineResponseBodyDataBody {
	s.InsufficientEcsCount = &v
	return s
}

func (s *BindAuthToMachineResponseBodyDataBody) SetRequestId(v string) *BindAuthToMachineResponseBodyDataBody {
	s.RequestId = &v
	return s
}

func (s *BindAuthToMachineResponseBodyDataBody) SetResultCode(v int32) *BindAuthToMachineResponseBodyDataBody {
	s.ResultCode = &v
	return s
}

func (s *BindAuthToMachineResponseBodyDataBody) SetUnBindCount(v int32) *BindAuthToMachineResponseBodyDataBody {
	s.UnBindCount = &v
	return s
}

func (s *BindAuthToMachineResponseBodyDataBody) Validate() error {
	return dara.Validate(s)
}
