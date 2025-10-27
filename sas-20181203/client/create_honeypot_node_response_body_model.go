// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHoneypotNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateHoneypotNodeResponseBody
	GetCode() *string
	SetHoneypotNode(v *CreateHoneypotNodeResponseBodyHoneypotNode) *CreateHoneypotNodeResponseBody
	GetHoneypotNode() *CreateHoneypotNodeResponseBodyHoneypotNode
	SetHttpStatusCode(v int32) *CreateHoneypotNodeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateHoneypotNodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateHoneypotNodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateHoneypotNodeResponseBody
	GetSuccess() *bool
}

type CreateHoneypotNodeResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	HoneypotNode *CreateHoneypotNodeResponseBodyHoneypotNode `json:"HoneypotNode,omitempty" xml:"HoneypotNode,omitempty" type:"Struct"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 028CF634-5268-5660-9575-48C9ED6BF880
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateHoneypotNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHoneypotNodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHoneypotNodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateHoneypotNodeResponseBody) GetHoneypotNode() *CreateHoneypotNodeResponseBodyHoneypotNode {
	return s.HoneypotNode
}

func (s *CreateHoneypotNodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateHoneypotNodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateHoneypotNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHoneypotNodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateHoneypotNodeResponseBody) SetCode(v string) *CreateHoneypotNodeResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHoneypotNodeResponseBody) SetHoneypotNode(v *CreateHoneypotNodeResponseBodyHoneypotNode) *CreateHoneypotNodeResponseBody {
	s.HoneypotNode = v
	return s
}

func (s *CreateHoneypotNodeResponseBody) SetHttpStatusCode(v int32) *CreateHoneypotNodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateHoneypotNodeResponseBody) SetMessage(v string) *CreateHoneypotNodeResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHoneypotNodeResponseBody) SetRequestId(v string) *CreateHoneypotNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHoneypotNodeResponseBody) SetSuccess(v bool) *CreateHoneypotNodeResponseBody {
	s.Success = &v
	return s
}

func (s *CreateHoneypotNodeResponseBody) Validate() error {
	if s.HoneypotNode != nil {
		if err := s.HoneypotNode.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateHoneypotNodeResponseBodyHoneypotNode struct {
	// The ID of the management node.
	//
	// example:
	//
	// 37a15ff1-3475-4897-aa6c-f7fd9122****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s CreateHoneypotNodeResponseBodyHoneypotNode) String() string {
	return dara.Prettify(s)
}

func (s CreateHoneypotNodeResponseBodyHoneypotNode) GoString() string {
	return s.String()
}

func (s *CreateHoneypotNodeResponseBodyHoneypotNode) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateHoneypotNodeResponseBodyHoneypotNode) SetNodeId(v string) *CreateHoneypotNodeResponseBodyHoneypotNode {
	s.NodeId = &v
	return s
}

func (s *CreateHoneypotNodeResponseBodyHoneypotNode) Validate() error {
	return dara.Validate(s)
}
