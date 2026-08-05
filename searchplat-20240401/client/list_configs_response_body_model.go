// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v int32) *ListConfigsResponseBody
	GetPage() *int32
	SetPageSize(v int32) *ListConfigsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListConfigsResponseBody
	GetRequestId() *string
	SetResult(v []*ListConfigsResponseBodyResult) *ListConfigsResponseBody
	GetResult() []*ListConfigsResponseBodyResult
	SetTotal(v int32) *ListConfigsResponseBody
	GetTotal() *int32
}

type ListConfigsResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 65150BD6-1622-4177-9D30-65B33A9F6969
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The configuration list.
	Result []*ListConfigsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of configurations.
	//
	// example:
	//
	// 1
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigsResponseBody) GetPage() *int32 {
	return s.Page
}

func (s *ListConfigsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConfigsResponseBody) GetResult() []*ListConfigsResponseBodyResult {
	return s.Result
}

func (s *ListConfigsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListConfigsResponseBody) SetPage(v int32) *ListConfigsResponseBody {
	s.Page = &v
	return s
}

func (s *ListConfigsResponseBody) SetPageSize(v int32) *ListConfigsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListConfigsResponseBody) SetRequestId(v string) *ListConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConfigsResponseBody) SetResult(v []*ListConfigsResponseBodyResult) *ListConfigsResponseBody {
	s.Result = v
	return s
}

func (s *ListConfigsResponseBody) SetTotal(v int32) *ListConfigsResponseBody {
	s.Total = &v
	return s
}

func (s *ListConfigsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListConfigsResponseBodyResult struct {
	// The configuration content.
	ConfigData map[string]interface{} `json:"configData,omitempty" xml:"configData,omitempty"`
	// The configuration type. Valid values:
	//
	//  	- prompt: Prompt configuration.
	//
	//  	- lark: Lark configuration.
	//
	// example:
	//
	// prompt
	ConfigType *string `json:"configType,omitempty" xml:"configType,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-001
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListConfigsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListConfigsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListConfigsResponseBodyResult) GetConfigData() map[string]interface{} {
	return s.ConfigData
}

func (s *ListConfigsResponseBodyResult) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListConfigsResponseBodyResult) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListConfigsResponseBodyResult) SetConfigData(v map[string]interface{}) *ListConfigsResponseBodyResult {
	s.ConfigData = v
	return s
}

func (s *ListConfigsResponseBodyResult) SetConfigType(v string) *ListConfigsResponseBodyResult {
	s.ConfigType = &v
	return s
}

func (s *ListConfigsResponseBodyResult) SetWorkspaceId(v string) *ListConfigsResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *ListConfigsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
