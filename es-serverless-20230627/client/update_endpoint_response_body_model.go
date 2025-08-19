// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateEndpointResponseBody
	GetRequestId() *string
	SetResult(v *UpdateEndpointResponseBodyResult) *UpdateEndpointResponseBody
	GetResult() *UpdateEndpointResponseBodyResult
}

type UpdateEndpointResponseBody struct {
	// example:
	//
	// FBAD8493-87FA-583E-8A4C-D487F2DE90FC
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *UpdateEndpointResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEndpointResponseBody) GetResult() *UpdateEndpointResponseBodyResult {
	return s.Result
}

func (s *UpdateEndpointResponseBody) SetRequestId(v string) *UpdateEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEndpointResponseBody) SetResult(v *UpdateEndpointResponseBodyResult) *UpdateEndpointResponseBody {
	s.Result = v
	return s
}

func (s *UpdateEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateEndpointResponseBodyResult struct {
	// example:
	//
	// ep-bp1i98bcbb1540d0****
	EndpointId *string `json:"endpointId,omitempty" xml:"endpointId,omitempty"`
}

func (s UpdateEndpointResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateEndpointResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateEndpointResponseBodyResult) GetEndpointId() *string {
	return s.EndpointId
}

func (s *UpdateEndpointResponseBodyResult) SetEndpointId(v string) *UpdateEndpointResponseBodyResult {
	s.EndpointId = &v
	return s
}

func (s *UpdateEndpointResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
