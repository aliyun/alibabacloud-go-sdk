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
	SetSuccess(v bool) *UpdateDataSourceResponseBody
	GetSuccess() *bool
}

type UpdateDataSourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 102E8E24-0387-531D-8A75-1C0AE7DD03E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the data source has been modified:
	//
	// - true: Yes
	//
	// - false: no
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

func (s *UpdateDataSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataSourceResponseBody) SetRequestId(v string) *UpdateDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataSourceResponseBody) SetSuccess(v bool) *UpdateDataSourceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
