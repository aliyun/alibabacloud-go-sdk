// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSubInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveSubInstanceResponseBody
	GetRequestId() *string
}

type RemoveSubInstanceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5D622714-AEDD-4609-9167-F5DDD3D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveSubInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveSubInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSubInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveSubInstanceResponseBody) SetRequestId(v string) *RemoveSubInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveSubInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
