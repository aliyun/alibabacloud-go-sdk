// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMseFeatureSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMseFeatureSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMseFeatureSwitchResponse
	GetStatusCode() *int32
	SetBody(v *GetMseFeatureSwitchResponseBody) *GetMseFeatureSwitchResponse
	GetBody() *GetMseFeatureSwitchResponseBody
}

type GetMseFeatureSwitchResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMseFeatureSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMseFeatureSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMseFeatureSwitchResponse) GoString() string {
	return s.String()
}

func (s *GetMseFeatureSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMseFeatureSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMseFeatureSwitchResponse) GetBody() *GetMseFeatureSwitchResponseBody {
	return s.Body
}

func (s *GetMseFeatureSwitchResponse) SetHeaders(v map[string]*string) *GetMseFeatureSwitchResponse {
	s.Headers = v
	return s
}

func (s *GetMseFeatureSwitchResponse) SetStatusCode(v int32) *GetMseFeatureSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMseFeatureSwitchResponse) SetBody(v *GetMseFeatureSwitchResponseBody) *GetMseFeatureSwitchResponse {
	s.Body = v
	return s
}

func (s *GetMseFeatureSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
