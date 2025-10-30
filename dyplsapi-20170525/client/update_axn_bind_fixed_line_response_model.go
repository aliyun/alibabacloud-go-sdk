// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAxnBindFixedLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAxnBindFixedLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAxnBindFixedLineResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAxnBindFixedLineResponseBody) *UpdateAxnBindFixedLineResponse
	GetBody() *UpdateAxnBindFixedLineResponseBody
}

type UpdateAxnBindFixedLineResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAxnBindFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAxnBindFixedLineResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAxnBindFixedLineResponse) GoString() string {
	return s.String()
}

func (s *UpdateAxnBindFixedLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAxnBindFixedLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAxnBindFixedLineResponse) GetBody() *UpdateAxnBindFixedLineResponseBody {
	return s.Body
}

func (s *UpdateAxnBindFixedLineResponse) SetHeaders(v map[string]*string) *UpdateAxnBindFixedLineResponse {
	s.Headers = v
	return s
}

func (s *UpdateAxnBindFixedLineResponse) SetStatusCode(v int32) *UpdateAxnBindFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAxnBindFixedLineResponse) SetBody(v *UpdateAxnBindFixedLineResponseBody) *UpdateAxnBindFixedLineResponse {
	s.Body = v
	return s
}

func (s *UpdateAxnBindFixedLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
