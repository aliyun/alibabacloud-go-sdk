// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpsNoticeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOpsNoticeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOpsNoticeResponse
	GetStatusCode() *int32
	SetBody(v *CreateOpsNoticeResponseBody) *CreateOpsNoticeResponse
	GetBody() *CreateOpsNoticeResponseBody
}

type CreateOpsNoticeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOpsNoticeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOpsNoticeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOpsNoticeResponse) GoString() string {
	return s.String()
}

func (s *CreateOpsNoticeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOpsNoticeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOpsNoticeResponse) GetBody() *CreateOpsNoticeResponseBody {
	return s.Body
}

func (s *CreateOpsNoticeResponse) SetHeaders(v map[string]*string) *CreateOpsNoticeResponse {
	s.Headers = v
	return s
}

func (s *CreateOpsNoticeResponse) SetStatusCode(v int32) *CreateOpsNoticeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOpsNoticeResponse) SetBody(v *CreateOpsNoticeResponseBody) *CreateOpsNoticeResponse {
	s.Body = v
	return s
}

func (s *CreateOpsNoticeResponse) Validate() error {
	return dara.Validate(s)
}
