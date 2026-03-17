// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *GetApiKeyResponseBody) *GetApiKeyResponse
	GetBody() *GetApiKeyResponseBody
}

type GetApiKeyResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApiKeyResponse) GoString() string {
	return s.String()
}

func (s *GetApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApiKeyResponse) GetBody() *GetApiKeyResponseBody {
	return s.Body
}

func (s *GetApiKeyResponse) SetHeaders(v map[string]*string) *GetApiKeyResponse {
	s.Headers = v
	return s
}

func (s *GetApiKeyResponse) SetStatusCode(v int32) *GetApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApiKeyResponse) SetBody(v *GetApiKeyResponseBody) *GetApiKeyResponse {
	s.Body = v
	return s
}

func (s *GetApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
