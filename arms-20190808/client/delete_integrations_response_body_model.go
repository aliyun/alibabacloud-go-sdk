// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIntegrationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsSuccess(v bool) *DeleteIntegrationsResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteIntegrationsResponseBody
	GetRequestId() *string
}

type DeleteIntegrationsResponseBody struct {
	// Indicates whether the alert integration is deleted. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 34ED024E-9E31-434A-9E4E-D9D15C3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIntegrationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIntegrationsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIntegrationsResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteIntegrationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIntegrationsResponseBody) SetIsSuccess(v bool) *DeleteIntegrationsResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteIntegrationsResponseBody) SetRequestId(v string) *DeleteIntegrationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIntegrationsResponseBody) Validate() error {
	return dara.Validate(s)
}
