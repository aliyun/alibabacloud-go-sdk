// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRouteTableAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRouteTableAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRouteTableAttributesResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRouteTableAttributesResponseBody) *ModifyRouteTableAttributesResponse
	GetBody() *ModifyRouteTableAttributesResponseBody
}

type ModifyRouteTableAttributesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRouteTableAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRouteTableAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRouteTableAttributesResponse) GoString() string {
	return s.String()
}

func (s *ModifyRouteTableAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRouteTableAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRouteTableAttributesResponse) GetBody() *ModifyRouteTableAttributesResponseBody {
	return s.Body
}

func (s *ModifyRouteTableAttributesResponse) SetHeaders(v map[string]*string) *ModifyRouteTableAttributesResponse {
	s.Headers = v
	return s
}

func (s *ModifyRouteTableAttributesResponse) SetStatusCode(v int32) *ModifyRouteTableAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRouteTableAttributesResponse) SetBody(v *ModifyRouteTableAttributesResponseBody) *ModifyRouteTableAttributesResponse {
	s.Body = v
	return s
}

func (s *ModifyRouteTableAttributesResponse) Validate() error {
	return dara.Validate(s)
}
