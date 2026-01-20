// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSharedTargetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListSharedTargetsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSharedTargetsResponseBody
	GetRequestId() *string
	SetSharedTargets(v []*ListSharedTargetsResponseBodySharedTargets) *ListSharedTargetsResponseBody
	GetSharedTargets() []*ListSharedTargetsResponseBodySharedTargets
}

type ListSharedTargetsResponseBody struct {
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of `NextToken`.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04677DCA-7C33-464B-8811-1B1DA3C3D197
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the principals.
	SharedTargets []*ListSharedTargetsResponseBodySharedTargets `json:"SharedTargets,omitempty" xml:"SharedTargets,omitempty" type:"Repeated"`
}

func (s ListSharedTargetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSharedTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSharedTargetsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSharedTargetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSharedTargetsResponseBody) GetSharedTargets() []*ListSharedTargetsResponseBodySharedTargets {
	return s.SharedTargets
}

func (s *ListSharedTargetsResponseBody) SetNextToken(v string) *ListSharedTargetsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSharedTargetsResponseBody) SetRequestId(v string) *ListSharedTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSharedTargetsResponseBody) SetSharedTargets(v []*ListSharedTargetsResponseBodySharedTargets) *ListSharedTargetsResponseBody {
	s.SharedTargets = v
	return s
}

func (s *ListSharedTargetsResponseBody) Validate() error {
	if s.SharedTargets != nil {
		for _, item := range s.SharedTargets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSharedTargetsResponseBodySharedTargets struct {
	// The time when the principal was associated with the resource share.
	//
	// example:
	//
	// 2020-12-07T09:16:59.905Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the principal is outside the resource directory. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	External *bool `json:"External,omitempty" xml:"External,omitempty"`
	// The ID of the resource share.
	//
	// example:
	//
	// rs-6GRmdD3X****
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The ID of the principal or resource owner.
	//
	// 	- If the value of `ResourceOwner` is `Self`, the value of this parameter is the ID of a principal.
	//
	// 	- If the value of `ResourceOwner` is `OtherAccounts`, the value of this parameter is the ID of a resource owner.
	//
	// example:
	//
	// 114240524784****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The properties of the principal, such as the time range within which the resource is shared.
	//
	// >  This parameter is returned only if the principal is an Alibaba Cloud service.
	//
	// example:
	//
	// {
	//
	//     "timeRange":{
	//
	//         "timeRangeType":"timeRange",
	//
	//         "beginAtTime":"00:00",
	//
	//         "timezone":"UTC+8",
	//
	//         "endAtTime":"19:59"
	//
	//     }
	//
	// }
	TargetProperty *string `json:"TargetProperty,omitempty" xml:"TargetProperty,omitempty"`
	// The time when the association of the principal was updated.
	//
	// example:
	//
	// 2020-12-07T09:16:59.905Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListSharedTargetsResponseBodySharedTargets) String() string {
	return dara.Prettify(s)
}

func (s ListSharedTargetsResponseBodySharedTargets) GoString() string {
	return s.String()
}

func (s *ListSharedTargetsResponseBodySharedTargets) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSharedTargetsResponseBodySharedTargets) GetExternal() *bool {
	return s.External
}

func (s *ListSharedTargetsResponseBodySharedTargets) GetResourceShareId() *string {
	return s.ResourceShareId
}

func (s *ListSharedTargetsResponseBodySharedTargets) GetTargetId() *string {
	return s.TargetId
}

func (s *ListSharedTargetsResponseBodySharedTargets) GetTargetProperty() *string {
	return s.TargetProperty
}

func (s *ListSharedTargetsResponseBodySharedTargets) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListSharedTargetsResponseBodySharedTargets) SetCreateTime(v string) *ListSharedTargetsResponseBodySharedTargets {
	s.CreateTime = &v
	return s
}

func (s *ListSharedTargetsResponseBodySharedTargets) SetExternal(v bool) *ListSharedTargetsResponseBodySharedTargets {
	s.External = &v
	return s
}

func (s *ListSharedTargetsResponseBodySharedTargets) SetResourceShareId(v string) *ListSharedTargetsResponseBodySharedTargets {
	s.ResourceShareId = &v
	return s
}

func (s *ListSharedTargetsResponseBodySharedTargets) SetTargetId(v string) *ListSharedTargetsResponseBodySharedTargets {
	s.TargetId = &v
	return s
}

func (s *ListSharedTargetsResponseBodySharedTargets) SetTargetProperty(v string) *ListSharedTargetsResponseBodySharedTargets {
	s.TargetProperty = &v
	return s
}

func (s *ListSharedTargetsResponseBodySharedTargets) SetUpdateTime(v string) *ListSharedTargetsResponseBodySharedTargets {
	s.UpdateTime = &v
	return s
}

func (s *ListSharedTargetsResponseBodySharedTargets) Validate() error {
	return dara.Validate(s)
}
