// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeAccountPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeAccountPasswordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChangeAccountPasswordResponseBody
	GetSuccess() *bool
}

type ChangeAccountPasswordResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// DSSDF-SEWE-*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeAccountPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeAccountPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeAccountPasswordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeAccountPasswordResponseBody) SetRequestId(v string) *ChangeAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeAccountPasswordResponseBody) SetSuccess(v bool) *ChangeAccountPasswordResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeAccountPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
