// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareFaceVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompareFaceVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompareFaceVerifyResponse
	GetStatusCode() *int32
	SetBody(v *CompareFaceVerifyResponseBody) *CompareFaceVerifyResponse
	GetBody() *CompareFaceVerifyResponseBody
}

type CompareFaceVerifyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompareFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompareFaceVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s CompareFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *CompareFaceVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompareFaceVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompareFaceVerifyResponse) GetBody() *CompareFaceVerifyResponseBody {
	return s.Body
}

func (s *CompareFaceVerifyResponse) SetHeaders(v map[string]*string) *CompareFaceVerifyResponse {
	s.Headers = v
	return s
}

func (s *CompareFaceVerifyResponse) SetStatusCode(v int32) *CompareFaceVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *CompareFaceVerifyResponse) SetBody(v *CompareFaceVerifyResponseBody) *CompareFaceVerifyResponse {
	s.Body = v
	return s
}

func (s *CompareFaceVerifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
