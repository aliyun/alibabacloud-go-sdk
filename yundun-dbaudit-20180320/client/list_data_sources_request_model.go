// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v int32) *ListDataSourcesRequest
	GetDbId() *int32
	SetInstanceId(v string) *ListDataSourcesRequest
	GetInstanceId() *string
	SetLang(v string) *ListDataSourcesRequest
	GetLang() *string
	SetRegionId(v string) *ListDataSourcesRequest
	GetRegionId() *string
}

type ListDataSourcesRequest struct {
	// example:
	//
	// 1
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dbaudit-cn-78v1gc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListDataSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourcesRequest) GetDbId() *int32 {
	return s.DbId
}

func (s *ListDataSourcesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDataSourcesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDataSourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDataSourcesRequest) SetDbId(v int32) *ListDataSourcesRequest {
	s.DbId = &v
	return s
}

func (s *ListDataSourcesRequest) SetInstanceId(v string) *ListDataSourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDataSourcesRequest) SetLang(v string) *ListDataSourcesRequest {
	s.Lang = &v
	return s
}

func (s *ListDataSourcesRequest) SetRegionId(v string) *ListDataSourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListDataSourcesRequest) Validate() error {
	return dara.Validate(s)
}
