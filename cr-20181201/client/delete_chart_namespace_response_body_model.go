// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChartNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteChartNamespaceResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *DeleteChartNamespaceResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteChartNamespaceResponseBody
	GetRequestId() *string
}

type DeleteChartNamespaceResponseBody struct {
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
	// FEC62DF1-1394-467F-A69F-4BC1BA29F383
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteChartNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteChartNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChartNamespaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteChartNamespaceResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteChartNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteChartNamespaceResponseBody) SetCode(v string) *DeleteChartNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteChartNamespaceResponseBody) SetIsSuccess(v bool) *DeleteChartNamespaceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteChartNamespaceResponseBody) SetRequestId(v string) *DeleteChartNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteChartNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}
