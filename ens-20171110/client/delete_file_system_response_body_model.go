// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileSystemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFileSystemResponseBody
	GetRequestId() *string
}

type DeleteFileSystemResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFileSystemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileSystemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFileSystemResponseBody) SetRequestId(v string) *DeleteFileSystemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFileSystemResponseBody) Validate() error {
	return dara.Validate(s)
}
