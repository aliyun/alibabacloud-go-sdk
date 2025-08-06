// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaLifecycleRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListMediaLifecycleRuleResponseBody
	GetRequestId() *string
	SetRuleList(v []*ListMediaLifecycleRuleResponseBodyRuleList) *ListMediaLifecycleRuleResponseBody
	GetRuleList() []*ListMediaLifecycleRuleResponseBodyRuleList
	SetTotal(v int64) *ListMediaLifecycleRuleResponseBody
	GetTotal() *int64
}

type ListMediaLifecycleRuleResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleList  []*ListMediaLifecycleRuleResponseBodyRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
	Total     *int64                                        `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMediaLifecycleRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLifecycleRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListMediaLifecycleRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMediaLifecycleRuleResponseBody) GetRuleList() []*ListMediaLifecycleRuleResponseBodyRuleList {
	return s.RuleList
}

func (s *ListMediaLifecycleRuleResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListMediaLifecycleRuleResponseBody) SetRequestId(v string) *ListMediaLifecycleRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMediaLifecycleRuleResponseBody) SetRuleList(v []*ListMediaLifecycleRuleResponseBodyRuleList) *ListMediaLifecycleRuleResponseBody {
	s.RuleList = v
	return s
}

func (s *ListMediaLifecycleRuleResponseBody) SetTotal(v int64) *ListMediaLifecycleRuleResponseBody {
	s.Total = &v
	return s
}

func (s *ListMediaLifecycleRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMediaLifecycleRuleResponseBodyRuleList struct {
	AppId            *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreationTime     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	RuleContent      *string `json:"RuleContent,omitempty" xml:"RuleContent,omitempty"`
	RuleId           *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleType         *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListMediaLifecycleRuleResponseBodyRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLifecycleRuleResponseBodyRuleList) GoString() string {
	return s.String()
}

func (s *ListMediaLifecycleRuleResponseBodyRuleList) GetAppId() *string {
	return s.AppId
}

func (s *ListMediaLifecycleRuleResponseBodyRuleList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListMediaLifecycleRuleResponseBodyRuleList) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *ListMediaLifecycleRuleResponseBodyRuleList) GetRuleContent() *string {
	return s.RuleContent
}

func (s *ListMediaLifecycleRuleResponseBodyRuleList) GetRuleId() *string {
	return s.RuleId
}

func (s *ListMediaLifecycleRuleResponseBodyRuleList) GetRuleType() *string {
	return s.RuleType
}

func (s *ListMediaLifecycleRuleResponseBodyRuleList) GetStatus() *string {
	return s.Status
}

func (s *ListMediaLifecycleRuleResponseBodyRuleList) SetAppId(v string) *ListMediaLifecycleRuleResponseBodyRuleList {
	s.AppId = &v
	return s
}

func (s *ListMediaLifecycleRuleResponseBodyRuleList) SetCreationTime(v string) *ListMediaLifecycleRuleResponseBodyRuleList {
	s.CreationTime = &v
	return s
}

func (s *ListMediaLifecycleRuleResponseBodyRuleList) SetModificationTime(v string) *ListMediaLifecycleRuleResponseBodyRuleList {
	s.ModificationTime = &v
	return s
}

func (s *ListMediaLifecycleRuleResponseBodyRuleList) SetRuleContent(v string) *ListMediaLifecycleRuleResponseBodyRuleList {
	s.RuleContent = &v
	return s
}

func (s *ListMediaLifecycleRuleResponseBodyRuleList) SetRuleId(v string) *ListMediaLifecycleRuleResponseBodyRuleList {
	s.RuleId = &v
	return s
}

func (s *ListMediaLifecycleRuleResponseBodyRuleList) SetRuleType(v string) *ListMediaLifecycleRuleResponseBodyRuleList {
	s.RuleType = &v
	return s
}

func (s *ListMediaLifecycleRuleResponseBodyRuleList) SetStatus(v string) *ListMediaLifecycleRuleResponseBodyRuleList {
	s.Status = &v
	return s
}

func (s *ListMediaLifecycleRuleResponseBodyRuleList) Validate() error {
	return dara.Validate(s)
}
