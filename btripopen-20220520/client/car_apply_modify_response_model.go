// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarApplyModifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CarApplyModifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CarApplyModifyResponse
	GetStatusCode() *int32
	SetBody(v *CarApplyModifyResponseBody) *CarApplyModifyResponse
	GetBody() *CarApplyModifyResponseBody
}

type CarApplyModifyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CarApplyModifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CarApplyModifyResponse) String() string {
	return dara.Prettify(s)
}

func (s CarApplyModifyResponse) GoString() string {
	return s.String()
}

func (s *CarApplyModifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CarApplyModifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CarApplyModifyResponse) GetBody() *CarApplyModifyResponseBody {
	return s.Body
}

func (s *CarApplyModifyResponse) SetHeaders(v map[string]*string) *CarApplyModifyResponse {
	s.Headers = v
	return s
}

func (s *CarApplyModifyResponse) SetStatusCode(v int32) *CarApplyModifyResponse {
	s.StatusCode = &v
	return s
}

func (s *CarApplyModifyResponse) SetBody(v *CarApplyModifyResponseBody) *CarApplyModifyResponse {
	s.Body = v
	return s
}

func (s *CarApplyModifyResponse) Validate() error {
	return dara.Validate(s)
}
