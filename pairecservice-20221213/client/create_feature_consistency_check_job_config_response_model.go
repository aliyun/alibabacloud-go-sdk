// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFeatureConsistencyCheckJobConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFeatureConsistencyCheckJobConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFeatureConsistencyCheckJobConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateFeatureConsistencyCheckJobConfigResponseBody) *CreateFeatureConsistencyCheckJobConfigResponse
	GetBody() *CreateFeatureConsistencyCheckJobConfigResponseBody
}

type CreateFeatureConsistencyCheckJobConfigResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFeatureConsistencyCheckJobConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFeatureConsistencyCheckJobConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFeatureConsistencyCheckJobConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateFeatureConsistencyCheckJobConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFeatureConsistencyCheckJobConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFeatureConsistencyCheckJobConfigResponse) GetBody() *CreateFeatureConsistencyCheckJobConfigResponseBody {
	return s.Body
}

func (s *CreateFeatureConsistencyCheckJobConfigResponse) SetHeaders(v map[string]*string) *CreateFeatureConsistencyCheckJobConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigResponse) SetStatusCode(v int32) *CreateFeatureConsistencyCheckJobConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigResponse) SetBody(v *CreateFeatureConsistencyCheckJobConfigResponseBody) *CreateFeatureConsistencyCheckJobConfigResponse {
	s.Body = v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
