// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RenameResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RenameResourceResponseBody
	GetSuccess() *bool
}

type RenameResourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4CDF7B72-020B-542A-8465-21CFFA8XXXXX
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

func (s RenameResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenameResourceResponseBody) GoString() string {
	return s.String()
}

func (s *RenameResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenameResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RenameResourceResponseBody) SetRequestId(v string) *RenameResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameResourceResponseBody) SetSuccess(v bool) *RenameResourceResponseBody {
	s.Success = &v
	return s
}

func (s *RenameResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
