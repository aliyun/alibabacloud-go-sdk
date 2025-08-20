// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChartNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateChartNamespaceResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *UpdateChartNamespaceResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *UpdateChartNamespaceResponseBody
	GetRequestId() *string
}

type UpdateChartNamespaceResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the API request is successful. Valid values:
	//
	// 	- `true`: successful
	//
	// 	- `false`: failed
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6877B80A-2895-44C4-BC9E-703B157DEE66
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateChartNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateChartNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateChartNamespaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateChartNamespaceResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *UpdateChartNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateChartNamespaceResponseBody) SetCode(v string) *UpdateChartNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateChartNamespaceResponseBody) SetIsSuccess(v bool) *UpdateChartNamespaceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateChartNamespaceResponseBody) SetRequestId(v string) *UpdateChartNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateChartNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}
