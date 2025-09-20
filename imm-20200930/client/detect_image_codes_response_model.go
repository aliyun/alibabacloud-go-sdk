// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageCodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectImageCodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectImageCodesResponse
	GetStatusCode() *int32
	SetBody(v *DetectImageCodesResponseBody) *DetectImageCodesResponse
	GetBody() *DetectImageCodesResponseBody
}

type DetectImageCodesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectImageCodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectImageCodesResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectImageCodesResponse) GoString() string {
	return s.String()
}

func (s *DetectImageCodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectImageCodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectImageCodesResponse) GetBody() *DetectImageCodesResponseBody {
	return s.Body
}

func (s *DetectImageCodesResponse) SetHeaders(v map[string]*string) *DetectImageCodesResponse {
	s.Headers = v
	return s
}

func (s *DetectImageCodesResponse) SetStatusCode(v int32) *DetectImageCodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectImageCodesResponse) SetBody(v *DetectImageCodesResponseBody) *DetectImageCodesResponse {
	s.Body = v
	return s
}

func (s *DetectImageCodesResponse) Validate() error {
	return dara.Validate(s)
}
