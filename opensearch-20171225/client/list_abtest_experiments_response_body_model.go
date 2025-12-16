// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListABTestExperimentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListABTestExperimentsResponseBody
	GetRequestId() *string
	SetResult(v []*ListABTestExperimentsResponseBodyResult) *ListABTestExperimentsResponseBody
	GetResult() []*ListABTestExperimentsResponseBodyResult
}

type ListABTestExperimentsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The experiment details.\\
	//
	// For more information, see [ABTestExperiment](https://help.aliyun.com/document_detail/173617.html).
	Result []*ListABTestExperimentsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListABTestExperimentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListABTestExperimentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListABTestExperimentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListABTestExperimentsResponseBody) GetResult() []*ListABTestExperimentsResponseBodyResult {
	return s.Result
}

func (s *ListABTestExperimentsResponseBody) SetRequestId(v string) *ListABTestExperimentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListABTestExperimentsResponseBody) SetResult(v []*ListABTestExperimentsResponseBodyResult) *ListABTestExperimentsResponseBody {
	s.Result = v
	return s
}

func (s *ListABTestExperimentsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListABTestExperimentsResponseBodyResult struct {
	// The time when the experiment was created.
	//
	// example:
	//
	// 1588842080
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The experiment ID.
	//
	// example:
	//
	// 12888
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The group alias.
	//
	// example:
	//
	// test1
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
	// 1
	Params map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// The percentage of traffic that is routed to the experiment.
	//
	// Valid values: [0,100]
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

func (s ListABTestExperimentsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListABTestExperimentsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListABTestExperimentsResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *ListABTestExperimentsResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ListABTestExperimentsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListABTestExperimentsResponseBodyResult) GetOnline() *bool {
	return s.Online
}

func (s *ListABTestExperimentsResponseBodyResult) GetParams() map[string]interface{} {
	return s.Params
}

func (s *ListABTestExperimentsResponseBodyResult) GetTraffic() *int32 {
	return s.Traffic
}

func (s *ListABTestExperimentsResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *ListABTestExperimentsResponseBodyResult) SetCreated(v int32) *ListABTestExperimentsResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListABTestExperimentsResponseBodyResult) SetId(v string) *ListABTestExperimentsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListABTestExperimentsResponseBodyResult) SetName(v string) *ListABTestExperimentsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListABTestExperimentsResponseBodyResult) SetOnline(v bool) *ListABTestExperimentsResponseBodyResult {
	s.Online = &v
	return s
}

func (s *ListABTestExperimentsResponseBodyResult) SetParams(v map[string]interface{}) *ListABTestExperimentsResponseBodyResult {
	s.Params = v
	return s
}

func (s *ListABTestExperimentsResponseBodyResult) SetTraffic(v int32) *ListABTestExperimentsResponseBodyResult {
	s.Traffic = &v
	return s
}

func (s *ListABTestExperimentsResponseBodyResult) SetUpdated(v int32) *ListABTestExperimentsResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListABTestExperimentsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
