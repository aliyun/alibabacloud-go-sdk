// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageBasicInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectImageBasicInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectImageBasicInfoResponse
	GetStatusCode() *int32
	SetBody(v *DetectImageBasicInfoResponseBody) *DetectImageBasicInfoResponse
	GetBody() *DetectImageBasicInfoResponseBody
}

type DetectImageBasicInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectImageBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectImageBasicInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectImageBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *DetectImageBasicInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectImageBasicInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectImageBasicInfoResponse) GetBody() *DetectImageBasicInfoResponseBody {
	return s.Body
}

func (s *DetectImageBasicInfoResponse) SetHeaders(v map[string]*string) *DetectImageBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *DetectImageBasicInfoResponse) SetStatusCode(v int32) *DetectImageBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectImageBasicInfoResponse) SetBody(v *DetectImageBasicInfoResponseBody) *DetectImageBasicInfoResponse {
	s.Body = v
	return s
}

func (s *DetectImageBasicInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
