// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFileResponse
	GetStatusCode() *int32
	SetBody(v *AddFileResponseBody) *AddFileResponse
	GetBody() *AddFileResponseBody
}

type AddFileResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFileResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFileResponse) GoString() string {
	return s.String()
}

func (s *AddFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFileResponse) GetBody() *AddFileResponseBody {
	return s.Body
}

func (s *AddFileResponse) SetHeaders(v map[string]*string) *AddFileResponse {
	s.Headers = v
	return s
}

func (s *AddFileResponse) SetStatusCode(v int32) *AddFileResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFileResponse) SetBody(v *AddFileResponseBody) *AddFileResponse {
	s.Body = v
	return s
}

func (s *AddFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
