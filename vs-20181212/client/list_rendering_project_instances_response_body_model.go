// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRenderingProjectInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRenderingInstances(v []*ListRenderingProjectInstancesResponseBodyRenderingInstances) *ListRenderingProjectInstancesResponseBody
	GetRenderingInstances() []*ListRenderingProjectInstancesResponseBodyRenderingInstances
	SetRequestId(v string) *ListRenderingProjectInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListRenderingProjectInstancesResponseBody
	GetTotalCount() *int64
}

type ListRenderingProjectInstancesResponseBody struct {
	// List of cloud application service instances
	RenderingInstances []*ListRenderingProjectInstancesResponseBodyRenderingInstances `json:"RenderingInstances,omitempty" xml:"RenderingInstances,omitempty" type:"Repeated"`
	// Request ID
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of cloud application service instances
	//
	// example:
	//
	// 8
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRenderingProjectInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingProjectInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRenderingProjectInstancesResponseBody) GetRenderingInstances() []*ListRenderingProjectInstancesResponseBodyRenderingInstances {
	return s.RenderingInstances
}

func (s *ListRenderingProjectInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRenderingProjectInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRenderingProjectInstancesResponseBody) SetRenderingInstances(v []*ListRenderingProjectInstancesResponseBodyRenderingInstances) *ListRenderingProjectInstancesResponseBody {
	s.RenderingInstances = v
	return s
}

func (s *ListRenderingProjectInstancesResponseBody) SetRequestId(v string) *ListRenderingProjectInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRenderingProjectInstancesResponseBody) SetTotalCount(v int64) *ListRenderingProjectInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRenderingProjectInstancesResponseBody) Validate() error {
	if s.RenderingInstances != nil {
		for _, item := range s.RenderingInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRenderingProjectInstancesResponseBodyRenderingInstances struct {
	// Time when the instance was associated with the project
	//
	// example:
	//
	// 2024-09-11T18:19:04+08:00
	AssociationTime *string `json:"AssociationTime,omitempty" xml:"AssociationTime,omitempty"`
	// Cloud application service instance
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// Status information for the project instance
	StateInfo *ListRenderingProjectInstancesResponseBodyRenderingInstancesStateInfo `json:"StateInfo,omitempty" xml:"StateInfo,omitempty" type:"Struct"`
}

func (s ListRenderingProjectInstancesResponseBodyRenderingInstances) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingProjectInstancesResponseBodyRenderingInstances) GoString() string {
	return s.String()
}

func (s *ListRenderingProjectInstancesResponseBodyRenderingInstances) GetAssociationTime() *string {
	return s.AssociationTime
}

func (s *ListRenderingProjectInstancesResponseBodyRenderingInstances) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ListRenderingProjectInstancesResponseBodyRenderingInstances) GetStateInfo() *ListRenderingProjectInstancesResponseBodyRenderingInstancesStateInfo {
	return s.StateInfo
}

func (s *ListRenderingProjectInstancesResponseBodyRenderingInstances) SetAssociationTime(v string) *ListRenderingProjectInstancesResponseBodyRenderingInstances {
	s.AssociationTime = &v
	return s
}

func (s *ListRenderingProjectInstancesResponseBodyRenderingInstances) SetRenderingInstanceId(v string) *ListRenderingProjectInstancesResponseBodyRenderingInstances {
	s.RenderingInstanceId = &v
	return s
}

func (s *ListRenderingProjectInstancesResponseBodyRenderingInstances) SetStateInfo(v *ListRenderingProjectInstancesResponseBodyRenderingInstancesStateInfo) *ListRenderingProjectInstancesResponseBodyRenderingInstances {
	s.StateInfo = v
	return s
}

func (s *ListRenderingProjectInstancesResponseBodyRenderingInstances) Validate() error {
	if s.StateInfo != nil {
		if err := s.StateInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRenderingProjectInstancesResponseBodyRenderingInstancesStateInfo struct {
	// Description of the current status
	//
	// example:
	//
	// 正在会话中
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Instance status
	//
	// example:
	//
	// InUse
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Time when the status was last updated
	//
	// example:
	//
	// 2024-11-11T18:19:04+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListRenderingProjectInstancesResponseBodyRenderingInstancesStateInfo) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingProjectInstancesResponseBodyRenderingInstancesStateInfo) GoString() string {
	return s.String()
}

func (s *ListRenderingProjectInstancesResponseBodyRenderingInstancesStateInfo) GetComment() *string {
	return s.Comment
}

func (s *ListRenderingProjectInstancesResponseBodyRenderingInstancesStateInfo) GetState() *string {
	return s.State
}

func (s *ListRenderingProjectInstancesResponseBodyRenderingInstancesStateInfo) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListRenderingProjectInstancesResponseBodyRenderingInstancesStateInfo) SetComment(v string) *ListRenderingProjectInstancesResponseBodyRenderingInstancesStateInfo {
	s.Comment = &v
	return s
}

func (s *ListRenderingProjectInstancesResponseBodyRenderingInstancesStateInfo) SetState(v string) *ListRenderingProjectInstancesResponseBodyRenderingInstancesStateInfo {
	s.State = &v
	return s
}

func (s *ListRenderingProjectInstancesResponseBodyRenderingInstancesStateInfo) SetUpdateTime(v string) *ListRenderingProjectInstancesResponseBodyRenderingInstancesStateInfo {
	s.UpdateTime = &v
	return s
}

func (s *ListRenderingProjectInstancesResponseBodyRenderingInstancesStateInfo) Validate() error {
	return dara.Validate(s)
}
