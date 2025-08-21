// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveTagResponseBody
	GetRequestId() *string
}

type RemoveTagResponseBody struct {
	// example:
	//
	// 23000F3C-0EFE-4C89-82EE-E04F42D37B3C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveTagResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveTagResponseBody) SetRequestId(v string) *RemoveTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveTagResponseBody) Validate() error {
	return dara.Validate(s)
}
