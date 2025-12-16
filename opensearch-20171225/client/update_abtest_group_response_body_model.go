// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateABTestGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateABTestGroupResponseBody
	GetRequestId() *string
	SetResult(v *UpdateABTestGroupResponseBodyResult) *UpdateABTestGroupResponseBody
	GetResult() *UpdateABTestGroupResponseBodyResult
}

type UpdateABTestGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// \\"\\"1111\\"\\"
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the test group.
	Result *UpdateABTestGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateABTestGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateABTestGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateABTestGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateABTestGroupResponseBody) GetResult() *UpdateABTestGroupResponseBodyResult {
	return s.Result
}

func (s *UpdateABTestGroupResponseBody) SetRequestId(v string) *UpdateABTestGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateABTestGroupResponseBody) SetResult(v *UpdateABTestGroupResponseBodyResult) *UpdateABTestGroupResponseBody {
	s.Result = v
	return s
}

func (s *UpdateABTestGroupResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateABTestGroupResponseBodyResult struct {
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

func (s UpdateABTestGroupResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateABTestGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateABTestGroupResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *UpdateABTestGroupResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *UpdateABTestGroupResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *UpdateABTestGroupResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *UpdateABTestGroupResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *UpdateABTestGroupResponseBodyResult) SetCreated(v int32) *UpdateABTestGroupResponseBodyResult {
	s.Created = &v
	return s
}

func (s *UpdateABTestGroupResponseBodyResult) SetId(v string) *UpdateABTestGroupResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateABTestGroupResponseBodyResult) SetName(v string) *UpdateABTestGroupResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateABTestGroupResponseBodyResult) SetStatus(v int32) *UpdateABTestGroupResponseBodyResult {
	s.Status = &v
	return s
}

func (s *UpdateABTestGroupResponseBodyResult) SetUpdated(v int32) *UpdateABTestGroupResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *UpdateABTestGroupResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
