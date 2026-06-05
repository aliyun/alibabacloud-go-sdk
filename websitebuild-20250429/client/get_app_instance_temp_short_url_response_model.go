// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppInstanceTempShortUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppInstanceTempShortUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppInstanceTempShortUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetAppInstanceTempShortUrlResponseBody) *GetAppInstanceTempShortUrlResponse
	GetBody() *GetAppInstanceTempShortUrlResponseBody
}

type GetAppInstanceTempShortUrlResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppInstanceTempShortUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppInstanceTempShortUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceTempShortUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAppInstanceTempShortUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppInstanceTempShortUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppInstanceTempShortUrlResponse) GetBody() *GetAppInstanceTempShortUrlResponseBody {
	return s.Body
}

func (s *GetAppInstanceTempShortUrlResponse) SetHeaders(v map[string]*string) *GetAppInstanceTempShortUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAppInstanceTempShortUrlResponse) SetStatusCode(v int32) *GetAppInstanceTempShortUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppInstanceTempShortUrlResponse) SetBody(v *GetAppInstanceTempShortUrlResponseBody) *GetAppInstanceTempShortUrlResponse {
	s.Body = v
	return s
}

func (s *GetAppInstanceTempShortUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
