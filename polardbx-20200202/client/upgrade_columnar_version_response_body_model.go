// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeColumnarVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpgradeColumnarVersionResponseBodyData) *UpgradeColumnarVersionResponseBody
	GetData() *UpgradeColumnarVersionResponseBodyData
	SetMessage(v string) *UpgradeColumnarVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpgradeColumnarVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpgradeColumnarVersionResponseBody
	GetSuccess() *bool
}

type UpgradeColumnarVersionResponseBody struct {
	Data *UpgradeColumnarVersionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// AE4F6C34-****-45AA-B5DC-4B8D816F6305
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpgradeColumnarVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeColumnarVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeColumnarVersionResponseBody) GetData() *UpgradeColumnarVersionResponseBodyData {
	return s.Data
}

func (s *UpgradeColumnarVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpgradeColumnarVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeColumnarVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpgradeColumnarVersionResponseBody) SetData(v *UpgradeColumnarVersionResponseBodyData) *UpgradeColumnarVersionResponseBody {
	s.Data = v
	return s
}

func (s *UpgradeColumnarVersionResponseBody) SetMessage(v string) *UpgradeColumnarVersionResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeColumnarVersionResponseBody) SetRequestId(v string) *UpgradeColumnarVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeColumnarVersionResponseBody) SetSuccess(v bool) *UpgradeColumnarVersionResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeColumnarVersionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpgradeColumnarVersionResponseBodyData struct {
	// example:
	//
	// polardb-2.4.0_5.4.19-20250116_xcluster5.4.20-20241213
	MinorVersion *string                                           `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	TaskList     []*UpgradeColumnarVersionResponseBodyDataTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
}

func (s UpgradeColumnarVersionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpgradeColumnarVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpgradeColumnarVersionResponseBodyData) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *UpgradeColumnarVersionResponseBodyData) GetTaskList() []*UpgradeColumnarVersionResponseBodyDataTaskList {
	return s.TaskList
}

func (s *UpgradeColumnarVersionResponseBodyData) SetMinorVersion(v string) *UpgradeColumnarVersionResponseBodyData {
	s.MinorVersion = &v
	return s
}

func (s *UpgradeColumnarVersionResponseBodyData) SetTaskList(v []*UpgradeColumnarVersionResponseBodyDataTaskList) *UpgradeColumnarVersionResponseBodyData {
	s.TaskList = v
	return s
}

func (s *UpgradeColumnarVersionResponseBodyData) Validate() error {
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

type UpgradeColumnarVersionResponseBodyDataTaskList struct {
	// example:
	//
	// rm-uf68f345****88zf8
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	// example:
	//
	// 1861190497624654848
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpgradeColumnarVersionResponseBodyDataTaskList) String() string {
	return dara.Prettify(s)
}

func (s UpgradeColumnarVersionResponseBodyDataTaskList) GoString() string {
	return s.String()
}

func (s *UpgradeColumnarVersionResponseBodyDataTaskList) GetDbInstanceName() *string {
	return s.DbInstanceName
}

func (s *UpgradeColumnarVersionResponseBodyDataTaskList) GetTaskId() *int32 {
	return s.TaskId
}

func (s *UpgradeColumnarVersionResponseBodyDataTaskList) SetDbInstanceName(v string) *UpgradeColumnarVersionResponseBodyDataTaskList {
	s.DbInstanceName = &v
	return s
}

func (s *UpgradeColumnarVersionResponseBodyDataTaskList) SetTaskId(v int32) *UpgradeColumnarVersionResponseBodyDataTaskList {
	s.TaskId = &v
	return s
}

func (s *UpgradeColumnarVersionResponseBodyDataTaskList) Validate() error {
	return dara.Validate(s)
}
