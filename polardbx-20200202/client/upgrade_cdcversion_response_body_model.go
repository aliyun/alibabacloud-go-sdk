// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeCDCVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpgradeCDCVersionResponseBodyData) *UpgradeCDCVersionResponseBody
	GetData() *UpgradeCDCVersionResponseBodyData
	SetMessage(v string) *UpgradeCDCVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpgradeCDCVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpgradeCDCVersionResponseBody
	GetSuccess() *bool
}

type UpgradeCDCVersionResponseBody struct {
	Data *UpgradeCDCVersionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CB4307F5-3D04-51E8-ABAD-49E0B3F962FC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpgradeCDCVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeCDCVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeCDCVersionResponseBody) GetData() *UpgradeCDCVersionResponseBodyData {
	return s.Data
}

func (s *UpgradeCDCVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpgradeCDCVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeCDCVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpgradeCDCVersionResponseBody) SetData(v *UpgradeCDCVersionResponseBodyData) *UpgradeCDCVersionResponseBody {
	s.Data = v
	return s
}

func (s *UpgradeCDCVersionResponseBody) SetMessage(v string) *UpgradeCDCVersionResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeCDCVersionResponseBody) SetRequestId(v string) *UpgradeCDCVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeCDCVersionResponseBody) SetSuccess(v bool) *UpgradeCDCVersionResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeCDCVersionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpgradeCDCVersionResponseBodyData struct {
	// example:
	//
	// polardb-2.4.0_5.4.19-20250116_xcluster5.4.20-20241213
	MinorVersion *string                                      `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	TaskList     []*UpgradeCDCVersionResponseBodyDataTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
}

func (s UpgradeCDCVersionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpgradeCDCVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpgradeCDCVersionResponseBodyData) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *UpgradeCDCVersionResponseBodyData) GetTaskList() []*UpgradeCDCVersionResponseBodyDataTaskList {
	return s.TaskList
}

func (s *UpgradeCDCVersionResponseBodyData) SetMinorVersion(v string) *UpgradeCDCVersionResponseBodyData {
	s.MinorVersion = &v
	return s
}

func (s *UpgradeCDCVersionResponseBodyData) SetTaskList(v []*UpgradeCDCVersionResponseBodyDataTaskList) *UpgradeCDCVersionResponseBodyData {
	s.TaskList = v
	return s
}

func (s *UpgradeCDCVersionResponseBodyData) Validate() error {
	if s.TaskList != nil {
		for _, item := range s.TaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpgradeCDCVersionResponseBodyDataTaskList struct {
	// example:
	//
	// rm-uf68f345****88zf8
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	// example:
	//
	// 1861190497624654848
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpgradeCDCVersionResponseBodyDataTaskList) String() string {
	return dara.Prettify(s)
}

func (s UpgradeCDCVersionResponseBodyDataTaskList) GoString() string {
	return s.String()
}

func (s *UpgradeCDCVersionResponseBodyDataTaskList) GetDbInstanceName() *string {
	return s.DbInstanceName
}

func (s *UpgradeCDCVersionResponseBodyDataTaskList) GetTaskId() *int32 {
	return s.TaskId
}

func (s *UpgradeCDCVersionResponseBodyDataTaskList) SetDbInstanceName(v string) *UpgradeCDCVersionResponseBodyDataTaskList {
	s.DbInstanceName = &v
	return s
}

func (s *UpgradeCDCVersionResponseBodyDataTaskList) SetTaskId(v int32) *UpgradeCDCVersionResponseBodyDataTaskList {
	s.TaskId = &v
	return s
}

func (s *UpgradeCDCVersionResponseBodyDataTaskList) Validate() error {
	return dara.Validate(s)
}
