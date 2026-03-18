// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateAgentResourceResponseBodyData) *CreateAgentResourceResponseBody
	GetData() *CreateAgentResourceResponseBodyData
	SetErrCode(v string) *CreateAgentResourceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateAgentResourceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *CreateAgentResourceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateAgentResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAgentResourceResponseBody
	GetSuccess() *bool
}

type CreateAgentResourceResponseBody struct {
	Data *CreateAgentResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// B67D142D-D54E-184F-A306-22BDC01B2XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAgentResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentResourceResponseBody) GetData() *CreateAgentResourceResponseBodyData {
	return s.Data
}

func (s *CreateAgentResourceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateAgentResourceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateAgentResourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateAgentResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgentResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAgentResourceResponseBody) SetData(v *CreateAgentResourceResponseBodyData) *CreateAgentResourceResponseBody {
	s.Data = v
	return s
}

func (s *CreateAgentResourceResponseBody) SetErrCode(v string) *CreateAgentResourceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateAgentResourceResponseBody) SetErrMessage(v string) *CreateAgentResourceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateAgentResourceResponseBody) SetHttpStatusCode(v int32) *CreateAgentResourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAgentResourceResponseBody) SetRequestId(v string) *CreateAgentResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentResourceResponseBody) SetSuccess(v bool) *CreateAgentResourceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAgentResourceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentResourceResponseBodyData struct {
	// example:
	//
	// ng-5e2ba600fee3****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// example:
	//
	// 241526000650XXX
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateAgentResourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAgentResourceResponseBodyData) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *CreateAgentResourceResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateAgentResourceResponseBodyData) SetNodeGroupId(v string) *CreateAgentResourceResponseBodyData {
	s.NodeGroupId = &v
	return s
}

func (s *CreateAgentResourceResponseBodyData) SetOrderId(v int64) *CreateAgentResourceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateAgentResourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
