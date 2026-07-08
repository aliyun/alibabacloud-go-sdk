// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotNewsWithTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListHotNewsWithTypeRequest
	GetAgentKey() *string
	SetCurrent(v int32) *ListHotNewsWithTypeRequest
	GetCurrent() *int32
	SetNewsType(v string) *ListHotNewsWithTypeRequest
	GetNewsType() *string
	SetNewsTypes(v []*string) *ListHotNewsWithTypeRequest
	GetNewsTypes() []*string
	SetSize(v int32) *ListHotNewsWithTypeRequest
	GetSize() *int32
}

type ListHotNewsWithTypeRequest struct {
	// The unique identifier of the business space: [AgentKey](https://help.aliyun.com/document_detail/3027170.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// c160c841c8e54295bf2f441432785944_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// The news category. Valid values:
	//
	// - society: social current affairs.
	//
	// - person: people news.
	//
	// - government: government affairs.
	//
	// example:
	//
	// society
	NewsType *string `json:"NewsType,omitempty" xml:"NewsType,omitempty"`
	// The list of news categories for multi-selection. Valid values:
	//
	// - society: social current affairs.
	//
	// - person: people news.
	//
	// - government: government affairs.
	NewsTypes []*string `json:"NewsTypes,omitempty" xml:"NewsTypes,omitempty" type:"Repeated"`
	// The number of records per page.
	//
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListHotNewsWithTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotNewsWithTypeRequest) GoString() string {
	return s.String()
}

func (s *ListHotNewsWithTypeRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListHotNewsWithTypeRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListHotNewsWithTypeRequest) GetNewsType() *string {
	return s.NewsType
}

func (s *ListHotNewsWithTypeRequest) GetNewsTypes() []*string {
	return s.NewsTypes
}

func (s *ListHotNewsWithTypeRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListHotNewsWithTypeRequest) SetAgentKey(v string) *ListHotNewsWithTypeRequest {
	s.AgentKey = &v
	return s
}

func (s *ListHotNewsWithTypeRequest) SetCurrent(v int32) *ListHotNewsWithTypeRequest {
	s.Current = &v
	return s
}

func (s *ListHotNewsWithTypeRequest) SetNewsType(v string) *ListHotNewsWithTypeRequest {
	s.NewsType = &v
	return s
}

func (s *ListHotNewsWithTypeRequest) SetNewsTypes(v []*string) *ListHotNewsWithTypeRequest {
	s.NewsTypes = v
	return s
}

func (s *ListHotNewsWithTypeRequest) SetSize(v int32) *ListHotNewsWithTypeRequest {
	s.Size = &v
	return s
}

func (s *ListHotNewsWithTypeRequest) Validate() error {
	return dara.Validate(s)
}
