// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MoveNodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MoveNodeResponseBody
	GetSuccess() *bool
}

type MoveNodeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C99E2BE6-9DEA-5C2E-8F51-1DDCFEADXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MoveNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveNodeResponseBody) GoString() string {
	return s.String()
}

func (s *MoveNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveNodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MoveNodeResponseBody) SetRequestId(v string) *MoveNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveNodeResponseBody) SetSuccess(v bool) *MoveNodeResponseBody {
	s.Success = &v
	return s
}

func (s *MoveNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
