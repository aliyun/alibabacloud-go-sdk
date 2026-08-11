// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelLimitsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateModelLimitsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateModelLimitsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateModelLimitsResponseBody) *UpdateModelLimitsResponse
	GetBody() *UpdateModelLimitsResponseBody
}

type UpdateModelLimitsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModelLimitsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelLimitsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelLimitsResponse) GoString() string {
	return s.String()
}

func (s *UpdateModelLimitsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateModelLimitsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateModelLimitsResponse) GetBody() *UpdateModelLimitsResponseBody {
	return s.Body
}

func (s *UpdateModelLimitsResponse) SetHeaders(v map[string]*string) *UpdateModelLimitsResponse {
	s.Headers = v
	return s
}

func (s *UpdateModelLimitsResponse) SetStatusCode(v int32) *UpdateModelLimitsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModelLimitsResponse) SetBody(v *UpdateModelLimitsResponseBody) *UpdateModelLimitsResponse {
	s.Body = v
	return s
}

func (s *UpdateModelLimitsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
