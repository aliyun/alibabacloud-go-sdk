// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetNodeConfigResponseBody
	GetRequestId() *string
	SetResult(v *GetNodeConfigResponseBodyResult) *GetNodeConfigResponseBody
	GetResult() *GetNodeConfigResponseBodyResult
}

type GetNodeConfigResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 2AE63638-5420-56DC-BF59-37D8174039A0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result set.
	Result *GetNodeConfigResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetNodeConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNodeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNodeConfigResponseBody) GetResult() *GetNodeConfigResponseBodyResult {
	return s.Result
}

func (s *GetNodeConfigResponseBody) SetRequestId(v string) *GetNodeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeConfigResponseBody) SetResult(v *GetNodeConfigResponseBodyResult) *GetNodeConfigResponseBody {
	s.Result = v
	return s
}

func (s *GetNodeConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetNodeConfigResponseBodyResult struct {
	// Indicates whether the index is effective online.
	//
	// example:
	//
	// 1
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The number of data replicas.
	//
	// example:
	//
	// 2
	DataDuplicateNumber *int32 `json:"dataDuplicateNumber,omitempty" xml:"dataDuplicateNumber,omitempty"`
	// The number of data shards.
	//
	// example:
	//
	// 2
	DataFragmentNumber *int32 `json:"dataFragmentNumber,omitempty" xml:"dataFragmentNumber,omitempty"`
	// The traffic percentage.
	//
	// example:
	//
	// 0
	FlowRatio *int32 `json:"flowRatio,omitempty" xml:"flowRatio,omitempty"`
	// The minimum service ratio.
	//
	// example:
	//
	// 100
	MinServicePercent *int32 `json:"minServicePercent,omitempty" xml:"minServicePercent,omitempty"`
	// Indicates whether the cluster is mounted.
	//
	// example:
	//
	// true
	Published *bool `json:"published,omitempty" xml:"published,omitempty"`
}

func (s GetNodeConfigResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetNodeConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetNodeConfigResponseBodyResult) GetActive() *bool {
	return s.Active
}

func (s *GetNodeConfigResponseBodyResult) GetDataDuplicateNumber() *int32 {
	return s.DataDuplicateNumber
}

func (s *GetNodeConfigResponseBodyResult) GetDataFragmentNumber() *int32 {
	return s.DataFragmentNumber
}

func (s *GetNodeConfigResponseBodyResult) GetFlowRatio() *int32 {
	return s.FlowRatio
}

func (s *GetNodeConfigResponseBodyResult) GetMinServicePercent() *int32 {
	return s.MinServicePercent
}

func (s *GetNodeConfigResponseBodyResult) GetPublished() *bool {
	return s.Published
}

func (s *GetNodeConfigResponseBodyResult) SetActive(v bool) *GetNodeConfigResponseBodyResult {
	s.Active = &v
	return s
}

func (s *GetNodeConfigResponseBodyResult) SetDataDuplicateNumber(v int32) *GetNodeConfigResponseBodyResult {
	s.DataDuplicateNumber = &v
	return s
}

func (s *GetNodeConfigResponseBodyResult) SetDataFragmentNumber(v int32) *GetNodeConfigResponseBodyResult {
	s.DataFragmentNumber = &v
	return s
}

func (s *GetNodeConfigResponseBodyResult) SetFlowRatio(v int32) *GetNodeConfigResponseBodyResult {
	s.FlowRatio = &v
	return s
}

func (s *GetNodeConfigResponseBodyResult) SetMinServicePercent(v int32) *GetNodeConfigResponseBodyResult {
	s.MinServicePercent = &v
	return s
}

func (s *GetNodeConfigResponseBodyResult) SetPublished(v bool) *GetNodeConfigResponseBodyResult {
	s.Published = &v
	return s
}

func (s *GetNodeConfigResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
