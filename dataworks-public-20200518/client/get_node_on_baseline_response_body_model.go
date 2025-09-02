// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeOnBaselineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetNodeOnBaselineResponseBodyData) *GetNodeOnBaselineResponseBody
	GetData() []*GetNodeOnBaselineResponseBodyData
	SetErrorCode(v string) *GetNodeOnBaselineResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetNodeOnBaselineResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetNodeOnBaselineResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetNodeOnBaselineResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetNodeOnBaselineResponseBody
	GetSuccess() *string
}

type GetNodeOnBaselineResponseBody struct {
	// The list of nodes.
	Data []*GetNodeOnBaselineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// The request ID.
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

func (s GetNodeOnBaselineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNodeOnBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeOnBaselineResponseBody) GetData() []*GetNodeOnBaselineResponseBodyData {
	return s.Data
}

func (s *GetNodeOnBaselineResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetNodeOnBaselineResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetNodeOnBaselineResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetNodeOnBaselineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNodeOnBaselineResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetNodeOnBaselineResponseBody) SetData(v []*GetNodeOnBaselineResponseBodyData) *GetNodeOnBaselineResponseBody {
	s.Data = v
	return s
}

func (s *GetNodeOnBaselineResponseBody) SetErrorCode(v string) *GetNodeOnBaselineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetNodeOnBaselineResponseBody) SetErrorMessage(v string) *GetNodeOnBaselineResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetNodeOnBaselineResponseBody) SetHttpStatusCode(v int32) *GetNodeOnBaselineResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetNodeOnBaselineResponseBody) SetRequestId(v string) *GetNodeOnBaselineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeOnBaselineResponseBody) SetSuccess(v string) *GetNodeOnBaselineResponseBody {
	s.Success = &v
	return s
}

func (s *GetNodeOnBaselineResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetNodeOnBaselineResponseBodyData struct {
	// The node ID.
	//
	// example:
	//
	// 1234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
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

func (s GetNodeOnBaselineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetNodeOnBaselineResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNodeOnBaselineResponseBodyData) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetNodeOnBaselineResponseBodyData) GetNodeName() *string {
	return s.NodeName
}

func (s *GetNodeOnBaselineResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *GetNodeOnBaselineResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetNodeOnBaselineResponseBodyData) SetNodeId(v int64) *GetNodeOnBaselineResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *GetNodeOnBaselineResponseBodyData) SetNodeName(v string) *GetNodeOnBaselineResponseBodyData {
	s.NodeName = &v
	return s
}

func (s *GetNodeOnBaselineResponseBodyData) SetOwner(v string) *GetNodeOnBaselineResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetNodeOnBaselineResponseBodyData) SetProjectId(v int64) *GetNodeOnBaselineResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetNodeOnBaselineResponseBodyData) Validate() error {
	return dara.Validate(s)
}
