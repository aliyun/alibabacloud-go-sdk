// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsRegionListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OnsRegionListResponseBodyData) *OnsRegionListResponseBody
	GetData() *OnsRegionListResponseBodyData
	SetRequestId(v string) *OnsRegionListResponseBody
	GetRequestId() *string
}

type OnsRegionListResponseBody struct {
	Data *OnsRegionListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 72D14A84-45E5-4E01-A6DB-F63C4721****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsRegionListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsRegionListResponseBody) GoString() string {
	return s.String()
}

func (s *OnsRegionListResponseBody) GetData() *OnsRegionListResponseBodyData {
	return s.Data
}

func (s *OnsRegionListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsRegionListResponseBody) SetData(v *OnsRegionListResponseBodyData) *OnsRegionListResponseBody {
	s.Data = v
	return s
}

func (s *OnsRegionListResponseBody) SetRequestId(v string) *OnsRegionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsRegionListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsRegionListResponseBodyData struct {
	RegionDo []*OnsRegionListResponseBodyDataRegionDo `json:"RegionDo,omitempty" xml:"RegionDo,omitempty" type:"Repeated"`
}

func (s OnsRegionListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OnsRegionListResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsRegionListResponseBodyData) GetRegionDo() []*OnsRegionListResponseBodyDataRegionDo {
	return s.RegionDo
}

func (s *OnsRegionListResponseBodyData) SetRegionDo(v []*OnsRegionListResponseBodyDataRegionDo) *OnsRegionListResponseBodyData {
	s.RegionDo = v
	return s
}

func (s *OnsRegionListResponseBodyData) Validate() error {
	if s.RegionDo != nil {
		for _, item := range s.RegionDo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsRegionListResponseBodyDataRegionDo struct {
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	OnsRegionId *string `json:"OnsRegionId,omitempty" xml:"OnsRegionId,omitempty"`
	RegionName  *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	UpdateTime  *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s OnsRegionListResponseBodyDataRegionDo) String() string {
	return dara.Prettify(s)
}

func (s OnsRegionListResponseBodyDataRegionDo) GoString() string {
	return s.String()
}

func (s *OnsRegionListResponseBodyDataRegionDo) GetChannelName() *string {
	return s.ChannelName
}

func (s *OnsRegionListResponseBodyDataRegionDo) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *OnsRegionListResponseBodyDataRegionDo) GetId() *int64 {
	return s.Id
}

func (s *OnsRegionListResponseBodyDataRegionDo) GetOnsRegionId() *string {
	return s.OnsRegionId
}

func (s *OnsRegionListResponseBodyDataRegionDo) GetRegionName() *string {
	return s.RegionName
}

func (s *OnsRegionListResponseBodyDataRegionDo) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *OnsRegionListResponseBodyDataRegionDo) SetChannelName(v string) *OnsRegionListResponseBodyDataRegionDo {
	s.ChannelName = &v
	return s
}

func (s *OnsRegionListResponseBodyDataRegionDo) SetCreateTime(v int64) *OnsRegionListResponseBodyDataRegionDo {
	s.CreateTime = &v
	return s
}

func (s *OnsRegionListResponseBodyDataRegionDo) SetId(v int64) *OnsRegionListResponseBodyDataRegionDo {
	s.Id = &v
	return s
}

func (s *OnsRegionListResponseBodyDataRegionDo) SetOnsRegionId(v string) *OnsRegionListResponseBodyDataRegionDo {
	s.OnsRegionId = &v
	return s
}

func (s *OnsRegionListResponseBodyDataRegionDo) SetRegionName(v string) *OnsRegionListResponseBodyDataRegionDo {
	s.RegionName = &v
	return s
}

func (s *OnsRegionListResponseBodyDataRegionDo) SetUpdateTime(v int64) *OnsRegionListResponseBodyDataRegionDo {
	s.UpdateTime = &v
	return s
}

func (s *OnsRegionListResponseBodyDataRegionDo) Validate() error {
	return dara.Validate(s)
}
