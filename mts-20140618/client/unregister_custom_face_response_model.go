// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnregisterCustomFaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnregisterCustomFaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnregisterCustomFaceResponse
	GetStatusCode() *int32
	SetBody(v *UnregisterCustomFaceResponseBody) *UnregisterCustomFaceResponse
	GetBody() *UnregisterCustomFaceResponseBody
}

type UnregisterCustomFaceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnregisterCustomFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnregisterCustomFaceResponse) String() string {
	return dara.Prettify(s)
}

func (s UnregisterCustomFaceResponse) GoString() string {
	return s.String()
}

func (s *UnregisterCustomFaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnregisterCustomFaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnregisterCustomFaceResponse) GetBody() *UnregisterCustomFaceResponseBody {
	return s.Body
}

func (s *UnregisterCustomFaceResponse) SetHeaders(v map[string]*string) *UnregisterCustomFaceResponse {
	s.Headers = v
	return s
}

func (s *UnregisterCustomFaceResponse) SetStatusCode(v int32) *UnregisterCustomFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnregisterCustomFaceResponse) SetBody(v *UnregisterCustomFaceResponseBody) *UnregisterCustomFaceResponse {
	s.Body = v
	return s
}

func (s *UnregisterCustomFaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
