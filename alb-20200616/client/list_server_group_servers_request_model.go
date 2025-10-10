// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServerGroupServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServerGroupServersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServerGroupServersRequest
	GetNextToken() *string
	SetServerGroupId(v string) *ListServerGroupServersRequest
	GetServerGroupId() *string
	SetServerIds(v []*string) *ListServerGroupServersRequest
	GetServerIds() []*string
	SetTag(v []*ListServerGroupServersRequestTag) *ListServerGroupServersRequest
	GetTag() []*ListServerGroupServersRequestTag
}

type ListServerGroupServersRequest struct {
	// The maximum number of entries to return. Valid values: **1*	- to **100**. If you do not specify a value, the default value **20*	- is used.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of **NextToken**.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXG****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The server group ID.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The IDs of the servers.
	ServerIds []*string `json:"ServerIds,omitempty" xml:"ServerIds,omitempty" type:"Repeated"`
	// The tags that are added to the server group. You can specify up to 10 tags in each call.
	Tag []*ListServerGroupServersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListServerGroupServersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServerGroupServersRequest) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServerGroupServersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServerGroupServersRequest) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *ListServerGroupServersRequest) GetServerIds() []*string {
	return s.ServerIds
}

func (s *ListServerGroupServersRequest) GetTag() []*ListServerGroupServersRequestTag {
	return s.Tag
}

func (s *ListServerGroupServersRequest) SetMaxResults(v int32) *ListServerGroupServersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupServersRequest) SetNextToken(v string) *ListServerGroupServersRequest {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupServersRequest) SetServerGroupId(v string) *ListServerGroupServersRequest {
	s.ServerGroupId = &v
	return s
}

func (s *ListServerGroupServersRequest) SetServerIds(v []*string) *ListServerGroupServersRequest {
	s.ServerIds = v
	return s
}

func (s *ListServerGroupServersRequest) SetTag(v []*ListServerGroupServersRequestTag) *ListServerGroupServersRequest {
	s.Tag = v
	return s
}

func (s *ListServerGroupServersRequest) Validate() error {
	return dara.Validate(s)
}

type ListServerGroupServersRequestTag struct {
	// The tag key. You can specify up to 10 tag keys.
	//
	// The tag key can be up to 64 characters in length, and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// Test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify up to 10 tag values.
	//
	// The tag value can be up to 128 characters in length, and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// Test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServerGroupServersRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListServerGroupServersRequestTag) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListServerGroupServersRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListServerGroupServersRequestTag) SetKey(v string) *ListServerGroupServersRequestTag {
	s.Key = &v
	return s
}

func (s *ListServerGroupServersRequestTag) SetValue(v string) *ListServerGroupServersRequestTag {
	s.Value = &v
	return s
}

func (s *ListServerGroupServersRequestTag) Validate() error {
	return dara.Validate(s)
}
