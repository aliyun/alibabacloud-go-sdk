// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFeatureConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFeatureConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFeatureConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetFeatureConfigResponseBody) *GetFeatureConfigResponse
	GetBody() *GetFeatureConfigResponseBody
}

type GetFeatureConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFeatureConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFeatureConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureConfigResponse) GoString() string {
	return s.String()
}

func (s *GetFeatureConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFeatureConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFeatureConfigResponse) GetBody() *GetFeatureConfigResponseBody {
	return s.Body
}

func (s *GetFeatureConfigResponse) SetHeaders(v map[string]*string) *GetFeatureConfigResponse {
	s.Headers = v
	return s
}

func (s *GetFeatureConfigResponse) SetStatusCode(v int32) *GetFeatureConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFeatureConfigResponse) SetBody(v *GetFeatureConfigResponseBody) *GetFeatureConfigResponse {
	s.Body = v
	return s
}

func (s *GetFeatureConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
