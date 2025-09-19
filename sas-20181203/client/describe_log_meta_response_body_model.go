// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogMetaList(v []*DescribeLogMetaResponseBodyLogMetaList) *DescribeLogMetaResponseBody
	GetLogMetaList() []*DescribeLogMetaResponseBodyLogMetaList
	SetRequestId(v string) *DescribeLogMetaResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeLogMetaResponseBody
	GetTotalCount() *int32
}

type DescribeLogMetaResponseBody struct {
	// An array that consists of the configurations of the log analysis feature.
	LogMetaList []*DescribeLogMetaResponseBodyLogMetaList `json:"LogMetaList,omitempty" xml:"LogMetaList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D0D6E6E4-CB8C-4897-B852-46AEFDA04B21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLogMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogMetaResponseBody) GetLogMetaList() []*DescribeLogMetaResponseBodyLogMetaList {
	return s.LogMetaList
}

func (s *DescribeLogMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogMetaResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeLogMetaResponseBody) SetLogMetaList(v []*DescribeLogMetaResponseBodyLogMetaList) *DescribeLogMetaResponseBody {
	s.LogMetaList = v
	return s
}

func (s *DescribeLogMetaResponseBody) SetRequestId(v string) *DescribeLogMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogMetaResponseBody) SetTotalCount(v int32) *DescribeLogMetaResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLogMetaResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLogMetaResponseBodyLogMetaList struct {
	// The category of logs. Valid values:
	//
	// 	- **host**
	//
	// 	- **network**
	//
	// 	- **security**
	//
	// example:
	//
	// host
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time period after which logs in hot storage are moved to cold storage. Unit: days.
	//
	// >  If the value is -1, logs that are stored in hot storage are not moved to cold storage.
	//
	// example:
	//
	// -1
	HotTtl *int32 `json:"HotTtl,omitempty" xml:"HotTtl,omitempty"`
	// The name of the log type.
	//
	// example:
	//
	// Login
	LogDesc *string `json:"LogDesc,omitempty" xml:"LogDesc,omitempty"`
	// The name of the dedicated Logstore in which logs are stored.
	//
	// example:
	//
	// aegis-log-login
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// aegis-log
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The status of the log analysis feature. Valid values:
	//
	// 	- **disabled**
	//
	// 	- **enabled**
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The topic of logs that are delivered.
	//
	// example:
	//
	// aegis-log-login
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The number of days during which logs can be retained.
	//
	// example:
	//
	// 180
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The name of the dedicated Logstore in which user logs are stored.
	//
	// example:
	//
	// aegis-log-network
	UserLogStore *string `json:"UserLogStore,omitempty" xml:"UserLogStore,omitempty"`
	// The name of the dedicated project in which logs are stored.
	//
	// example:
	//
	// sasnew-log-XXXX-cn-hangzhou
	UserProject *string `json:"UserProject,omitempty" xml:"UserProject,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	UserRegion *string `json:"UserRegion,omitempty" xml:"UserRegion,omitempty"`
}

func (s DescribeLogMetaResponseBodyLogMetaList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogMetaResponseBodyLogMetaList) GoString() string {
	return s.String()
}

func (s *DescribeLogMetaResponseBodyLogMetaList) GetCategory() *string {
	return s.Category
}

func (s *DescribeLogMetaResponseBodyLogMetaList) GetHotTtl() *int32 {
	return s.HotTtl
}

func (s *DescribeLogMetaResponseBodyLogMetaList) GetLogDesc() *string {
	return s.LogDesc
}

func (s *DescribeLogMetaResponseBodyLogMetaList) GetLogStore() *string {
	return s.LogStore
}

func (s *DescribeLogMetaResponseBodyLogMetaList) GetProject() *string {
	return s.Project
}

func (s *DescribeLogMetaResponseBodyLogMetaList) GetStatus() *string {
	return s.Status
}

func (s *DescribeLogMetaResponseBodyLogMetaList) GetTopic() *string {
	return s.Topic
}

func (s *DescribeLogMetaResponseBodyLogMetaList) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeLogMetaResponseBodyLogMetaList) GetUserLogStore() *string {
	return s.UserLogStore
}

func (s *DescribeLogMetaResponseBodyLogMetaList) GetUserProject() *string {
	return s.UserProject
}

func (s *DescribeLogMetaResponseBodyLogMetaList) GetUserRegion() *string {
	return s.UserRegion
}

func (s *DescribeLogMetaResponseBodyLogMetaList) SetCategory(v string) *DescribeLogMetaResponseBodyLogMetaList {
	s.Category = &v
	return s
}

func (s *DescribeLogMetaResponseBodyLogMetaList) SetHotTtl(v int32) *DescribeLogMetaResponseBodyLogMetaList {
	s.HotTtl = &v
	return s
}

func (s *DescribeLogMetaResponseBodyLogMetaList) SetLogDesc(v string) *DescribeLogMetaResponseBodyLogMetaList {
	s.LogDesc = &v
	return s
}

func (s *DescribeLogMetaResponseBodyLogMetaList) SetLogStore(v string) *DescribeLogMetaResponseBodyLogMetaList {
	s.LogStore = &v
	return s
}

func (s *DescribeLogMetaResponseBodyLogMetaList) SetProject(v string) *DescribeLogMetaResponseBodyLogMetaList {
	s.Project = &v
	return s
}

func (s *DescribeLogMetaResponseBodyLogMetaList) SetStatus(v string) *DescribeLogMetaResponseBodyLogMetaList {
	s.Status = &v
	return s
}

func (s *DescribeLogMetaResponseBodyLogMetaList) SetTopic(v string) *DescribeLogMetaResponseBodyLogMetaList {
	s.Topic = &v
	return s
}

func (s *DescribeLogMetaResponseBodyLogMetaList) SetTtl(v int32) *DescribeLogMetaResponseBodyLogMetaList {
	s.Ttl = &v
	return s
}

func (s *DescribeLogMetaResponseBodyLogMetaList) SetUserLogStore(v string) *DescribeLogMetaResponseBodyLogMetaList {
	s.UserLogStore = &v
	return s
}

func (s *DescribeLogMetaResponseBodyLogMetaList) SetUserProject(v string) *DescribeLogMetaResponseBodyLogMetaList {
	s.UserProject = &v
	return s
}

func (s *DescribeLogMetaResponseBodyLogMetaList) SetUserRegion(v string) *DescribeLogMetaResponseBodyLogMetaList {
	s.UserRegion = &v
	return s
}

func (s *DescribeLogMetaResponseBodyLogMetaList) Validate() error {
	return dara.Validate(s)
}
