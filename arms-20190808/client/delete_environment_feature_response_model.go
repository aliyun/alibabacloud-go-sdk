// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnvironmentFeatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEnvironmentFeatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEnvironmentFeatureResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEnvironmentFeatureResponseBody) *DeleteEnvironmentFeatureResponse
	GetBody() *DeleteEnvironmentFeatureResponseBody
}

type DeleteEnvironmentFeatureResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEnvironmentFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEnvironmentFeatureResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnvironmentFeatureResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnvironmentFeatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEnvironmentFeatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEnvironmentFeatureResponse) GetBody() *DeleteEnvironmentFeatureResponseBody {
	return s.Body
}

func (s *DeleteEnvironmentFeatureResponse) SetHeaders(v map[string]*string) *DeleteEnvironmentFeatureResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnvironmentFeatureResponse) SetStatusCode(v int32) *DeleteEnvironmentFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEnvironmentFeatureResponse) SetBody(v *DeleteEnvironmentFeatureResponseBody) *DeleteEnvironmentFeatureResponse {
	s.Body = v
	return s
}

func (s *DeleteEnvironmentFeatureResponse) Validate() error {
	return dara.Validate(s)
}
