// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNormalServiceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNormalServiceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNormalServiceConfigResponse
	GetStatusCode() *int32
	SetBody(v *ServiceConfigResult) *GetNormalServiceConfigResponse
	GetBody() *ServiceConfigResult
}

type GetNormalServiceConfigResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ServiceConfigResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNormalServiceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNormalServiceConfigResponse) GoString() string {
	return s.String()
}

func (s *GetNormalServiceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNormalServiceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNormalServiceConfigResponse) GetBody() *ServiceConfigResult {
	return s.Body
}

func (s *GetNormalServiceConfigResponse) SetHeaders(v map[string]*string) *GetNormalServiceConfigResponse {
	s.Headers = v
	return s
}

func (s *GetNormalServiceConfigResponse) SetStatusCode(v int32) *GetNormalServiceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNormalServiceConfigResponse) SetBody(v *ServiceConfigResult) *GetNormalServiceConfigResponse {
	s.Body = v
	return s
}

func (s *GetNormalServiceConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
