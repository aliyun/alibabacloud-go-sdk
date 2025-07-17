// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteNodeResponseBody
	GetSuccess() *bool
}

type DeleteNodeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A1E54497-5122-505E-91C6-BAC14980XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// true\\
	//
	// false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteNodeResponseBody) SetRequestId(v string) *DeleteNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNodeResponseBody) SetSuccess(v bool) *DeleteNodeResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
