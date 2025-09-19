// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCheckInstanceResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyCheckInstanceResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyCheckInstanceResultResponse
	GetStatusCode() *int32
	SetBody(v *VerifyCheckInstanceResultResponseBody) *VerifyCheckInstanceResultResponse
	GetBody() *VerifyCheckInstanceResultResponseBody
}

type VerifyCheckInstanceResultResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyCheckInstanceResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyCheckInstanceResultResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyCheckInstanceResultResponse) GoString() string {
	return s.String()
}

func (s *VerifyCheckInstanceResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyCheckInstanceResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyCheckInstanceResultResponse) GetBody() *VerifyCheckInstanceResultResponseBody {
	return s.Body
}

func (s *VerifyCheckInstanceResultResponse) SetHeaders(v map[string]*string) *VerifyCheckInstanceResultResponse {
	s.Headers = v
	return s
}

func (s *VerifyCheckInstanceResultResponse) SetStatusCode(v int32) *VerifyCheckInstanceResultResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyCheckInstanceResultResponse) SetBody(v *VerifyCheckInstanceResultResponseBody) *VerifyCheckInstanceResultResponse {
	s.Body = v
	return s
}

func (s *VerifyCheckInstanceResultResponse) Validate() error {
	return dara.Validate(s)
}
