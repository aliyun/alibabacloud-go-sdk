// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iComparePlaybooksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompareResult(v *ComparePlaybooksResponseBodyCompareResult) *ComparePlaybooksResponseBody
	GetCompareResult() *ComparePlaybooksResponseBodyCompareResult
	SetRequestId(v string) *ComparePlaybooksResponseBody
	GetRequestId() *string
}

type ComparePlaybooksResponseBody struct {
	// The comparison result.
	CompareResult *ComparePlaybooksResponseBodyCompareResult `json:"CompareResult,omitempty" xml:"CompareResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2EC05B06-BF3C-5F3E-8FE8-3B1FAD76087A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ComparePlaybooksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ComparePlaybooksResponseBody) GoString() string {
	return s.String()
}

func (s *ComparePlaybooksResponseBody) GetCompareResult() *ComparePlaybooksResponseBodyCompareResult {
	return s.CompareResult
}

func (s *ComparePlaybooksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ComparePlaybooksResponseBody) SetCompareResult(v *ComparePlaybooksResponseBodyCompareResult) *ComparePlaybooksResponseBody {
	s.CompareResult = v
	return s
}

func (s *ComparePlaybooksResponseBody) SetRequestId(v string) *ComparePlaybooksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ComparePlaybooksResponseBody) Validate() error {
	if s.CompareResult != nil {
		if err := s.CompareResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ComparePlaybooksResponseBodyCompareResult struct {
	// The description of the comparison result.
	//
	// example:
	//
	// The first version adds one node compared to the second version
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the second version provides more information than the first version. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	New *bool `json:"New,omitempty" xml:"New,omitempty"`
	// Indicates whether the configurations of the two versions are the same. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Same *bool `json:"Same,omitempty" xml:"Same,omitempty"`
}

func (s ComparePlaybooksResponseBodyCompareResult) String() string {
	return dara.Prettify(s)
}

func (s ComparePlaybooksResponseBodyCompareResult) GoString() string {
	return s.String()
}

func (s *ComparePlaybooksResponseBodyCompareResult) GetDescription() *string {
	return s.Description
}

func (s *ComparePlaybooksResponseBodyCompareResult) GetNew() *bool {
	return s.New
}

func (s *ComparePlaybooksResponseBodyCompareResult) GetSame() *bool {
	return s.Same
}

func (s *ComparePlaybooksResponseBodyCompareResult) SetDescription(v string) *ComparePlaybooksResponseBodyCompareResult {
	s.Description = &v
	return s
}

func (s *ComparePlaybooksResponseBodyCompareResult) SetNew(v bool) *ComparePlaybooksResponseBodyCompareResult {
	s.New = &v
	return s
}

func (s *ComparePlaybooksResponseBodyCompareResult) SetSame(v bool) *ComparePlaybooksResponseBodyCompareResult {
	s.Same = &v
	return s
}

func (s *ComparePlaybooksResponseBodyCompareResult) Validate() error {
	return dara.Validate(s)
}
