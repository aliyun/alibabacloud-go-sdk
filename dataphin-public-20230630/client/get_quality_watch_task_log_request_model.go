// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityWatchTaskLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetQualityWatchTaskLogRequest
	GetOpTenantId() *int64
	SetWatchTaskId(v int64) *GetQualityWatchTaskLogRequest
	GetWatchTaskId() *int64
}

type GetQualityWatchTaskLogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11
	WatchTaskId *int64 `json:"WatchTaskId,omitempty" xml:"WatchTaskId,omitempty"`
}

func (s GetQualityWatchTaskLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchTaskLogRequest) GoString() string {
	return s.String()
}

func (s *GetQualityWatchTaskLogRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetQualityWatchTaskLogRequest) GetWatchTaskId() *int64 {
	return s.WatchTaskId
}

func (s *GetQualityWatchTaskLogRequest) SetOpTenantId(v int64) *GetQualityWatchTaskLogRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetQualityWatchTaskLogRequest) SetWatchTaskId(v int64) *GetQualityWatchTaskLogRequest {
	s.WatchTaskId = &v
	return s
}

func (s *GetQualityWatchTaskLogRequest) Validate() error {
	return dara.Validate(s)
}
