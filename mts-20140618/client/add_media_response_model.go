// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMediaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddMediaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddMediaResponse
	GetStatusCode() *int32
	SetBody(v *AddMediaResponseBody) *AddMediaResponse
	GetBody() *AddMediaResponseBody
}

type AddMediaResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMediaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMediaResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMediaResponse) GoString() string {
	return s.String()
}

func (s *AddMediaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddMediaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddMediaResponse) GetBody() *AddMediaResponseBody {
	return s.Body
}

func (s *AddMediaResponse) SetHeaders(v map[string]*string) *AddMediaResponse {
	s.Headers = v
	return s
}

func (s *AddMediaResponse) SetStatusCode(v int32) *AddMediaResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMediaResponse) SetBody(v *AddMediaResponseBody) *AddMediaResponse {
	s.Body = v
	return s
}

func (s *AddMediaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
