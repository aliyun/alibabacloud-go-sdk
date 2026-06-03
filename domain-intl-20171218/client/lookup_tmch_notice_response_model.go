// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLookupTmchNoticeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LookupTmchNoticeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LookupTmchNoticeResponse
	GetStatusCode() *int32
	SetBody(v *LookupTmchNoticeResponseBody) *LookupTmchNoticeResponse
	GetBody() *LookupTmchNoticeResponseBody
}

type LookupTmchNoticeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LookupTmchNoticeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LookupTmchNoticeResponse) String() string {
	return dara.Prettify(s)
}

func (s LookupTmchNoticeResponse) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LookupTmchNoticeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LookupTmchNoticeResponse) GetBody() *LookupTmchNoticeResponseBody {
	return s.Body
}

func (s *LookupTmchNoticeResponse) SetHeaders(v map[string]*string) *LookupTmchNoticeResponse {
	s.Headers = v
	return s
}

func (s *LookupTmchNoticeResponse) SetStatusCode(v int32) *LookupTmchNoticeResponse {
	s.StatusCode = &v
	return s
}

func (s *LookupTmchNoticeResponse) SetBody(v *LookupTmchNoticeResponseBody) *LookupTmchNoticeResponse {
	s.Body = v
	return s
}

func (s *LookupTmchNoticeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
