// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMemoryResponseBody
	GetCode() *string
	SetData(v *CreateMemoryResponseBodyData) *CreateMemoryResponseBody
	GetData() *CreateMemoryResponseBodyData
	SetHttpStatusCode(v int32) *CreateMemoryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateMemoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMemoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMemoryResponseBody
	GetSuccess() *bool
}

type CreateMemoryResponseBody struct {
	// example:
	//
	// 200
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateMemoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance instance-002\\"`curl h33E1En5.popscan.xaliyun.com` does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 31033EC0-6968-5610-8328-708B59508E5A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMemoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMemoryResponseBody) GetData() *CreateMemoryResponseBodyData {
	return s.Data
}

func (s *CreateMemoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateMemoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMemoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMemoryResponseBody) SetCode(v string) *CreateMemoryResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMemoryResponseBody) SetData(v *CreateMemoryResponseBodyData) *CreateMemoryResponseBody {
	s.Data = v
	return s
}

func (s *CreateMemoryResponseBody) SetHttpStatusCode(v int32) *CreateMemoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateMemoryResponseBody) SetMessage(v string) *CreateMemoryResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMemoryResponseBody) SetRequestId(v string) *CreateMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMemoryResponseBody) SetSuccess(v bool) *CreateMemoryResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMemoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMemoryResponseBodyData struct {
	MemoryNodes []*CreateMemoryResponseBodyDataMemoryNodes `json:"MemoryNodes,omitempty" xml:"MemoryNodes,omitempty" type:"Repeated"`
}

func (s CreateMemoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMemoryResponseBodyData) GetMemoryNodes() []*CreateMemoryResponseBodyDataMemoryNodes {
	return s.MemoryNodes
}

func (s *CreateMemoryResponseBodyData) SetMemoryNodes(v []*CreateMemoryResponseBodyDataMemoryNodes) *CreateMemoryResponseBodyData {
	s.MemoryNodes = v
	return s
}

func (s *CreateMemoryResponseBodyData) Validate() error {
	if s.MemoryNodes != nil {
		for _, item := range s.MemoryNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateMemoryResponseBodyDataMemoryNodes struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// ADD
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 384dc4786b9d4f5a8cab0d83112cd5a8
	MemoryNodeId *string `json:"MemoryNodeId,omitempty" xml:"MemoryNodeId,omitempty"`
}

func (s CreateMemoryResponseBodyDataMemoryNodes) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryResponseBodyDataMemoryNodes) GoString() string {
	return s.String()
}

func (s *CreateMemoryResponseBodyDataMemoryNodes) GetContent() *string {
	return s.Content
}

func (s *CreateMemoryResponseBodyDataMemoryNodes) GetEvent() *string {
	return s.Event
}

func (s *CreateMemoryResponseBodyDataMemoryNodes) GetMemoryNodeId() *string {
	return s.MemoryNodeId
}

func (s *CreateMemoryResponseBodyDataMemoryNodes) SetContent(v string) *CreateMemoryResponseBodyDataMemoryNodes {
	s.Content = &v
	return s
}

func (s *CreateMemoryResponseBodyDataMemoryNodes) SetEvent(v string) *CreateMemoryResponseBodyDataMemoryNodes {
	s.Event = &v
	return s
}

func (s *CreateMemoryResponseBodyDataMemoryNodes) SetMemoryNodeId(v string) *CreateMemoryResponseBodyDataMemoryNodes {
	s.MemoryNodeId = &v
	return s
}

func (s *CreateMemoryResponseBodyDataMemoryNodes) Validate() error {
	return dara.Validate(s)
}
