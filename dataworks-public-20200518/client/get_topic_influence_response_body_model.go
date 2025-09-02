// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicInfluenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetTopicInfluenceResponseBodyData) *GetTopicInfluenceResponseBody
	GetData() *GetTopicInfluenceResponseBodyData
	SetErrorCode(v string) *GetTopicInfluenceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetTopicInfluenceResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetTopicInfluenceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetTopicInfluenceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTopicInfluenceResponseBody
	GetSuccess() *bool
}

type GetTopicInfluenceResponseBody struct {
	// The data returned.
	Data *GetTopicInfluenceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned.
	//
	// example:
	//
	// 1031203110005
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The specified parameters are invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTopicInfluenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTopicInfluenceResponseBody) GoString() string {
	return s.String()
}

func (s *GetTopicInfluenceResponseBody) GetData() *GetTopicInfluenceResponseBodyData {
	return s.Data
}

func (s *GetTopicInfluenceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTopicInfluenceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTopicInfluenceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTopicInfluenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTopicInfluenceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTopicInfluenceResponseBody) SetData(v *GetTopicInfluenceResponseBodyData) *GetTopicInfluenceResponseBody {
	s.Data = v
	return s
}

func (s *GetTopicInfluenceResponseBody) SetErrorCode(v string) *GetTopicInfluenceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTopicInfluenceResponseBody) SetErrorMessage(v string) *GetTopicInfluenceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetTopicInfluenceResponseBody) SetHttpStatusCode(v int32) *GetTopicInfluenceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTopicInfluenceResponseBody) SetRequestId(v string) *GetTopicInfluenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTopicInfluenceResponseBody) SetSuccess(v bool) *GetTopicInfluenceResponseBody {
	s.Success = &v
	return s
}

func (s *GetTopicInfluenceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTopicInfluenceResponseBodyData struct {
	// The affected baseline instances.
	Influences []*GetTopicInfluenceResponseBodyDataInfluences `json:"Influences,omitempty" xml:"Influences,omitempty" type:"Repeated"`
	// The ID of the event.
	//
	// example:
	//
	// 1234
	TopicId *int64 `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
}

func (s GetTopicInfluenceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTopicInfluenceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTopicInfluenceResponseBodyData) GetInfluences() []*GetTopicInfluenceResponseBodyDataInfluences {
	return s.Influences
}

func (s *GetTopicInfluenceResponseBodyData) GetTopicId() *int64 {
	return s.TopicId
}

func (s *GetTopicInfluenceResponseBodyData) SetInfluences(v []*GetTopicInfluenceResponseBodyDataInfluences) *GetTopicInfluenceResponseBodyData {
	s.Influences = v
	return s
}

func (s *GetTopicInfluenceResponseBodyData) SetTopicId(v int64) *GetTopicInfluenceResponseBodyData {
	s.TopicId = &v
	return s
}

func (s *GetTopicInfluenceResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetTopicInfluenceResponseBodyDataInfluences struct {
	// The ID of the baseline.
	//
	// example:
	//
	// 12345
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The name of the baseline.
	//
	// example:
	//
	// Baseline name
	BaselineName *string `json:"BaselineName,omitempty" xml:"BaselineName,omitempty"`
	// The data timestamp of the baseline instance.
	//
	// example:
	//
	// 1553356800000
	Bizdate *int64 `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The margin of the baseline instance. Unit: seconds.
	//
	// example:
	//
	// 360
	Buffer *int64 `json:"Buffer,omitempty" xml:"Buffer,omitempty"`
	// The ID of the cycle of the baseline instance. For a baseline instance that is scheduled by day, the field value is 1. For a baseline instance that is scheduled by hour, the field value ranges from 1 to 24.
	//
	// example:
	//
	// 1
	InGroupId *int32 `json:"InGroupId,omitempty" xml:"InGroupId,omitempty"`
	// The ID of the Alibaba Cloud account used by the baseline owner. Multiple IDs are separated by commas (,).
	//
	// example:
	//
	// 952795****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The priority of the baseline. Valid values: 1, 2, 5, 7, and 8.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the workspace to which the baseline belongs.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The status of the baseline. Valid values: ERROR, SAFE, DANGROUS, and OVER. The value ERROR indicates that no nodes are associated with the baseline, or all nodes associated with the baseline are suspended. The value SAFE indicates that nodes are run before the alert duration begins. The value DANGROUS indicates that nodes are still running after the alert duration ends but the committed time does not arrive. The value OVER indicates that nodes are still running after the committed time.
	//
	// example:
	//
	// SAFE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetTopicInfluenceResponseBodyDataInfluences) String() string {
	return dara.Prettify(s)
}

func (s GetTopicInfluenceResponseBodyDataInfluences) GoString() string {
	return s.String()
}

func (s *GetTopicInfluenceResponseBodyDataInfluences) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *GetTopicInfluenceResponseBodyDataInfluences) GetBaselineName() *string {
	return s.BaselineName
}

func (s *GetTopicInfluenceResponseBodyDataInfluences) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *GetTopicInfluenceResponseBodyDataInfluences) GetBuffer() *int64 {
	return s.Buffer
}

func (s *GetTopicInfluenceResponseBodyDataInfluences) GetInGroupId() *int32 {
	return s.InGroupId
}

func (s *GetTopicInfluenceResponseBodyDataInfluences) GetOwner() *string {
	return s.Owner
}

func (s *GetTopicInfluenceResponseBodyDataInfluences) GetPriority() *int32 {
	return s.Priority
}

func (s *GetTopicInfluenceResponseBodyDataInfluences) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetTopicInfluenceResponseBodyDataInfluences) GetStatus() *string {
	return s.Status
}

func (s *GetTopicInfluenceResponseBodyDataInfluences) SetBaselineId(v int64) *GetTopicInfluenceResponseBodyDataInfluences {
	s.BaselineId = &v
	return s
}

func (s *GetTopicInfluenceResponseBodyDataInfluences) SetBaselineName(v string) *GetTopicInfluenceResponseBodyDataInfluences {
	s.BaselineName = &v
	return s
}

func (s *GetTopicInfluenceResponseBodyDataInfluences) SetBizdate(v int64) *GetTopicInfluenceResponseBodyDataInfluences {
	s.Bizdate = &v
	return s
}

func (s *GetTopicInfluenceResponseBodyDataInfluences) SetBuffer(v int64) *GetTopicInfluenceResponseBodyDataInfluences {
	s.Buffer = &v
	return s
}

func (s *GetTopicInfluenceResponseBodyDataInfluences) SetInGroupId(v int32) *GetTopicInfluenceResponseBodyDataInfluences {
	s.InGroupId = &v
	return s
}

func (s *GetTopicInfluenceResponseBodyDataInfluences) SetOwner(v string) *GetTopicInfluenceResponseBodyDataInfluences {
	s.Owner = &v
	return s
}

func (s *GetTopicInfluenceResponseBodyDataInfluences) SetPriority(v int32) *GetTopicInfluenceResponseBodyDataInfluences {
	s.Priority = &v
	return s
}

func (s *GetTopicInfluenceResponseBodyDataInfluences) SetProjectId(v int64) *GetTopicInfluenceResponseBodyDataInfluences {
	s.ProjectId = &v
	return s
}

func (s *GetTopicInfluenceResponseBodyDataInfluences) SetStatus(v string) *GetTopicInfluenceResponseBodyDataInfluences {
	s.Status = &v
	return s
}

func (s *GetTopicInfluenceResponseBodyDataInfluences) Validate() error {
	return dara.Validate(s)
}
