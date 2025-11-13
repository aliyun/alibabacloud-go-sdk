// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSearchLawQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RunSearchLawQueryRequest
	GetAppId() *string
	SetFilterCondition(v *RunSearchLawQueryRequestFilterCondition) *RunSearchLawQueryRequest
	GetFilterCondition() *RunSearchLawQueryRequestFilterCondition
	SetPageParam(v *RunSearchLawQueryRequestPageParam) *RunSearchLawQueryRequest
	GetPageParam() *RunSearchLawQueryRequestPageParam
	SetQuery(v string) *RunSearchLawQueryRequest
	GetQuery() *string
	SetQueryKeywords(v []*string) *RunSearchLawQueryRequest
	GetQueryKeywords() []*string
	SetThread(v *RunSearchLawQueryRequestThread) *RunSearchLawQueryRequest
	GetThread() *RunSearchLawQueryRequestThread
}

type RunSearchLawQueryRequest struct {
	// example:
	//
	// farui
	AppId           *string                                  `json:"appId,omitempty" xml:"appId,omitempty"`
	FilterCondition *RunSearchLawQueryRequestFilterCondition `json:"filterCondition,omitempty" xml:"filterCondition,omitempty" type:"Struct"`
	PageParam       *RunSearchLawQueryRequestPageParam       `json:"pageParam,omitempty" xml:"pageParam,omitempty" type:"Struct"`
	// This parameter is required.
	Query         *string                         `json:"query,omitempty" xml:"query,omitempty"`
	QueryKeywords []*string                       `json:"queryKeywords,omitempty" xml:"queryKeywords,omitempty" type:"Repeated"`
	Thread        *RunSearchLawQueryRequestThread `json:"thread,omitempty" xml:"thread,omitempty" type:"Struct"`
}

func (s RunSearchLawQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s RunSearchLawQueryRequest) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryRequest) GetAppId() *string {
	return s.AppId
}

func (s *RunSearchLawQueryRequest) GetFilterCondition() *RunSearchLawQueryRequestFilterCondition {
	return s.FilterCondition
}

func (s *RunSearchLawQueryRequest) GetPageParam() *RunSearchLawQueryRequestPageParam {
	return s.PageParam
}

func (s *RunSearchLawQueryRequest) GetQuery() *string {
	return s.Query
}

func (s *RunSearchLawQueryRequest) GetQueryKeywords() []*string {
	return s.QueryKeywords
}

func (s *RunSearchLawQueryRequest) GetThread() *RunSearchLawQueryRequestThread {
	return s.Thread
}

func (s *RunSearchLawQueryRequest) SetAppId(v string) *RunSearchLawQueryRequest {
	s.AppId = &v
	return s
}

func (s *RunSearchLawQueryRequest) SetFilterCondition(v *RunSearchLawQueryRequestFilterCondition) *RunSearchLawQueryRequest {
	s.FilterCondition = v
	return s
}

func (s *RunSearchLawQueryRequest) SetPageParam(v *RunSearchLawQueryRequestPageParam) *RunSearchLawQueryRequest {
	s.PageParam = v
	return s
}

func (s *RunSearchLawQueryRequest) SetQuery(v string) *RunSearchLawQueryRequest {
	s.Query = &v
	return s
}

func (s *RunSearchLawQueryRequest) SetQueryKeywords(v []*string) *RunSearchLawQueryRequest {
	s.QueryKeywords = v
	return s
}

func (s *RunSearchLawQueryRequest) SetThread(v *RunSearchLawQueryRequestThread) *RunSearchLawQueryRequest {
	s.Thread = v
	return s
}

func (s *RunSearchLawQueryRequest) Validate() error {
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

type RunSearchLawQueryRequestFilterCondition struct {
	LawName *string `json:"lawName,omitempty" xml:"lawName,omitempty"`
}

func (s RunSearchLawQueryRequestFilterCondition) String() string {
	return dara.Prettify(s)
}

func (s RunSearchLawQueryRequestFilterCondition) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryRequestFilterCondition) GetLawName() *string {
	return s.LawName
}

func (s *RunSearchLawQueryRequestFilterCondition) SetLawName(v string) *RunSearchLawQueryRequestFilterCondition {
	s.LawName = &v
	return s
}

func (s *RunSearchLawQueryRequestFilterCondition) Validate() error {
	return dara.Validate(s)
}

type RunSearchLawQueryRequestPageParam struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s RunSearchLawQueryRequestPageParam) String() string {
	return dara.Prettify(s)
}

func (s RunSearchLawQueryRequestPageParam) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryRequestPageParam) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *RunSearchLawQueryRequestPageParam) GetPageSize() *int32 {
	return s.PageSize
}

func (s *RunSearchLawQueryRequestPageParam) SetPageNumber(v int32) *RunSearchLawQueryRequestPageParam {
	s.PageNumber = &v
	return s
}

func (s *RunSearchLawQueryRequestPageParam) SetPageSize(v int32) *RunSearchLawQueryRequestPageParam {
	s.PageSize = &v
	return s
}

func (s *RunSearchLawQueryRequestPageParam) Validate() error {
	return dara.Validate(s)
}

type RunSearchLawQueryRequestThread struct {
	Messages []*RunSearchLawQueryRequestThreadMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
}

func (s RunSearchLawQueryRequestThread) String() string {
	return dara.Prettify(s)
}

func (s RunSearchLawQueryRequestThread) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryRequestThread) GetMessages() []*RunSearchLawQueryRequestThreadMessages {
	return s.Messages
}

func (s *RunSearchLawQueryRequestThread) SetMessages(v []*RunSearchLawQueryRequestThreadMessages) *RunSearchLawQueryRequestThread {
	s.Messages = v
	return s
}

func (s *RunSearchLawQueryRequestThread) Validate() error {
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

type RunSearchLawQueryRequestThreadMessages struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s RunSearchLawQueryRequestThreadMessages) String() string {
	return dara.Prettify(s)
}

func (s RunSearchLawQueryRequestThreadMessages) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryRequestThreadMessages) GetContent() *string {
	return s.Content
}

func (s *RunSearchLawQueryRequestThreadMessages) GetRole() *string {
	return s.Role
}

func (s *RunSearchLawQueryRequestThreadMessages) SetContent(v string) *RunSearchLawQueryRequestThreadMessages {
	s.Content = &v
	return s
}

func (s *RunSearchLawQueryRequestThreadMessages) SetRole(v string) *RunSearchLawQueryRequestThreadMessages {
	s.Role = &v
	return s
}

func (s *RunSearchLawQueryRequestThreadMessages) Validate() error {
	return dara.Validate(s)
}
