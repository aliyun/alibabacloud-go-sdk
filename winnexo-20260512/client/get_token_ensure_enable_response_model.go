// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenEnsureEnableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTokenEnsureEnableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTokenEnsureEnableResponse
	GetStatusCode() *int32
	SetBody(v *GetTokenEnsureEnableResponseBody) *GetTokenEnsureEnableResponse
	GetBody() *GetTokenEnsureEnableResponseBody
}

type GetTokenEnsureEnableResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTokenEnsureEnableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTokenEnsureEnableResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTokenEnsureEnableResponse) GoString() string {
	return s.String()
}

func (s *GetTokenEnsureEnableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTokenEnsureEnableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTokenEnsureEnableResponse) GetBody() *GetTokenEnsureEnableResponseBody {
	return s.Body
}

func (s *GetTokenEnsureEnableResponse) SetHeaders(v map[string]*string) *GetTokenEnsureEnableResponse {
	s.Headers = v
	return s
}

func (s *GetTokenEnsureEnableResponse) SetStatusCode(v int32) *GetTokenEnsureEnableResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTokenEnsureEnableResponse) SetBody(v *GetTokenEnsureEnableResponseBody) *GetTokenEnsureEnableResponse {
	s.Body = v
	return s
}

func (s *GetTokenEnsureEnableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
