// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCrowdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCrowdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCrowdResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCrowdResponseBody) *UpdateCrowdResponse
	GetBody() *UpdateCrowdResponseBody
}

type UpdateCrowdResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCrowdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCrowdResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCrowdResponse) GoString() string {
	return s.String()
}

func (s *UpdateCrowdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCrowdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCrowdResponse) GetBody() *UpdateCrowdResponseBody {
	return s.Body
}

func (s *UpdateCrowdResponse) SetHeaders(v map[string]*string) *UpdateCrowdResponse {
	s.Headers = v
	return s
}

func (s *UpdateCrowdResponse) SetStatusCode(v int32) *UpdateCrowdResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCrowdResponse) SetBody(v *UpdateCrowdResponseBody) *UpdateCrowdResponse {
	s.Body = v
	return s
}

func (s *UpdateCrowdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
