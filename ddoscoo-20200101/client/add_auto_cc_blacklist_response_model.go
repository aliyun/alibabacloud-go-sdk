// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAutoCcBlacklistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAutoCcBlacklistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAutoCcBlacklistResponse
	GetStatusCode() *int32
	SetBody(v *AddAutoCcBlacklistResponseBody) *AddAutoCcBlacklistResponse
	GetBody() *AddAutoCcBlacklistResponseBody
}

type AddAutoCcBlacklistResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAutoCcBlacklistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAutoCcBlacklistResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAutoCcBlacklistResponse) GoString() string {
	return s.String()
}

func (s *AddAutoCcBlacklistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAutoCcBlacklistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAutoCcBlacklistResponse) GetBody() *AddAutoCcBlacklistResponseBody {
	return s.Body
}

func (s *AddAutoCcBlacklistResponse) SetHeaders(v map[string]*string) *AddAutoCcBlacklistResponse {
	s.Headers = v
	return s
}

func (s *AddAutoCcBlacklistResponse) SetStatusCode(v int32) *AddAutoCcBlacklistResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAutoCcBlacklistResponse) SetBody(v *AddAutoCcBlacklistResponseBody) *AddAutoCcBlacklistResponse {
	s.Body = v
	return s
}

func (s *AddAutoCcBlacklistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
