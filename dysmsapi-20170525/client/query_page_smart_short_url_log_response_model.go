// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPageSmartShortUrlLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPageSmartShortUrlLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPageSmartShortUrlLogResponse
	GetStatusCode() *int32
	SetBody(v *QueryPageSmartShortUrlLogResponseBody) *QueryPageSmartShortUrlLogResponse
	GetBody() *QueryPageSmartShortUrlLogResponseBody
}

type QueryPageSmartShortUrlLogResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPageSmartShortUrlLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPageSmartShortUrlLogResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPageSmartShortUrlLogResponse) GoString() string {
	return s.String()
}

func (s *QueryPageSmartShortUrlLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPageSmartShortUrlLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPageSmartShortUrlLogResponse) GetBody() *QueryPageSmartShortUrlLogResponseBody {
	return s.Body
}

func (s *QueryPageSmartShortUrlLogResponse) SetHeaders(v map[string]*string) *QueryPageSmartShortUrlLogResponse {
	s.Headers = v
	return s
}

func (s *QueryPageSmartShortUrlLogResponse) SetStatusCode(v int32) *QueryPageSmartShortUrlLogResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponse) SetBody(v *QueryPageSmartShortUrlLogResponseBody) *QueryPageSmartShortUrlLogResponse {
	s.Body = v
	return s
}

func (s *QueryPageSmartShortUrlLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
