// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityWatchTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetQualityWatchTaskRequest
	GetOpTenantId() *int64
	SetWatchTaskId(v int64) *GetQualityWatchTaskRequest
	GetWatchTaskId() *int64
}

type GetQualityWatchTaskRequest struct {
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

func (s GetQualityWatchTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchTaskRequest) GoString() string {
	return s.String()
}

func (s *GetQualityWatchTaskRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetQualityWatchTaskRequest) GetWatchTaskId() *int64 {
	return s.WatchTaskId
}

func (s *GetQualityWatchTaskRequest) SetOpTenantId(v int64) *GetQualityWatchTaskRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetQualityWatchTaskRequest) SetWatchTaskId(v int64) *GetQualityWatchTaskRequest {
	s.WatchTaskId = &v
	return s
}

func (s *GetQualityWatchTaskRequest) Validate() error {
	return dara.Validate(s)
}
