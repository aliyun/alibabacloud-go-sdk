// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrafficMirrorServiceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTrafficMirrorServiceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTrafficMirrorServiceStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetTrafficMirrorServiceStatusResponseBody) *GetTrafficMirrorServiceStatusResponse
	GetBody() *GetTrafficMirrorServiceStatusResponseBody
}

type GetTrafficMirrorServiceStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrafficMirrorServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrafficMirrorServiceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTrafficMirrorServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetTrafficMirrorServiceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTrafficMirrorServiceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTrafficMirrorServiceStatusResponse) GetBody() *GetTrafficMirrorServiceStatusResponseBody {
	return s.Body
}

func (s *GetTrafficMirrorServiceStatusResponse) SetHeaders(v map[string]*string) *GetTrafficMirrorServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetTrafficMirrorServiceStatusResponse) SetStatusCode(v int32) *GetTrafficMirrorServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrafficMirrorServiceStatusResponse) SetBody(v *GetTrafficMirrorServiceStatusResponseBody) *GetTrafficMirrorServiceStatusResponse {
	s.Body = v
	return s
}

func (s *GetTrafficMirrorServiceStatusResponse) Validate() error {
	return dara.Validate(s)
}
