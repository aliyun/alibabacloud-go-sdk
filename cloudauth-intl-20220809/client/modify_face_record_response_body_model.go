// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFaceRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyFaceRecordResponseBody
	GetCode() *string
	SetData(v int32) *ModifyFaceRecordResponseBody
	GetData() *int32
	SetMessage(v string) *ModifyFaceRecordResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyFaceRecordResponseBody
	GetRequestId() *string
}

type ModifyFaceRecordResponseBody struct {
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

func (s ModifyFaceRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyFaceRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFaceRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyFaceRecordResponseBody) GetData() *int32 {
	return s.Data
}

func (s *ModifyFaceRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyFaceRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyFaceRecordResponseBody) SetCode(v string) *ModifyFaceRecordResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyFaceRecordResponseBody) SetData(v int32) *ModifyFaceRecordResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyFaceRecordResponseBody) SetMessage(v string) *ModifyFaceRecordResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyFaceRecordResponseBody) SetRequestId(v string) *ModifyFaceRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFaceRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
