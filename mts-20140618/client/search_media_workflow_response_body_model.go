// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaWorkflowList(v *SearchMediaWorkflowResponseBodyMediaWorkflowList) *SearchMediaWorkflowResponseBody
	GetMediaWorkflowList() *SearchMediaWorkflowResponseBodyMediaWorkflowList
	SetPageNumber(v int64) *SearchMediaWorkflowResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *SearchMediaWorkflowResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *SearchMediaWorkflowResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *SearchMediaWorkflowResponseBody
	GetTotalCount() *int64
}

type SearchMediaWorkflowResponseBody struct {
	// The details of the media workflows.
	MediaWorkflowList *SearchMediaWorkflowResponseBodyMediaWorkflowList `json:"MediaWorkflowList,omitempty" xml:"MediaWorkflowList,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 16CD0CDD-457E-420D-9755-8385075A1234
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchMediaWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *SearchMediaWorkflowResponseBody) GetMediaWorkflowList() *SearchMediaWorkflowResponseBodyMediaWorkflowList {
	return s.MediaWorkflowList
}

func (s *SearchMediaWorkflowResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *SearchMediaWorkflowResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *SearchMediaWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchMediaWorkflowResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *SearchMediaWorkflowResponseBody) SetMediaWorkflowList(v *SearchMediaWorkflowResponseBodyMediaWorkflowList) *SearchMediaWorkflowResponseBody {
	s.MediaWorkflowList = v
	return s
}

func (s *SearchMediaWorkflowResponseBody) SetPageNumber(v int64) *SearchMediaWorkflowResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchMediaWorkflowResponseBody) SetPageSize(v int64) *SearchMediaWorkflowResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchMediaWorkflowResponseBody) SetRequestId(v string) *SearchMediaWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchMediaWorkflowResponseBody) SetTotalCount(v int64) *SearchMediaWorkflowResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchMediaWorkflowResponseBody) Validate() error {
	if s.MediaWorkflowList != nil {
		if err := s.MediaWorkflowList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchMediaWorkflowResponseBodyMediaWorkflowList struct {
	MediaWorkflow []*SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow `json:"MediaWorkflow,omitempty" xml:"MediaWorkflow,omitempty" type:"Repeated"`
}

func (s SearchMediaWorkflowResponseBodyMediaWorkflowList) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaWorkflowResponseBodyMediaWorkflowList) GoString() string {
	return s.String()
}

func (s *SearchMediaWorkflowResponseBodyMediaWorkflowList) GetMediaWorkflow() []*SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow {
	return s.MediaWorkflow
}

func (s *SearchMediaWorkflowResponseBodyMediaWorkflowList) SetMediaWorkflow(v []*SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow) *SearchMediaWorkflowResponseBodyMediaWorkflowList {
	s.MediaWorkflow = v
	return s
}

func (s *SearchMediaWorkflowResponseBodyMediaWorkflowList) Validate() error {
	if s.MediaWorkflow != nil {
		for _, item := range s.MediaWorkflow {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow struct {
	// The time when the media workflow was created.
	//
	// example:
	//
	// 2016-04-01T05:38:41Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the media workflow.
	//
	// example:
	//
	// 88c6ca184c0e4578645b665e2a12****
	MediaWorkflowId *string `json:"MediaWorkflowId,omitempty" xml:"MediaWorkflowId,omitempty"`
	// The name of the media workflow.
	//
	// example:
	//
	// example-mediaworkflow-****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the media workflow. Valid values:
	//
	// 	- **Inactive**: The media workflow is deactivated.
	//
	// 	- **Active**: The media workflow is activated.
	//
	// 	- **Deleted**: The media workflow is deleted.
	//
	// example:
	//
	// Active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The topology of the media workflow.
	//
	// example:
	//
	// {"MediaWorkflowList":{"MediaWorkflow":[{"CreationTime":"2016-04-01T05:29:38Z","Name":"example-mediaworkflow-****","State":"Active","Topology":"{\\"Activities\\":{\\"Act-Start\\":{\\"Parameters\\":{\\"PipelineId\\":\\"130266f58161436a80bf07cb12c8****\\",\\"InputFile\\":\\"{\\\\\\"Bucket\\\\\\": \\\\\\"example-bucket-****\\\\\\",\\\\\\"Location\\\\\\": \\\\\\"example-location\\\\\\"}\\"},\\"Type\\":\\"Start\\"},\\"Act-Report\\":{\\"Parameters\\":{},\\"Type\\":\\"Report\\"},\\"Act-Transcode-M3U8\\":{\\"Parameters\\":{\\"Outputs\\":\\"[{\\\\\\"Object\\\\\\":\\\\\\"transcode/{ObjectPrefix}{FileName}\\\\\\",\\\\\\"TemplateId\\\\\\": \\\\\\"957d1719ee85ed6527b90cf62726****\\\\\\"}]\\",\\"OutputBucket\\":\\"example-bucket-****\\",\\"OutputLocation\\":\\"example-location\\"},\\"Type\\":\\"Transcode\\"}},\\"Dependencies\\":{\\"Act-Start\\":[\\"Act-Transcode-M3U8\\"],\\"Act-Report\\":[],\\"Act-Transcode-M3U8\\":[\\"Act-Report\\"]}}","MediaWorkflowId":"93ab850b4f6f44eab54b6e91d24d****"}]},"RequestId":"16CD0CDD-457E-420D-9755-8385075A1234"}
	Topology *string `json:"Topology,omitempty" xml:"Topology,omitempty"`
	// The trigger mode of the media workflow. Valid values:
	//
	// 	- **OssAutoTrigger**: The media workflow is automatically triggered.
	//
	// 	- **NotInAuto**: The media workflow is not automatically triggered.
	//
	// example:
	//
	// OssAutoTrigger
	TriggerMode *string `json:"TriggerMode,omitempty" xml:"TriggerMode,omitempty"`
}

func (s SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow) GoString() string {
	return s.String()
}

func (s *SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow) GetCreationTime() *string {
	return s.CreationTime
}

func (s *SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow) GetMediaWorkflowId() *string {
	return s.MediaWorkflowId
}

func (s *SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow) GetName() *string {
	return s.Name
}

func (s *SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow) GetState() *string {
	return s.State
}

func (s *SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow) GetTopology() *string {
	return s.Topology
}

func (s *SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow) GetTriggerMode() *string {
	return s.TriggerMode
}

func (s *SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow) SetCreationTime(v string) *SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow {
	s.CreationTime = &v
	return s
}

func (s *SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow) SetMediaWorkflowId(v string) *SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow {
	s.MediaWorkflowId = &v
	return s
}

func (s *SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow) SetName(v string) *SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow {
	s.Name = &v
	return s
}

func (s *SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow) SetState(v string) *SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow {
	s.State = &v
	return s
}

func (s *SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow) SetTopology(v string) *SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow {
	s.Topology = &v
	return s
}

func (s *SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow) SetTriggerMode(v string) *SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow {
	s.TriggerMode = &v
	return s
}

func (s *SearchMediaWorkflowResponseBodyMediaWorkflowListMediaWorkflow) Validate() error {
	return dara.Validate(s)
}
