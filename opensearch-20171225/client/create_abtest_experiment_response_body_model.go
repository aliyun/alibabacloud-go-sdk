// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateABTestExperimentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateABTestExperimentResponseBody
	GetRequestId() *string
	SetResult(v *CreateABTestExperimentResponseBodyResult) *CreateABTestExperimentResponseBody
	GetResult() *CreateABTestExperimentResponseBodyResult
}

type CreateABTestExperimentResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The experiment details.
	Result *CreateABTestExperimentResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateABTestExperimentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateABTestExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateABTestExperimentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateABTestExperimentResponseBody) GetResult() *CreateABTestExperimentResponseBodyResult {
	return s.Result
}

func (s *CreateABTestExperimentResponseBody) SetRequestId(v string) *CreateABTestExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateABTestExperimentResponseBody) SetResult(v *CreateABTestExperimentResponseBodyResult) *CreateABTestExperimentResponseBody {
	s.Result = v
	return s
}

func (s *CreateABTestExperimentResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateABTestExperimentResponseBodyResult struct {
	// The time when the experiment was created.
	//
	// example:
	//
	// 0
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The experiment ID.
	//
	// example:
	//
	// 12889
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The experiment alias.
	//
	// example:
	//
	// test3
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Indicates whether the experiment is in effect. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Online *bool `json:"online,omitempty" xml:"online,omitempty"`
	// The experiment parameters.
	//
	// example:
	//
	// {"firstFormulaName": "default"}
	Params map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// The percentage of traffic that is routed to the experiment.
	//
	// example:
	//
	// 30
	Traffic *int32 `json:"traffic,omitempty" xml:"traffic,omitempty"`
	// The time when the experiment was last modified.
	//
	// example:
	//
	// 1589017861
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s CreateABTestExperimentResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateABTestExperimentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateABTestExperimentResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *CreateABTestExperimentResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *CreateABTestExperimentResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *CreateABTestExperimentResponseBodyResult) GetOnline() *bool {
	return s.Online
}

func (s *CreateABTestExperimentResponseBodyResult) GetParams() map[string]interface{} {
	return s.Params
}

func (s *CreateABTestExperimentResponseBodyResult) GetTraffic() *int32 {
	return s.Traffic
}

func (s *CreateABTestExperimentResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *CreateABTestExperimentResponseBodyResult) SetCreated(v int32) *CreateABTestExperimentResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateABTestExperimentResponseBodyResult) SetId(v string) *CreateABTestExperimentResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateABTestExperimentResponseBodyResult) SetName(v string) *CreateABTestExperimentResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateABTestExperimentResponseBodyResult) SetOnline(v bool) *CreateABTestExperimentResponseBodyResult {
	s.Online = &v
	return s
}

func (s *CreateABTestExperimentResponseBodyResult) SetParams(v map[string]interface{}) *CreateABTestExperimentResponseBodyResult {
	s.Params = v
	return s
}

func (s *CreateABTestExperimentResponseBodyResult) SetTraffic(v int32) *CreateABTestExperimentResponseBodyResult {
	s.Traffic = &v
	return s
}

func (s *CreateABTestExperimentResponseBodyResult) SetUpdated(v int32) *CreateABTestExperimentResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *CreateABTestExperimentResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
