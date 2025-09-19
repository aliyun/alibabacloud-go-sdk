// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyncAssetTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSyncAssetTaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSyncAssetTaskListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSyncAssetTaskListResponseBody) *DescribeSyncAssetTaskListResponse
	GetBody() *DescribeSyncAssetTaskListResponseBody
}

type DescribeSyncAssetTaskListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSyncAssetTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSyncAssetTaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncAssetTaskListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSyncAssetTaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSyncAssetTaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSyncAssetTaskListResponse) GetBody() *DescribeSyncAssetTaskListResponseBody {
	return s.Body
}

func (s *DescribeSyncAssetTaskListResponse) SetHeaders(v map[string]*string) *DescribeSyncAssetTaskListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSyncAssetTaskListResponse) SetStatusCode(v int32) *DescribeSyncAssetTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSyncAssetTaskListResponse) SetBody(v *DescribeSyncAssetTaskListResponseBody) *DescribeSyncAssetTaskListResponse {
	s.Body = v
	return s
}

func (s *DescribeSyncAssetTaskListResponse) Validate() error {
	return dara.Validate(s)
}
