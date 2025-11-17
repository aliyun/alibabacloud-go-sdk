// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEmbeddedStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEmbeddedStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEmbeddedStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEmbeddedStatusResponseBody) *UpdateEmbeddedStatusResponse
	GetBody() *UpdateEmbeddedStatusResponseBody
}

type UpdateEmbeddedStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEmbeddedStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEmbeddedStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEmbeddedStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateEmbeddedStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEmbeddedStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEmbeddedStatusResponse) GetBody() *UpdateEmbeddedStatusResponseBody {
	return s.Body
}

func (s *UpdateEmbeddedStatusResponse) SetHeaders(v map[string]*string) *UpdateEmbeddedStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateEmbeddedStatusResponse) SetStatusCode(v int32) *UpdateEmbeddedStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEmbeddedStatusResponse) SetBody(v *UpdateEmbeddedStatusResponseBody) *UpdateEmbeddedStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateEmbeddedStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
