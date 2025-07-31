// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssCheckStatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOssCheckStatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOssCheckStatResponse
	GetStatusCode() *int32
	SetBody(v *GetOssCheckStatResponseBody) *GetOssCheckStatResponse
	GetBody() *GetOssCheckStatResponseBody
}

type GetOssCheckStatResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOssCheckStatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOssCheckStatResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckStatResponse) GoString() string {
	return s.String()
}

func (s *GetOssCheckStatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOssCheckStatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOssCheckStatResponse) GetBody() *GetOssCheckStatResponseBody {
	return s.Body
}

func (s *GetOssCheckStatResponse) SetHeaders(v map[string]*string) *GetOssCheckStatResponse {
	s.Headers = v
	return s
}

func (s *GetOssCheckStatResponse) SetStatusCode(v int32) *GetOssCheckStatResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssCheckStatResponse) SetBody(v *GetOssCheckStatResponseBody) *GetOssCheckStatResponse {
	s.Body = v
	return s
}

func (s *GetOssCheckStatResponse) Validate() error {
	return dara.Validate(s)
}
