// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyLastUsedInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccessKeyLastUsedInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccessKeyLastUsedInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetAccessKeyLastUsedInfoResponseBody) *GetAccessKeyLastUsedInfoResponse
	GetBody() *GetAccessKeyLastUsedInfoResponseBody
}

type GetAccessKeyLastUsedInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessKeyLastUsedInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessKeyLastUsedInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyLastUsedInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccessKeyLastUsedInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccessKeyLastUsedInfoResponse) GetBody() *GetAccessKeyLastUsedInfoResponseBody {
	return s.Body
}

func (s *GetAccessKeyLastUsedInfoResponse) SetHeaders(v map[string]*string) *GetAccessKeyLastUsedInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponse) SetStatusCode(v int32) *GetAccessKeyLastUsedInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponse) SetBody(v *GetAccessKeyLastUsedInfoResponseBody) *GetAccessKeyLastUsedInfoResponse {
	s.Body = v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponse) Validate() error {
	return dara.Validate(s)
}
