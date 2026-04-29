// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolarClawConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPolarClawConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPolarClawConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetPolarClawConfigResponseBody) *GetPolarClawConfigResponse
	GetBody() *GetPolarClawConfigResponseBody
}

type GetPolarClawConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPolarClawConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPolarClawConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPolarClawConfigResponse) GoString() string {
	return s.String()
}

func (s *GetPolarClawConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPolarClawConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPolarClawConfigResponse) GetBody() *GetPolarClawConfigResponseBody {
	return s.Body
}

func (s *GetPolarClawConfigResponse) SetHeaders(v map[string]*string) *GetPolarClawConfigResponse {
	s.Headers = v
	return s
}

func (s *GetPolarClawConfigResponse) SetStatusCode(v int32) *GetPolarClawConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPolarClawConfigResponse) SetBody(v *GetPolarClawConfigResponseBody) *GetPolarClawConfigResponse {
	s.Body = v
	return s
}

func (s *GetPolarClawConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
