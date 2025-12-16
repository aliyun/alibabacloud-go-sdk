// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeABTestGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeABTestGroupResponseBody
	GetRequestId() *string
	SetResult(v *DescribeABTestGroupResponseBodyResult) *DescribeABTestGroupResponseBody
	GetResult() *DescribeABTestGroupResponseBodyResult
}

type DescribeABTestGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the test group.
	Result *DescribeABTestGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeABTestGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeABTestGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeABTestGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeABTestGroupResponseBody) GetResult() *DescribeABTestGroupResponseBodyResult {
	return s.Result
}

func (s *DescribeABTestGroupResponseBody) SetRequestId(v string) *DescribeABTestGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeABTestGroupResponseBody) SetResult(v *DescribeABTestGroupResponseBodyResult) *DescribeABTestGroupResponseBody {
	s.Result = v
	return s
}

func (s *DescribeABTestGroupResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeABTestGroupResponseBodyResult struct {
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
	// The alias of the test group.
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
	// The time when the test group was last updated.
	//
	// example:
	//
	// 1588839490
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeABTestGroupResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeABTestGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeABTestGroupResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *DescribeABTestGroupResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *DescribeABTestGroupResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *DescribeABTestGroupResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeABTestGroupResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *DescribeABTestGroupResponseBodyResult) SetCreated(v int32) *DescribeABTestGroupResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeABTestGroupResponseBodyResult) SetId(v string) *DescribeABTestGroupResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeABTestGroupResponseBodyResult) SetName(v string) *DescribeABTestGroupResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeABTestGroupResponseBodyResult) SetStatus(v int32) *DescribeABTestGroupResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeABTestGroupResponseBodyResult) SetUpdated(v int32) *DescribeABTestGroupResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *DescribeABTestGroupResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
