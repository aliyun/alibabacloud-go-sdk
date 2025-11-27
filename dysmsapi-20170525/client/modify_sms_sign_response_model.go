// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySmsSignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySmsSignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySmsSignResponse
	GetStatusCode() *int32
	SetBody(v *ModifySmsSignResponseBody) *ModifySmsSignResponse
	GetBody() *ModifySmsSignResponseBody
}

type ModifySmsSignResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySmsSignResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySmsSignResponse) GoString() string {
	return s.String()
}

func (s *ModifySmsSignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySmsSignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySmsSignResponse) GetBody() *ModifySmsSignResponseBody {
	return s.Body
}

func (s *ModifySmsSignResponse) SetHeaders(v map[string]*string) *ModifySmsSignResponse {
	s.Headers = v
	return s
}

func (s *ModifySmsSignResponse) SetStatusCode(v int32) *ModifySmsSignResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySmsSignResponse) SetBody(v *ModifySmsSignResponseBody) *ModifySmsSignResponse {
	s.Body = v
	return s
}

func (s *ModifySmsSignResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
