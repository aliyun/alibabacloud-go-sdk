// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckExpandStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CheckExpandStatusResponseBodyData) *CheckExpandStatusResponseBody
	GetData() *CheckExpandStatusResponseBodyData
	SetRequestId(v string) *CheckExpandStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckExpandStatusResponseBody
	GetSuccess() *bool
}

type CheckExpandStatusResponseBody struct {
	// The result of the verification.
	Data *CheckExpandStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 7CDBA7D5-8D62-4D24-9C65-510D62******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckExpandStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckExpandStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckExpandStatusResponseBody) GetData() *CheckExpandStatusResponseBodyData {
	return s.Data
}

func (s *CheckExpandStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckExpandStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckExpandStatusResponseBody) SetData(v *CheckExpandStatusResponseBodyData) *CheckExpandStatusResponseBody {
	s.Data = v
	return s
}

func (s *CheckExpandStatusResponseBody) SetRequestId(v string) *CheckExpandStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckExpandStatusResponseBody) SetSuccess(v bool) *CheckExpandStatusResponseBody {
	s.Success = &v
	return s
}

func (s *CheckExpandStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type CheckExpandStatusResponseBodyData struct {
	// Indicates whether scale-out operations can be performed on the database.
	//
	// example:
	//
	// true
	IsActive *bool `json:"IsActive,omitempty" xml:"IsActive,omitempty"`
	// The additional information.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
}

func (s CheckExpandStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CheckExpandStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckExpandStatusResponseBodyData) GetIsActive() *bool {
	return s.IsActive
}

func (s *CheckExpandStatusResponseBodyData) GetMsg() *string {
	return s.Msg
}

func (s *CheckExpandStatusResponseBodyData) SetIsActive(v bool) *CheckExpandStatusResponseBodyData {
	s.IsActive = &v
	return s
}

func (s *CheckExpandStatusResponseBodyData) SetMsg(v string) *CheckExpandStatusResponseBodyData {
	s.Msg = &v
	return s
}

func (s *CheckExpandStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
