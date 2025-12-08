// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareFaceWithMaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompareFaceWithMaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompareFaceWithMaskResponse
	GetStatusCode() *int32
	SetBody(v *CompareFaceWithMaskResponseBody) *CompareFaceWithMaskResponse
	GetBody() *CompareFaceWithMaskResponseBody
}

type CompareFaceWithMaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompareFaceWithMaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompareFaceWithMaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CompareFaceWithMaskResponse) GoString() string {
	return s.String()
}

func (s *CompareFaceWithMaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompareFaceWithMaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompareFaceWithMaskResponse) GetBody() *CompareFaceWithMaskResponseBody {
	return s.Body
}

func (s *CompareFaceWithMaskResponse) SetHeaders(v map[string]*string) *CompareFaceWithMaskResponse {
	s.Headers = v
	return s
}

func (s *CompareFaceWithMaskResponse) SetStatusCode(v int32) *CompareFaceWithMaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CompareFaceWithMaskResponse) SetBody(v *CompareFaceWithMaskResponseBody) *CompareFaceWithMaskResponse {
	s.Body = v
	return s
}

func (s *CompareFaceWithMaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
