// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbIds(v []*string) *ListDataSourceAttributeRequest
	GetDbIds() []*string
	SetInstanceId(v string) *ListDataSourceAttributeRequest
	GetInstanceId() *string
	SetLang(v string) *ListDataSourceAttributeRequest
	GetLang() *string
	SetRegionId(v string) *ListDataSourceAttributeRequest
	GetRegionId() *string
}

type ListDataSourceAttributeRequest struct {
	// example:
	//
	// 1
	DbIds []*string `json:"DbIds,omitempty" xml:"DbIds,omitempty" type:"Repeated"`
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

func (s ListDataSourceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourceAttributeRequest) GetDbIds() []*string {
	return s.DbIds
}

func (s *ListDataSourceAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDataSourceAttributeRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDataSourceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDataSourceAttributeRequest) SetDbIds(v []*string) *ListDataSourceAttributeRequest {
	s.DbIds = v
	return s
}

func (s *ListDataSourceAttributeRequest) SetInstanceId(v string) *ListDataSourceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDataSourceAttributeRequest) SetLang(v string) *ListDataSourceAttributeRequest {
	s.Lang = &v
	return s
}

func (s *ListDataSourceAttributeRequest) SetRegionId(v string) *ListDataSourceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ListDataSourceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
