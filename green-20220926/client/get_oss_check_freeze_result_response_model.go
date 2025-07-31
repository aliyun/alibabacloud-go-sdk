// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssCheckFreezeResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOssCheckFreezeResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOssCheckFreezeResultResponse
	GetStatusCode() *int32
	SetBody(v *GetOssCheckFreezeResultResponseBody) *GetOssCheckFreezeResultResponse
	GetBody() *GetOssCheckFreezeResultResponseBody
}

type GetOssCheckFreezeResultResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOssCheckFreezeResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOssCheckFreezeResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckFreezeResultResponse) GoString() string {
	return s.String()
}

func (s *GetOssCheckFreezeResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOssCheckFreezeResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOssCheckFreezeResultResponse) GetBody() *GetOssCheckFreezeResultResponseBody {
	return s.Body
}

func (s *GetOssCheckFreezeResultResponse) SetHeaders(v map[string]*string) *GetOssCheckFreezeResultResponse {
	s.Headers = v
	return s
}

func (s *GetOssCheckFreezeResultResponse) SetStatusCode(v int32) *GetOssCheckFreezeResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssCheckFreezeResultResponse) SetBody(v *GetOssCheckFreezeResultResponseBody) *GetOssCheckFreezeResultResponse {
	s.Body = v
	return s
}

func (s *GetOssCheckFreezeResultResponse) Validate() error {
	return dara.Validate(s)
}
