// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateImageLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateImageLibResponse
	GetStatusCode() *int32
	SetBody(v *CreateImageLibResponseBody) *CreateImageLibResponse
	GetBody() *CreateImageLibResponseBody
}

type CreateImageLibResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImageLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateImageLibResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateImageLibResponse) GoString() string {
	return s.String()
}

func (s *CreateImageLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateImageLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateImageLibResponse) GetBody() *CreateImageLibResponseBody {
	return s.Body
}

func (s *CreateImageLibResponse) SetHeaders(v map[string]*string) *CreateImageLibResponse {
	s.Headers = v
	return s
}

func (s *CreateImageLibResponse) SetStatusCode(v int32) *CreateImageLibResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageLibResponse) SetBody(v *CreateImageLibResponseBody) *CreateImageLibResponse {
	s.Body = v
	return s
}

func (s *CreateImageLibResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
