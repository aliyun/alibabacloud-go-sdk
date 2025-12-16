// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateABTestSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateABTestSceneResponseBody
	GetRequestId() *string
	SetResult(v *CreateABTestSceneResponseBodyResult) *CreateABTestSceneResponseBody
	GetResult() *CreateABTestSceneResponseBodyResult
}

type CreateABTestSceneResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned data.
	Result *CreateABTestSceneResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateABTestSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateABTestSceneResponseBody) GoString() string {
	return s.String()
}

func (s *CreateABTestSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateABTestSceneResponseBody) GetResult() *CreateABTestSceneResponseBodyResult {
	return s.Result
}

func (s *CreateABTestSceneResponseBody) SetRequestId(v string) *CreateABTestSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateABTestSceneResponseBody) SetResult(v *CreateABTestSceneResponseBodyResult) *CreateABTestSceneResponseBody {
	s.Result = v
	return s
}

func (s *CreateABTestSceneResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateABTestSceneResponseBodyResult struct {
	// The time when the test scenario was created.
	//
	// example:
	//
	// 0
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test group.
	//
	// example:
	//
	// 20405
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the A/B test group.
	//
	// example:
	//
	// kevintest_2020-5-7_15:21:48
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status. Valid values:
	//
	// 	- 0: not in effect
	//
	// 	- 1: in effect
	//
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the test scenario was last updated.
	//
	// example:
	//
	// 1589012351
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
	// The ID of the test scenario
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s CreateABTestSceneResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateABTestSceneResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateABTestSceneResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *CreateABTestSceneResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *CreateABTestSceneResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *CreateABTestSceneResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *CreateABTestSceneResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *CreateABTestSceneResponseBodyResult) GetValues() []*string {
	return s.Values
}

func (s *CreateABTestSceneResponseBodyResult) SetCreated(v int32) *CreateABTestSceneResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateABTestSceneResponseBodyResult) SetId(v string) *CreateABTestSceneResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateABTestSceneResponseBodyResult) SetName(v string) *CreateABTestSceneResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateABTestSceneResponseBodyResult) SetStatus(v int32) *CreateABTestSceneResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CreateABTestSceneResponseBodyResult) SetUpdated(v int32) *CreateABTestSceneResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *CreateABTestSceneResponseBodyResult) SetValues(v []*string) *CreateABTestSceneResponseBodyResult {
	s.Values = v
	return s
}

func (s *CreateABTestSceneResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
