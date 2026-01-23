// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityAlertOfAllRuleScopeByWatchIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody
	GetMessage() *string
	SetQualityAlertInfo(v *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo) *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody
	GetQualityAlertInfo() *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo
	SetRequestId(v string) *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody
	GetSuccess() *bool
}

type GetQualityAlertOfAllRuleScopeByWatchIdResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message          *string                                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	QualityAlertInfo *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo `json:"QualityAlertInfo,omitempty" xml:"QualityAlertInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualityAlertOfAllRuleScopeByWatchIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualityAlertOfAllRuleScopeByWatchIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody) GetQualityAlertInfo() *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo {
	return s.QualityAlertInfo
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody) SetCode(v string) *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody) SetHttpStatusCode(v int32) *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody) SetMessage(v string) *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody) SetQualityAlertInfo(v *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo) *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody {
	s.QualityAlertInfo = v
	return s
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody) SetRequestId(v string) *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody) SetSuccess(v bool) *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody) Validate() error {
	if s.QualityAlertInfo != nil {
		if err := s.QualityAlertInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo struct {
	AlertDutyChannelList         []*string                                                                          `json:"AlertDutyChannelList,omitempty" xml:"AlertDutyChannelList,omitempty" type:"Repeated"`
	AlertDutyList                []*GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertDutyList `json:"AlertDutyList,omitempty" xml:"AlertDutyList,omitempty" type:"Repeated"`
	AlertQualityOwnerChannelList []*string                                                                          `json:"AlertQualityOwnerChannelList,omitempty" xml:"AlertQualityOwnerChannelList,omitempty" type:"Repeated"`
	AlertUserChannelList         []*string                                                                          `json:"AlertUserChannelList,omitempty" xml:"AlertUserChannelList,omitempty" type:"Repeated"`
	AlertUserList                []*GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertUserList `json:"AlertUserList,omitempty" xml:"AlertUserList,omitempty" type:"Repeated"`
	EnableAlertQualityOwner      *bool                                                                              `json:"EnableAlertQualityOwner,omitempty" xml:"EnableAlertQualityOwner,omitempty"`
	// example:
	//
	// 1
	WatchId *int64 `json:"WatchId,omitempty" xml:"WatchId,omitempty"`
}

func (s GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo) String() string {
	return dara.Prettify(s)
}

func (s GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo) GoString() string {
	return s.String()
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo) GetAlertDutyChannelList() []*string {
	return s.AlertDutyChannelList
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo) GetAlertDutyList() []*GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertDutyList {
	return s.AlertDutyList
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo) GetAlertQualityOwnerChannelList() []*string {
	return s.AlertQualityOwnerChannelList
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo) GetAlertUserChannelList() []*string {
	return s.AlertUserChannelList
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo) GetAlertUserList() []*GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertUserList {
	return s.AlertUserList
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo) GetEnableAlertQualityOwner() *bool {
	return s.EnableAlertQualityOwner
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo) GetWatchId() *int64 {
	return s.WatchId
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo) SetAlertDutyChannelList(v []*string) *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo {
	s.AlertDutyChannelList = v
	return s
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo) SetAlertDutyList(v []*GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertDutyList) *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo {
	s.AlertDutyList = v
	return s
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo) SetAlertQualityOwnerChannelList(v []*string) *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo {
	s.AlertQualityOwnerChannelList = v
	return s
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo) SetAlertUserChannelList(v []*string) *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo {
	s.AlertUserChannelList = v
	return s
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo) SetAlertUserList(v []*GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertUserList) *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo {
	s.AlertUserList = v
	return s
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo) SetEnableAlertQualityOwner(v bool) *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo {
	s.EnableAlertQualityOwner = &v
	return s
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo) SetWatchId(v int64) *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo {
	s.WatchId = &v
	return s
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfo) Validate() error {
	if s.AlertDutyList != nil {
		for _, item := range s.AlertDutyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AlertUserList != nil {
		for _, item := range s.AlertUserList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertDutyList struct {
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertDutyList) String() string {
	return dara.Prettify(s)
}

func (s GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertDutyList) GoString() string {
	return s.String()
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertDutyList) GetId() *string {
	return s.Id
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertDutyList) GetName() *string {
	return s.Name
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertDutyList) SetId(v string) *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertDutyList {
	s.Id = &v
	return s
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertDutyList) SetName(v string) *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertDutyList {
	s.Name = &v
	return s
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertDutyList) Validate() error {
	return dara.Validate(s)
}

type GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertUserList struct {
	// example:
	//
	// 30012011
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertUserList) String() string {
	return dara.Prettify(s)
}

func (s GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertUserList) GoString() string {
	return s.String()
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertUserList) GetId() *string {
	return s.Id
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertUserList) GetName() *string {
	return s.Name
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertUserList) SetId(v string) *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertUserList {
	s.Id = &v
	return s
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertUserList) SetName(v string) *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertUserList {
	s.Name = &v
	return s
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponseBodyQualityAlertInfoAlertUserList) Validate() error {
	return dara.Validate(s)
}
