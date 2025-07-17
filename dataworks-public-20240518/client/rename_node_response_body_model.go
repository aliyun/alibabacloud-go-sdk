// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RenameNodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RenameNodeResponseBody
	GetSuccess() *bool
}

type RenameNodeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4CDF7B72-020B-542A-8465-21CFFA81XXXX
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

func (s RenameNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenameNodeResponseBody) GoString() string {
	return s.String()
}

func (s *RenameNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenameNodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RenameNodeResponseBody) SetRequestId(v string) *RenameNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameNodeResponseBody) SetSuccess(v bool) *RenameNodeResponseBody {
	s.Success = &v
	return s
}

func (s *RenameNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
