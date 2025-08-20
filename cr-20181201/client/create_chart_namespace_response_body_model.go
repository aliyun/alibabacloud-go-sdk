// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChartNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateChartNamespaceResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CreateChartNamespaceResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *CreateChartNamespaceResponseBody
	GetRequestId() *string
}

type CreateChartNamespaceResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 724402D0-75CD-4794-BC20-7D3720823AE0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateChartNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateChartNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChartNamespaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateChartNamespaceResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CreateChartNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateChartNamespaceResponseBody) SetCode(v string) *CreateChartNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateChartNamespaceResponseBody) SetIsSuccess(v bool) *CreateChartNamespaceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateChartNamespaceResponseBody) SetRequestId(v string) *CreateChartNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateChartNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}
