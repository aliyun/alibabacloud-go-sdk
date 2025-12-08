// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPedestrianDetectAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PedestrianDetectAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PedestrianDetectAttributeResponse
	GetStatusCode() *int32
	SetBody(v *PedestrianDetectAttributeResponseBody) *PedestrianDetectAttributeResponse
	GetBody() *PedestrianDetectAttributeResponseBody
}

type PedestrianDetectAttributeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PedestrianDetectAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PedestrianDetectAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s PedestrianDetectAttributeResponse) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PedestrianDetectAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PedestrianDetectAttributeResponse) GetBody() *PedestrianDetectAttributeResponseBody {
	return s.Body
}

func (s *PedestrianDetectAttributeResponse) SetHeaders(v map[string]*string) *PedestrianDetectAttributeResponse {
	s.Headers = v
	return s
}

func (s *PedestrianDetectAttributeResponse) SetStatusCode(v int32) *PedestrianDetectAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *PedestrianDetectAttributeResponse) SetBody(v *PedestrianDetectAttributeResponseBody) *PedestrianDetectAttributeResponse {
	s.Body = v
	return s
}

func (s *PedestrianDetectAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
