// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListABTestScenesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListABTestScenesResponseBody
	GetRequestId() *string
	SetResult(v []*ListABTestScenesResponseBodyResult) *ListABTestScenesResponseBody
	GetResult() []*ListABTestScenesResponseBodyResult
}

type ListABTestScenesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the test scenario.
	//
	// For more information, see [ABTestScene](https://help.aliyun.com/document_detail/173618.html).
	Result []*ListABTestScenesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListABTestScenesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListABTestScenesResponseBody) GoString() string {
	return s.String()
}

func (s *ListABTestScenesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListABTestScenesResponseBody) GetResult() []*ListABTestScenesResponseBodyResult {
	return s.Result
}

func (s *ListABTestScenesResponseBody) SetRequestId(v string) *ListABTestScenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListABTestScenesResponseBody) SetResult(v []*ListABTestScenesResponseBodyResult) *ListABTestScenesResponseBody {
	s.Result = v
	return s
}

func (s *ListABTestScenesResponseBody) Validate() error {
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

type ListABTestScenesResponseBodyResult struct {
	// The time when the test group was created.
	//
	// example:
	//
	// 1588836130
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test group.
	//
	// example:
	//
	// 20404
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The alias of the test group.
	//
	// example:
	//
	// kevintest_2020-5-7_15:21:482
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
	// The time when the test group was last updated.
	//
	// example:
	//
	// 1588836129
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
	// The name of the test scenario.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ListABTestScenesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListABTestScenesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListABTestScenesResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *ListABTestScenesResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ListABTestScenesResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListABTestScenesResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *ListABTestScenesResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *ListABTestScenesResponseBodyResult) GetValues() []*string {
	return s.Values
}

func (s *ListABTestScenesResponseBodyResult) SetCreated(v int32) *ListABTestScenesResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListABTestScenesResponseBodyResult) SetId(v string) *ListABTestScenesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListABTestScenesResponseBodyResult) SetName(v string) *ListABTestScenesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListABTestScenesResponseBodyResult) SetStatus(v int32) *ListABTestScenesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListABTestScenesResponseBodyResult) SetUpdated(v int32) *ListABTestScenesResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListABTestScenesResponseBodyResult) SetValues(v []*string) *ListABTestScenesResponseBodyResult {
	s.Values = v
	return s
}

func (s *ListABTestScenesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
