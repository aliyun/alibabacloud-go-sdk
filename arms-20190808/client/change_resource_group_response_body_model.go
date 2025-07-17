// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ChangeResourceGroupResponseBody
	GetCode() *string
	SetData(v *ChangeResourceGroupResponseBodyData) *ChangeResourceGroupResponseBody
	GetData() *ChangeResourceGroupResponseBodyData
	SetMessage(v string) *ChangeResourceGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChangeResourceGroupResponseBody
	GetRequestId() *string
}

type ChangeResourceGroupResponseBody struct {
	// The status code or error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data *ChangeResourceGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 46355DD8-FC56-40C5-BFC6-269DE4F9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChangeResourceGroupResponseBody) GetData() *ChangeResourceGroupResponseBodyData {
	return s.Data
}

func (s *ChangeResourceGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChangeResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeResourceGroupResponseBody) SetCode(v string) *ChangeResourceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetData(v *ChangeResourceGroupResponseBodyData) *ChangeResourceGroupResponseBody {
	s.Data = v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetMessage(v string) *ChangeResourceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type ChangeResourceGroupResponseBodyData struct {
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2vezare****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// ggxw4lnjuz@cfd34a78f******
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s ChangeResourceGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ChangeResourceGroupResponseBodyData) GetResourceId() *string {
	return s.ResourceId
}

func (s *ChangeResourceGroupResponseBodyData) SetResourceGroupId(v string) *ChangeResourceGroupResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupResponseBodyData) SetResourceId(v string) *ChangeResourceGroupResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
