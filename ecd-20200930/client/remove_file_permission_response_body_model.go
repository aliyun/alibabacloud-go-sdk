// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveFilePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveFilePermissionResponseBody
	GetRequestId() *string
}

type RemoveFilePermissionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4636DBE0-BBB4-4076-8B8E-94D21A9A3CFB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveFilePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveFilePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveFilePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveFilePermissionResponseBody) SetRequestId(v string) *RemoveFilePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveFilePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
