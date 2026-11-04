// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceUserByNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsolationId(v string) *QueryFaceUserByNameRequest
	GetIsolationId() *string
	SetName(v string) *QueryFaceUserByNameRequest
	GetName() *string
	SetPageNo(v int32) *QueryFaceUserByNameRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryFaceUserByNameRequest
	GetPageSize() *int32
	SetParams(v string) *QueryFaceUserByNameRequest
	GetParams() *string
}

type QueryFaceUserByNameRequest struct {
	// The business isolation ID, which is used to isolate data between different business applications under the same Alibaba Cloud account. Make sure that this parameter is consistent with the **IsolationId*	- set when you call the [AddFaceUser](https://help.aliyun.com/document_detail/288338.html) operation to add a user.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZheJiangHZ
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	// The username. The name can contain Chinese characters, uppercase and lowercase letters, digits, underscores (_), and hyphens (-), and can be up to 64 characters in length.
	//
	// Fuzzy match is supported. If this parameter is not specified, all records are matched.
	//
	// example:
	//
	// Tony
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number of the page to return. Minimum value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page. Maximum value: 50. Minimum value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The custom user parameters.
	//
	// Exact match is supported, but fuzzy match is not supported.
	//
	// example:
	//
	// {"key1":"value1", "key2":"value2"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s QueryFaceUserByNameRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceUserByNameRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceUserByNameRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *QueryFaceUserByNameRequest) GetName() *string {
	return s.Name
}

func (s *QueryFaceUserByNameRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryFaceUserByNameRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryFaceUserByNameRequest) GetParams() *string {
	return s.Params
}

func (s *QueryFaceUserByNameRequest) SetIsolationId(v string) *QueryFaceUserByNameRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceUserByNameRequest) SetName(v string) *QueryFaceUserByNameRequest {
	s.Name = &v
	return s
}

func (s *QueryFaceUserByNameRequest) SetPageNo(v int32) *QueryFaceUserByNameRequest {
	s.PageNo = &v
	return s
}

func (s *QueryFaceUserByNameRequest) SetPageSize(v int32) *QueryFaceUserByNameRequest {
	s.PageSize = &v
	return s
}

func (s *QueryFaceUserByNameRequest) SetParams(v string) *QueryFaceUserByNameRequest {
	s.Params = &v
	return s
}

func (s *QueryFaceUserByNameRequest) Validate() error {
	return dara.Validate(s)
}
