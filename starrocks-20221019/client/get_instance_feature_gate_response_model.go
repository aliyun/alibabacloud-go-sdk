// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceFeatureGateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceFeatureGateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceFeatureGateResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceFeatureGateResponseBody) *GetInstanceFeatureGateResponse
	GetBody() *GetInstanceFeatureGateResponseBody
}

type GetInstanceFeatureGateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceFeatureGateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceFeatureGateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceFeatureGateResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceFeatureGateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceFeatureGateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceFeatureGateResponse) GetBody() *GetInstanceFeatureGateResponseBody {
	return s.Body
}

func (s *GetInstanceFeatureGateResponse) SetHeaders(v map[string]*string) *GetInstanceFeatureGateResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceFeatureGateResponse) SetStatusCode(v int32) *GetInstanceFeatureGateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceFeatureGateResponse) SetBody(v *GetInstanceFeatureGateResponseBody) *GetInstanceFeatureGateResponse {
	s.Body = v
	return s
}

func (s *GetInstanceFeatureGateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
