// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceCompareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FaceCompareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FaceCompareResponse
	GetStatusCode() *int32
	SetBody(v *FaceCompareResponseBody) *FaceCompareResponse
	GetBody() *FaceCompareResponseBody
}

type FaceCompareResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FaceCompareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FaceCompareResponse) String() string {
	return dara.Prettify(s)
}

func (s FaceCompareResponse) GoString() string {
	return s.String()
}

func (s *FaceCompareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FaceCompareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FaceCompareResponse) GetBody() *FaceCompareResponseBody {
	return s.Body
}

func (s *FaceCompareResponse) SetHeaders(v map[string]*string) *FaceCompareResponse {
	s.Headers = v
	return s
}

func (s *FaceCompareResponse) SetStatusCode(v int32) *FaceCompareResponse {
	s.StatusCode = &v
	return s
}

func (s *FaceCompareResponse) SetBody(v *FaceCompareResponseBody) *FaceCompareResponse {
	s.Body = v
	return s
}

func (s *FaceCompareResponse) Validate() error {
	return dara.Validate(s)
}
