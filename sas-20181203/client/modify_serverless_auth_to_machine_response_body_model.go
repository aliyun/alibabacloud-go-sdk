// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyServerlessAuthToMachineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyServerlessAuthToMachineResponseBodyData) *ModifyServerlessAuthToMachineResponseBody
	GetData() *ModifyServerlessAuthToMachineResponseBodyData
	SetRequestId(v string) *ModifyServerlessAuthToMachineResponseBody
	GetRequestId() *string
}

type ModifyServerlessAuthToMachineResponseBody struct {
	// Details of the returned data.
	Data *ModifyServerlessAuthToMachineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 本次调用请求的ID，是由阿里云为该请求生成的唯一标识符，可用于排查和定位问题。
	//
	// example:
	//
	// A47F77A1***8CD37050E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyServerlessAuthToMachineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyServerlessAuthToMachineResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyServerlessAuthToMachineResponseBody) GetData() *ModifyServerlessAuthToMachineResponseBodyData {
	return s.Data
}

func (s *ModifyServerlessAuthToMachineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyServerlessAuthToMachineResponseBody) SetData(v *ModifyServerlessAuthToMachineResponseBodyData) *ModifyServerlessAuthToMachineResponseBody {
	s.Data = v
	return s
}

func (s *ModifyServerlessAuthToMachineResponseBody) SetRequestId(v string) *ModifyServerlessAuthToMachineResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyServerlessAuthToMachineResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyServerlessAuthToMachineResponseBodyData struct {
	// Result code. Values:
	//
	// - **0**: Success
	//
	// - **1**: Parameter error
	//
	// example:
	//
	// 0
	ResultCode *int32 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
}

func (s ModifyServerlessAuthToMachineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyServerlessAuthToMachineResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyServerlessAuthToMachineResponseBodyData) GetResultCode() *int32 {
	return s.ResultCode
}

func (s *ModifyServerlessAuthToMachineResponseBodyData) SetResultCode(v int32) *ModifyServerlessAuthToMachineResponseBodyData {
	s.ResultCode = &v
	return s
}

func (s *ModifyServerlessAuthToMachineResponseBodyData) Validate() error {
	return dara.Validate(s)
}
