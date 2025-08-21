// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAuthCodeBindForExtResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckAuthCodeBindForExtResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckAuthCodeBindForExtResponse
	GetStatusCode() *int32
	SetBody(v *CheckAuthCodeBindForExtResponseBody) *CheckAuthCodeBindForExtResponse
	GetBody() *CheckAuthCodeBindForExtResponseBody
}

type CheckAuthCodeBindForExtResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAuthCodeBindForExtResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAuthCodeBindForExtResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckAuthCodeBindForExtResponse) GoString() string {
	return s.String()
}

func (s *CheckAuthCodeBindForExtResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckAuthCodeBindForExtResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckAuthCodeBindForExtResponse) GetBody() *CheckAuthCodeBindForExtResponseBody {
	return s.Body
}

func (s *CheckAuthCodeBindForExtResponse) SetHeaders(v map[string]*string) *CheckAuthCodeBindForExtResponse {
	s.Headers = v
	return s
}

func (s *CheckAuthCodeBindForExtResponse) SetStatusCode(v int32) *CheckAuthCodeBindForExtResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAuthCodeBindForExtResponse) SetBody(v *CheckAuthCodeBindForExtResponseBody) *CheckAuthCodeBindForExtResponse {
	s.Body = v
	return s
}

func (s *CheckAuthCodeBindForExtResponse) Validate() error {
	return dara.Validate(s)
}
