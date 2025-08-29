// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFeatureConsistencyCheckJobConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFeatureConsistencyCheckJobConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFeatureConsistencyCheckJobConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFeatureConsistencyCheckJobConfigResponseBody) *UpdateFeatureConsistencyCheckJobConfigResponse
	GetBody() *UpdateFeatureConsistencyCheckJobConfigResponseBody
}

type UpdateFeatureConsistencyCheckJobConfigResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFeatureConsistencyCheckJobConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFeatureConsistencyCheckJobConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFeatureConsistencyCheckJobConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateFeatureConsistencyCheckJobConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFeatureConsistencyCheckJobConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFeatureConsistencyCheckJobConfigResponse) GetBody() *UpdateFeatureConsistencyCheckJobConfigResponseBody {
	return s.Body
}

func (s *UpdateFeatureConsistencyCheckJobConfigResponse) SetHeaders(v map[string]*string) *UpdateFeatureConsistencyCheckJobConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigResponse) SetStatusCode(v int32) *UpdateFeatureConsistencyCheckJobConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigResponse) SetBody(v *UpdateFeatureConsistencyCheckJobConfigResponseBody) *UpdateFeatureConsistencyCheckJobConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigResponse) Validate() error {
	return dara.Validate(s)
}
