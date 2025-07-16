// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFeatureEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFeatureEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFeatureEntityResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFeatureEntityResponseBody) *DeleteFeatureEntityResponse
	GetBody() *DeleteFeatureEntityResponseBody
}

type DeleteFeatureEntityResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFeatureEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFeatureEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFeatureEntityResponse) GoString() string {
	return s.String()
}

func (s *DeleteFeatureEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFeatureEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFeatureEntityResponse) GetBody() *DeleteFeatureEntityResponseBody {
	return s.Body
}

func (s *DeleteFeatureEntityResponse) SetHeaders(v map[string]*string) *DeleteFeatureEntityResponse {
	s.Headers = v
	return s
}

func (s *DeleteFeatureEntityResponse) SetStatusCode(v int32) *DeleteFeatureEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFeatureEntityResponse) SetBody(v *DeleteFeatureEntityResponseBody) *DeleteFeatureEntityResponse {
	s.Body = v
	return s
}

func (s *DeleteFeatureEntityResponse) Validate() error {
	return dara.Validate(s)
}
