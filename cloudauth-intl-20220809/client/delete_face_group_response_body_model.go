// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteFaceGroupResponseBody
	GetCode() *string
	SetData(v int32) *DeleteFaceGroupResponseBody
	GetData() *int32
	SetMessage(v string) *DeleteFaceGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteFaceGroupResponseBody
	GetRequestId() *string
}

type DeleteFaceGroupResponseBody struct {
	// Return code
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return result.
	//
	// example:
	//
	// 1
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// Return message
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 595E387B-3F0E-5C52-BD02-8EFE63D41FD5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFaceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteFaceGroupResponseBody) GetData() *int32 {
	return s.Data
}

func (s *DeleteFaceGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteFaceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFaceGroupResponseBody) SetCode(v string) *DeleteFaceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFaceGroupResponseBody) SetData(v int32) *DeleteFaceGroupResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteFaceGroupResponseBody) SetMessage(v string) *DeleteFaceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFaceGroupResponseBody) SetRequestId(v string) *DeleteFaceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
