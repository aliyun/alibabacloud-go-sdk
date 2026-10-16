// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MoveResourcesResponseBody
	GetRequestId() *string
	SetResponses(v []*MoveResourcesResponseBodyResponses) *MoveResourcesResponseBody
	GetResponses() []*MoveResourcesResponseBodyResponses
}

type MoveResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C00B89D3-3247-11DE-95D8-A7C01FB0AB4F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response results.
	Responses []*MoveResourcesResponseBodyResponses `json:"Responses,omitempty" xml:"Responses,omitempty" type:"Repeated"`
}

func (s MoveResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveResourcesResponseBody) GetResponses() []*MoveResourcesResponseBodyResponses {
	return s.Responses
}

func (s *MoveResourcesResponseBody) SetRequestId(v string) *MoveResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveResourcesResponseBody) SetResponses(v []*MoveResourcesResponseBodyResponses) *MoveResourcesResponseBody {
	s.Responses = v
	return s
}

func (s *MoveResourcesResponseBody) Validate() error {
	if s.Responses != nil {
		for _, item := range s.Responses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MoveResourcesResponseBodyResponses struct {
	// The error code.
	//
	// > This parameter is returned only when the resource transfer fails.
	//
	// example:
	//
	// NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// > This parameter is returned only when the resource transfer fails.
	//
	// example:
	//
	// No permission
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C00B89D3-3247-11DE-95D8-A7C01FB0AB4F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// i-23v38****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The cloud service code.
	//
	// example:
	//
	// ecs
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The status. Valid values:
	//
	// - SUCCESS: Succeeded.
	//
	// - FAIL: Failed.
	//
	// example:
	//
	// FAIL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s MoveResourcesResponseBodyResponses) String() string {
	return dara.Prettify(s)
}

func (s MoveResourcesResponseBodyResponses) GoString() string {
	return s.String()
}

func (s *MoveResourcesResponseBodyResponses) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *MoveResourcesResponseBodyResponses) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *MoveResourcesResponseBodyResponses) GetRegionId() *string {
	return s.RegionId
}

func (s *MoveResourcesResponseBodyResponses) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveResourcesResponseBodyResponses) GetResourceId() *string {
	return s.ResourceId
}

func (s *MoveResourcesResponseBodyResponses) GetResourceType() *string {
	return s.ResourceType
}

func (s *MoveResourcesResponseBodyResponses) GetService() *string {
	return s.Service
}

func (s *MoveResourcesResponseBodyResponses) GetStatus() *string {
	return s.Status
}

func (s *MoveResourcesResponseBodyResponses) SetErrorCode(v string) *MoveResourcesResponseBodyResponses {
	s.ErrorCode = &v
	return s
}

func (s *MoveResourcesResponseBodyResponses) SetErrorMsg(v string) *MoveResourcesResponseBodyResponses {
	s.ErrorMsg = &v
	return s
}

func (s *MoveResourcesResponseBodyResponses) SetRegionId(v string) *MoveResourcesResponseBodyResponses {
	s.RegionId = &v
	return s
}

func (s *MoveResourcesResponseBodyResponses) SetRequestId(v string) *MoveResourcesResponseBodyResponses {
	s.RequestId = &v
	return s
}

func (s *MoveResourcesResponseBodyResponses) SetResourceId(v string) *MoveResourcesResponseBodyResponses {
	s.ResourceId = &v
	return s
}

func (s *MoveResourcesResponseBodyResponses) SetResourceType(v string) *MoveResourcesResponseBodyResponses {
	s.ResourceType = &v
	return s
}

func (s *MoveResourcesResponseBodyResponses) SetService(v string) *MoveResourcesResponseBodyResponses {
	s.Service = &v
	return s
}

func (s *MoveResourcesResponseBodyResponses) SetStatus(v string) *MoveResourcesResponseBodyResponses {
	s.Status = &v
	return s
}

func (s *MoveResourcesResponseBodyResponses) Validate() error {
	return dara.Validate(s)
}
