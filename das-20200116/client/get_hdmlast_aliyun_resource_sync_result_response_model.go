// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHDMLastAliyunResourceSyncResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHDMLastAliyunResourceSyncResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHDMLastAliyunResourceSyncResultResponse
	GetStatusCode() *int32
	SetBody(v *GetHDMLastAliyunResourceSyncResultResponseBody) *GetHDMLastAliyunResourceSyncResultResponse
	GetBody() *GetHDMLastAliyunResourceSyncResultResponseBody
}

type GetHDMLastAliyunResourceSyncResultResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHDMLastAliyunResourceSyncResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHDMLastAliyunResourceSyncResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHDMLastAliyunResourceSyncResultResponse) GoString() string {
	return s.String()
}

func (s *GetHDMLastAliyunResourceSyncResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHDMLastAliyunResourceSyncResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHDMLastAliyunResourceSyncResultResponse) GetBody() *GetHDMLastAliyunResourceSyncResultResponseBody {
	return s.Body
}

func (s *GetHDMLastAliyunResourceSyncResultResponse) SetHeaders(v map[string]*string) *GetHDMLastAliyunResourceSyncResultResponse {
	s.Headers = v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponse) SetStatusCode(v int32) *GetHDMLastAliyunResourceSyncResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponse) SetBody(v *GetHDMLastAliyunResourceSyncResultResponseBody) *GetHDMLastAliyunResourceSyncResultResponse {
	s.Body = v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponse) Validate() error {
	return dara.Validate(s)
}
