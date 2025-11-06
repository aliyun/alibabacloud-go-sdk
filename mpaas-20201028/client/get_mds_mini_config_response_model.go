// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMdsMiniConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMdsMiniConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMdsMiniConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetMdsMiniConfigResponseBody) *GetMdsMiniConfigResponse
	GetBody() *GetMdsMiniConfigResponseBody
}

type GetMdsMiniConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMdsMiniConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMdsMiniConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMdsMiniConfigResponse) GoString() string {
	return s.String()
}

func (s *GetMdsMiniConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMdsMiniConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMdsMiniConfigResponse) GetBody() *GetMdsMiniConfigResponseBody {
	return s.Body
}

func (s *GetMdsMiniConfigResponse) SetHeaders(v map[string]*string) *GetMdsMiniConfigResponse {
	s.Headers = v
	return s
}

func (s *GetMdsMiniConfigResponse) SetStatusCode(v int32) *GetMdsMiniConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMdsMiniConfigResponse) SetBody(v *GetMdsMiniConfigResponseBody) *GetMdsMiniConfigResponse {
	s.Body = v
	return s
}

func (s *GetMdsMiniConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
