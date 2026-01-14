// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportNumberResponse
	GetStatusCode() *int32
	SetBody(v *ImportNumberResponseBody) *ImportNumberResponse
	GetBody() *ImportNumberResponseBody
}

type ImportNumberResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportNumberResponse) GoString() string {
	return s.String()
}

func (s *ImportNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportNumberResponse) GetBody() *ImportNumberResponseBody {
	return s.Body
}

func (s *ImportNumberResponse) SetHeaders(v map[string]*string) *ImportNumberResponse {
	s.Headers = v
	return s
}

func (s *ImportNumberResponse) SetStatusCode(v int32) *ImportNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportNumberResponse) SetBody(v *ImportNumberResponseBody) *ImportNumberResponse {
	s.Body = v
	return s
}

func (s *ImportNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
