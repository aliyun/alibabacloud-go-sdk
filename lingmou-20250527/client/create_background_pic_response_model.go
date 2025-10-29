// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackgroundPicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBackgroundPicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBackgroundPicResponse
	GetStatusCode() *int32
	SetBody(v *CreateBackgroundPicResponseBody) *CreateBackgroundPicResponse
	GetBody() *CreateBackgroundPicResponseBody
}

type CreateBackgroundPicResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBackgroundPicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBackgroundPicResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBackgroundPicResponse) GoString() string {
	return s.String()
}

func (s *CreateBackgroundPicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBackgroundPicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBackgroundPicResponse) GetBody() *CreateBackgroundPicResponseBody {
	return s.Body
}

func (s *CreateBackgroundPicResponse) SetHeaders(v map[string]*string) *CreateBackgroundPicResponse {
	s.Headers = v
	return s
}

func (s *CreateBackgroundPicResponse) SetStatusCode(v int32) *CreateBackgroundPicResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBackgroundPicResponse) SetBody(v *CreateBackgroundPicResponseBody) *CreateBackgroundPicResponse {
	s.Body = v
	return s
}

func (s *CreateBackgroundPicResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
