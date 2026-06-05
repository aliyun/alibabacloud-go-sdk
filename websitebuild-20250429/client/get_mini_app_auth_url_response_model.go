// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMiniAppAuthUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMiniAppAuthUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMiniAppAuthUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetMiniAppAuthUrlResponseBody) *GetMiniAppAuthUrlResponse
	GetBody() *GetMiniAppAuthUrlResponseBody
}

type GetMiniAppAuthUrlResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMiniAppAuthUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMiniAppAuthUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMiniAppAuthUrlResponse) GoString() string {
	return s.String()
}

func (s *GetMiniAppAuthUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMiniAppAuthUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMiniAppAuthUrlResponse) GetBody() *GetMiniAppAuthUrlResponseBody {
	return s.Body
}

func (s *GetMiniAppAuthUrlResponse) SetHeaders(v map[string]*string) *GetMiniAppAuthUrlResponse {
	s.Headers = v
	return s
}

func (s *GetMiniAppAuthUrlResponse) SetStatusCode(v int32) *GetMiniAppAuthUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMiniAppAuthUrlResponse) SetBody(v *GetMiniAppAuthUrlResponseBody) *GetMiniAppAuthUrlResponse {
	s.Body = v
	return s
}

func (s *GetMiniAppAuthUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
