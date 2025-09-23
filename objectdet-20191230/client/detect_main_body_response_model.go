// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectMainBodyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectMainBodyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectMainBodyResponse
	GetStatusCode() *int32
	SetBody(v *DetectMainBodyResponseBody) *DetectMainBodyResponse
	GetBody() *DetectMainBodyResponseBody
}

type DetectMainBodyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectMainBodyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectMainBodyResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectMainBodyResponse) GoString() string {
	return s.String()
}

func (s *DetectMainBodyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectMainBodyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectMainBodyResponse) GetBody() *DetectMainBodyResponseBody {
	return s.Body
}

func (s *DetectMainBodyResponse) SetHeaders(v map[string]*string) *DetectMainBodyResponse {
	s.Headers = v
	return s
}

func (s *DetectMainBodyResponse) SetStatusCode(v int32) *DetectMainBodyResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectMainBodyResponse) SetBody(v *DetectMainBodyResponseBody) *DetectMainBodyResponse {
	s.Body = v
	return s
}

func (s *DetectMainBodyResponse) Validate() error {
	return dara.Validate(s)
}
