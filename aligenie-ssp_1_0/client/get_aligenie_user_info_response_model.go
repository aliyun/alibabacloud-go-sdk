// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAligenieUserInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAligenieUserInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAligenieUserInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetAligenieUserInfoResponseBody) *GetAligenieUserInfoResponse
	GetBody() *GetAligenieUserInfoResponseBody
}

type GetAligenieUserInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAligenieUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAligenieUserInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAligenieUserInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAligenieUserInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAligenieUserInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAligenieUserInfoResponse) GetBody() *GetAligenieUserInfoResponseBody {
	return s.Body
}

func (s *GetAligenieUserInfoResponse) SetHeaders(v map[string]*string) *GetAligenieUserInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAligenieUserInfoResponse) SetStatusCode(v int32) *GetAligenieUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAligenieUserInfoResponse) SetBody(v *GetAligenieUserInfoResponseBody) *GetAligenieUserInfoResponse {
	s.Body = v
	return s
}

func (s *GetAligenieUserInfoResponse) Validate() error {
	return dara.Validate(s)
}
