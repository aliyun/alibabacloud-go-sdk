// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSourceContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSourceContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSourceContentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSourceContentResponseBody) *UpdateSourceContentResponse
	GetBody() *UpdateSourceContentResponseBody
}

type UpdateSourceContentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSourceContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSourceContentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSourceContentResponse) GoString() string {
	return s.String()
}

func (s *UpdateSourceContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSourceContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSourceContentResponse) GetBody() *UpdateSourceContentResponseBody {
	return s.Body
}

func (s *UpdateSourceContentResponse) SetHeaders(v map[string]*string) *UpdateSourceContentResponse {
	s.Headers = v
	return s
}

func (s *UpdateSourceContentResponse) SetStatusCode(v int32) *UpdateSourceContentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSourceContentResponse) SetBody(v *UpdateSourceContentResponseBody) *UpdateSourceContentResponse {
	s.Body = v
	return s
}

func (s *UpdateSourceContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
