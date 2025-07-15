// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamHistoryUserNumResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamHistoryUserNumResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamHistoryUserNumResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamHistoryUserNumResponseBody) *DescribeLiveStreamHistoryUserNumResponse
	GetBody() *DescribeLiveStreamHistoryUserNumResponseBody
}

type DescribeLiveStreamHistoryUserNumResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamHistoryUserNumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamHistoryUserNumResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamHistoryUserNumResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamHistoryUserNumResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamHistoryUserNumResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamHistoryUserNumResponse) GetBody() *DescribeLiveStreamHistoryUserNumResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamHistoryUserNumResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamHistoryUserNumResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamHistoryUserNumResponse) SetStatusCode(v int32) *DescribeLiveStreamHistoryUserNumResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamHistoryUserNumResponse) SetBody(v *DescribeLiveStreamHistoryUserNumResponseBody) *DescribeLiveStreamHistoryUserNumResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamHistoryUserNumResponse) Validate() error {
	return dara.Validate(s)
}
