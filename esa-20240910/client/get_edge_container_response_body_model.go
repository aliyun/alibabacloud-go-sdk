// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBillingMode(v string) *GetEdgeContainerResponseBody
	GetBillingMode() *string
	SetInstanceId(v string) *GetEdgeContainerResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *GetEdgeContainerResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetEdgeContainerResponseBody
	GetStatus() *string
}

type GetEdgeContainerResponseBody struct {
	// example:
	//
	// container_95
	BillingMode *string `json:"BillingMode,omitempty" xml:"BillingMode,omitempty"`
	// example:
	//
	// esa-cn-jea67jfbs0x
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 9BEB8659-9CDE-5F2C-83E9-50F55277E844
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetEdgeContainerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerResponseBody) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerResponseBody) GetBillingMode() *string {
	return s.BillingMode
}

func (s *GetEdgeContainerResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetEdgeContainerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEdgeContainerResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetEdgeContainerResponseBody) SetBillingMode(v string) *GetEdgeContainerResponseBody {
	s.BillingMode = &v
	return s
}

func (s *GetEdgeContainerResponseBody) SetInstanceId(v string) *GetEdgeContainerResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetEdgeContainerResponseBody) SetRequestId(v string) *GetEdgeContainerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEdgeContainerResponseBody) SetStatus(v string) *GetEdgeContainerResponseBody {
	s.Status = &v
	return s
}

func (s *GetEdgeContainerResponseBody) Validate() error {
	return dara.Validate(s)
}
