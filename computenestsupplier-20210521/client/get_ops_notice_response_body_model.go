// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpsNoticeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttributes(v string) *GetOpsNoticeResponseBody
	GetAttributes() *string
	SetCategory(v string) *GetOpsNoticeResponseBody
	GetCategory() *string
	SetContent(v string) *GetOpsNoticeResponseBody
	GetContent() *string
	SetNoticeId(v string) *GetOpsNoticeResponseBody
	GetNoticeId() *string
	SetRequestId(v string) *GetOpsNoticeResponseBody
	GetRequestId() *string
	SetServiceId(v string) *GetOpsNoticeResponseBody
	GetServiceId() *string
	SetServiceInstanceCount(v string) *GetOpsNoticeResponseBody
	GetServiceInstanceCount() *string
	SetServiceName(v string) *GetOpsNoticeResponseBody
	GetServiceName() *string
	SetServiceVersions(v []*string) *GetOpsNoticeResponseBody
	GetServiceVersions() []*string
	SetSeverity(v string) *GetOpsNoticeResponseBody
	GetSeverity() *string
	SetSolutions(v string) *GetOpsNoticeResponseBody
	GetSolutions() *string
	SetStartTime(v string) *GetOpsNoticeResponseBody
	GetStartTime() *string
	SetSuccess(v string) *GetOpsNoticeResponseBody
	GetSuccess() *string
	SetType(v string) *GetOpsNoticeResponseBody
	GetType() *string
	SetUserCount(v string) *GetOpsNoticeResponseBody
	GetUserCount() *string
}

type GetOpsNoticeResponseBody struct {
	// example:
	//
	// {\\"cveId\\":\\"CVE-2021-4034\\"}
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// example:
	//
	// Availability
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// message
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// notice-2338d3835bxxxxx
	NoticeId *string `json:"NoticeId,omitempty" xml:"NoticeId,omitempty"`
	// example:
	//
	// 1B3AD3CC-E938-5042-A771-7FD9A2FE03F3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// service-c2d118c9193e49xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// 2
	ServiceInstanceCount *string   `json:"ServiceInstanceCount,omitempty" xml:"ServiceInstanceCount,omitempty"`
	ServiceName          *string   `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceVersions      []*string `json:"ServiceVersions,omitempty" xml:"ServiceVersions,omitempty" type:"Repeated"`
	// example:
	//
	// Critical
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// example:
	//
	// Solutions
	Solutions *string `json:"Solutions,omitempty" xml:"Solutions,omitempty"`
	// example:
	//
	// 2024-11-18T02:05:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// ServiceInstanceUpgrade
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1
	UserCount *string `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
}

func (s GetOpsNoticeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOpsNoticeResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpsNoticeResponseBody) GetAttributes() *string {
	return s.Attributes
}

func (s *GetOpsNoticeResponseBody) GetCategory() *string {
	return s.Category
}

func (s *GetOpsNoticeResponseBody) GetContent() *string {
	return s.Content
}

func (s *GetOpsNoticeResponseBody) GetNoticeId() *string {
	return s.NoticeId
}

func (s *GetOpsNoticeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOpsNoticeResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetOpsNoticeResponseBody) GetServiceInstanceCount() *string {
	return s.ServiceInstanceCount
}

func (s *GetOpsNoticeResponseBody) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetOpsNoticeResponseBody) GetServiceVersions() []*string {
	return s.ServiceVersions
}

func (s *GetOpsNoticeResponseBody) GetSeverity() *string {
	return s.Severity
}

func (s *GetOpsNoticeResponseBody) GetSolutions() *string {
	return s.Solutions
}

func (s *GetOpsNoticeResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetOpsNoticeResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetOpsNoticeResponseBody) GetType() *string {
	return s.Type
}

func (s *GetOpsNoticeResponseBody) GetUserCount() *string {
	return s.UserCount
}

func (s *GetOpsNoticeResponseBody) SetAttributes(v string) *GetOpsNoticeResponseBody {
	s.Attributes = &v
	return s
}

func (s *GetOpsNoticeResponseBody) SetCategory(v string) *GetOpsNoticeResponseBody {
	s.Category = &v
	return s
}

func (s *GetOpsNoticeResponseBody) SetContent(v string) *GetOpsNoticeResponseBody {
	s.Content = &v
	return s
}

func (s *GetOpsNoticeResponseBody) SetNoticeId(v string) *GetOpsNoticeResponseBody {
	s.NoticeId = &v
	return s
}

func (s *GetOpsNoticeResponseBody) SetRequestId(v string) *GetOpsNoticeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpsNoticeResponseBody) SetServiceId(v string) *GetOpsNoticeResponseBody {
	s.ServiceId = &v
	return s
}

func (s *GetOpsNoticeResponseBody) SetServiceInstanceCount(v string) *GetOpsNoticeResponseBody {
	s.ServiceInstanceCount = &v
	return s
}

func (s *GetOpsNoticeResponseBody) SetServiceName(v string) *GetOpsNoticeResponseBody {
	s.ServiceName = &v
	return s
}

func (s *GetOpsNoticeResponseBody) SetServiceVersions(v []*string) *GetOpsNoticeResponseBody {
	s.ServiceVersions = v
	return s
}

func (s *GetOpsNoticeResponseBody) SetSeverity(v string) *GetOpsNoticeResponseBody {
	s.Severity = &v
	return s
}

func (s *GetOpsNoticeResponseBody) SetSolutions(v string) *GetOpsNoticeResponseBody {
	s.Solutions = &v
	return s
}

func (s *GetOpsNoticeResponseBody) SetStartTime(v string) *GetOpsNoticeResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetOpsNoticeResponseBody) SetSuccess(v string) *GetOpsNoticeResponseBody {
	s.Success = &v
	return s
}

func (s *GetOpsNoticeResponseBody) SetType(v string) *GetOpsNoticeResponseBody {
	s.Type = &v
	return s
}

func (s *GetOpsNoticeResponseBody) SetUserCount(v string) *GetOpsNoticeResponseBody {
	s.UserCount = &v
	return s
}

func (s *GetOpsNoticeResponseBody) Validate() error {
	return dara.Validate(s)
}
