// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLayerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLayerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLayerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLayerResponseBody) *UpdateLayerResponse
	GetBody() *UpdateLayerResponseBody
}

type UpdateLayerResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLayerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLayerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLayerResponse) GoString() string {
	return s.String()
}

func (s *UpdateLayerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLayerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLayerResponse) GetBody() *UpdateLayerResponseBody {
	return s.Body
}

func (s *UpdateLayerResponse) SetHeaders(v map[string]*string) *UpdateLayerResponse {
	s.Headers = v
	return s
}

func (s *UpdateLayerResponse) SetStatusCode(v int32) *UpdateLayerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLayerResponse) SetBody(v *UpdateLayerResponseBody) *UpdateLayerResponse {
	s.Body = v
	return s
}

func (s *UpdateLayerResponse) Validate() error {
	return dara.Validate(s)
}
