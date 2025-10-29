// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v *ListPipelineResponseBodyHeaders) *ListPipelineResponseBody
	GetHeaders() *ListPipelineResponseBodyHeaders
	SetRequestId(v string) *ListPipelineResponseBody
	GetRequestId() *string
	SetResult(v []*ListPipelineResponseBodyResult) *ListPipelineResponseBody
	GetResult() []*ListPipelineResponseBodyResult
}

type ListPipelineResponseBody struct {
	// The ID of the ApsaraVideo Media Processing (MPS) queue that is used to run the job.
	Headers *ListPipelineResponseBodyHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// The response.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the pipeline was created.
	Result []*ListPipelineResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListPipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineResponseBody) GetHeaders() *ListPipelineResponseBodyHeaders {
	return s.Headers
}

func (s *ListPipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPipelineResponseBody) GetResult() []*ListPipelineResponseBodyResult {
	return s.Result
}

func (s *ListPipelineResponseBody) SetHeaders(v *ListPipelineResponseBodyHeaders) *ListPipelineResponseBody {
	s.Headers = v
	return s
}

func (s *ListPipelineResponseBody) SetRequestId(v string) *ListPipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelineResponseBody) SetResult(v []*ListPipelineResponseBodyResult) *ListPipelineResponseBody {
	s.Result = v
	return s
}

func (s *ListPipelineResponseBody) Validate() error {
	if s.Headers != nil {
		if err := s.Headers.Validate(); err != nil {
			return err
		}
	}
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

type ListPipelineResponseBodyHeaders struct {
	// The time when the pipeline was updated.
	//
	// example:
	//
	// 2
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListPipelineResponseBodyHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListPipelineResponseBodyHeaders) GetXTotalCount() *int32 {
	return s.XTotalCount
}

func (s *ListPipelineResponseBodyHeaders) SetXTotalCount(v int32) *ListPipelineResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

func (s *ListPipelineResponseBodyHeaders) Validate() error {
	return dara.Validate(s)
}

type ListPipelineResponseBodyResult struct {
	// example:
	//
	// 2020-08-05T03:10:38.188Z
	GmtCreatedTime *string `json:"gmtCreatedTime,omitempty" xml:"gmtCreatedTime,omitempty"`
	// example:
	//
	// 2020-08-05T08:43:31.757Z
	GmtUpdateTime *string `json:"gmtUpdateTime,omitempty" xml:"gmtUpdateTime,omitempty"`
	// The status of the pipeline. Supported:
	//
	// 	- NOT_DEPLOYED: The node is not deployed.
	//
	// 	- RUNNING
	//
	// 	- DELETED: Deleted. The console does not display this status.
	//
	// example:
	//
	// pipeline_test
	PipelineId *string `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
	// example:
	//
	// NOT_DEPLOYED
	PipelineStatus *string `json:"pipelineStatus,omitempty" xml:"pipelineStatus,omitempty"`
}

func (s ListPipelineResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListPipelineResponseBodyResult) GetGmtCreatedTime() *string {
	return s.GmtCreatedTime
}

func (s *ListPipelineResponseBodyResult) GetGmtUpdateTime() *string {
	return s.GmtUpdateTime
}

func (s *ListPipelineResponseBodyResult) GetPipelineId() *string {
	return s.PipelineId
}

func (s *ListPipelineResponseBodyResult) GetPipelineStatus() *string {
	return s.PipelineStatus
}

func (s *ListPipelineResponseBodyResult) SetGmtCreatedTime(v string) *ListPipelineResponseBodyResult {
	s.GmtCreatedTime = &v
	return s
}

func (s *ListPipelineResponseBodyResult) SetGmtUpdateTime(v string) *ListPipelineResponseBodyResult {
	s.GmtUpdateTime = &v
	return s
}

func (s *ListPipelineResponseBodyResult) SetPipelineId(v string) *ListPipelineResponseBodyResult {
	s.PipelineId = &v
	return s
}

func (s *ListPipelineResponseBodyResult) SetPipelineStatus(v string) *ListPipelineResponseBodyResult {
	s.PipelineStatus = &v
	return s
}

func (s *ListPipelineResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
