// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncClusterNodePoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncClusterNodePoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncClusterNodePoolResponse
	GetStatusCode() *int32
	SetBody(v *SyncClusterNodePoolResponseBody) *SyncClusterNodePoolResponse
	GetBody() *SyncClusterNodePoolResponseBody
}

type SyncClusterNodePoolResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncClusterNodePoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncClusterNodePoolResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncClusterNodePoolResponse) GoString() string {
	return s.String()
}

func (s *SyncClusterNodePoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncClusterNodePoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncClusterNodePoolResponse) GetBody() *SyncClusterNodePoolResponseBody {
	return s.Body
}

func (s *SyncClusterNodePoolResponse) SetHeaders(v map[string]*string) *SyncClusterNodePoolResponse {
	s.Headers = v
	return s
}

func (s *SyncClusterNodePoolResponse) SetStatusCode(v int32) *SyncClusterNodePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncClusterNodePoolResponse) SetBody(v *SyncClusterNodePoolResponseBody) *SyncClusterNodePoolResponse {
	s.Body = v
	return s
}

func (s *SyncClusterNodePoolResponse) Validate() error {
	return dara.Validate(s)
}
