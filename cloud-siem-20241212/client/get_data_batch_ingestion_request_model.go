// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataBatchIngestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetDataBatchIngestionRequest
	GetLang() *string
	SetRegionId(v string) *GetDataBatchIngestionRequest
	GetRegionId() *string
	SetRoleFor(v int64) *GetDataBatchIngestionRequest
	GetRoleFor() *int64
}

type GetDataBatchIngestionRequest struct {
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s GetDataBatchIngestionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataBatchIngestionRequest) GoString() string {
	return s.String()
}

func (s *GetDataBatchIngestionRequest) GetLang() *string {
	return s.Lang
}

func (s *GetDataBatchIngestionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDataBatchIngestionRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *GetDataBatchIngestionRequest) SetLang(v string) *GetDataBatchIngestionRequest {
	s.Lang = &v
	return s
}

func (s *GetDataBatchIngestionRequest) SetRegionId(v string) *GetDataBatchIngestionRequest {
	s.RegionId = &v
	return s
}

func (s *GetDataBatchIngestionRequest) SetRoleFor(v int64) *GetDataBatchIngestionRequest {
	s.RoleFor = &v
	return s
}

func (s *GetDataBatchIngestionRequest) Validate() error {
	return dara.Validate(s)
}
