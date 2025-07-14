// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserTagMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUserTagMetaResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateUserTagMetaResponseBody
	GetResult() *bool
	SetSuccess(v bool) *UpdateUserTagMetaResponseBody
	GetSuccess() *bool
}

type UpdateUserTagMetaResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the interface was executed successfully. Possible values:
	//
	// - true: Execution succeeded
	//
	// - false: Execution failed
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: The request succeeded - false: The request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateUserTagMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserTagMetaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserTagMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserTagMetaResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateUserTagMetaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateUserTagMetaResponseBody) SetRequestId(v string) *UpdateUserTagMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserTagMetaResponseBody) SetResult(v bool) *UpdateUserTagMetaResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateUserTagMetaResponseBody) SetSuccess(v bool) *UpdateUserTagMetaResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateUserTagMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
