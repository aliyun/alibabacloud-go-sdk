// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceUpgradeStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDeviceUpgradeStatusResponseBody
	GetCode() *string
	SetData(v *GetDeviceUpgradeStatusResponseBodyData) *GetDeviceUpgradeStatusResponseBody
	GetData() *GetDeviceUpgradeStatusResponseBodyData
	SetMessage(v string) *GetDeviceUpgradeStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDeviceUpgradeStatusResponseBody
	GetRequestId() *string
}

type GetDeviceUpgradeStatusResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetDeviceUpgradeStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceUpgradeStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceUpgradeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceUpgradeStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDeviceUpgradeStatusResponseBody) GetData() *GetDeviceUpgradeStatusResponseBodyData {
	return s.Data
}

func (s *GetDeviceUpgradeStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDeviceUpgradeStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeviceUpgradeStatusResponseBody) SetCode(v string) *GetDeviceUpgradeStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBody) SetData(v *GetDeviceUpgradeStatusResponseBodyData) *GetDeviceUpgradeStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBody) SetMessage(v string) *GetDeviceUpgradeStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBody) SetRequestId(v string) *GetDeviceUpgradeStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDeviceUpgradeStatusResponseBodyData struct {
	AppOtaStatusDTOList []*GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList `json:"AppOtaStatusDTOList,omitempty" xml:"AppOtaStatusDTOList,omitempty" type:"Repeated"`
}

func (s GetDeviceUpgradeStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceUpgradeStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceUpgradeStatusResponseBodyData) GetAppOtaStatusDTOList() []*GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList {
	return s.AppOtaStatusDTOList
}

func (s *GetDeviceUpgradeStatusResponseBodyData) SetAppOtaStatusDTOList(v []*GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) *GetDeviceUpgradeStatusResponseBodyData {
	s.AppOtaStatusDTOList = v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList struct {
	BaseVersion   *string `json:"BaseVersion,omitempty" xml:"BaseVersion,omitempty"`
	ClientType    *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	ClientUid     *string `json:"ClientUid,omitempty" xml:"ClientUid,omitempty"`
	Note          *string `json:"Note,omitempty" xml:"Note,omitempty"`
	OsType        *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	Project       *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetVersion *string `json:"TargetVersion,omitempty" xml:"TargetVersion,omitempty"`
	TaskUid       *string `json:"TaskUid,omitempty" xml:"TaskUid,omitempty"`
}

func (s GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) GoString() string {
	return s.String()
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) GetBaseVersion() *string {
	return s.BaseVersion
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) GetClientType() *int32 {
	return s.ClientType
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) GetClientUid() *string {
	return s.ClientUid
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) GetNote() *string {
	return s.Note
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) GetOsType() *string {
	return s.OsType
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) GetProject() *string {
	return s.Project
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) GetStatus() *int32 {
	return s.Status
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) GetTaskUid() *string {
	return s.TaskUid
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) SetBaseVersion(v string) *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList {
	s.BaseVersion = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) SetClientType(v int32) *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList {
	s.ClientType = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) SetClientUid(v string) *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList {
	s.ClientUid = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) SetNote(v string) *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList {
	s.Note = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) SetOsType(v string) *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList {
	s.OsType = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) SetProject(v string) *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList {
	s.Project = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) SetStatus(v int32) *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList {
	s.Status = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) SetTargetVersion(v string) *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList {
	s.TargetVersion = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) SetTaskUid(v string) *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList {
	s.TaskUid = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponseBodyDataAppOtaStatusDTOList) Validate() error {
	return dara.Validate(s)
}
