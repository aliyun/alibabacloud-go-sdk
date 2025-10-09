// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComponentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteComponentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteComponentResponseBody
	GetSuccess() *bool
}

type DeleteComponentResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 0000-ABCD-EF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteComponentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteComponentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteComponentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteComponentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteComponentResponseBody) SetRequestId(v string) *DeleteComponentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteComponentResponseBody) SetSuccess(v bool) *DeleteComponentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteComponentResponseBody) Validate() error {
	return dara.Validate(s)
}
