// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkbenchEventResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWorkbenchEventResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateWorkbenchEventResultResponseBody
	GetSuccess() *bool
}

type UpdateWorkbenchEventResultResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0bc1ec92159376****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWorkbenchEventResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkbenchEventResultResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkbenchEventResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkbenchEventResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateWorkbenchEventResultResponseBody) SetRequestId(v string) *UpdateWorkbenchEventResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkbenchEventResultResponseBody) SetSuccess(v bool) *UpdateWorkbenchEventResultResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWorkbenchEventResultResponseBody) Validate() error {
	return dara.Validate(s)
}
