// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlgorithmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAlgorithmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAlgorithmResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAlgorithmResponseBody) *UpdateAlgorithmResponse
	GetBody() *UpdateAlgorithmResponseBody
}

type UpdateAlgorithmResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAlgorithmResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlgorithmResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlgorithmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAlgorithmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAlgorithmResponse) GetBody() *UpdateAlgorithmResponseBody {
	return s.Body
}

func (s *UpdateAlgorithmResponse) SetHeaders(v map[string]*string) *UpdateAlgorithmResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlgorithmResponse) SetStatusCode(v int32) *UpdateAlgorithmResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlgorithmResponse) SetBody(v *UpdateAlgorithmResponseBody) *UpdateAlgorithmResponse {
	s.Body = v
	return s
}

func (s *UpdateAlgorithmResponse) Validate() error {
	return dara.Validate(s)
}
