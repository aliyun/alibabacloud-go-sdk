// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateABTestExperimentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateABTestExperimentResponseBody
	GetRequestId() *string
	SetResult(v *UpdateABTestExperimentResponseBodyResult) *UpdateABTestExperimentResponseBody
	GetResult() *UpdateABTestExperimentResponseBodyResult
}

type UpdateABTestExperimentResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the experiment.
	Result *UpdateABTestExperimentResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateABTestExperimentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateABTestExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateABTestExperimentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateABTestExperimentResponseBody) GetResult() *UpdateABTestExperimentResponseBodyResult {
	return s.Result
}

func (s *UpdateABTestExperimentResponseBody) SetRequestId(v string) *UpdateABTestExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateABTestExperimentResponseBody) SetResult(v *UpdateABTestExperimentResponseBodyResult) *UpdateABTestExperimentResponseBody {
	s.Result = v
	return s
}

func (s *UpdateABTestExperimentResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateABTestExperimentResponseBodyResult struct {
	// The time when the experiment was created.
	//
	// example:
	//
	// 1588842080
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the experiment.
	//
	// example:
	//
	// 12888
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the experiment.
	//
	// example:
	//
	// test1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the experiment.
	//
	// - true: The experiment is online.
	//
	// - false: The experiment is offline.
	//
	// example:
	//
	// true
	Online *bool `json:"online,omitempty" xml:"online,omitempty"`
	// The parameters of the experiment.
	//
	// example:
	//
	// {}
	Params map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// The percentage of traffic allocated to the experiment. Valid values: 0 to 100.
	//
	// example:
	//
	// 30
	Traffic *int32 `json:"traffic,omitempty" xml:"traffic,omitempty"`
	// The time when the experiment was last modified.
	//
	// example:
	//
	// 1588842080
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateABTestExperimentResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateABTestExperimentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateABTestExperimentResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *UpdateABTestExperimentResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *UpdateABTestExperimentResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *UpdateABTestExperimentResponseBodyResult) GetOnline() *bool {
	return s.Online
}

func (s *UpdateABTestExperimentResponseBodyResult) GetParams() map[string]interface{} {
	return s.Params
}

func (s *UpdateABTestExperimentResponseBodyResult) GetTraffic() *int32 {
	return s.Traffic
}

func (s *UpdateABTestExperimentResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *UpdateABTestExperimentResponseBodyResult) SetCreated(v int32) *UpdateABTestExperimentResponseBodyResult {
	s.Created = &v
	return s
}

func (s *UpdateABTestExperimentResponseBodyResult) SetId(v string) *UpdateABTestExperimentResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateABTestExperimentResponseBodyResult) SetName(v string) *UpdateABTestExperimentResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateABTestExperimentResponseBodyResult) SetOnline(v bool) *UpdateABTestExperimentResponseBodyResult {
	s.Online = &v
	return s
}

func (s *UpdateABTestExperimentResponseBodyResult) SetParams(v map[string]interface{}) *UpdateABTestExperimentResponseBodyResult {
	s.Params = v
	return s
}

func (s *UpdateABTestExperimentResponseBodyResult) SetTraffic(v int32) *UpdateABTestExperimentResponseBodyResult {
	s.Traffic = &v
	return s
}

func (s *UpdateABTestExperimentResponseBodyResult) SetUpdated(v int32) *UpdateABTestExperimentResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *UpdateABTestExperimentResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
