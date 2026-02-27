// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstanceRecordsResponseBody
	GetCode() *string
	SetMessage(v string) *ListInstanceRecordsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInstanceRecordsResponseBody
	GetRequestId() *string
	SetRoot(v *ListInstanceRecordsResponseBodyRoot) *ListInstanceRecordsResponseBody
	GetRoot() *ListInstanceRecordsResponseBodyRoot
	SetSuccess(v bool) *ListInstanceRecordsResponseBody
	GetSuccess() *bool
}

type ListInstanceRecordsResponseBody struct {
	// example:
	//
	// InvalidParamter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter is null or invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Root      *ListInstanceRecordsResponseBodyRoot `json:"Root,omitempty" xml:"Root,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInstanceRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceRecordsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstanceRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstanceRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceRecordsResponseBody) GetRoot() *ListInstanceRecordsResponseBodyRoot {
	return s.Root
}

func (s *ListInstanceRecordsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstanceRecordsResponseBody) SetCode(v string) *ListInstanceRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceRecordsResponseBody) SetMessage(v string) *ListInstanceRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceRecordsResponseBody) SetRequestId(v string) *ListInstanceRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceRecordsResponseBody) SetRoot(v *ListInstanceRecordsResponseBodyRoot) *ListInstanceRecordsResponseBody {
	s.Root = v
	return s
}

func (s *ListInstanceRecordsResponseBody) SetSuccess(v bool) *ListInstanceRecordsResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstanceRecordsResponseBody) Validate() error {
	if s.Root != nil {
		if err := s.Root.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstanceRecordsResponseBodyRoot struct {
	RecordList []*ListInstanceRecordsResponseBodyRootRecordList `json:"RecordList,omitempty" xml:"RecordList,omitempty" type:"Repeated"`
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstanceRecordsResponseBodyRoot) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRecordsResponseBodyRoot) GoString() string {
	return s.String()
}

func (s *ListInstanceRecordsResponseBodyRoot) GetRecordList() []*ListInstanceRecordsResponseBodyRootRecordList {
	return s.RecordList
}

func (s *ListInstanceRecordsResponseBodyRoot) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInstanceRecordsResponseBodyRoot) SetRecordList(v []*ListInstanceRecordsResponseBodyRootRecordList) *ListInstanceRecordsResponseBodyRoot {
	s.RecordList = v
	return s
}

func (s *ListInstanceRecordsResponseBodyRoot) SetTotalCount(v int32) *ListInstanceRecordsResponseBodyRoot {
	s.TotalCount = &v
	return s
}

func (s *ListInstanceRecordsResponseBodyRoot) Validate() error {
	if s.RecordList != nil {
		for _, item := range s.RecordList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceRecordsResponseBodyRootRecordList struct {
	// example:
	//
	// 1234
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// 2023-11-16T02:59:39Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 2023-04-10T12:41:28Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// i-xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// https://ecs-workbench.aliyun.com/view/instance/record/replay/abc
	InstanceRecordUrl *string `json:"InstanceRecordUrl,omitempty" xml:"InstanceRecordUrl,omitempty"`
	// example:
	//
	// 123
	RecordDurationMillis *int64 `json:"RecordDurationMillis,omitempty" xml:"RecordDurationMillis,omitempty"`
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// abc
	TerminalSessionToken *string `json:"TerminalSessionToken,omitempty" xml:"TerminalSessionToken,omitempty"`
}

func (s ListInstanceRecordsResponseBodyRootRecordList) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRecordsResponseBodyRootRecordList) GoString() string {
	return s.String()
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) GetInstanceRecordUrl() *string {
	return s.InstanceRecordUrl
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) GetRecordDurationMillis() *int64 {
	return s.RecordDurationMillis
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) GetStatus() *string {
	return s.Status
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) GetTerminalSessionToken() *string {
	return s.TerminalSessionToken
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) SetAccountId(v int64) *ListInstanceRecordsResponseBodyRootRecordList {
	s.AccountId = &v
	return s
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) SetExpireTime(v string) *ListInstanceRecordsResponseBodyRootRecordList {
	s.ExpireTime = &v
	return s
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) SetGmtCreate(v string) *ListInstanceRecordsResponseBodyRootRecordList {
	s.GmtCreate = &v
	return s
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) SetInstanceId(v string) *ListInstanceRecordsResponseBodyRootRecordList {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) SetInstanceRecordUrl(v string) *ListInstanceRecordsResponseBodyRootRecordList {
	s.InstanceRecordUrl = &v
	return s
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) SetRecordDurationMillis(v int64) *ListInstanceRecordsResponseBodyRootRecordList {
	s.RecordDurationMillis = &v
	return s
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) SetStatus(v string) *ListInstanceRecordsResponseBodyRootRecordList {
	s.Status = &v
	return s
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) SetTerminalSessionToken(v string) *ListInstanceRecordsResponseBodyRootRecordList {
	s.TerminalSessionToken = &v
	return s
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) Validate() error {
	return dara.Validate(s)
}
