// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTraceAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteTraceAppRequest
	GetAppId() *string
	SetDeleteReason(v *DeleteTraceAppRequestDeleteReason) *DeleteTraceAppRequest
	GetDeleteReason() *DeleteTraceAppRequestDeleteReason
	SetPid(v string) *DeleteTraceAppRequest
	GetPid() *string
	SetRegionId(v string) *DeleteTraceAppRequest
	GetRegionId() *string
	SetType(v string) *DeleteTraceAppRequest
	GetType() *string
}

type DeleteTraceAppRequest struct {
	// The ID of the application that you want to delete. You can call the SearchTraceAppByName operation to query the application ID. For more information, see [SearchTraceAppByName](https://help.aliyun.com/document_detail/130676.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 5406**
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The reason(s) to delete application.
	DeleteReason *DeleteTraceAppRequestDeleteReason `json:"DeleteReason,omitempty" xml:"DeleteReason,omitempty" type:"Struct"`
	// The PID of the application. For more information about how to query the PID, see [QueryMetricByPage](https://www.alibabacloud.com/help/zh/doc-detail/186100.htm?spm=a2cdw.13409063.0.0.7a72281f0bkTfx#title-imy-7gj-qhr).
	//
	// This parameter is required.
	//
	// example:
	//
	// 9w0sc5gxxz@edcsd447c2f****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The ID of the region in which the application is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the application that you want to delete. You can call the SearchTraceAppByName operation to query the application type. For more information, see [SearchTraceAppByName](https://help.aliyun.com/document_detail/130676.html). Valid values:
	//
	// 	- `TRACE`: Application Monitoring
	//
	// 	- `RETCODE`: frontend monitoring
	//
	// This parameter is required.
	//
	// example:
	//
	// TRACE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeleteTraceAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTraceAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteTraceAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteTraceAppRequest) GetDeleteReason() *DeleteTraceAppRequestDeleteReason {
	return s.DeleteReason
}

func (s *DeleteTraceAppRequest) GetPid() *string {
	return s.Pid
}

func (s *DeleteTraceAppRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteTraceAppRequest) GetType() *string {
	return s.Type
}

func (s *DeleteTraceAppRequest) SetAppId(v string) *DeleteTraceAppRequest {
	s.AppId = &v
	return s
}

func (s *DeleteTraceAppRequest) SetDeleteReason(v *DeleteTraceAppRequestDeleteReason) *DeleteTraceAppRequest {
	s.DeleteReason = v
	return s
}

func (s *DeleteTraceAppRequest) SetPid(v string) *DeleteTraceAppRequest {
	s.Pid = &v
	return s
}

func (s *DeleteTraceAppRequest) SetRegionId(v string) *DeleteTraceAppRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTraceAppRequest) SetType(v string) *DeleteTraceAppRequest {
	s.Type = &v
	return s
}

func (s *DeleteTraceAppRequest) Validate() error {
	if s.DeleteReason != nil {
		if err := s.DeleteReason.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteTraceAppRequestDeleteReason struct {
	// Reasons Ids.
	ReasonIds []*DeleteTraceAppRequestDeleteReasonReasonIds `json:"ReasonIds,omitempty" xml:"ReasonIds,omitempty" type:"Repeated"`
	// Additional remarks when none of the reasons for removal provided are met.
	//
	// example:
	//
	// The business scenario cannot be satisfied.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s DeleteTraceAppRequestDeleteReason) String() string {
	return dara.Prettify(s)
}

func (s DeleteTraceAppRequestDeleteReason) GoString() string {
	return s.String()
}

func (s *DeleteTraceAppRequestDeleteReason) GetReasonIds() []*DeleteTraceAppRequestDeleteReasonReasonIds {
	return s.ReasonIds
}

func (s *DeleteTraceAppRequestDeleteReason) GetRemark() *string {
	return s.Remark
}

func (s *DeleteTraceAppRequestDeleteReason) SetReasonIds(v []*DeleteTraceAppRequestDeleteReasonReasonIds) *DeleteTraceAppRequestDeleteReason {
	s.ReasonIds = v
	return s
}

func (s *DeleteTraceAppRequestDeleteReason) SetRemark(v string) *DeleteTraceAppRequestDeleteReason {
	s.Remark = &v
	return s
}

func (s *DeleteTraceAppRequestDeleteReason) Validate() error {
	if s.ReasonIds != nil {
		for _, item := range s.ReasonIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteTraceAppRequestDeleteReasonReasonIds struct {
	// The ID of the reason for deletion.
	//
	// example:
	//
	// 0
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// A description of the reason for removal.
	//
	// example:
	//
	// The function is not perfect, and the root cause of the problem cannot be located.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteTraceAppRequestDeleteReasonReasonIds) String() string {
	return dara.Prettify(s)
}

func (s DeleteTraceAppRequestDeleteReasonReasonIds) GoString() string {
	return s.String()
}

func (s *DeleteTraceAppRequestDeleteReasonReasonIds) GetId() *int32 {
	return s.Id
}

func (s *DeleteTraceAppRequestDeleteReasonReasonIds) GetName() *string {
	return s.Name
}

func (s *DeleteTraceAppRequestDeleteReasonReasonIds) SetId(v int32) *DeleteTraceAppRequestDeleteReasonReasonIds {
	s.Id = &v
	return s
}

func (s *DeleteTraceAppRequestDeleteReasonReasonIds) SetName(v string) *DeleteTraceAppRequestDeleteReasonReasonIds {
	s.Name = &v
	return s
}

func (s *DeleteTraceAppRequestDeleteReasonReasonIds) Validate() error {
	return dara.Validate(s)
}
