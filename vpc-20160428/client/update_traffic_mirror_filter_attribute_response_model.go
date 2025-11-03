// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrafficMirrorFilterAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTrafficMirrorFilterAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTrafficMirrorFilterAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTrafficMirrorFilterAttributeResponseBody) *UpdateTrafficMirrorFilterAttributeResponse
	GetBody() *UpdateTrafficMirrorFilterAttributeResponseBody
}

type UpdateTrafficMirrorFilterAttributeResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTrafficMirrorFilterAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTrafficMirrorFilterAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficMirrorFilterAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateTrafficMirrorFilterAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTrafficMirrorFilterAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTrafficMirrorFilterAttributeResponse) GetBody() *UpdateTrafficMirrorFilterAttributeResponseBody {
	return s.Body
}

func (s *UpdateTrafficMirrorFilterAttributeResponse) SetHeaders(v map[string]*string) *UpdateTrafficMirrorFilterAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateTrafficMirrorFilterAttributeResponse) SetStatusCode(v int32) *UpdateTrafficMirrorFilterAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTrafficMirrorFilterAttributeResponse) SetBody(v *UpdateTrafficMirrorFilterAttributeResponseBody) *UpdateTrafficMirrorFilterAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateTrafficMirrorFilterAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
