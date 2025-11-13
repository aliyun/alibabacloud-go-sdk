// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSearchCaseFullTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RunSearchCaseFullTextRequest
	GetAppId() *string
	SetFilterCondition(v *RunSearchCaseFullTextRequestFilterCondition) *RunSearchCaseFullTextRequest
	GetFilterCondition() *RunSearchCaseFullTextRequestFilterCondition
	SetPageParam(v *RunSearchCaseFullTextRequestPageParam) *RunSearchCaseFullTextRequest
	GetPageParam() *RunSearchCaseFullTextRequestPageParam
	SetQuery(v string) *RunSearchCaseFullTextRequest
	GetQuery() *string
	SetQueryKeywords(v []*string) *RunSearchCaseFullTextRequest
	GetQueryKeywords() []*string
	SetReferLevel(v string) *RunSearchCaseFullTextRequest
	GetReferLevel() *string
	SetSortKeyAndDirection(v map[string]*string) *RunSearchCaseFullTextRequest
	GetSortKeyAndDirection() map[string]*string
	SetThread(v *RunSearchCaseFullTextRequestThread) *RunSearchCaseFullTextRequest
	GetThread() *RunSearchCaseFullTextRequestThread
}

type RunSearchCaseFullTextRequest struct {
	// example:
	//
	// farui
	AppId           *string                                      `json:"appId,omitempty" xml:"appId,omitempty"`
	FilterCondition *RunSearchCaseFullTextRequestFilterCondition `json:"filterCondition,omitempty" xml:"filterCondition,omitempty" type:"Struct"`
	// This parameter is required.
	PageParam *RunSearchCaseFullTextRequestPageParam `json:"pageParam,omitempty" xml:"pageParam,omitempty" type:"Struct"`
	// This parameter is required.
	Query               *string                             `json:"query,omitempty" xml:"query,omitempty"`
	QueryKeywords       []*string                           `json:"queryKeywords,omitempty" xml:"queryKeywords,omitempty" type:"Repeated"`
	ReferLevel          *string                             `json:"referLevel,omitempty" xml:"referLevel,omitempty"`
	SortKeyAndDirection map[string]*string                  `json:"sortKeyAndDirection,omitempty" xml:"sortKeyAndDirection,omitempty"`
	Thread              *RunSearchCaseFullTextRequestThread `json:"thread,omitempty" xml:"thread,omitempty" type:"Struct"`
}

func (s RunSearchCaseFullTextRequest) String() string {
	return dara.Prettify(s)
}

func (s RunSearchCaseFullTextRequest) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextRequest) GetAppId() *string {
	return s.AppId
}

func (s *RunSearchCaseFullTextRequest) GetFilterCondition() *RunSearchCaseFullTextRequestFilterCondition {
	return s.FilterCondition
}

func (s *RunSearchCaseFullTextRequest) GetPageParam() *RunSearchCaseFullTextRequestPageParam {
	return s.PageParam
}

func (s *RunSearchCaseFullTextRequest) GetQuery() *string {
	return s.Query
}

func (s *RunSearchCaseFullTextRequest) GetQueryKeywords() []*string {
	return s.QueryKeywords
}

func (s *RunSearchCaseFullTextRequest) GetReferLevel() *string {
	return s.ReferLevel
}

func (s *RunSearchCaseFullTextRequest) GetSortKeyAndDirection() map[string]*string {
	return s.SortKeyAndDirection
}

func (s *RunSearchCaseFullTextRequest) GetThread() *RunSearchCaseFullTextRequestThread {
	return s.Thread
}

func (s *RunSearchCaseFullTextRequest) SetAppId(v string) *RunSearchCaseFullTextRequest {
	s.AppId = &v
	return s
}

func (s *RunSearchCaseFullTextRequest) SetFilterCondition(v *RunSearchCaseFullTextRequestFilterCondition) *RunSearchCaseFullTextRequest {
	s.FilterCondition = v
	return s
}

func (s *RunSearchCaseFullTextRequest) SetPageParam(v *RunSearchCaseFullTextRequestPageParam) *RunSearchCaseFullTextRequest {
	s.PageParam = v
	return s
}

func (s *RunSearchCaseFullTextRequest) SetQuery(v string) *RunSearchCaseFullTextRequest {
	s.Query = &v
	return s
}

func (s *RunSearchCaseFullTextRequest) SetQueryKeywords(v []*string) *RunSearchCaseFullTextRequest {
	s.QueryKeywords = v
	return s
}

func (s *RunSearchCaseFullTextRequest) SetReferLevel(v string) *RunSearchCaseFullTextRequest {
	s.ReferLevel = &v
	return s
}

func (s *RunSearchCaseFullTextRequest) SetSortKeyAndDirection(v map[string]*string) *RunSearchCaseFullTextRequest {
	s.SortKeyAndDirection = v
	return s
}

func (s *RunSearchCaseFullTextRequest) SetThread(v *RunSearchCaseFullTextRequestThread) *RunSearchCaseFullTextRequest {
	s.Thread = v
	return s
}

func (s *RunSearchCaseFullTextRequest) Validate() error {
	if s.FilterCondition != nil {
		if err := s.FilterCondition.Validate(); err != nil {
			return err
		}
	}
	if s.PageParam != nil {
		if err := s.PageParam.Validate(); err != nil {
			return err
		}
	}
	if s.Thread != nil {
		if err := s.Thread.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchCaseFullTextRequestFilterCondition struct {
	CaseNo    *string `json:"caseNo,omitempty" xml:"caseNo,omitempty"`
	CaseTitle *string `json:"caseTitle,omitempty" xml:"caseTitle,omitempty"`
}

func (s RunSearchCaseFullTextRequestFilterCondition) String() string {
	return dara.Prettify(s)
}

func (s RunSearchCaseFullTextRequestFilterCondition) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextRequestFilterCondition) GetCaseNo() *string {
	return s.CaseNo
}

func (s *RunSearchCaseFullTextRequestFilterCondition) GetCaseTitle() *string {
	return s.CaseTitle
}

func (s *RunSearchCaseFullTextRequestFilterCondition) SetCaseNo(v string) *RunSearchCaseFullTextRequestFilterCondition {
	s.CaseNo = &v
	return s
}

func (s *RunSearchCaseFullTextRequestFilterCondition) SetCaseTitle(v string) *RunSearchCaseFullTextRequestFilterCondition {
	s.CaseTitle = &v
	return s
}

func (s *RunSearchCaseFullTextRequestFilterCondition) Validate() error {
	return dara.Validate(s)
}

type RunSearchCaseFullTextRequestPageParam struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s RunSearchCaseFullTextRequestPageParam) String() string {
	return dara.Prettify(s)
}

func (s RunSearchCaseFullTextRequestPageParam) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextRequestPageParam) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *RunSearchCaseFullTextRequestPageParam) GetPageSize() *int32 {
	return s.PageSize
}

func (s *RunSearchCaseFullTextRequestPageParam) SetPageNumber(v int32) *RunSearchCaseFullTextRequestPageParam {
	s.PageNumber = &v
	return s
}

func (s *RunSearchCaseFullTextRequestPageParam) SetPageSize(v int32) *RunSearchCaseFullTextRequestPageParam {
	s.PageSize = &v
	return s
}

func (s *RunSearchCaseFullTextRequestPageParam) Validate() error {
	return dara.Validate(s)
}

type RunSearchCaseFullTextRequestThread struct {
	Messages []*RunSearchCaseFullTextRequestThreadMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
}

func (s RunSearchCaseFullTextRequestThread) String() string {
	return dara.Prettify(s)
}

func (s RunSearchCaseFullTextRequestThread) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextRequestThread) GetMessages() []*RunSearchCaseFullTextRequestThreadMessages {
	return s.Messages
}

func (s *RunSearchCaseFullTextRequestThread) SetMessages(v []*RunSearchCaseFullTextRequestThreadMessages) *RunSearchCaseFullTextRequestThread {
	s.Messages = v
	return s
}

func (s *RunSearchCaseFullTextRequestThread) Validate() error {
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchCaseFullTextRequestThreadMessages struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s RunSearchCaseFullTextRequestThreadMessages) String() string {
	return dara.Prettify(s)
}

func (s RunSearchCaseFullTextRequestThreadMessages) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextRequestThreadMessages) GetContent() *string {
	return s.Content
}

func (s *RunSearchCaseFullTextRequestThreadMessages) GetRole() *string {
	return s.Role
}

func (s *RunSearchCaseFullTextRequestThreadMessages) SetContent(v string) *RunSearchCaseFullTextRequestThreadMessages {
	s.Content = &v
	return s
}

func (s *RunSearchCaseFullTextRequestThreadMessages) SetRole(v string) *RunSearchCaseFullTextRequestThreadMessages {
	s.Role = &v
	return s
}

func (s *RunSearchCaseFullTextRequestThreadMessages) Validate() error {
	return dara.Validate(s)
}
