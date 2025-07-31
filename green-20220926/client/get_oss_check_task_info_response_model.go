// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssCheckTaskInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOssCheckTaskInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOssCheckTaskInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetOssCheckTaskInfoResponseBody) *GetOssCheckTaskInfoResponse
	GetBody() *GetOssCheckTaskInfoResponseBody
}

type GetOssCheckTaskInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOssCheckTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOssCheckTaskInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *GetOssCheckTaskInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOssCheckTaskInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOssCheckTaskInfoResponse) GetBody() *GetOssCheckTaskInfoResponseBody {
	return s.Body
}

func (s *GetOssCheckTaskInfoResponse) SetHeaders(v map[string]*string) *GetOssCheckTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *GetOssCheckTaskInfoResponse) SetStatusCode(v int32) *GetOssCheckTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssCheckTaskInfoResponse) SetBody(v *GetOssCheckTaskInfoResponseBody) *GetOssCheckTaskInfoResponse {
	s.Body = v
	return s
}

func (s *GetOssCheckTaskInfoResponse) Validate() error {
	return dara.Validate(s)
}
