// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateMacSecKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateMacSecKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateMacSecKeyResponse
	GetStatusCode() *int32
	SetBody(v *AssociateMacSecKeyResponseBody) *AssociateMacSecKeyResponse
	GetBody() *AssociateMacSecKeyResponseBody
}

type AssociateMacSecKeyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateMacSecKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateMacSecKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateMacSecKeyResponse) GoString() string {
	return s.String()
}

func (s *AssociateMacSecKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateMacSecKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateMacSecKeyResponse) GetBody() *AssociateMacSecKeyResponseBody {
	return s.Body
}

func (s *AssociateMacSecKeyResponse) SetHeaders(v map[string]*string) *AssociateMacSecKeyResponse {
	s.Headers = v
	return s
}

func (s *AssociateMacSecKeyResponse) SetStatusCode(v int32) *AssociateMacSecKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateMacSecKeyResponse) SetBody(v *AssociateMacSecKeyResponseBody) *AssociateMacSecKeyResponse {
	s.Body = v
	return s
}

func (s *AssociateMacSecKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
