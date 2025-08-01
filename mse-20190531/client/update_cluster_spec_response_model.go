// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateClusterSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateClusterSpecResponse
	GetStatusCode() *int32
	SetBody(v *UpdateClusterSpecResponseBody) *UpdateClusterSpecResponse
	GetBody() *UpdateClusterSpecResponseBody
}

type UpdateClusterSpecResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateClusterSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateClusterSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterSpecResponse) GoString() string {
	return s.String()
}

func (s *UpdateClusterSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateClusterSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateClusterSpecResponse) GetBody() *UpdateClusterSpecResponseBody {
	return s.Body
}

func (s *UpdateClusterSpecResponse) SetHeaders(v map[string]*string) *UpdateClusterSpecResponse {
	s.Headers = v
	return s
}

func (s *UpdateClusterSpecResponse) SetStatusCode(v int32) *UpdateClusterSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClusterSpecResponse) SetBody(v *UpdateClusterSpecResponseBody) *UpdateClusterSpecResponse {
	s.Body = v
	return s
}

func (s *UpdateClusterSpecResponse) Validate() error {
	return dara.Validate(s)
}
