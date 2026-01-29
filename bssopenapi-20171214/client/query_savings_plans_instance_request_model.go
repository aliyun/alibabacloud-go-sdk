// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySavingsPlansInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *QuerySavingsPlansInstanceRequest
	GetCommodityCode() *string
	SetEndTime(v string) *QuerySavingsPlansInstanceRequest
	GetEndTime() *string
	SetInstanceId(v string) *QuerySavingsPlansInstanceRequest
	GetInstanceId() *string
	SetLocale(v string) *QuerySavingsPlansInstanceRequest
	GetLocale() *string
	SetPageNum(v int32) *QuerySavingsPlansInstanceRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QuerySavingsPlansInstanceRequest
	GetPageSize() *int32
	SetStartTime(v string) *QuerySavingsPlansInstanceRequest
	GetStartTime() *string
	SetStatus(v string) *QuerySavingsPlansInstanceRequest
	GetStatus() *string
	SetTag(v []*QuerySavingsPlansInstanceRequestTag) *QuerySavingsPlansInstanceRequest
	GetTag() []*QuerySavingsPlansInstanceRequestTag
}

type QuerySavingsPlansInstanceRequest struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The end of the time range to query. Specify the time in the format of yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2021-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the savings plan instance.
	//
	// example:
	//
	// spn-xxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the return data. Valid values:
	//
	// 	- ZH: Chinese
	//
	// 	- EN: English
	//
	// example:
	//
	// ZH
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Specify the time in the format of yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2020-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the Instance.
	//
	// 	- NORMAL
	//
	// 	- RELEASE
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tag []*QuerySavingsPlansInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s QuerySavingsPlansInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySavingsPlansInstanceRequest) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansInstanceRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *QuerySavingsPlansInstanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QuerySavingsPlansInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QuerySavingsPlansInstanceRequest) GetLocale() *string {
	return s.Locale
}

func (s *QuerySavingsPlansInstanceRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QuerySavingsPlansInstanceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySavingsPlansInstanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QuerySavingsPlansInstanceRequest) GetStatus() *string {
	return s.Status
}

func (s *QuerySavingsPlansInstanceRequest) GetTag() []*QuerySavingsPlansInstanceRequestTag {
	return s.Tag
}

func (s *QuerySavingsPlansInstanceRequest) SetCommodityCode(v string) *QuerySavingsPlansInstanceRequest {
	s.CommodityCode = &v
	return s
}

func (s *QuerySavingsPlansInstanceRequest) SetEndTime(v string) *QuerySavingsPlansInstanceRequest {
	s.EndTime = &v
	return s
}

func (s *QuerySavingsPlansInstanceRequest) SetInstanceId(v string) *QuerySavingsPlansInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *QuerySavingsPlansInstanceRequest) SetLocale(v string) *QuerySavingsPlansInstanceRequest {
	s.Locale = &v
	return s
}

func (s *QuerySavingsPlansInstanceRequest) SetPageNum(v int32) *QuerySavingsPlansInstanceRequest {
	s.PageNum = &v
	return s
}

func (s *QuerySavingsPlansInstanceRequest) SetPageSize(v int32) *QuerySavingsPlansInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySavingsPlansInstanceRequest) SetStartTime(v string) *QuerySavingsPlansInstanceRequest {
	s.StartTime = &v
	return s
}

func (s *QuerySavingsPlansInstanceRequest) SetStatus(v string) *QuerySavingsPlansInstanceRequest {
	s.Status = &v
	return s
}

func (s *QuerySavingsPlansInstanceRequest) SetTag(v []*QuerySavingsPlansInstanceRequestTag) *QuerySavingsPlansInstanceRequest {
	s.Tag = v
	return s
}

func (s *QuerySavingsPlansInstanceRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySavingsPlansInstanceRequestTag struct {
	// The key of the tag to query.
	//
	// example:
	//
	// ecs
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag to query.
	//
	// example:
	//
	// 001
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QuerySavingsPlansInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s QuerySavingsPlansInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *QuerySavingsPlansInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *QuerySavingsPlansInstanceRequestTag) SetKey(v string) *QuerySavingsPlansInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *QuerySavingsPlansInstanceRequestTag) SetValue(v string) *QuerySavingsPlansInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *QuerySavingsPlansInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}
