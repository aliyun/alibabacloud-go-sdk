// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineIdsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListPipelineIdsResponseBody
	GetRequestId() *string
	SetResult(v []*ListPipelineIdsResponseBodyResult) *ListPipelineIdsResponseBody
	GetResult() []*ListPipelineIdsResponseBodyResult
}

type ListPipelineIdsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of results.
	Result []*ListPipelineIdsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListPipelineIdsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineIdsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPipelineIdsResponseBody) GetResult() []*ListPipelineIdsResponseBodyResult {
	return s.Result
}

func (s *ListPipelineIdsResponseBody) SetRequestId(v string) *ListPipelineIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelineIdsResponseBody) SetResult(v []*ListPipelineIdsResponseBodyResult) *ListPipelineIdsResponseBody {
	s.Result = v
	return s
}

func (s *ListPipelineIdsResponseBody) Validate() error {
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

type ListPipelineIdsResponseBodyResult struct {
	// Indicates whether the pipeline is available. Valid values:
	//
	// - true: Available.
	//
	// - false: Not available.
	//
	// example:
	//
	// true
	Available *bool `json:"available,omitempty" xml:"available,omitempty"`
	// The error code returned when the pipeline is unavailable.
	//
	// example:
	//
	// OK
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error message returned when the pipeline is unavailable.
	//
	// example:
	//
	// OK
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pipeline ID created in Kibana.
	//
	// example:
	//
	// testKibanaManagement
	PipelineId *string `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
}

func (s ListPipelineIdsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineIdsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListPipelineIdsResponseBodyResult) GetAvailable() *bool {
	return s.Available
}

func (s *ListPipelineIdsResponseBodyResult) GetCode() *string {
	return s.Code
}

func (s *ListPipelineIdsResponseBodyResult) GetMessage() *string {
	return s.Message
}

func (s *ListPipelineIdsResponseBodyResult) GetPipelineId() *string {
	return s.PipelineId
}

func (s *ListPipelineIdsResponseBodyResult) SetAvailable(v bool) *ListPipelineIdsResponseBodyResult {
	s.Available = &v
	return s
}

func (s *ListPipelineIdsResponseBodyResult) SetCode(v string) *ListPipelineIdsResponseBodyResult {
	s.Code = &v
	return s
}

func (s *ListPipelineIdsResponseBodyResult) SetMessage(v string) *ListPipelineIdsResponseBodyResult {
	s.Message = &v
	return s
}

func (s *ListPipelineIdsResponseBodyResult) SetPipelineId(v string) *ListPipelineIdsResponseBodyResult {
	s.PipelineId = &v
	return s
}

func (s *ListPipelineIdsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
