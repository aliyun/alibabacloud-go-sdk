// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaLifecycleRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForbiddenRuleIds(v []*string) *GetMediaLifecycleRuleResponseBody
	GetForbiddenRuleIds() []*string
	SetNonExistRuleIds(v []*string) *GetMediaLifecycleRuleResponseBody
	GetNonExistRuleIds() []*string
	SetRequestId(v string) *GetMediaLifecycleRuleResponseBody
	GetRequestId() *string
	SetRuleList(v []*GetMediaLifecycleRuleResponseBodyRuleList) *GetMediaLifecycleRuleResponseBody
	GetRuleList() []*GetMediaLifecycleRuleResponseBodyRuleList
}

type GetMediaLifecycleRuleResponseBody struct {
	ForbiddenRuleIds []*string                                    `json:"ForbiddenRuleIds,omitempty" xml:"ForbiddenRuleIds,omitempty" type:"Repeated"`
	NonExistRuleIds  []*string                                    `json:"NonExistRuleIds,omitempty" xml:"NonExistRuleIds,omitempty" type:"Repeated"`
	RequestId        *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleList         []*GetMediaLifecycleRuleResponseBodyRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s GetMediaLifecycleRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLifecycleRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaLifecycleRuleResponseBody) GetForbiddenRuleIds() []*string {
	return s.ForbiddenRuleIds
}

func (s *GetMediaLifecycleRuleResponseBody) GetNonExistRuleIds() []*string {
	return s.NonExistRuleIds
}

func (s *GetMediaLifecycleRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaLifecycleRuleResponseBody) GetRuleList() []*GetMediaLifecycleRuleResponseBodyRuleList {
	return s.RuleList
}

func (s *GetMediaLifecycleRuleResponseBody) SetForbiddenRuleIds(v []*string) *GetMediaLifecycleRuleResponseBody {
	s.ForbiddenRuleIds = v
	return s
}

func (s *GetMediaLifecycleRuleResponseBody) SetNonExistRuleIds(v []*string) *GetMediaLifecycleRuleResponseBody {
	s.NonExistRuleIds = v
	return s
}

func (s *GetMediaLifecycleRuleResponseBody) SetRequestId(v string) *GetMediaLifecycleRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaLifecycleRuleResponseBody) SetRuleList(v []*GetMediaLifecycleRuleResponseBodyRuleList) *GetMediaLifecycleRuleResponseBody {
	s.RuleList = v
	return s
}

func (s *GetMediaLifecycleRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMediaLifecycleRuleResponseBodyRuleList struct {
	AppId            *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreationTime     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	RuleContent      *string `json:"RuleContent,omitempty" xml:"RuleContent,omitempty"`
	RuleId           *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleType         *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetMediaLifecycleRuleResponseBodyRuleList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLifecycleRuleResponseBodyRuleList) GoString() string {
	return s.String()
}

func (s *GetMediaLifecycleRuleResponseBodyRuleList) GetAppId() *string {
	return s.AppId
}

func (s *GetMediaLifecycleRuleResponseBodyRuleList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetMediaLifecycleRuleResponseBodyRuleList) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *GetMediaLifecycleRuleResponseBodyRuleList) GetRuleContent() *string {
	return s.RuleContent
}

func (s *GetMediaLifecycleRuleResponseBodyRuleList) GetRuleId() *string {
	return s.RuleId
}

func (s *GetMediaLifecycleRuleResponseBodyRuleList) GetRuleType() *string {
	return s.RuleType
}

func (s *GetMediaLifecycleRuleResponseBodyRuleList) GetStatus() *string {
	return s.Status
}

func (s *GetMediaLifecycleRuleResponseBodyRuleList) SetAppId(v string) *GetMediaLifecycleRuleResponseBodyRuleList {
	s.AppId = &v
	return s
}

func (s *GetMediaLifecycleRuleResponseBodyRuleList) SetCreationTime(v string) *GetMediaLifecycleRuleResponseBodyRuleList {
	s.CreationTime = &v
	return s
}

func (s *GetMediaLifecycleRuleResponseBodyRuleList) SetModificationTime(v string) *GetMediaLifecycleRuleResponseBodyRuleList {
	s.ModificationTime = &v
	return s
}

func (s *GetMediaLifecycleRuleResponseBodyRuleList) SetRuleContent(v string) *GetMediaLifecycleRuleResponseBodyRuleList {
	s.RuleContent = &v
	return s
}

func (s *GetMediaLifecycleRuleResponseBodyRuleList) SetRuleId(v string) *GetMediaLifecycleRuleResponseBodyRuleList {
	s.RuleId = &v
	return s
}

func (s *GetMediaLifecycleRuleResponseBodyRuleList) SetRuleType(v string) *GetMediaLifecycleRuleResponseBodyRuleList {
	s.RuleType = &v
	return s
}

func (s *GetMediaLifecycleRuleResponseBodyRuleList) SetStatus(v string) *GetMediaLifecycleRuleResponseBodyRuleList {
	s.Status = &v
	return s
}

func (s *GetMediaLifecycleRuleResponseBodyRuleList) Validate() error {
	return dara.Validate(s)
}
