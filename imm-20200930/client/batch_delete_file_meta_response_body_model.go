// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteFileMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchDeleteFileMetaResponseBody
	GetRequestId() *string
}

type BatchDeleteFileMetaResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3A82F6C9-5AC0-38F9-914F-F02623B3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDeleteFileMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteFileMetaResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteFileMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteFileMetaResponseBody) SetRequestId(v string) *BatchDeleteFileMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteFileMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
