// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetAppConfigResponseBody) *GetAppConfigResponse
	GetBody() *GetAppConfigResponseBody
}

type GetAppConfigResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAppConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppConfigResponse) GetBody() *GetAppConfigResponseBody {
	return s.Body
}

func (s *GetAppConfigResponse) SetHeaders(v map[string]*string) *GetAppConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAppConfigResponse) SetStatusCode(v int32) *GetAppConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppConfigResponse) SetBody(v *GetAppConfigResponseBody) *GetAppConfigResponse {
	s.Body = v
	return s
}

func (s *GetAppConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
