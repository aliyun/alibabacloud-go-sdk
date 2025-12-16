// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListABTestGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListABTestGroupsResponseBody
	GetRequestId() *string
	SetResult(v []*ListABTestGroupsResponseBodyResult) *ListABTestGroupsResponseBody
	GetResult() []*ListABTestGroupsResponseBodyResult
}

type ListABTestGroupsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The test groups.
	//
	// For more information, see [ABTestGroup](https://help.aliyun.com/document_detail/178935.html).
	Result []*ListABTestGroupsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListABTestGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListABTestGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListABTestGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListABTestGroupsResponseBody) GetResult() []*ListABTestGroupsResponseBodyResult {
	return s.Result
}

func (s *ListABTestGroupsResponseBody) SetRequestId(v string) *ListABTestGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListABTestGroupsResponseBody) SetResult(v []*ListABTestGroupsResponseBodyResult) *ListABTestGroupsResponseBody {
	s.Result = v
	return s
}

func (s *ListABTestGroupsResponseBody) Validate() error {
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

type ListABTestGroupsResponseBodyResult struct {
	// The time when the test group was created.
	//
	// example:
	//
	// 1588839490
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test group.
	//
	// example:
	//
	// 13466
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the test group.
	//
	// example:
	//
	// Group_2020-5-7_15:23:3
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test group. Valid values:
	//
	// 	- 0: not in effect
	//
	// 	- 1: in effect
	//
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the test group was last modified.
	//
	// example:
	//
	// 1588839490
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListABTestGroupsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListABTestGroupsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListABTestGroupsResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *ListABTestGroupsResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ListABTestGroupsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListABTestGroupsResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *ListABTestGroupsResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *ListABTestGroupsResponseBodyResult) SetCreated(v int32) *ListABTestGroupsResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListABTestGroupsResponseBodyResult) SetId(v string) *ListABTestGroupsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListABTestGroupsResponseBodyResult) SetName(v string) *ListABTestGroupsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListABTestGroupsResponseBodyResult) SetStatus(v int32) *ListABTestGroupsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListABTestGroupsResponseBodyResult) SetUpdated(v int32) *ListABTestGroupsResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListABTestGroupsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
