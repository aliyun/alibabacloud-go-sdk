// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNetworkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNetworkResponseBody) *UpdateNetworkResponse
	GetBody() *UpdateNetworkResponseBody
}

type UpdateNetworkResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkResponse) GoString() string {
	return s.String()
}

func (s *UpdateNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNetworkResponse) GetBody() *UpdateNetworkResponseBody {
	return s.Body
}

func (s *UpdateNetworkResponse) SetHeaders(v map[string]*string) *UpdateNetworkResponse {
	s.Headers = v
	return s
}

func (s *UpdateNetworkResponse) SetStatusCode(v int32) *UpdateNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNetworkResponse) SetBody(v *UpdateNetworkResponseBody) *UpdateNetworkResponse {
	s.Body = v
	return s
}

func (s *UpdateNetworkResponse) Validate() error {
	return dara.Validate(s)
}
