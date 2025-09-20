// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFileMetaResponseBody
	GetRequestId() *string
}

type DeleteFileMetaResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7F82D6C9-5AC0-49F9-914D-F02678F3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFileMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFileMetaResponseBody) SetRequestId(v string) *DeleteFileMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFileMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
