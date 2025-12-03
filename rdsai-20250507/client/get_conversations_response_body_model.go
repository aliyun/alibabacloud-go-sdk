// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConversationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetConversationsResponseBodyData) *GetConversationsResponseBody
	GetData() []*GetConversationsResponseBodyData
	SetHasMore(v string) *GetConversationsResponseBody
	GetHasMore() *string
	SetLimit(v int64) *GetConversationsResponseBody
	GetLimit() *int64
	SetRequestId(v string) *GetConversationsResponseBody
	GetRequestId() *string
}

type GetConversationsResponseBody struct {
	Data []*GetConversationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// true
	HasMore *string `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// example:
	//
	// 100
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetConversationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConversationsResponseBody) GoString() string {
	return s.String()
}

func (s *GetConversationsResponseBody) GetData() []*GetConversationsResponseBodyData {
	return s.Data
}

func (s *GetConversationsResponseBody) GetHasMore() *string {
	return s.HasMore
}

func (s *GetConversationsResponseBody) GetLimit() *int64 {
	return s.Limit
}

func (s *GetConversationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConversationsResponseBody) SetData(v []*GetConversationsResponseBodyData) *GetConversationsResponseBody {
	s.Data = v
	return s
}

func (s *GetConversationsResponseBody) SetHasMore(v string) *GetConversationsResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetConversationsResponseBody) SetLimit(v int64) *GetConversationsResponseBody {
	s.Limit = &v
	return s
}

func (s *GetConversationsResponseBody) SetRequestId(v string) *GetConversationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConversationsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetConversationsResponseBodyData struct {
	// example:
	//
	// 1764055092
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// 60b335ca-124d-4ee1-864b-de554987****
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetConversationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetConversationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetConversationsResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetConversationsResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetConversationsResponseBodyData) GetIntroduction() *string {
	return s.Introduction
}

func (s *GetConversationsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetConversationsResponseBodyData) SetCreatedAt(v string) *GetConversationsResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetConversationsResponseBodyData) SetId(v string) *GetConversationsResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetConversationsResponseBodyData) SetIntroduction(v string) *GetConversationsResponseBodyData {
	s.Introduction = &v
	return s
}

func (s *GetConversationsResponseBodyData) SetName(v string) *GetConversationsResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetConversationsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
