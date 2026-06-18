// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitUploadTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SubmitUploadTaskResponseBody
	GetRequestId() *string
	SetUploadId(v int64) *SubmitUploadTaskResponseBody
	GetUploadId() *int64
}

type SubmitUploadTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the file upload task.
	//
	// example:
	//
	// 1593805857882113
	UploadId *int64 `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
}

func (s SubmitUploadTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitUploadTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitUploadTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitUploadTaskResponseBody) GetUploadId() *int64 {
	return s.UploadId
}

func (s *SubmitUploadTaskResponseBody) SetRequestId(v string) *SubmitUploadTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitUploadTaskResponseBody) SetUploadId(v int64) *SubmitUploadTaskResponseBody {
	s.UploadId = &v
	return s
}

func (s *SubmitUploadTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
