// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRectifyFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateRectifyFlowResponseBody
	GetCode() *string
	SetData(v *CreateRectifyFlowResponseBodyData) *CreateRectifyFlowResponseBody
	GetData() *CreateRectifyFlowResponseBodyData
	SetHttpStatusCode(v int32) *CreateRectifyFlowResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateRectifyFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateRectifyFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateRectifyFlowResponseBody
	GetSuccess() *bool
}

type CreateRectifyFlowResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The operation result.
	Data *CreateRectifyFlowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRectifyFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRectifyFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRectifyFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateRectifyFlowResponseBody) GetData() *CreateRectifyFlowResponseBodyData {
	return s.Data
}

func (s *CreateRectifyFlowResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateRectifyFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateRectifyFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRectifyFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateRectifyFlowResponseBody) SetCode(v string) *CreateRectifyFlowResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRectifyFlowResponseBody) SetData(v *CreateRectifyFlowResponseBodyData) *CreateRectifyFlowResponseBody {
	s.Data = v
	return s
}

func (s *CreateRectifyFlowResponseBody) SetHttpStatusCode(v int32) *CreateRectifyFlowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateRectifyFlowResponseBody) SetMessage(v string) *CreateRectifyFlowResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRectifyFlowResponseBody) SetRequestId(v string) *CreateRectifyFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRectifyFlowResponseBody) SetSuccess(v bool) *CreateRectifyFlowResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRectifyFlowResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRectifyFlowResponseBodyData struct {
	// The actual number of deduplicated issues successfully included in this operation. Skipped issues and historical issues of the original process are excluded. Issues brought in by scope expansion are included.
	//
	// example:
	//
	// 3
	ProcessedIssueCount *int64 `json:"ProcessedIssueCount,omitempty" xml:"ProcessedIssueCount,omitempty"`
	// The list of process IDs for this operation.
	RectifyIds []*int64 `json:"RectifyIds,omitempty" xml:"RectifyIds,omitempty" type:"Repeated"`
}

func (s CreateRectifyFlowResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateRectifyFlowResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateRectifyFlowResponseBodyData) GetProcessedIssueCount() *int64 {
	return s.ProcessedIssueCount
}

func (s *CreateRectifyFlowResponseBodyData) GetRectifyIds() []*int64 {
	return s.RectifyIds
}

func (s *CreateRectifyFlowResponseBodyData) SetProcessedIssueCount(v int64) *CreateRectifyFlowResponseBodyData {
	s.ProcessedIssueCount = &v
	return s
}

func (s *CreateRectifyFlowResponseBodyData) SetRectifyIds(v []*int64) *CreateRectifyFlowResponseBodyData {
	s.RectifyIds = v
	return s
}

func (s *CreateRectifyFlowResponseBodyData) Validate() error {
	return dara.Validate(s)
}
