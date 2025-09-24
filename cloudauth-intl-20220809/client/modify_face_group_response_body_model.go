// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFaceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyFaceGroupResponseBody
	GetCode() *string
	SetData(v int32) *ModifyFaceGroupResponseBody
	GetData() *int32
	SetMessage(v string) *ModifyFaceGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyFaceGroupResponseBody
	GetRequestId() *string
}

type ModifyFaceGroupResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFaceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyFaceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFaceGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyFaceGroupResponseBody) GetData() *int32 {
	return s.Data
}

func (s *ModifyFaceGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyFaceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyFaceGroupResponseBody) SetCode(v string) *ModifyFaceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyFaceGroupResponseBody) SetData(v int32) *ModifyFaceGroupResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyFaceGroupResponseBody) SetMessage(v string) *ModifyFaceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyFaceGroupResponseBody) SetRequestId(v string) *ModifyFaceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFaceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
