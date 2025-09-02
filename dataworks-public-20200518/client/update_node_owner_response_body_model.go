// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodeOwnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateNodeOwnerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateNodeOwnerResponseBody
	GetSuccess() *bool
}

type UpdateNodeOwnerResponseBody struct {
	// The request ID. You can use the request ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// E6F0DBDD-5AD****
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

func (s UpdateNodeOwnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNodeOwnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNodeOwnerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateNodeOwnerResponseBody) SetRequestId(v string) *UpdateNodeOwnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNodeOwnerResponseBody) SetSuccess(v bool) *UpdateNodeOwnerResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateNodeOwnerResponseBody) Validate() error {
	return dara.Validate(s)
}
