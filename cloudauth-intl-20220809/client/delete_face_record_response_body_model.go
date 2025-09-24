// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteFaceRecordResponseBody
	GetCode() *string
	SetData(v int32) *DeleteFaceRecordResponseBody
	GetData() *int32
	SetMessage(v string) *DeleteFaceRecordResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteFaceRecordResponseBody
	GetRequestId() *string
}

type DeleteFaceRecordResponseBody struct {
	// Return code.
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
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFaceRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteFaceRecordResponseBody) GetData() *int32 {
	return s.Data
}

func (s *DeleteFaceRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteFaceRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFaceRecordResponseBody) SetCode(v string) *DeleteFaceRecordResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFaceRecordResponseBody) SetData(v int32) *DeleteFaceRecordResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteFaceRecordResponseBody) SetMessage(v string) *DeleteFaceRecordResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFaceRecordResponseBody) SetRequestId(v string) *DeleteFaceRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
