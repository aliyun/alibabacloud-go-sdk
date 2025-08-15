// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyLastUsedIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccessKeyLastUsedIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccessKeyLastUsedIpsResponse
	GetStatusCode() *int32
	SetBody(v *GetAccessKeyLastUsedIpsResponseBody) *GetAccessKeyLastUsedIpsResponse
	GetBody() *GetAccessKeyLastUsedIpsResponseBody
}

type GetAccessKeyLastUsedIpsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessKeyLastUsedIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessKeyLastUsedIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyLastUsedIpsResponse) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccessKeyLastUsedIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccessKeyLastUsedIpsResponse) GetBody() *GetAccessKeyLastUsedIpsResponseBody {
	return s.Body
}

func (s *GetAccessKeyLastUsedIpsResponse) SetHeaders(v map[string]*string) *GetAccessKeyLastUsedIpsResponse {
	s.Headers = v
	return s
}

func (s *GetAccessKeyLastUsedIpsResponse) SetStatusCode(v int32) *GetAccessKeyLastUsedIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessKeyLastUsedIpsResponse) SetBody(v *GetAccessKeyLastUsedIpsResponseBody) *GetAccessKeyLastUsedIpsResponse {
	s.Body = v
	return s
}

func (s *GetAccessKeyLastUsedIpsResponse) Validate() error {
	return dara.Validate(s)
}
