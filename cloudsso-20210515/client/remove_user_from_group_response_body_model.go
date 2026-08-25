// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserFromGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveUserFromGroupResponseBody
	GetRequestId() *string
}

type RemoveUserFromGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F723DE01-6276-5DC4-9B1F-9CBE3E1748B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveUserFromGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserFromGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserFromGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveUserFromGroupResponseBody) SetRequestId(v string) *RemoveUserFromGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUserFromGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
