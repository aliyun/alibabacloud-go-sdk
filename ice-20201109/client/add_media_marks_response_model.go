// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMediaMarksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddMediaMarksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddMediaMarksResponse
	GetStatusCode() *int32
	SetBody(v *AddMediaMarksResponseBody) *AddMediaMarksResponse
	GetBody() *AddMediaMarksResponseBody
}

type AddMediaMarksResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMediaMarksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMediaMarksResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMediaMarksResponse) GoString() string {
	return s.String()
}

func (s *AddMediaMarksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddMediaMarksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddMediaMarksResponse) GetBody() *AddMediaMarksResponseBody {
	return s.Body
}

func (s *AddMediaMarksResponse) SetHeaders(v map[string]*string) *AddMediaMarksResponse {
	s.Headers = v
	return s
}

func (s *AddMediaMarksResponse) SetStatusCode(v int32) *AddMediaMarksResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMediaMarksResponse) SetBody(v *AddMediaMarksResponseBody) *AddMediaMarksResponse {
	s.Body = v
	return s
}

func (s *AddMediaMarksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
