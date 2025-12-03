// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProxyByTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProxyByTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProxyByTypeResponse
	GetStatusCode() *int32
	SetBody(v *GetProxyByTypeResponseBody) *GetProxyByTypeResponse
	GetBody() *GetProxyByTypeResponseBody
}

type GetProxyByTypeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProxyByTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProxyByTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProxyByTypeResponse) GoString() string {
	return s.String()
}

func (s *GetProxyByTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProxyByTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProxyByTypeResponse) GetBody() *GetProxyByTypeResponseBody {
	return s.Body
}

func (s *GetProxyByTypeResponse) SetHeaders(v map[string]*string) *GetProxyByTypeResponse {
	s.Headers = v
	return s
}

func (s *GetProxyByTypeResponse) SetStatusCode(v int32) *GetProxyByTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProxyByTypeResponse) SetBody(v *GetProxyByTypeResponseBody) *GetProxyByTypeResponse {
	s.Body = v
	return s
}

func (s *GetProxyByTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
