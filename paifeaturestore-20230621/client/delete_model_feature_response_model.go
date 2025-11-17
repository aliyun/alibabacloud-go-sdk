// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelFeatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteModelFeatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteModelFeatureResponse
	GetStatusCode() *int32
	SetBody(v *DeleteModelFeatureResponseBody) *DeleteModelFeatureResponse
	GetBody() *DeleteModelFeatureResponseBody
}

type DeleteModelFeatureResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModelFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModelFeatureResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelFeatureResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelFeatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteModelFeatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteModelFeatureResponse) GetBody() *DeleteModelFeatureResponseBody {
	return s.Body
}

func (s *DeleteModelFeatureResponse) SetHeaders(v map[string]*string) *DeleteModelFeatureResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelFeatureResponse) SetStatusCode(v int32) *DeleteModelFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModelFeatureResponse) SetBody(v *DeleteModelFeatureResponseBody) *DeleteModelFeatureResponse {
	s.Body = v
	return s
}

func (s *DeleteModelFeatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
