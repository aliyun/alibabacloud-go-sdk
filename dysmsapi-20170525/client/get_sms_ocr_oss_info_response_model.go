// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmsOcrOssInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSmsOcrOssInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSmsOcrOssInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetSmsOcrOssInfoResponseBody) *GetSmsOcrOssInfoResponse
	GetBody() *GetSmsOcrOssInfoResponseBody
}

type GetSmsOcrOssInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSmsOcrOssInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSmsOcrOssInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSmsOcrOssInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSmsOcrOssInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSmsOcrOssInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSmsOcrOssInfoResponse) GetBody() *GetSmsOcrOssInfoResponseBody {
	return s.Body
}

func (s *GetSmsOcrOssInfoResponse) SetHeaders(v map[string]*string) *GetSmsOcrOssInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSmsOcrOssInfoResponse) SetStatusCode(v int32) *GetSmsOcrOssInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSmsOcrOssInfoResponse) SetBody(v *GetSmsOcrOssInfoResponseBody) *GetSmsOcrOssInfoResponse {
	s.Body = v
	return s
}

func (s *GetSmsOcrOssInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
