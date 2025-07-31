// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBackupBucketsListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBackupBucketsListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBackupBucketsListResponse
	GetStatusCode() *int32
	SetBody(v *GetBackupBucketsListResponseBody) *GetBackupBucketsListResponse
	GetBody() *GetBackupBucketsListResponseBody
}

type GetBackupBucketsListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBackupBucketsListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBackupBucketsListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBackupBucketsListResponse) GoString() string {
	return s.String()
}

func (s *GetBackupBucketsListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBackupBucketsListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBackupBucketsListResponse) GetBody() *GetBackupBucketsListResponseBody {
	return s.Body
}

func (s *GetBackupBucketsListResponse) SetHeaders(v map[string]*string) *GetBackupBucketsListResponse {
	s.Headers = v
	return s
}

func (s *GetBackupBucketsListResponse) SetStatusCode(v int32) *GetBackupBucketsListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBackupBucketsListResponse) SetBody(v *GetBackupBucketsListResponseBody) *GetBackupBucketsListResponse {
	s.Body = v
	return s
}

func (s *GetBackupBucketsListResponse) Validate() error {
	return dara.Validate(s)
}
