// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateStoryResponseBody
	GetRequestId() *string
}

type UpdateStoryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6E93D6C9-5AC0-49F9-914D-E02678D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateStoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateStoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateStoryResponseBody) SetRequestId(v string) *UpdateStoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStoryResponseBody) Validate() error {
	return dara.Validate(s)
}
