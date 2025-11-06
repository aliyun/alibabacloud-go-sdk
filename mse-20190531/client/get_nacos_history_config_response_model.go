// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNacosHistoryConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNacosHistoryConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNacosHistoryConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetNacosHistoryConfigResponseBody) *GetNacosHistoryConfigResponse
	GetBody() *GetNacosHistoryConfigResponseBody
}

type GetNacosHistoryConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNacosHistoryConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNacosHistoryConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNacosHistoryConfigResponse) GoString() string {
	return s.String()
}

func (s *GetNacosHistoryConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNacosHistoryConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNacosHistoryConfigResponse) GetBody() *GetNacosHistoryConfigResponseBody {
	return s.Body
}

func (s *GetNacosHistoryConfigResponse) SetHeaders(v map[string]*string) *GetNacosHistoryConfigResponse {
	s.Headers = v
	return s
}

func (s *GetNacosHistoryConfigResponse) SetStatusCode(v int32) *GetNacosHistoryConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNacosHistoryConfigResponse) SetBody(v *GetNacosHistoryConfigResponseBody) *GetNacosHistoryConfigResponse {
	s.Body = v
	return s
}

func (s *GetNacosHistoryConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
