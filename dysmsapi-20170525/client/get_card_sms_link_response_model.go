// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCardSmsLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCardSmsLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCardSmsLinkResponse
	GetStatusCode() *int32
	SetBody(v *GetCardSmsLinkResponseBody) *GetCardSmsLinkResponse
	GetBody() *GetCardSmsLinkResponseBody
}

type GetCardSmsLinkResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCardSmsLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCardSmsLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCardSmsLinkResponse) GoString() string {
	return s.String()
}

func (s *GetCardSmsLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCardSmsLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCardSmsLinkResponse) GetBody() *GetCardSmsLinkResponseBody {
	return s.Body
}

func (s *GetCardSmsLinkResponse) SetHeaders(v map[string]*string) *GetCardSmsLinkResponse {
	s.Headers = v
	return s
}

func (s *GetCardSmsLinkResponse) SetStatusCode(v int32) *GetCardSmsLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCardSmsLinkResponse) SetBody(v *GetCardSmsLinkResponseBody) *GetCardSmsLinkResponse {
	s.Body = v
	return s
}

func (s *GetCardSmsLinkResponse) Validate() error {
	return dara.Validate(s)
}
