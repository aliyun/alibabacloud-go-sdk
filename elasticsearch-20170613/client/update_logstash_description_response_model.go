// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLogstashDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLogstashDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLogstashDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLogstashDescriptionResponseBody) *UpdateLogstashDescriptionResponse
	GetBody() *UpdateLogstashDescriptionResponseBody
}

type UpdateLogstashDescriptionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLogstashDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLogstashDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogstashDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateLogstashDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLogstashDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLogstashDescriptionResponse) GetBody() *UpdateLogstashDescriptionResponseBody {
	return s.Body
}

func (s *UpdateLogstashDescriptionResponse) SetHeaders(v map[string]*string) *UpdateLogstashDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateLogstashDescriptionResponse) SetStatusCode(v int32) *UpdateLogstashDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLogstashDescriptionResponse) SetBody(v *UpdateLogstashDescriptionResponseBody) *UpdateLogstashDescriptionResponse {
	s.Body = v
	return s
}

func (s *UpdateLogstashDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
