// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageCroppingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectImageCroppingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectImageCroppingResponse
	GetStatusCode() *int32
	SetBody(v *DetectImageCroppingResponseBody) *DetectImageCroppingResponse
	GetBody() *DetectImageCroppingResponseBody
}

type DetectImageCroppingResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectImageCroppingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectImageCroppingResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectImageCroppingResponse) GoString() string {
	return s.String()
}

func (s *DetectImageCroppingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectImageCroppingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectImageCroppingResponse) GetBody() *DetectImageCroppingResponseBody {
	return s.Body
}

func (s *DetectImageCroppingResponse) SetHeaders(v map[string]*string) *DetectImageCroppingResponse {
	s.Headers = v
	return s
}

func (s *DetectImageCroppingResponse) SetStatusCode(v int32) *DetectImageCroppingResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectImageCroppingResponse) SetBody(v *DetectImageCroppingResponseBody) *DetectImageCroppingResponse {
	s.Body = v
	return s
}

func (s *DetectImageCroppingResponse) Validate() error {
	return dara.Validate(s)
}
