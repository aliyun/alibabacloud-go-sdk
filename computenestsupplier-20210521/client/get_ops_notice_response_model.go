// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpsNoticeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOpsNoticeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOpsNoticeResponse
	GetStatusCode() *int32
	SetBody(v *GetOpsNoticeResponseBody) *GetOpsNoticeResponse
	GetBody() *GetOpsNoticeResponseBody
}

type GetOpsNoticeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOpsNoticeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOpsNoticeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOpsNoticeResponse) GoString() string {
	return s.String()
}

func (s *GetOpsNoticeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOpsNoticeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOpsNoticeResponse) GetBody() *GetOpsNoticeResponseBody {
	return s.Body
}

func (s *GetOpsNoticeResponse) SetHeaders(v map[string]*string) *GetOpsNoticeResponse {
	s.Headers = v
	return s
}

func (s *GetOpsNoticeResponse) SetStatusCode(v int32) *GetOpsNoticeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpsNoticeResponse) SetBody(v *GetOpsNoticeResponseBody) *GetOpsNoticeResponse {
	s.Body = v
	return s
}

func (s *GetOpsNoticeResponse) Validate() error {
	return dara.Validate(s)
}
