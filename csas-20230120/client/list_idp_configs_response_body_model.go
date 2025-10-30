// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdpConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListIdpConfigsResponseBodyData) *ListIdpConfigsResponseBody
	GetData() *ListIdpConfigsResponseBodyData
	SetRequestId(v string) *ListIdpConfigsResponseBody
	GetRequestId() *string
}

type ListIdpConfigsResponseBody struct {
	Data *ListIdpConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// FD724DBC-CD76-5235-BF76-59C51B73296D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListIdpConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIdpConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIdpConfigsResponseBody) GetData() *ListIdpConfigsResponseBodyData {
	return s.Data
}

func (s *ListIdpConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIdpConfigsResponseBody) SetData(v *ListIdpConfigsResponseBodyData) *ListIdpConfigsResponseBody {
	s.Data = v
	return s
}

func (s *ListIdpConfigsResponseBody) SetRequestId(v string) *ListIdpConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIdpConfigsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListIdpConfigsResponseBodyData struct {
	DataList []*ListIdpConfigsResponseBodyDataDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListIdpConfigsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListIdpConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIdpConfigsResponseBodyData) GetDataList() []*ListIdpConfigsResponseBodyDataDataList {
	return s.DataList
}

func (s *ListIdpConfigsResponseBodyData) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *ListIdpConfigsResponseBodyData) SetDataList(v []*ListIdpConfigsResponseBodyDataDataList) *ListIdpConfigsResponseBodyData {
	s.DataList = v
	return s
}

func (s *ListIdpConfigsResponseBodyData) SetTotalNum(v int64) *ListIdpConfigsResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *ListIdpConfigsResponseBodyData) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIdpConfigsResponseBodyDataDataList struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 277
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// totp
	Mfa *string `json:"Mfa,omitempty" xml:"Mfa,omitempty"`
	// example:
	//
	// password
	MobileLoginType *string `json:"MobileLoginType,omitempty" xml:"MobileLoginType,omitempty"`
	// example:
	//
	// password
	MobileMfaConfigType *string `json:"MobileMfaConfigType,omitempty" xml:"MobileMfaConfigType,omitempty"`
	// example:
	//
	// 1482,1355
	MultiIdpInfo *string `json:"MultiIdpInfo,omitempty" xml:"MultiIdpInfo,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// password
	PcLoginType *string `json:"PcLoginType,omitempty" xml:"PcLoginType,omitempty"`
	// example:
	//
	// Disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// DingTalk
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2023-05-09T02:22:41.430Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListIdpConfigsResponseBodyDataDataList) String() string {
	return dara.Prettify(s)
}

func (s ListIdpConfigsResponseBodyDataDataList) GoString() string {
	return s.String()
}

func (s *ListIdpConfigsResponseBodyDataDataList) GetDescription() *string {
	return s.Description
}

func (s *ListIdpConfigsResponseBodyDataDataList) GetId() *string {
	return s.Id
}

func (s *ListIdpConfigsResponseBodyDataDataList) GetMfa() *string {
	return s.Mfa
}

func (s *ListIdpConfigsResponseBodyDataDataList) GetMobileLoginType() *string {
	return s.MobileLoginType
}

func (s *ListIdpConfigsResponseBodyDataDataList) GetMobileMfaConfigType() *string {
	return s.MobileMfaConfigType
}

func (s *ListIdpConfigsResponseBodyDataDataList) GetMultiIdpInfo() *string {
	return s.MultiIdpInfo
}

func (s *ListIdpConfigsResponseBodyDataDataList) GetName() *string {
	return s.Name
}

func (s *ListIdpConfigsResponseBodyDataDataList) GetPcLoginType() *string {
	return s.PcLoginType
}

func (s *ListIdpConfigsResponseBodyDataDataList) GetStatus() *string {
	return s.Status
}

func (s *ListIdpConfigsResponseBodyDataDataList) GetType() *string {
	return s.Type
}

func (s *ListIdpConfigsResponseBodyDataDataList) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListIdpConfigsResponseBodyDataDataList) SetDescription(v string) *ListIdpConfigsResponseBodyDataDataList {
	s.Description = &v
	return s
}

func (s *ListIdpConfigsResponseBodyDataDataList) SetId(v string) *ListIdpConfigsResponseBodyDataDataList {
	s.Id = &v
	return s
}

func (s *ListIdpConfigsResponseBodyDataDataList) SetMfa(v string) *ListIdpConfigsResponseBodyDataDataList {
	s.Mfa = &v
	return s
}

func (s *ListIdpConfigsResponseBodyDataDataList) SetMobileLoginType(v string) *ListIdpConfigsResponseBodyDataDataList {
	s.MobileLoginType = &v
	return s
}

func (s *ListIdpConfigsResponseBodyDataDataList) SetMobileMfaConfigType(v string) *ListIdpConfigsResponseBodyDataDataList {
	s.MobileMfaConfigType = &v
	return s
}

func (s *ListIdpConfigsResponseBodyDataDataList) SetMultiIdpInfo(v string) *ListIdpConfigsResponseBodyDataDataList {
	s.MultiIdpInfo = &v
	return s
}

func (s *ListIdpConfigsResponseBodyDataDataList) SetName(v string) *ListIdpConfigsResponseBodyDataDataList {
	s.Name = &v
	return s
}

func (s *ListIdpConfigsResponseBodyDataDataList) SetPcLoginType(v string) *ListIdpConfigsResponseBodyDataDataList {
	s.PcLoginType = &v
	return s
}

func (s *ListIdpConfigsResponseBodyDataDataList) SetStatus(v string) *ListIdpConfigsResponseBodyDataDataList {
	s.Status = &v
	return s
}

func (s *ListIdpConfigsResponseBodyDataDataList) SetType(v string) *ListIdpConfigsResponseBodyDataDataList {
	s.Type = &v
	return s
}

func (s *ListIdpConfigsResponseBodyDataDataList) SetUpdateTime(v string) *ListIdpConfigsResponseBodyDataDataList {
	s.UpdateTime = &v
	return s
}

func (s *ListIdpConfigsResponseBodyDataDataList) Validate() error {
	return dara.Validate(s)
}
