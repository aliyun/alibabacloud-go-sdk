// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateABTestGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateABTestGroupResponseBody
	GetRequestId() *string
	SetResult(v *CreateABTestGroupResponseBodyResult) *CreateABTestGroupResponseBody
	GetResult() *CreateABTestGroupResponseBodyResult
}

type CreateABTestGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned results.
	Result *CreateABTestGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateABTestGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateABTestGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateABTestGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateABTestGroupResponseBody) GetResult() *CreateABTestGroupResponseBodyResult {
	return s.Result
}

func (s *CreateABTestGroupResponseBody) SetRequestId(v string) *CreateABTestGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateABTestGroupResponseBody) SetResult(v *CreateABTestGroupResponseBodyResult) *CreateABTestGroupResponseBody {
	s.Result = v
	return s
}

func (s *CreateABTestGroupResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateABTestGroupResponseBodyResult struct {
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
	// The status of the test group.
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

func (s CreateABTestGroupResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateABTestGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateABTestGroupResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *CreateABTestGroupResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *CreateABTestGroupResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *CreateABTestGroupResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *CreateABTestGroupResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *CreateABTestGroupResponseBodyResult) SetCreated(v int32) *CreateABTestGroupResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateABTestGroupResponseBodyResult) SetId(v string) *CreateABTestGroupResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateABTestGroupResponseBodyResult) SetName(v string) *CreateABTestGroupResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateABTestGroupResponseBodyResult) SetStatus(v int32) *CreateABTestGroupResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CreateABTestGroupResponseBodyResult) SetUpdated(v int32) *CreateABTestGroupResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *CreateABTestGroupResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
