// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodeRunModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateNodeRunModeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateNodeRunModeResponseBody
	GetSuccess() *bool
}

type UpdateNodeRunModeResponseBody struct {
	// The request ID.
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

func (s UpdateNodeRunModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeRunModeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNodeRunModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNodeRunModeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateNodeRunModeResponseBody) SetRequestId(v string) *UpdateNodeRunModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNodeRunModeResponseBody) SetSuccess(v bool) *UpdateNodeRunModeResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateNodeRunModeResponseBody) Validate() error {
	return dara.Validate(s)
}
