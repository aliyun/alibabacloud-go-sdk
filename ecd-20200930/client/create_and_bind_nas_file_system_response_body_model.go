// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndBindNasFileSystemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAndBindNasFileSystemResponseBody
	GetRequestId() *string
}

type CreateAndBindNasFileSystemResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F7E4322D-D679-5ACB-A909-490D2F0E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAndBindNasFileSystemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAndBindNasFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAndBindNasFileSystemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAndBindNasFileSystemResponseBody) SetRequestId(v string) *CreateAndBindNasFileSystemResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAndBindNasFileSystemResponseBody) Validate() error {
	return dara.Validate(s)
}
