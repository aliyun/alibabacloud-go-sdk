// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectFaceAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectFaceAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectFaceAttributesResponse
	GetStatusCode() *int32
	SetBody(v *DetectFaceAttributesResponseBody) *DetectFaceAttributesResponse
	GetBody() *DetectFaceAttributesResponseBody
}

type DetectFaceAttributesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectFaceAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectFaceAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectFaceAttributesResponse) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectFaceAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectFaceAttributesResponse) GetBody() *DetectFaceAttributesResponseBody {
	return s.Body
}

func (s *DetectFaceAttributesResponse) SetHeaders(v map[string]*string) *DetectFaceAttributesResponse {
	s.Headers = v
	return s
}

func (s *DetectFaceAttributesResponse) SetStatusCode(v int32) *DetectFaceAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectFaceAttributesResponse) SetBody(v *DetectFaceAttributesResponseBody) *DetectFaceAttributesResponse {
	s.Body = v
	return s
}

func (s *DetectFaceAttributesResponse) Validate() error {
	return dara.Validate(s)
}
