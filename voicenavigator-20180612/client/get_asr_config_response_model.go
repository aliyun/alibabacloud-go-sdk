// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsrConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAsrConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAsrConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetAsrConfigResponseBody) *GetAsrConfigResponse
	GetBody() *GetAsrConfigResponseBody
}

type GetAsrConfigResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAsrConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAsrConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAsrConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAsrConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAsrConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAsrConfigResponse) GetBody() *GetAsrConfigResponseBody {
	return s.Body
}

func (s *GetAsrConfigResponse) SetHeaders(v map[string]*string) *GetAsrConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAsrConfigResponse) SetStatusCode(v int32) *GetAsrConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsrConfigResponse) SetBody(v *GetAsrConfigResponseBody) *GetAsrConfigResponse {
	s.Body = v
	return s
}

func (s *GetAsrConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
