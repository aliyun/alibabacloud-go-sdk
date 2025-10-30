// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetActiveIdpConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetActiveIdpConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetActiveIdpConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetActiveIdpConfigResponseBody) *GetActiveIdpConfigResponse
	GetBody() *GetActiveIdpConfigResponseBody
}

type GetActiveIdpConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetActiveIdpConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetActiveIdpConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetActiveIdpConfigResponse) GoString() string {
	return s.String()
}

func (s *GetActiveIdpConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetActiveIdpConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetActiveIdpConfigResponse) GetBody() *GetActiveIdpConfigResponseBody {
	return s.Body
}

func (s *GetActiveIdpConfigResponse) SetHeaders(v map[string]*string) *GetActiveIdpConfigResponse {
	s.Headers = v
	return s
}

func (s *GetActiveIdpConfigResponse) SetStatusCode(v int32) *GetActiveIdpConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetActiveIdpConfigResponse) SetBody(v *GetActiveIdpConfigResponseBody) *GetActiveIdpConfigResponse {
	s.Body = v
	return s
}

func (s *GetActiveIdpConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
