// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateSourceResponseBody) *CreateSourceResponse
	GetBody() *CreateSourceResponseBody
}

type CreateSourceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSourceResponse) GetBody() *CreateSourceResponseBody {
	return s.Body
}

func (s *CreateSourceResponse) SetHeaders(v map[string]*string) *CreateSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateSourceResponse) SetStatusCode(v int32) *CreateSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSourceResponse) SetBody(v *CreateSourceResponseBody) *CreateSourceResponse {
	s.Body = v
	return s
}

func (s *CreateSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
