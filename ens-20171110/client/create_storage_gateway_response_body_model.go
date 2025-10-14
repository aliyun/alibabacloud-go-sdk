// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStorageGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationId(v []*CreateStorageGatewayResponseBodyAllocationId) *CreateStorageGatewayResponseBody
	GetAllocationId() []*CreateStorageGatewayResponseBodyAllocationId
	SetBizStatusCode(v string) *CreateStorageGatewayResponseBody
	GetBizStatusCode() *string
	SetRequestId(v string) *CreateStorageGatewayResponseBody
	GetRequestId() *string
	SetUnAllocationId(v []*CreateStorageGatewayResponseBodyUnAllocationId) *CreateStorageGatewayResponseBody
	GetUnAllocationId() []*CreateStorageGatewayResponseBodyUnAllocationId
}

type CreateStorageGatewayResponseBody struct {
	// The list of created nodes.
	AllocationId []*CreateStorageGatewayResponseBodyAllocationId `json:"AllocationId,omitempty" xml:"AllocationId,omitempty" type:"Repeated"`
	// The success status code.
	//
	// 	- **PartSuccess**: partially succeeded.
	//
	// 	- **AllSuccess**: all succeeded.
	//
	// example:
	//
	// AllSuccess
	BizStatusCode *string `json:"BizStatusCode,omitempty" xml:"BizStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EA3758E0-8899-17D3-9526-5F62CF33A586
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of nodes that are not created.
	UnAllocationId []*CreateStorageGatewayResponseBodyUnAllocationId `json:"UnAllocationId,omitempty" xml:"UnAllocationId,omitempty" type:"Repeated"`
}

func (s CreateStorageGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStorageGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStorageGatewayResponseBody) GetAllocationId() []*CreateStorageGatewayResponseBodyAllocationId {
	return s.AllocationId
}

func (s *CreateStorageGatewayResponseBody) GetBizStatusCode() *string {
	return s.BizStatusCode
}

func (s *CreateStorageGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStorageGatewayResponseBody) GetUnAllocationId() []*CreateStorageGatewayResponseBodyUnAllocationId {
	return s.UnAllocationId
}

func (s *CreateStorageGatewayResponseBody) SetAllocationId(v []*CreateStorageGatewayResponseBodyAllocationId) *CreateStorageGatewayResponseBody {
	s.AllocationId = v
	return s
}

func (s *CreateStorageGatewayResponseBody) SetBizStatusCode(v string) *CreateStorageGatewayResponseBody {
	s.BizStatusCode = &v
	return s
}

func (s *CreateStorageGatewayResponseBody) SetRequestId(v string) *CreateStorageGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStorageGatewayResponseBody) SetUnAllocationId(v []*CreateStorageGatewayResponseBodyUnAllocationId) *CreateStorageGatewayResponseBody {
	s.UnAllocationId = v
	return s
}

func (s *CreateStorageGatewayResponseBody) Validate() error {
	if s.AllocationId != nil {
		for _, item := range s.AllocationId {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UnAllocationId != nil {
		for _, item := range s.UnAllocationId {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateStorageGatewayResponseBodyAllocationId struct {
	// The ID of the node.
	//
	// example:
	//
	// cn-chengdu-telecom-4
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// e426409223
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateStorageGatewayResponseBodyAllocationId) String() string {
	return dara.Prettify(s)
}

func (s CreateStorageGatewayResponseBodyAllocationId) GoString() string {
	return s.String()
}

func (s *CreateStorageGatewayResponseBodyAllocationId) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CreateStorageGatewayResponseBodyAllocationId) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateStorageGatewayResponseBodyAllocationId) SetEnsRegionId(v string) *CreateStorageGatewayResponseBodyAllocationId {
	s.EnsRegionId = &v
	return s
}

func (s *CreateStorageGatewayResponseBodyAllocationId) SetInstanceId(v string) *CreateStorageGatewayResponseBodyAllocationId {
	s.InstanceId = &v
	return s
}

func (s *CreateStorageGatewayResponseBodyAllocationId) Validate() error {
	return dara.Validate(s)
}

type CreateStorageGatewayResponseBodyUnAllocationId struct {
	// The ID of the node.
	//
	// example:
	//
	// cn-chengdu-26
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// e426409258
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateStorageGatewayResponseBodyUnAllocationId) String() string {
	return dara.Prettify(s)
}

func (s CreateStorageGatewayResponseBodyUnAllocationId) GoString() string {
	return s.String()
}

func (s *CreateStorageGatewayResponseBodyUnAllocationId) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CreateStorageGatewayResponseBodyUnAllocationId) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateStorageGatewayResponseBodyUnAllocationId) SetEnsRegionId(v string) *CreateStorageGatewayResponseBodyUnAllocationId {
	s.EnsRegionId = &v
	return s
}

func (s *CreateStorageGatewayResponseBodyUnAllocationId) SetInstanceId(v string) *CreateStorageGatewayResponseBodyUnAllocationId {
	s.InstanceId = &v
	return s
}

func (s *CreateStorageGatewayResponseBodyUnAllocationId) Validate() error {
	return dara.Validate(s)
}
