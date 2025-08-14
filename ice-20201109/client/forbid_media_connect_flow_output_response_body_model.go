// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForbidMediaConnectFlowOutputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *ForbidMediaConnectFlowOutputResponseBody
	GetContent() *string
	SetDescription(v string) *ForbidMediaConnectFlowOutputResponseBody
	GetDescription() *string
	SetRequestId(v string) *ForbidMediaConnectFlowOutputResponseBody
	GetRequestId() *string
	SetRetCode(v int32) *ForbidMediaConnectFlowOutputResponseBody
	GetRetCode() *int32
}

type ForbidMediaConnectFlowOutputResponseBody struct {
	// example:
	//
	// ""
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// OK
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1BCA0CFC-CBD4-5656-9D04-21B1FADBB92A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
}

func (s ForbidMediaConnectFlowOutputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ForbidMediaConnectFlowOutputResponseBody) GoString() string {
	return s.String()
}

func (s *ForbidMediaConnectFlowOutputResponseBody) GetContent() *string {
	return s.Content
}

func (s *ForbidMediaConnectFlowOutputResponseBody) GetDescription() *string {
	return s.Description
}

func (s *ForbidMediaConnectFlowOutputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ForbidMediaConnectFlowOutputResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *ForbidMediaConnectFlowOutputResponseBody) SetContent(v string) *ForbidMediaConnectFlowOutputResponseBody {
	s.Content = &v
	return s
}

func (s *ForbidMediaConnectFlowOutputResponseBody) SetDescription(v string) *ForbidMediaConnectFlowOutputResponseBody {
	s.Description = &v
	return s
}

func (s *ForbidMediaConnectFlowOutputResponseBody) SetRequestId(v string) *ForbidMediaConnectFlowOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *ForbidMediaConnectFlowOutputResponseBody) SetRetCode(v int32) *ForbidMediaConnectFlowOutputResponseBody {
	s.RetCode = &v
	return s
}

func (s *ForbidMediaConnectFlowOutputResponseBody) Validate() error {
	return dara.Validate(s)
}
