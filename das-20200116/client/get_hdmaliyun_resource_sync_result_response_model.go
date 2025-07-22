// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHDMAliyunResourceSyncResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHDMAliyunResourceSyncResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHDMAliyunResourceSyncResultResponse
	GetStatusCode() *int32
	SetBody(v *GetHDMAliyunResourceSyncResultResponseBody) *GetHDMAliyunResourceSyncResultResponse
	GetBody() *GetHDMAliyunResourceSyncResultResponseBody
}

type GetHDMAliyunResourceSyncResultResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHDMAliyunResourceSyncResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHDMAliyunResourceSyncResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHDMAliyunResourceSyncResultResponse) GoString() string {
	return s.String()
}

func (s *GetHDMAliyunResourceSyncResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHDMAliyunResourceSyncResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHDMAliyunResourceSyncResultResponse) GetBody() *GetHDMAliyunResourceSyncResultResponseBody {
	return s.Body
}

func (s *GetHDMAliyunResourceSyncResultResponse) SetHeaders(v map[string]*string) *GetHDMAliyunResourceSyncResultResponse {
	s.Headers = v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponse) SetStatusCode(v int32) *GetHDMAliyunResourceSyncResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponse) SetBody(v *GetHDMAliyunResourceSyncResultResponseBody) *GetHDMAliyunResourceSyncResultResponse {
	s.Body = v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponse) Validate() error {
	return dara.Validate(s)
}
