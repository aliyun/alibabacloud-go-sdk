// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssCheckStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOssCheckStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOssCheckStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetOssCheckStatusResponseBody) *GetOssCheckStatusResponse
	GetBody() *GetOssCheckStatusResponseBody
}

type GetOssCheckStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOssCheckStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOssCheckStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckStatusResponse) GoString() string {
	return s.String()
}

func (s *GetOssCheckStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOssCheckStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOssCheckStatusResponse) GetBody() *GetOssCheckStatusResponseBody {
	return s.Body
}

func (s *GetOssCheckStatusResponse) SetHeaders(v map[string]*string) *GetOssCheckStatusResponse {
	s.Headers = v
	return s
}

func (s *GetOssCheckStatusResponse) SetStatusCode(v int32) *GetOssCheckStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssCheckStatusResponse) SetBody(v *GetOssCheckStatusResponseBody) *GetOssCheckStatusResponse {
	s.Body = v
	return s
}

func (s *GetOssCheckStatusResponse) Validate() error {
	return dara.Validate(s)
}
