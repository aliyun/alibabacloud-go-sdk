// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetTableTokenResponseBody) *GetTableTokenResponse
	GetBody() *GetTableTokenResponseBody
}

type GetTableTokenResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableTokenResponse) GoString() string {
	return s.String()
}

func (s *GetTableTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableTokenResponse) GetBody() *GetTableTokenResponseBody {
	return s.Body
}

func (s *GetTableTokenResponse) SetHeaders(v map[string]*string) *GetTableTokenResponse {
	s.Headers = v
	return s
}

func (s *GetTableTokenResponse) SetStatusCode(v int32) *GetTableTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableTokenResponse) SetBody(v *GetTableTokenResponseBody) *GetTableTokenResponse {
	s.Body = v
	return s
}

func (s *GetTableTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
