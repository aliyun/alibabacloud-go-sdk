// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEndpointAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEndpointAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEndpointAttributeResponse
	GetStatusCode() *int32
	SetBody(v *GetEndpointAttributeResponseBody) *GetEndpointAttributeResponse
	GetBody() *GetEndpointAttributeResponseBody
}

type GetEndpointAttributeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEndpointAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEndpointAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEndpointAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetEndpointAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEndpointAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEndpointAttributeResponse) GetBody() *GetEndpointAttributeResponseBody {
	return s.Body
}

func (s *GetEndpointAttributeResponse) SetHeaders(v map[string]*string) *GetEndpointAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetEndpointAttributeResponse) SetStatusCode(v int32) *GetEndpointAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEndpointAttributeResponse) SetBody(v *GetEndpointAttributeResponseBody) *GetEndpointAttributeResponse {
	s.Body = v
	return s
}

func (s *GetEndpointAttributeResponse) Validate() error {
	return dara.Validate(s)
}
