// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeABTestExperimentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeABTestExperimentResponseBody
	GetRequestId() *string
	SetResult(v *DescribeABTestExperimentResponseBodyResult) *DescribeABTestExperimentResponseBody
	GetResult() *DescribeABTestExperimentResponseBodyResult
}

type DescribeABTestExperimentResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the test.
	Result *DescribeABTestExperimentResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeABTestExperimentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeABTestExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeABTestExperimentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeABTestExperimentResponseBody) GetResult() *DescribeABTestExperimentResponseBodyResult {
	return s.Result
}

func (s *DescribeABTestExperimentResponseBody) SetRequestId(v string) *DescribeABTestExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeABTestExperimentResponseBody) SetResult(v *DescribeABTestExperimentResponseBodyResult) *DescribeABTestExperimentResponseBody {
	s.Result = v
	return s
}

func (s *DescribeABTestExperimentResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeABTestExperimentResponseBodyResult struct {
	// The time when the test was created.
	//
	// example:
	//
	// 1588842080
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test.
	//
	// example:
	//
	// 12888
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the test.
	//
	// example:
	//
	// test1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test. Valid values:
	//
	// 	- true: in effect
	//
	// 	- false: not in effect
	//
	// example:
	//
	// true
	Online *bool `json:"online,omitempty" xml:"online,omitempty"`
	// The parameters of the test.
	Params *DescribeABTestExperimentResponseBodyResultParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	// The percentage of traffic that is routed to the test.
	//
	// example:
	//
	// 30
	Traffic *int32 `json:"traffic,omitempty" xml:"traffic,omitempty"`
	// The time when the test was last modified.
	//
	// example:
	//
	// 1588842080
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeABTestExperimentResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeABTestExperimentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeABTestExperimentResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *DescribeABTestExperimentResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *DescribeABTestExperimentResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *DescribeABTestExperimentResponseBodyResult) GetOnline() *bool {
	return s.Online
}

func (s *DescribeABTestExperimentResponseBodyResult) GetParams() *DescribeABTestExperimentResponseBodyResultParams {
	return s.Params
}

func (s *DescribeABTestExperimentResponseBodyResult) GetTraffic() *int32 {
	return s.Traffic
}

func (s *DescribeABTestExperimentResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *DescribeABTestExperimentResponseBodyResult) SetCreated(v int32) *DescribeABTestExperimentResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeABTestExperimentResponseBodyResult) SetId(v string) *DescribeABTestExperimentResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeABTestExperimentResponseBodyResult) SetName(v string) *DescribeABTestExperimentResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeABTestExperimentResponseBodyResult) SetOnline(v bool) *DescribeABTestExperimentResponseBodyResult {
	s.Online = &v
	return s
}

func (s *DescribeABTestExperimentResponseBodyResult) SetParams(v *DescribeABTestExperimentResponseBodyResultParams) *DescribeABTestExperimentResponseBodyResult {
	s.Params = v
	return s
}

func (s *DescribeABTestExperimentResponseBodyResult) SetTraffic(v int32) *DescribeABTestExperimentResponseBodyResult {
	s.Traffic = &v
	return s
}

func (s *DescribeABTestExperimentResponseBodyResult) SetUpdated(v int32) *DescribeABTestExperimentResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *DescribeABTestExperimentResponseBodyResult) Validate() error {
	if s.Params != nil {
		if err := s.Params.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeABTestExperimentResponseBodyResultParams struct {
	// The name of the rough sort policy.
	//
	// example:
	//
	// default
	FirstFormulaName *string `json:"first_formula_name,omitempty" xml:"first_formula_name,omitempty"`
}

func (s DescribeABTestExperimentResponseBodyResultParams) String() string {
	return dara.Prettify(s)
}

func (s DescribeABTestExperimentResponseBodyResultParams) GoString() string {
	return s.String()
}

func (s *DescribeABTestExperimentResponseBodyResultParams) GetFirstFormulaName() *string {
	return s.FirstFormulaName
}

func (s *DescribeABTestExperimentResponseBodyResultParams) SetFirstFormulaName(v string) *DescribeABTestExperimentResponseBodyResultParams {
	s.FirstFormulaName = &v
	return s
}

func (s *DescribeABTestExperimentResponseBodyResultParams) Validate() error {
	return dara.Validate(s)
}
