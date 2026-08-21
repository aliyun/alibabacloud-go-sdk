// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppInfoList(v []*GetAppInfosResponseBodyAppInfoList) *GetAppInfosResponseBody
	GetAppInfoList() []*GetAppInfosResponseBodyAppInfoList
	SetCode(v string) *GetAppInfosResponseBody
	GetCode() *string
	SetNonExistAppIds(v []*string) *GetAppInfosResponseBody
	GetNonExistAppIds() []*string
	SetRequestId(v string) *GetAppInfosResponseBody
	GetRequestId() *string
}

type GetAppInfosResponseBody struct {
	// The list of application information.
	AppInfoList []*GetAppInfosResponseBodyAppInfoList `json:"AppInfoList,omitempty" xml:"AppInfoList,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of application IDs that do not exist.
	NonExistAppIds []*string `json:"NonExistAppIds,omitempty" xml:"NonExistAppIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4A13-4DC4-D7393642****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAppInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppInfosResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppInfosResponseBody) GetAppInfoList() []*GetAppInfosResponseBodyAppInfoList {
	return s.AppInfoList
}

func (s *GetAppInfosResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAppInfosResponseBody) GetNonExistAppIds() []*string {
	return s.NonExistAppIds
}

func (s *GetAppInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppInfosResponseBody) SetAppInfoList(v []*GetAppInfosResponseBodyAppInfoList) *GetAppInfosResponseBody {
	s.AppInfoList = v
	return s
}

func (s *GetAppInfosResponseBody) SetCode(v string) *GetAppInfosResponseBody {
	s.Code = &v
	return s
}

func (s *GetAppInfosResponseBody) SetNonExistAppIds(v []*string) *GetAppInfosResponseBody {
	s.NonExistAppIds = v
	return s
}

func (s *GetAppInfosResponseBody) SetRequestId(v string) *GetAppInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppInfosResponseBody) Validate() error {
	if s.AppInfoList != nil {
		for _, item := range s.AppInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAppInfosResponseBodyAppInfoList struct {
	// The application ID.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The creation time in UTC. The time follows the ISO 8601 standard in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-03-01T08:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The application description.
	//
	// example:
	//
	// my first app.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The modification time in UTC. The time follows the ISO 8601 standard in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-03-01T09:00:00Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekzko7fsuj****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The application status. Valid values:
	//
	// - **Normal**: normal.
	//
	// - **Disable**: disabled.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The application type. Valid values:
	//
	// - **System**: system default.
	//
	// - **Custom**: user-created.
	//
	// example:
	//
	// System
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAppInfosResponseBodyAppInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetAppInfosResponseBodyAppInfoList) GoString() string {
	return s.String()
}

func (s *GetAppInfosResponseBodyAppInfoList) GetAppId() *string {
	return s.AppId
}

func (s *GetAppInfosResponseBodyAppInfoList) GetAppName() *string {
	return s.AppName
}

func (s *GetAppInfosResponseBodyAppInfoList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetAppInfosResponseBodyAppInfoList) GetDescription() *string {
	return s.Description
}

func (s *GetAppInfosResponseBodyAppInfoList) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *GetAppInfosResponseBodyAppInfoList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetAppInfosResponseBodyAppInfoList) GetStatus() *string {
	return s.Status
}

func (s *GetAppInfosResponseBodyAppInfoList) GetType() *string {
	return s.Type
}

func (s *GetAppInfosResponseBodyAppInfoList) SetAppId(v string) *GetAppInfosResponseBodyAppInfoList {
	s.AppId = &v
	return s
}

func (s *GetAppInfosResponseBodyAppInfoList) SetAppName(v string) *GetAppInfosResponseBodyAppInfoList {
	s.AppName = &v
	return s
}

func (s *GetAppInfosResponseBodyAppInfoList) SetCreationTime(v string) *GetAppInfosResponseBodyAppInfoList {
	s.CreationTime = &v
	return s
}

func (s *GetAppInfosResponseBodyAppInfoList) SetDescription(v string) *GetAppInfosResponseBodyAppInfoList {
	s.Description = &v
	return s
}

func (s *GetAppInfosResponseBodyAppInfoList) SetModificationTime(v string) *GetAppInfosResponseBodyAppInfoList {
	s.ModificationTime = &v
	return s
}

func (s *GetAppInfosResponseBodyAppInfoList) SetResourceGroupId(v string) *GetAppInfosResponseBodyAppInfoList {
	s.ResourceGroupId = &v
	return s
}

func (s *GetAppInfosResponseBodyAppInfoList) SetStatus(v string) *GetAppInfosResponseBodyAppInfoList {
	s.Status = &v
	return s
}

func (s *GetAppInfosResponseBodyAppInfoList) SetType(v string) *GetAppInfosResponseBodyAppInfoList {
	s.Type = &v
	return s
}

func (s *GetAppInfosResponseBodyAppInfoList) Validate() error {
	return dara.Validate(s)
}
