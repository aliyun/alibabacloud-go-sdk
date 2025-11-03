// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenTrafficMirrorServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenTrafficMirrorServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenTrafficMirrorServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenTrafficMirrorServiceResponseBody) *OpenTrafficMirrorServiceResponse
	GetBody() *OpenTrafficMirrorServiceResponseBody
}

type OpenTrafficMirrorServiceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenTrafficMirrorServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenTrafficMirrorServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenTrafficMirrorServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenTrafficMirrorServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenTrafficMirrorServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenTrafficMirrorServiceResponse) GetBody() *OpenTrafficMirrorServiceResponseBody {
	return s.Body
}

func (s *OpenTrafficMirrorServiceResponse) SetHeaders(v map[string]*string) *OpenTrafficMirrorServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenTrafficMirrorServiceResponse) SetStatusCode(v int32) *OpenTrafficMirrorServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenTrafficMirrorServiceResponse) SetBody(v *OpenTrafficMirrorServiceResponseBody) *OpenTrafficMirrorServiceResponse {
	s.Body = v
	return s
}

func (s *OpenTrafficMirrorServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
