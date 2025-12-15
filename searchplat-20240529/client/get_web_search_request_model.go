// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWebSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *GetWebSearchRequest
	GetContentType() *string
	SetHistory(v []*GetWebSearchRequestHistory) *GetWebSearchRequest
	GetHistory() []*GetWebSearchRequestHistory
	SetQuery(v string) *GetWebSearchRequest
	GetQuery() *string
	SetQueryRewrite(v bool) *GetWebSearchRequest
	GetQueryRewrite() *bool
	SetTopK(v int64) *GetWebSearchRequest
	GetTopK() *int64
	SetWay(v string) *GetWebSearchRequest
	GetWay() *string
}

type GetWebSearchRequest struct {
	ContentType *string                       `json:"content_type,omitempty" xml:"content_type,omitempty"`
	History     []*GetWebSearchRequestHistory `json:"history,omitempty" xml:"history,omitempty" type:"Repeated"`
	// This parameter is required.
	Query        *string `json:"query,omitempty" xml:"query,omitempty"`
	QueryRewrite *bool   `json:"query_rewrite,omitempty" xml:"query_rewrite,omitempty"`
	TopK         *int64  `json:"top_k,omitempty" xml:"top_k,omitempty"`
	Way          *string `json:"way,omitempty" xml:"way,omitempty"`
}

func (s GetWebSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWebSearchRequest) GoString() string {
	return s.String()
}

func (s *GetWebSearchRequest) GetContentType() *string {
	return s.ContentType
}

func (s *GetWebSearchRequest) GetHistory() []*GetWebSearchRequestHistory {
	return s.History
}

func (s *GetWebSearchRequest) GetQuery() *string {
	return s.Query
}

func (s *GetWebSearchRequest) GetQueryRewrite() *bool {
	return s.QueryRewrite
}

func (s *GetWebSearchRequest) GetTopK() *int64 {
	return s.TopK
}

func (s *GetWebSearchRequest) GetWay() *string {
	return s.Way
}

func (s *GetWebSearchRequest) SetContentType(v string) *GetWebSearchRequest {
	s.ContentType = &v
	return s
}

func (s *GetWebSearchRequest) SetHistory(v []*GetWebSearchRequestHistory) *GetWebSearchRequest {
	s.History = v
	return s
}

func (s *GetWebSearchRequest) SetQuery(v string) *GetWebSearchRequest {
	s.Query = &v
	return s
}

func (s *GetWebSearchRequest) SetQueryRewrite(v bool) *GetWebSearchRequest {
	s.QueryRewrite = &v
	return s
}

func (s *GetWebSearchRequest) SetTopK(v int64) *GetWebSearchRequest {
	s.TopK = &v
	return s
}

func (s *GetWebSearchRequest) SetWay(v string) *GetWebSearchRequest {
	s.Way = &v
	return s
}

func (s *GetWebSearchRequest) Validate() error {
	if s.History != nil {
		for _, item := range s.History {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWebSearchRequestHistory struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Role    *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s GetWebSearchRequestHistory) String() string {
	return dara.Prettify(s)
}

func (s GetWebSearchRequestHistory) GoString() string {
	return s.String()
}

func (s *GetWebSearchRequestHistory) GetContent() *string {
	return s.Content
}

func (s *GetWebSearchRequestHistory) GetRole() *string {
	return s.Role
}

func (s *GetWebSearchRequestHistory) SetContent(v string) *GetWebSearchRequestHistory {
	s.Content = &v
	return s
}

func (s *GetWebSearchRequestHistory) SetRole(v string) *GetWebSearchRequestHistory {
	s.Role = &v
	return s
}

func (s *GetWebSearchRequestHistory) Validate() error {
	return dara.Validate(s)
}
