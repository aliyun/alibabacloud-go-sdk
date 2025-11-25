// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAutoCcWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAutoCcWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAutoCcWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *AddAutoCcWhitelistResponseBody) *AddAutoCcWhitelistResponse
	GetBody() *AddAutoCcWhitelistResponseBody
}

type AddAutoCcWhitelistResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAutoCcWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAutoCcWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAutoCcWhitelistResponse) GoString() string {
	return s.String()
}

func (s *AddAutoCcWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAutoCcWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAutoCcWhitelistResponse) GetBody() *AddAutoCcWhitelistResponseBody {
	return s.Body
}

func (s *AddAutoCcWhitelistResponse) SetHeaders(v map[string]*string) *AddAutoCcWhitelistResponse {
	s.Headers = v
	return s
}

func (s *AddAutoCcWhitelistResponse) SetStatusCode(v int32) *AddAutoCcWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAutoCcWhitelistResponse) SetBody(v *AddAutoCcWhitelistResponseBody) *AddAutoCcWhitelistResponse {
	s.Body = v
	return s
}

func (s *AddAutoCcWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
