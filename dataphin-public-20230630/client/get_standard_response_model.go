// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStandardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStandardResponse
	GetStatusCode() *int32
	SetBody(v *GetStandardResponseBody) *GetStandardResponse
	GetBody() *GetStandardResponseBody
}

type GetStandardResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStandardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStandardResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponse) GoString() string {
	return s.String()
}

func (s *GetStandardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStandardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStandardResponse) GetBody() *GetStandardResponseBody {
	return s.Body
}

func (s *GetStandardResponse) SetHeaders(v map[string]*string) *GetStandardResponse {
	s.Headers = v
	return s
}

func (s *GetStandardResponse) SetStatusCode(v int32) *GetStandardResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStandardResponse) SetBody(v *GetStandardResponseBody) *GetStandardResponse {
	s.Body = v
	return s
}

func (s *GetStandardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
