// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesByBaselineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListNodesByBaselineResponseBodyData) *ListNodesByBaselineResponseBody
	GetData() []*ListNodesByBaselineResponseBodyData
	SetErrorCode(v string) *ListNodesByBaselineResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListNodesByBaselineResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListNodesByBaselineResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListNodesByBaselineResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ListNodesByBaselineResponseBody
	GetSuccess() *string
}

type ListNodesByBaselineResponseBody struct {
	// The nodes in the baseline.
	Data []*ListNodesByBaselineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// 1031203110005
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameters are invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 0000-ABCD-E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNodesByBaselineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodesByBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesByBaselineResponseBody) GetData() []*ListNodesByBaselineResponseBodyData {
	return s.Data
}

func (s *ListNodesByBaselineResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListNodesByBaselineResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListNodesByBaselineResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListNodesByBaselineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNodesByBaselineResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ListNodesByBaselineResponseBody) SetData(v []*ListNodesByBaselineResponseBodyData) *ListNodesByBaselineResponseBody {
	s.Data = v
	return s
}

func (s *ListNodesByBaselineResponseBody) SetErrorCode(v string) *ListNodesByBaselineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListNodesByBaselineResponseBody) SetErrorMessage(v string) *ListNodesByBaselineResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListNodesByBaselineResponseBody) SetHttpStatusCode(v int32) *ListNodesByBaselineResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListNodesByBaselineResponseBody) SetRequestId(v string) *ListNodesByBaselineResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesByBaselineResponseBody) SetSuccess(v string) *ListNodesByBaselineResponseBody {
	s.Success = &v
	return s
}

func (s *ListNodesByBaselineResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListNodesByBaselineResponseBodyData struct {
	// The node ID.
	//
	// example:
	//
	// 1234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// Node name
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The ID of the Alibaba Cloud account used by the node owner.
	//
	// example:
	//
	// 9527952****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the workspace to which the node belongs.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListNodesByBaselineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListNodesByBaselineResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNodesByBaselineResponseBodyData) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListNodesByBaselineResponseBodyData) GetNodeName() *string {
	return s.NodeName
}

func (s *ListNodesByBaselineResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *ListNodesByBaselineResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListNodesByBaselineResponseBodyData) SetNodeId(v int64) *ListNodesByBaselineResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *ListNodesByBaselineResponseBodyData) SetNodeName(v string) *ListNodesByBaselineResponseBodyData {
	s.NodeName = &v
	return s
}

func (s *ListNodesByBaselineResponseBodyData) SetOwner(v string) *ListNodesByBaselineResponseBodyData {
	s.Owner = &v
	return s
}

func (s *ListNodesByBaselineResponseBodyData) SetProjectId(v int64) *ListNodesByBaselineResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *ListNodesByBaselineResponseBodyData) Validate() error {
	return dara.Validate(s)
}
