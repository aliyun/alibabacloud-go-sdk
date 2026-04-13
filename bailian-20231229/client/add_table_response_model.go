// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddTableResponse
	GetStatusCode() *int32
	SetBody(v *AddTableResponseBody) *AddTableResponse
	GetBody() *AddTableResponseBody
}

type AddTableResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTableResponse) String() string {
	return dara.Prettify(s)
}

func (s AddTableResponse) GoString() string {
	return s.String()
}

func (s *AddTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddTableResponse) GetBody() *AddTableResponseBody {
	return s.Body
}

func (s *AddTableResponse) SetHeaders(v map[string]*string) *AddTableResponse {
	s.Headers = v
	return s
}

func (s *AddTableResponse) SetStatusCode(v int32) *AddTableResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTableResponse) SetBody(v *AddTableResponseBody) *AddTableResponse {
	s.Body = v
	return s
}

func (s *AddTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
