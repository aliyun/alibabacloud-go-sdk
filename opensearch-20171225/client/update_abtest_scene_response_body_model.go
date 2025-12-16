// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateABTestSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateABTestSceneResponseBody
	GetRequestId() *string
	SetResult(v *UpdateABTestSceneResponseBodyResult) *UpdateABTestSceneResponseBody
	GetResult() *UpdateABTestSceneResponseBodyResult
}

type UpdateABTestSceneResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the test scenario. For more information, see [ABTestScene](https://help.aliyun.com/document_detail/173618.html).
	//
	// example:
	//
	// {}
	Result *UpdateABTestSceneResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateABTestSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateABTestSceneResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateABTestSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateABTestSceneResponseBody) GetResult() *UpdateABTestSceneResponseBodyResult {
	return s.Result
}

func (s *UpdateABTestSceneResponseBody) SetRequestId(v string) *UpdateABTestSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateABTestSceneResponseBody) SetResult(v *UpdateABTestSceneResponseBodyResult) *UpdateABTestSceneResponseBody {
	s.Result = v
	return s
}

func (s *UpdateABTestSceneResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateABTestSceneResponseBodyResult struct {
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
	// kevintest22
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test. Valid values:
	//
	// 	- true: The test is started.
	//
	// 	- false: The test is stopped.
	//
	// example:
	//
	// true
	Online *bool `json:"online,omitempty" xml:"online,omitempty"`
	// The parameters of the A/B test.
	//
	// example:
	//
	// {}
	Params map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// The percentage of traffic that is allocated to the A/B test. Valid values: 0 to 100.
	//
	// example:
	//
	// 111
	Traffic *int32 `json:"traffic,omitempty" xml:"traffic,omitempty"`
	// The time when the test scenario was last modified.
	//
	// example:
	//
	// 1596527691
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateABTestSceneResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateABTestSceneResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateABTestSceneResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *UpdateABTestSceneResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *UpdateABTestSceneResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *UpdateABTestSceneResponseBodyResult) GetOnline() *bool {
	return s.Online
}

func (s *UpdateABTestSceneResponseBodyResult) GetParams() map[string]interface{} {
	return s.Params
}

func (s *UpdateABTestSceneResponseBodyResult) GetTraffic() *int32 {
	return s.Traffic
}

func (s *UpdateABTestSceneResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *UpdateABTestSceneResponseBodyResult) SetCreated(v int32) *UpdateABTestSceneResponseBodyResult {
	s.Created = &v
	return s
}

func (s *UpdateABTestSceneResponseBodyResult) SetId(v string) *UpdateABTestSceneResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateABTestSceneResponseBodyResult) SetName(v string) *UpdateABTestSceneResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateABTestSceneResponseBodyResult) SetOnline(v bool) *UpdateABTestSceneResponseBodyResult {
	s.Online = &v
	return s
}

func (s *UpdateABTestSceneResponseBodyResult) SetParams(v map[string]interface{}) *UpdateABTestSceneResponseBodyResult {
	s.Params = v
	return s
}

func (s *UpdateABTestSceneResponseBodyResult) SetTraffic(v int32) *UpdateABTestSceneResponseBodyResult {
	s.Traffic = &v
	return s
}

func (s *UpdateABTestSceneResponseBodyResult) SetUpdated(v int32) *UpdateABTestSceneResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *UpdateABTestSceneResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
