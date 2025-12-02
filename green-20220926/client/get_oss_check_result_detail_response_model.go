// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssCheckResultDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOssCheckResultDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOssCheckResultDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetOssCheckResultDetailResponseBody) *GetOssCheckResultDetailResponse
	GetBody() *GetOssCheckResultDetailResponseBody
}

type GetOssCheckResultDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOssCheckResultDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOssCheckResultDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckResultDetailResponse) GoString() string {
	return s.String()
}

func (s *GetOssCheckResultDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOssCheckResultDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOssCheckResultDetailResponse) GetBody() *GetOssCheckResultDetailResponseBody {
	return s.Body
}

func (s *GetOssCheckResultDetailResponse) SetHeaders(v map[string]*string) *GetOssCheckResultDetailResponse {
	s.Headers = v
	return s
}

func (s *GetOssCheckResultDetailResponse) SetStatusCode(v int32) *GetOssCheckResultDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssCheckResultDetailResponse) SetBody(v *GetOssCheckResultDetailResponseBody) *GetOssCheckResultDetailResponse {
	s.Body = v
	return s
}

func (s *GetOssCheckResultDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
