// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalNodeContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPhysicalNodeContentResponseBody
	GetCode() *string
	SetData(v *GetPhysicalNodeContentResponseBodyData) *GetPhysicalNodeContentResponseBody
	GetData() *GetPhysicalNodeContentResponseBodyData
	SetHttpStatusCode(v int32) *GetPhysicalNodeContentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPhysicalNodeContentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPhysicalNodeContentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPhysicalNodeContentResponseBody
	GetSuccess() *bool
}

type GetPhysicalNodeContentResponseBody struct {
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetPhysicalNodeContentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPhysicalNodeContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeContentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPhysicalNodeContentResponseBody) GetData() *GetPhysicalNodeContentResponseBodyData {
	return s.Data
}

func (s *GetPhysicalNodeContentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPhysicalNodeContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPhysicalNodeContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPhysicalNodeContentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPhysicalNodeContentResponseBody) SetCode(v string) *GetPhysicalNodeContentResponseBody {
	s.Code = &v
	return s
}

func (s *GetPhysicalNodeContentResponseBody) SetData(v *GetPhysicalNodeContentResponseBodyData) *GetPhysicalNodeContentResponseBody {
	s.Data = v
	return s
}

func (s *GetPhysicalNodeContentResponseBody) SetHttpStatusCode(v int32) *GetPhysicalNodeContentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPhysicalNodeContentResponseBody) SetMessage(v string) *GetPhysicalNodeContentResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhysicalNodeContentResponseBody) SetRequestId(v string) *GetPhysicalNodeContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhysicalNodeContentResponseBody) SetSuccess(v bool) *GetPhysicalNodeContentResponseBody {
	s.Success = &v
	return s
}

func (s *GetPhysicalNodeContentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPhysicalNodeContentResponseBodyData struct {
	// example:
	//
	// select 1;
	CodeContent *string `json:"CodeContent,omitempty" xml:"CodeContent,omitempty"`
	// example:
	//
	// n_232411
	NodeId   *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s GetPhysicalNodeContentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeContentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeContentResponseBodyData) GetCodeContent() *string {
	return s.CodeContent
}

func (s *GetPhysicalNodeContentResponseBodyData) GetNodeId() *string {
	return s.NodeId
}

func (s *GetPhysicalNodeContentResponseBodyData) GetNodeName() *string {
	return s.NodeName
}

func (s *GetPhysicalNodeContentResponseBodyData) SetCodeContent(v string) *GetPhysicalNodeContentResponseBodyData {
	s.CodeContent = &v
	return s
}

func (s *GetPhysicalNodeContentResponseBodyData) SetNodeId(v string) *GetPhysicalNodeContentResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *GetPhysicalNodeContentResponseBodyData) SetNodeName(v string) *GetPhysicalNodeContentResponseBodyData {
	s.NodeName = &v
	return s
}

func (s *GetPhysicalNodeContentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
