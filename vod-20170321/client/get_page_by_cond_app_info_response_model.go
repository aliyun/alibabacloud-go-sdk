// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPageByCondAppInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPageByCondAppInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPageByCondAppInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetPageByCondAppInfoResponseBody) *GetPageByCondAppInfoResponse
	GetBody() *GetPageByCondAppInfoResponseBody
}

type GetPageByCondAppInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPageByCondAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPageByCondAppInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPageByCondAppInfoResponse) GoString() string {
	return s.String()
}

func (s *GetPageByCondAppInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPageByCondAppInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPageByCondAppInfoResponse) GetBody() *GetPageByCondAppInfoResponseBody {
	return s.Body
}

func (s *GetPageByCondAppInfoResponse) SetHeaders(v map[string]*string) *GetPageByCondAppInfoResponse {
	s.Headers = v
	return s
}

func (s *GetPageByCondAppInfoResponse) SetStatusCode(v int32) *GetPageByCondAppInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPageByCondAppInfoResponse) SetBody(v *GetPageByCondAppInfoResponseBody) *GetPageByCondAppInfoResponse {
	s.Body = v
	return s
}

func (s *GetPageByCondAppInfoResponse) Validate() error {
	return dara.Validate(s)
}
