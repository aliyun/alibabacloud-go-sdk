// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MoveResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MoveResourceResponseBody
	GetSuccess() *bool
}

type MoveResourceResponseBody struct {
	// The ID of the request. You can use the ID to troubleshoot errors.
	//
	// example:
	//
	// F332BED4-DD73-5972-A9C2-642BA6CFXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// - true: The call is successful.
	//
	// - false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MoveResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveResourceResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MoveResourceResponseBody) SetRequestId(v string) *MoveResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveResourceResponseBody) SetSuccess(v bool) *MoveResourceResponseBody {
	s.Success = &v
	return s
}

func (s *MoveResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
