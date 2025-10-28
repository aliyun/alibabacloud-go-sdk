// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleoutApplicationWithNewInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScaleoutApplicationWithNewInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScaleoutApplicationWithNewInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ScaleoutApplicationWithNewInstancesResponseBody) *ScaleoutApplicationWithNewInstancesResponse
	GetBody() *ScaleoutApplicationWithNewInstancesResponseBody
}

type ScaleoutApplicationWithNewInstancesResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScaleoutApplicationWithNewInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScaleoutApplicationWithNewInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ScaleoutApplicationWithNewInstancesResponse) GoString() string {
	return s.String()
}

func (s *ScaleoutApplicationWithNewInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScaleoutApplicationWithNewInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScaleoutApplicationWithNewInstancesResponse) GetBody() *ScaleoutApplicationWithNewInstancesResponseBody {
	return s.Body
}

func (s *ScaleoutApplicationWithNewInstancesResponse) SetHeaders(v map[string]*string) *ScaleoutApplicationWithNewInstancesResponse {
	s.Headers = v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesResponse) SetStatusCode(v int32) *ScaleoutApplicationWithNewInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesResponse) SetBody(v *ScaleoutApplicationWithNewInstancesResponseBody) *ScaleoutApplicationWithNewInstancesResponse {
	s.Body = v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
