// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityWatchByObjectIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetQualityWatchByObjectIdRequest
	GetOpTenantId() *int64
	SetWatchObjectId(v string) *GetQualityWatchByObjectIdRequest
	GetWatchObjectId() *string
	SetWatchType(v string) *GetQualityWatchByObjectIdRequest
	GetWatchType() *string
}

type GetQualityWatchByObjectIdRequest struct {
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
	// cc
	WatchObjectId *string `json:"WatchObjectId,omitempty" xml:"WatchObjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TABLE
	WatchType *string `json:"WatchType,omitempty" xml:"WatchType,omitempty"`
}

func (s GetQualityWatchByObjectIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchByObjectIdRequest) GoString() string {
	return s.String()
}

func (s *GetQualityWatchByObjectIdRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetQualityWatchByObjectIdRequest) GetWatchObjectId() *string {
	return s.WatchObjectId
}

func (s *GetQualityWatchByObjectIdRequest) GetWatchType() *string {
	return s.WatchType
}

func (s *GetQualityWatchByObjectIdRequest) SetOpTenantId(v int64) *GetQualityWatchByObjectIdRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetQualityWatchByObjectIdRequest) SetWatchObjectId(v string) *GetQualityWatchByObjectIdRequest {
	s.WatchObjectId = &v
	return s
}

func (s *GetQualityWatchByObjectIdRequest) SetWatchType(v string) *GetQualityWatchByObjectIdRequest {
	s.WatchType = &v
	return s
}

func (s *GetQualityWatchByObjectIdRequest) Validate() error {
	return dara.Validate(s)
}
