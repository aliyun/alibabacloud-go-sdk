// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrafficMirrorSessionAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTrafficMirrorSessionAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTrafficMirrorSessionAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTrafficMirrorSessionAttributeResponseBody) *UpdateTrafficMirrorSessionAttributeResponse
	GetBody() *UpdateTrafficMirrorSessionAttributeResponseBody
}

type UpdateTrafficMirrorSessionAttributeResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTrafficMirrorSessionAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTrafficMirrorSessionAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficMirrorSessionAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateTrafficMirrorSessionAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTrafficMirrorSessionAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTrafficMirrorSessionAttributeResponse) GetBody() *UpdateTrafficMirrorSessionAttributeResponseBody {
	return s.Body
}

func (s *UpdateTrafficMirrorSessionAttributeResponse) SetHeaders(v map[string]*string) *UpdateTrafficMirrorSessionAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateTrafficMirrorSessionAttributeResponse) SetStatusCode(v int32) *UpdateTrafficMirrorSessionAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTrafficMirrorSessionAttributeResponse) SetBody(v *UpdateTrafficMirrorSessionAttributeResponseBody) *UpdateTrafficMirrorSessionAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateTrafficMirrorSessionAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
