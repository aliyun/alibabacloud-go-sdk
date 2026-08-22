// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleInOpenSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScaleInOpenSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScaleInOpenSearchResponse
	GetStatusCode() *int32
	SetBody(v *ScaleInOpenSearchResponseBody) *ScaleInOpenSearchResponse
	GetBody() *ScaleInOpenSearchResponseBody
}

type ScaleInOpenSearchResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScaleInOpenSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScaleInOpenSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s ScaleInOpenSearchResponse) GoString() string {
	return s.String()
}

func (s *ScaleInOpenSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScaleInOpenSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScaleInOpenSearchResponse) GetBody() *ScaleInOpenSearchResponseBody {
	return s.Body
}

func (s *ScaleInOpenSearchResponse) SetHeaders(v map[string]*string) *ScaleInOpenSearchResponse {
	s.Headers = v
	return s
}

func (s *ScaleInOpenSearchResponse) SetStatusCode(v int32) *ScaleInOpenSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *ScaleInOpenSearchResponse) SetBody(v *ScaleInOpenSearchResponseBody) *ScaleInOpenSearchResponse {
	s.Body = v
	return s
}

func (s *ScaleInOpenSearchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
