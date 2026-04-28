// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateWorkspaceResponseBodyData) *UpdateWorkspaceResponseBody
	GetData() *UpdateWorkspaceResponseBodyData
	SetOrderId(v string) *UpdateWorkspaceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *UpdateWorkspaceResponseBody
	GetRequestId() *string
}

type UpdateWorkspaceResponseBody struct {
	Data *UpdateWorkspaceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 979071171373334529
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResponseBody) GetData() *UpdateWorkspaceResponseBodyData {
	return s.Data
}

func (s *UpdateWorkspaceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *UpdateWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkspaceResponseBody) SetData(v *UpdateWorkspaceResponseBodyData) *UpdateWorkspaceResponseBody {
	s.Data = v
	return s
}

func (s *UpdateWorkspaceResponseBody) SetOrderId(v string) *UpdateWorkspaceResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpdateWorkspaceResponseBody) SetRequestId(v string) *UpdateWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkspaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkspaceResponseBodyData struct {
	// example:
	//
	// 1005565802416783361
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s UpdateWorkspaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *UpdateWorkspaceResponseBodyData) SetOrderId(v string) *UpdateWorkspaceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
