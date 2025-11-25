// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoCcBlacklistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAutoCcBlacklistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAutoCcBlacklistResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAutoCcBlacklistResponseBody) *DeleteAutoCcBlacklistResponse
	GetBody() *DeleteAutoCcBlacklistResponseBody
}

type DeleteAutoCcBlacklistResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAutoCcBlacklistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAutoCcBlacklistResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoCcBlacklistResponse) GoString() string {
	return s.String()
}

func (s *DeleteAutoCcBlacklistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAutoCcBlacklistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAutoCcBlacklistResponse) GetBody() *DeleteAutoCcBlacklistResponseBody {
	return s.Body
}

func (s *DeleteAutoCcBlacklistResponse) SetHeaders(v map[string]*string) *DeleteAutoCcBlacklistResponse {
	s.Headers = v
	return s
}

func (s *DeleteAutoCcBlacklistResponse) SetStatusCode(v int32) *DeleteAutoCcBlacklistResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAutoCcBlacklistResponse) SetBody(v *DeleteAutoCcBlacklistResponseBody) *DeleteAutoCcBlacklistResponse {
	s.Body = v
	return s
}

func (s *DeleteAutoCcBlacklistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
