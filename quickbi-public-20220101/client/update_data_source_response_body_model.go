// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDataSourceResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateDataSourceResponseBody
	GetResult() *bool
	SetSuccess(v bool) *UpdateDataSourceResponseBody
	GetSuccess() *bool
}

type UpdateDataSourceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D787E1A***********5DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the operation. Valid values:
	//
	// - `true`: The operation was successful.
	//
	// - `false`: The operation failed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - `true`: The request was successful.
	//
	// - `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataSourceResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateDataSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataSourceResponseBody) SetRequestId(v string) *UpdateDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataSourceResponseBody) SetResult(v bool) *UpdateDataSourceResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateDataSourceResponseBody) SetSuccess(v bool) *UpdateDataSourceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
