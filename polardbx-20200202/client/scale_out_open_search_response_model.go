// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleOutOpenSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScaleOutOpenSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScaleOutOpenSearchResponse
	GetStatusCode() *int32
	SetBody(v *ScaleOutOpenSearchResponseBody) *ScaleOutOpenSearchResponse
	GetBody() *ScaleOutOpenSearchResponseBody
}

type ScaleOutOpenSearchResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScaleOutOpenSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScaleOutOpenSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s ScaleOutOpenSearchResponse) GoString() string {
	return s.String()
}

func (s *ScaleOutOpenSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScaleOutOpenSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScaleOutOpenSearchResponse) GetBody() *ScaleOutOpenSearchResponseBody {
	return s.Body
}

func (s *ScaleOutOpenSearchResponse) SetHeaders(v map[string]*string) *ScaleOutOpenSearchResponse {
	s.Headers = v
	return s
}

func (s *ScaleOutOpenSearchResponse) SetStatusCode(v int32) *ScaleOutOpenSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *ScaleOutOpenSearchResponse) SetBody(v *ScaleOutOpenSearchResponseBody) *ScaleOutOpenSearchResponse {
	s.Body = v
	return s
}

func (s *ScaleOutOpenSearchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
