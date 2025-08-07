// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshDBClusterStorageUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshDBClusterStorageUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshDBClusterStorageUsageResponse
	GetStatusCode() *int32
	SetBody(v *RefreshDBClusterStorageUsageResponseBody) *RefreshDBClusterStorageUsageResponse
	GetBody() *RefreshDBClusterStorageUsageResponseBody
}

type RefreshDBClusterStorageUsageResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshDBClusterStorageUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshDBClusterStorageUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshDBClusterStorageUsageResponse) GoString() string {
	return s.String()
}

func (s *RefreshDBClusterStorageUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshDBClusterStorageUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshDBClusterStorageUsageResponse) GetBody() *RefreshDBClusterStorageUsageResponseBody {
	return s.Body
}

func (s *RefreshDBClusterStorageUsageResponse) SetHeaders(v map[string]*string) *RefreshDBClusterStorageUsageResponse {
	s.Headers = v
	return s
}

func (s *RefreshDBClusterStorageUsageResponse) SetStatusCode(v int32) *RefreshDBClusterStorageUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshDBClusterStorageUsageResponse) SetBody(v *RefreshDBClusterStorageUsageResponseBody) *RefreshDBClusterStorageUsageResponse {
	s.Body = v
	return s
}

func (s *RefreshDBClusterStorageUsageResponse) Validate() error {
	return dara.Validate(s)
}
