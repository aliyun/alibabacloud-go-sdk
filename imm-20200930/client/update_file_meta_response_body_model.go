// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateFileMetaResponseBody
	GetRequestId() *string
}

type UpdateFileMetaResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6D53E6C9-5AC0-48F9-825F-D02678E3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFileMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileMetaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFileMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFileMetaResponseBody) SetRequestId(v string) *UpdateFileMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFileMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
