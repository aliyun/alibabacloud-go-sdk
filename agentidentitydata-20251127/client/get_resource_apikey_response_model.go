// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceAPIKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceAPIKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceAPIKeyResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceAPIKeyResponseBody) *GetResourceAPIKeyResponse
	GetBody() *GetResourceAPIKeyResponseBody
}

type GetResourceAPIKeyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceAPIKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceAPIKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceAPIKeyResponse) GoString() string {
	return s.String()
}

func (s *GetResourceAPIKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceAPIKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceAPIKeyResponse) GetBody() *GetResourceAPIKeyResponseBody {
	return s.Body
}

func (s *GetResourceAPIKeyResponse) SetHeaders(v map[string]*string) *GetResourceAPIKeyResponse {
	s.Headers = v
	return s
}

func (s *GetResourceAPIKeyResponse) SetStatusCode(v int32) *GetResourceAPIKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceAPIKeyResponse) SetBody(v *GetResourceAPIKeyResponseBody) *GetResourceAPIKeyResponse {
	s.Body = v
	return s
}

func (s *GetResourceAPIKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
