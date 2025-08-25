// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssessSharpnessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssessSharpnessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssessSharpnessResponse
	GetStatusCode() *int32
	SetBody(v *AssessSharpnessResponseBody) *AssessSharpnessResponse
	GetBody() *AssessSharpnessResponseBody
}

type AssessSharpnessResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssessSharpnessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssessSharpnessResponse) String() string {
	return dara.Prettify(s)
}

func (s AssessSharpnessResponse) GoString() string {
	return s.String()
}

func (s *AssessSharpnessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssessSharpnessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssessSharpnessResponse) GetBody() *AssessSharpnessResponseBody {
	return s.Body
}

func (s *AssessSharpnessResponse) SetHeaders(v map[string]*string) *AssessSharpnessResponse {
	s.Headers = v
	return s
}

func (s *AssessSharpnessResponse) SetStatusCode(v int32) *AssessSharpnessResponse {
	s.StatusCode = &v
	return s
}

func (s *AssessSharpnessResponse) SetBody(v *AssessSharpnessResponseBody) *AssessSharpnessResponse {
	s.Body = v
	return s
}

func (s *AssessSharpnessResponse) Validate() error {
	return dara.Validate(s)
}
