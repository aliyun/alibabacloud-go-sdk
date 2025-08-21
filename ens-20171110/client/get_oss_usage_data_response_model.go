// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssUsageDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOssUsageDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOssUsageDataResponse
	GetStatusCode() *int32
	SetBody(v *GetOssUsageDataResponseBody) *GetOssUsageDataResponse
	GetBody() *GetOssUsageDataResponseBody
}

type GetOssUsageDataResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOssUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOssUsageDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOssUsageDataResponse) GoString() string {
	return s.String()
}

func (s *GetOssUsageDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOssUsageDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOssUsageDataResponse) GetBody() *GetOssUsageDataResponseBody {
	return s.Body
}

func (s *GetOssUsageDataResponse) SetHeaders(v map[string]*string) *GetOssUsageDataResponse {
	s.Headers = v
	return s
}

func (s *GetOssUsageDataResponse) SetStatusCode(v int32) *GetOssUsageDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssUsageDataResponse) SetBody(v *GetOssUsageDataResponseBody) *GetOssUsageDataResponse {
	s.Body = v
	return s
}

func (s *GetOssUsageDataResponse) Validate() error {
	return dara.Validate(s)
}
