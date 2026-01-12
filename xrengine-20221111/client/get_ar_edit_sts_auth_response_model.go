// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArEditStsAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetArEditStsAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetArEditStsAuthResponse
	GetStatusCode() *int32
	SetBody(v *GetArEditStsAuthResponseBody) *GetArEditStsAuthResponse
	GetBody() *GetArEditStsAuthResponseBody
}

type GetArEditStsAuthResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetArEditStsAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArEditStsAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s GetArEditStsAuthResponse) GoString() string {
	return s.String()
}

func (s *GetArEditStsAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetArEditStsAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetArEditStsAuthResponse) GetBody() *GetArEditStsAuthResponseBody {
	return s.Body
}

func (s *GetArEditStsAuthResponse) SetHeaders(v map[string]*string) *GetArEditStsAuthResponse {
	s.Headers = v
	return s
}

func (s *GetArEditStsAuthResponse) SetStatusCode(v int32) *GetArEditStsAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArEditStsAuthResponse) SetBody(v *GetArEditStsAuthResponseBody) *GetArEditStsAuthResponse {
	s.Body = v
	return s
}

func (s *GetArEditStsAuthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
