// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScaleServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScaleServiceResponse
	GetStatusCode() *int32
	SetBody(v *ScaleServiceResponseBody) *ScaleServiceResponse
	GetBody() *ScaleServiceResponseBody
}

type ScaleServiceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScaleServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScaleServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ScaleServiceResponse) GoString() string {
	return s.String()
}

func (s *ScaleServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScaleServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScaleServiceResponse) GetBody() *ScaleServiceResponseBody {
	return s.Body
}

func (s *ScaleServiceResponse) SetHeaders(v map[string]*string) *ScaleServiceResponse {
	s.Headers = v
	return s
}

func (s *ScaleServiceResponse) SetStatusCode(v int32) *ScaleServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ScaleServiceResponse) SetBody(v *ScaleServiceResponseBody) *ScaleServiceResponse {
	s.Body = v
	return s
}

func (s *ScaleServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
