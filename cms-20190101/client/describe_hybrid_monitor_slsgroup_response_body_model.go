// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridMonitorSLSGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeHybridMonitorSLSGroupResponseBody
	GetCode() *string
	SetList(v []*DescribeHybridMonitorSLSGroupResponseBodyList) *DescribeHybridMonitorSLSGroupResponseBody
	GetList() []*DescribeHybridMonitorSLSGroupResponseBodyList
	SetMessage(v string) *DescribeHybridMonitorSLSGroupResponseBody
	GetMessage() *string
	SetPageNumber(v int64) *DescribeHybridMonitorSLSGroupResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeHybridMonitorSLSGroupResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeHybridMonitorSLSGroupResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeHybridMonitorSLSGroupResponseBody
	GetSuccess() *string
	SetTotal(v int64) *DescribeHybridMonitorSLSGroupResponseBody
	GetTotal() *int64
}

type DescribeHybridMonitorSLSGroupResponseBody struct {
	// The HTTP status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The queried Logstore groups.
	List []*DescribeHybridMonitorSLSGroupResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// NotFound.SLSGroup
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 66683237-7126-50F8-BBF8-D67ACC919A17
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeHybridMonitorSLSGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorSLSGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorSLSGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeHybridMonitorSLSGroupResponseBody) GetList() []*DescribeHybridMonitorSLSGroupResponseBodyList {
	return s.List
}

func (s *DescribeHybridMonitorSLSGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeHybridMonitorSLSGroupResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeHybridMonitorSLSGroupResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeHybridMonitorSLSGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridMonitorSLSGroupResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeHybridMonitorSLSGroupResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeHybridMonitorSLSGroupResponseBody) SetCode(v string) *DescribeHybridMonitorSLSGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHybridMonitorSLSGroupResponseBody) SetList(v []*DescribeHybridMonitorSLSGroupResponseBodyList) *DescribeHybridMonitorSLSGroupResponseBody {
	s.List = v
	return s
}

func (s *DescribeHybridMonitorSLSGroupResponseBody) SetMessage(v string) *DescribeHybridMonitorSLSGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHybridMonitorSLSGroupResponseBody) SetPageNumber(v int64) *DescribeHybridMonitorSLSGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHybridMonitorSLSGroupResponseBody) SetPageSize(v int64) *DescribeHybridMonitorSLSGroupResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHybridMonitorSLSGroupResponseBody) SetRequestId(v string) *DescribeHybridMonitorSLSGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridMonitorSLSGroupResponseBody) SetSuccess(v string) *DescribeHybridMonitorSLSGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHybridMonitorSLSGroupResponseBody) SetTotal(v int64) *DescribeHybridMonitorSLSGroupResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeHybridMonitorSLSGroupResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHybridMonitorSLSGroupResponseBodyList struct {
	// The time when the Logstore group was created.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1652845630000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The configurations of the Logstore group.
	SLSGroupConfig []*DescribeHybridMonitorSLSGroupResponseBodyListSLSGroupConfig `json:"SLSGroupConfig,omitempty" xml:"SLSGroupConfig,omitempty" type:"Repeated"`
	// The description of the Logstore group.
	SLSGroupDescription *string `json:"SLSGroupDescription,omitempty" xml:"SLSGroupDescription,omitempty"`
	// The name of the Logstore group.
	//
	// example:
	//
	// Logstore_test
	SLSGroupName *string `json:"SLSGroupName,omitempty" xml:"SLSGroupName,omitempty"`
	// The time when the Logstore group was modified.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1652845630000
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeHybridMonitorSLSGroupResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorSLSGroupResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorSLSGroupResponseBodyList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeHybridMonitorSLSGroupResponseBodyList) GetSLSGroupConfig() []*DescribeHybridMonitorSLSGroupResponseBodyListSLSGroupConfig {
	return s.SLSGroupConfig
}

func (s *DescribeHybridMonitorSLSGroupResponseBodyList) GetSLSGroupDescription() *string {
	return s.SLSGroupDescription
}

func (s *DescribeHybridMonitorSLSGroupResponseBodyList) GetSLSGroupName() *string {
	return s.SLSGroupName
}

func (s *DescribeHybridMonitorSLSGroupResponseBodyList) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeHybridMonitorSLSGroupResponseBodyList) SetCreateTime(v string) *DescribeHybridMonitorSLSGroupResponseBodyList {
	s.CreateTime = &v
	return s
}

func (s *DescribeHybridMonitorSLSGroupResponseBodyList) SetSLSGroupConfig(v []*DescribeHybridMonitorSLSGroupResponseBodyListSLSGroupConfig) *DescribeHybridMonitorSLSGroupResponseBodyList {
	s.SLSGroupConfig = v
	return s
}

func (s *DescribeHybridMonitorSLSGroupResponseBodyList) SetSLSGroupDescription(v string) *DescribeHybridMonitorSLSGroupResponseBodyList {
	s.SLSGroupDescription = &v
	return s
}

func (s *DescribeHybridMonitorSLSGroupResponseBodyList) SetSLSGroupName(v string) *DescribeHybridMonitorSLSGroupResponseBodyList {
	s.SLSGroupName = &v
	return s
}

func (s *DescribeHybridMonitorSLSGroupResponseBodyList) SetUpdateTime(v string) *DescribeHybridMonitorSLSGroupResponseBodyList {
	s.UpdateTime = &v
	return s
}

func (s *DescribeHybridMonitorSLSGroupResponseBodyList) Validate() error {
	if s.SLSGroupConfig != nil {
		for _, item := range s.SLSGroupConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHybridMonitorSLSGroupResponseBodyListSLSGroupConfig struct {
	// The Logstore.
	//
	// example:
	//
	// Logstore-aliyun-all
	SLSLogstore *string `json:"SLSLogstore,omitempty" xml:"SLSLogstore,omitempty"`
	// The Simple Log Service project.
	//
	// example:
	//
	// aliyun-project
	SLSProject *string `json:"SLSProject,omitempty" xml:"SLSProject,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	SLSRegion *string `json:"SLSRegion,omitempty" xml:"SLSRegion,omitempty"`
	// The member ID.
	//
	// **Description*	- This parameter is returned when you call the operation by using an administrative account.
	//
	// example:
	//
	// 120886317861****
	SLSUserId *string `json:"SLSUserId,omitempty" xml:"SLSUserId,omitempty"`
}

func (s DescribeHybridMonitorSLSGroupResponseBodyListSLSGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorSLSGroupResponseBodyListSLSGroupConfig) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorSLSGroupResponseBodyListSLSGroupConfig) GetSLSLogstore() *string {
	return s.SLSLogstore
}

func (s *DescribeHybridMonitorSLSGroupResponseBodyListSLSGroupConfig) GetSLSProject() *string {
	return s.SLSProject
}

func (s *DescribeHybridMonitorSLSGroupResponseBodyListSLSGroupConfig) GetSLSRegion() *string {
	return s.SLSRegion
}

func (s *DescribeHybridMonitorSLSGroupResponseBodyListSLSGroupConfig) GetSLSUserId() *string {
	return s.SLSUserId
}

func (s *DescribeHybridMonitorSLSGroupResponseBodyListSLSGroupConfig) SetSLSLogstore(v string) *DescribeHybridMonitorSLSGroupResponseBodyListSLSGroupConfig {
	s.SLSLogstore = &v
	return s
}

func (s *DescribeHybridMonitorSLSGroupResponseBodyListSLSGroupConfig) SetSLSProject(v string) *DescribeHybridMonitorSLSGroupResponseBodyListSLSGroupConfig {
	s.SLSProject = &v
	return s
}

func (s *DescribeHybridMonitorSLSGroupResponseBodyListSLSGroupConfig) SetSLSRegion(v string) *DescribeHybridMonitorSLSGroupResponseBodyListSLSGroupConfig {
	s.SLSRegion = &v
	return s
}

func (s *DescribeHybridMonitorSLSGroupResponseBodyListSLSGroupConfig) SetSLSUserId(v string) *DescribeHybridMonitorSLSGroupResponseBodyListSLSGroupConfig {
	s.SLSUserId = &v
	return s
}

func (s *DescribeHybridMonitorSLSGroupResponseBodyListSLSGroupConfig) Validate() error {
	return dara.Validate(s)
}
