// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodeOpsOwnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateNodeOpsOwnerResponseBody
	GetCode() *string
	SetData(v []*UpdateNodeOpsOwnerResponseBodyData) *UpdateNodeOpsOwnerResponseBody
	GetData() []*UpdateNodeOpsOwnerResponseBodyData
	SetHttpStatusCode(v int32) *UpdateNodeOpsOwnerResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateNodeOpsOwnerResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateNodeOpsOwnerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateNodeOpsOwnerResponseBody
	GetSuccess() *bool
}

type UpdateNodeOpsOwnerResponseBody struct {
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of per-node operation results.
	Data []*UpdateNodeOpsOwnerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateNodeOpsOwnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeOpsOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNodeOpsOwnerResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateNodeOpsOwnerResponseBody) GetData() []*UpdateNodeOpsOwnerResponseBodyData {
	return s.Data
}

func (s *UpdateNodeOpsOwnerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateNodeOpsOwnerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateNodeOpsOwnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNodeOpsOwnerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateNodeOpsOwnerResponseBody) SetCode(v string) *UpdateNodeOpsOwnerResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateNodeOpsOwnerResponseBody) SetData(v []*UpdateNodeOpsOwnerResponseBodyData) *UpdateNodeOpsOwnerResponseBody {
	s.Data = v
	return s
}

func (s *UpdateNodeOpsOwnerResponseBody) SetHttpStatusCode(v int32) *UpdateNodeOpsOwnerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateNodeOpsOwnerResponseBody) SetMessage(v string) *UpdateNodeOpsOwnerResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateNodeOpsOwnerResponseBody) SetRequestId(v string) *UpdateNodeOpsOwnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNodeOpsOwnerResponseBody) SetSuccess(v bool) *UpdateNodeOpsOwnerResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateNodeOpsOwnerResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateNodeOpsOwnerResponseBodyData struct {
	// The failure reason. This value is empty if the operation was successful.
	//
	// example:
	//
	// test
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The node ID. This corresponds to the Id in the NodeIdList request parameter.
	//
	// example:
	//
	// n_8198365584737107968
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The node name.
	//
	// example:
	//
	// demo_node
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node source type.
	//
	// example:
	//
	// DATA_PROCESS
	NodeFromType *string `json:"NodeFromType,omitempty" xml:"NodeFromType,omitempty"`
	// The change result status for the node.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateNodeOpsOwnerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeOpsOwnerResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateNodeOpsOwnerResponseBodyData) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *UpdateNodeOpsOwnerResponseBodyData) GetId() *string {
	return s.Id
}

func (s *UpdateNodeOpsOwnerResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateNodeOpsOwnerResponseBodyData) GetNodeFromType() *string {
	return s.NodeFromType
}

func (s *UpdateNodeOpsOwnerResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateNodeOpsOwnerResponseBodyData) SetErrorInfo(v string) *UpdateNodeOpsOwnerResponseBodyData {
	s.ErrorInfo = &v
	return s
}

func (s *UpdateNodeOpsOwnerResponseBodyData) SetId(v string) *UpdateNodeOpsOwnerResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateNodeOpsOwnerResponseBodyData) SetName(v string) *UpdateNodeOpsOwnerResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateNodeOpsOwnerResponseBodyData) SetNodeFromType(v string) *UpdateNodeOpsOwnerResponseBodyData {
	s.NodeFromType = &v
	return s
}

func (s *UpdateNodeOpsOwnerResponseBodyData) SetStatus(v string) *UpdateNodeOpsOwnerResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateNodeOpsOwnerResponseBodyData) Validate() error {
	return dara.Validate(s)
}
