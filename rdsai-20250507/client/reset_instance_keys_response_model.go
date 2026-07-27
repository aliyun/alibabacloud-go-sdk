// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetInstanceKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetInstanceKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetInstanceKeysResponse
	GetStatusCode() *int32
	SetBody(v *ResetInstanceKeysResponseBody) *ResetInstanceKeysResponse
	GetBody() *ResetInstanceKeysResponseBody
}

type ResetInstanceKeysResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetInstanceKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetInstanceKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetInstanceKeysResponse) GoString() string {
	return s.String()
}

func (s *ResetInstanceKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetInstanceKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetInstanceKeysResponse) GetBody() *ResetInstanceKeysResponseBody {
	return s.Body
}

func (s *ResetInstanceKeysResponse) SetHeaders(v map[string]*string) *ResetInstanceKeysResponse {
	s.Headers = v
	return s
}

func (s *ResetInstanceKeysResponse) SetStatusCode(v int32) *ResetInstanceKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetInstanceKeysResponse) SetBody(v *ResetInstanceKeysResponseBody) *ResetInstanceKeysResponse {
	s.Body = v
	return s
}

func (s *ResetInstanceKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
