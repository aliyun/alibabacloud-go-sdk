// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRegionConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRegionConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRegionConfigResponse
	GetStatusCode() *int32
	SetBody(v *QueryRegionConfigResponseBody) *QueryRegionConfigResponse
	GetBody() *QueryRegionConfigResponseBody
}

type QueryRegionConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRegionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRegionConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRegionConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryRegionConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRegionConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRegionConfigResponse) GetBody() *QueryRegionConfigResponseBody {
	return s.Body
}

func (s *QueryRegionConfigResponse) SetHeaders(v map[string]*string) *QueryRegionConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryRegionConfigResponse) SetStatusCode(v int32) *QueryRegionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRegionConfigResponse) SetBody(v *QueryRegionConfigResponseBody) *QueryRegionConfigResponse {
	s.Body = v
	return s
}

func (s *QueryRegionConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
