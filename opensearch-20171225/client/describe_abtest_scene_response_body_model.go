// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeABTestSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeABTestSceneResponseBody
	GetRequestId() *string
	SetResult(v *DescribeABTestSceneResponseBodyResult) *DescribeABTestSceneResponseBody
	GetResult() *DescribeABTestSceneResponseBodyResult
}

type DescribeABTestSceneResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the test scenario.
	Result *DescribeABTestSceneResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeABTestSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeABTestSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeABTestSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeABTestSceneResponseBody) GetResult() *DescribeABTestSceneResponseBodyResult {
	return s.Result
}

func (s *DescribeABTestSceneResponseBody) SetRequestId(v string) *DescribeABTestSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeABTestSceneResponseBody) SetResult(v *DescribeABTestSceneResponseBodyResult) *DescribeABTestSceneResponseBody {
	s.Result = v
	return s
}

func (s *DescribeABTestSceneResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeABTestSceneResponseBodyResult struct {
	// The time when the test scenario was created.
	//
	// example:
	//
	// 1596527691
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test scenario.
	//
	// example:
	//
	// 20614
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the test scenario.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test scenario. Valid values:
	//
	// 	- 0: The test is stopped.
	//
	// 	- 1: The test is started.
	//
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the test was last modified.
	//
	// example:
	//
	// 1596527691
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
	// The indicators of the test scenarios.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s DescribeABTestSceneResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeABTestSceneResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeABTestSceneResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *DescribeABTestSceneResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *DescribeABTestSceneResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *DescribeABTestSceneResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeABTestSceneResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *DescribeABTestSceneResponseBodyResult) GetValues() []*string {
	return s.Values
}

func (s *DescribeABTestSceneResponseBodyResult) SetCreated(v int32) *DescribeABTestSceneResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeABTestSceneResponseBodyResult) SetId(v string) *DescribeABTestSceneResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeABTestSceneResponseBodyResult) SetName(v string) *DescribeABTestSceneResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeABTestSceneResponseBodyResult) SetStatus(v int32) *DescribeABTestSceneResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeABTestSceneResponseBodyResult) SetUpdated(v int32) *DescribeABTestSceneResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *DescribeABTestSceneResponseBodyResult) SetValues(v []*string) *DescribeABTestSceneResponseBodyResult {
	s.Values = v
	return s
}

func (s *DescribeABTestSceneResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
