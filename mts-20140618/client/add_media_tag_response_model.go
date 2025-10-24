// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMediaTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddMediaTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddMediaTagResponse
	GetStatusCode() *int32
	SetBody(v *AddMediaTagResponseBody) *AddMediaTagResponse
	GetBody() *AddMediaTagResponseBody
}

type AddMediaTagResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMediaTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMediaTagResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMediaTagResponse) GoString() string {
	return s.String()
}

func (s *AddMediaTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddMediaTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddMediaTagResponse) GetBody() *AddMediaTagResponseBody {
	return s.Body
}

func (s *AddMediaTagResponse) SetHeaders(v map[string]*string) *AddMediaTagResponse {
	s.Headers = v
	return s
}

func (s *AddMediaTagResponse) SetStatusCode(v int32) *AddMediaTagResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMediaTagResponse) SetBody(v *AddMediaTagResponseBody) *AddMediaTagResponse {
	s.Body = v
	return s
}

func (s *AddMediaTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
