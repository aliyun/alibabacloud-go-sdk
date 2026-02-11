// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyBastionAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ApplyBastionAccountResponseBody
	GetCode() *int64
	SetRequestId(v string) *ApplyBastionAccountResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *ApplyBastionAccountResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *ApplyBastionAccountResponseBody
	GetSuccess() *bool
}

type ApplyBastionAccountResponseBody struct {
	// Status code. A return value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 4A91D2D1-AEC9-1658-ABCE-5A39DE66A5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Result object.
	//
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
	// Whether the call was successful.
	//
	// - **true**: Call succeeded.
	//
	// - **false**: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ApplyBastionAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyBastionAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyBastionAccountResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ApplyBastionAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyBastionAccountResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *ApplyBastionAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyBastionAccountResponseBody) SetCode(v int64) *ApplyBastionAccountResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyBastionAccountResponseBody) SetRequestId(v string) *ApplyBastionAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyBastionAccountResponseBody) SetResultObject(v bool) *ApplyBastionAccountResponseBody {
	s.ResultObject = &v
	return s
}

func (s *ApplyBastionAccountResponseBody) SetSuccess(v bool) *ApplyBastionAccountResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyBastionAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
