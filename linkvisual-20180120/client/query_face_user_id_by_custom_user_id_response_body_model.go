// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceUserIdByCustomUserIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryFaceUserIdByCustomUserIdResponseBody
	GetCode() *string
	SetData(v *QueryFaceUserIdByCustomUserIdResponseBodyData) *QueryFaceUserIdByCustomUserIdResponseBody
	GetData() *QueryFaceUserIdByCustomUserIdResponseBodyData
	SetErrorMessage(v string) *QueryFaceUserIdByCustomUserIdResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryFaceUserIdByCustomUserIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryFaceUserIdByCustomUserIdResponseBody
	GetSuccess() *bool
}

type QueryFaceUserIdByCustomUserIdResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return value of the API operation.
	Data *QueryFaceUserIdByCustomUserIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Indicates whether the API call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceUserIdByCustomUserIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceUserIdByCustomUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceUserIdByCustomUserIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryFaceUserIdByCustomUserIdResponseBody) GetData() *QueryFaceUserIdByCustomUserIdResponseBodyData {
	return s.Data
}

func (s *QueryFaceUserIdByCustomUserIdResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryFaceUserIdByCustomUserIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFaceUserIdByCustomUserIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryFaceUserIdByCustomUserIdResponseBody) SetCode(v string) *QueryFaceUserIdByCustomUserIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponseBody) SetData(v *QueryFaceUserIdByCustomUserIdResponseBodyData) *QueryFaceUserIdByCustomUserIdResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponseBody) SetErrorMessage(v string) *QueryFaceUserIdByCustomUserIdResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponseBody) SetRequestId(v string) *QueryFaceUserIdByCustomUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponseBody) SetSuccess(v bool) *QueryFaceUserIdByCustomUserIdResponseBody {
	s.Success = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryFaceUserIdByCustomUserIdResponseBodyData struct {
	// The user ID in the caller system.
	//
	// example:
	//
	// TestCustomUserId
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	// The username.
	//
	// example:
	//
	// TestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user parameters.
	//
	// example:
	//
	// {"age":28, "sex":"male"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The user ID.
	//
	// example:
	//
	// TestUserId
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryFaceUserIdByCustomUserIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceUserIdByCustomUserIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceUserIdByCustomUserIdResponseBodyData) GetCustomUserId() *string {
	return s.CustomUserId
}

func (s *QueryFaceUserIdByCustomUserIdResponseBodyData) GetName() *string {
	return s.Name
}

func (s *QueryFaceUserIdByCustomUserIdResponseBodyData) GetParams() *string {
	return s.Params
}

func (s *QueryFaceUserIdByCustomUserIdResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *QueryFaceUserIdByCustomUserIdResponseBodyData) SetCustomUserId(v string) *QueryFaceUserIdByCustomUserIdResponseBodyData {
	s.CustomUserId = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponseBodyData) SetName(v string) *QueryFaceUserIdByCustomUserIdResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponseBodyData) SetParams(v string) *QueryFaceUserIdByCustomUserIdResponseBodyData {
	s.Params = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponseBodyData) SetUserId(v string) *QueryFaceUserIdByCustomUserIdResponseBodyData {
	s.UserId = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
