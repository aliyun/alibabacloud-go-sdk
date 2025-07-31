// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFeatureConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFeatureConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFeatureConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFeatureConfigResponseBody) *DeleteFeatureConfigResponse
	GetBody() *DeleteFeatureConfigResponseBody
}

type DeleteFeatureConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFeatureConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFeatureConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFeatureConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteFeatureConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFeatureConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFeatureConfigResponse) GetBody() *DeleteFeatureConfigResponseBody {
	return s.Body
}

func (s *DeleteFeatureConfigResponse) SetHeaders(v map[string]*string) *DeleteFeatureConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteFeatureConfigResponse) SetStatusCode(v int32) *DeleteFeatureConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFeatureConfigResponse) SetBody(v *DeleteFeatureConfigResponseBody) *DeleteFeatureConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteFeatureConfigResponse) Validate() error {
	return dara.Validate(s)
}
