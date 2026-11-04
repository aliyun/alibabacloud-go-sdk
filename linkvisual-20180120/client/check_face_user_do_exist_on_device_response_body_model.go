// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckFaceUserDoExistOnDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckFaceUserDoExistOnDeviceResponseBody
	GetCode() *string
	SetData(v *CheckFaceUserDoExistOnDeviceResponseBodyData) *CheckFaceUserDoExistOnDeviceResponseBody
	GetData() *CheckFaceUserDoExistOnDeviceResponseBodyData
	SetErrorMessage(v string) *CheckFaceUserDoExistOnDeviceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CheckFaceUserDoExistOnDeviceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckFaceUserDoExistOnDeviceResponseBody
	GetSuccess() *bool
}

type CheckFaceUserDoExistOnDeviceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned check information.
	Data *CheckFaceUserDoExistOnDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0A13279A-5640-45E7-87AA-83750541AD0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckFaceUserDoExistOnDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckFaceUserDoExistOnDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *CheckFaceUserDoExistOnDeviceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckFaceUserDoExistOnDeviceResponseBody) GetData() *CheckFaceUserDoExistOnDeviceResponseBodyData {
	return s.Data
}

func (s *CheckFaceUserDoExistOnDeviceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CheckFaceUserDoExistOnDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckFaceUserDoExistOnDeviceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckFaceUserDoExistOnDeviceResponseBody) SetCode(v string) *CheckFaceUserDoExistOnDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceResponseBody) SetData(v *CheckFaceUserDoExistOnDeviceResponseBodyData) *CheckFaceUserDoExistOnDeviceResponseBody {
	s.Data = v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceResponseBody) SetErrorMessage(v string) *CheckFaceUserDoExistOnDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceResponseBody) SetRequestId(v string) *CheckFaceUserDoExistOnDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceResponseBody) SetSuccess(v bool) *CheckFaceUserDoExistOnDeviceResponseBody {
	s.Success = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckFaceUserDoExistOnDeviceResponseBodyData struct {
	// Indicates whether the face of the specified user is successfully synchronized.
	//
	// example:
	//
	// true
	DoExist *bool `json:"DoExist,omitempty" xml:"DoExist,omitempty"`
}

func (s CheckFaceUserDoExistOnDeviceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CheckFaceUserDoExistOnDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckFaceUserDoExistOnDeviceResponseBodyData) GetDoExist() *bool {
	return s.DoExist
}

func (s *CheckFaceUserDoExistOnDeviceResponseBodyData) SetDoExist(v bool) *CheckFaceUserDoExistOnDeviceResponseBodyData {
	s.DoExist = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
