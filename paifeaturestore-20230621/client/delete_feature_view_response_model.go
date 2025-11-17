// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFeatureViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFeatureViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFeatureViewResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFeatureViewResponseBody) *DeleteFeatureViewResponse
	GetBody() *DeleteFeatureViewResponseBody
}

type DeleteFeatureViewResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFeatureViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFeatureViewResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFeatureViewResponse) GoString() string {
	return s.String()
}

func (s *DeleteFeatureViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFeatureViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFeatureViewResponse) GetBody() *DeleteFeatureViewResponseBody {
	return s.Body
}

func (s *DeleteFeatureViewResponse) SetHeaders(v map[string]*string) *DeleteFeatureViewResponse {
	s.Headers = v
	return s
}

func (s *DeleteFeatureViewResponse) SetStatusCode(v int32) *DeleteFeatureViewResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFeatureViewResponse) SetBody(v *DeleteFeatureViewResponseBody) *DeleteFeatureViewResponse {
	s.Body = v
	return s
}

func (s *DeleteFeatureViewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
