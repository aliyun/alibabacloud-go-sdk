// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAxbBindFixedLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAxbBindFixedLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAxbBindFixedLineResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAxbBindFixedLineResponseBody) *UpdateAxbBindFixedLineResponse
	GetBody() *UpdateAxbBindFixedLineResponseBody
}

type UpdateAxbBindFixedLineResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAxbBindFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAxbBindFixedLineResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAxbBindFixedLineResponse) GoString() string {
	return s.String()
}

func (s *UpdateAxbBindFixedLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAxbBindFixedLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAxbBindFixedLineResponse) GetBody() *UpdateAxbBindFixedLineResponseBody {
	return s.Body
}

func (s *UpdateAxbBindFixedLineResponse) SetHeaders(v map[string]*string) *UpdateAxbBindFixedLineResponse {
	s.Headers = v
	return s
}

func (s *UpdateAxbBindFixedLineResponse) SetStatusCode(v int32) *UpdateAxbBindFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAxbBindFixedLineResponse) SetBody(v *UpdateAxbBindFixedLineResponseBody) *UpdateAxbBindFixedLineResponse {
	s.Body = v
	return s
}

func (s *UpdateAxbBindFixedLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
