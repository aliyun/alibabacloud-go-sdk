// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceSharesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListResourceSharesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListResourceSharesResponseBody
	GetRequestId() *string
	SetResourceShares(v []*ListResourceSharesResponseBodyResourceShares) *ListResourceSharesResponseBody
	GetResourceShares() []*ListResourceSharesResponseBodyResourceShares
}

type ListResourceSharesResponseBody struct {
	// The `token` that is used to initiate the next request if the response of the current request is truncated. You can use the token to initiate another request and obtain the remaining records.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2F23CFB6-A721-4E90-AC1E-0E30FA8B45DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resource shares.
	ResourceShares []*ListResourceSharesResponseBodyResourceShares `json:"ResourceShares,omitempty" xml:"ResourceShares,omitempty" type:"Repeated"`
}

func (s ListResourceSharesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceSharesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceSharesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceSharesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceSharesResponseBody) GetResourceShares() []*ListResourceSharesResponseBodyResourceShares {
	return s.ResourceShares
}

func (s *ListResourceSharesResponseBody) SetNextToken(v string) *ListResourceSharesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourceSharesResponseBody) SetRequestId(v string) *ListResourceSharesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceSharesResponseBody) SetResourceShares(v []*ListResourceSharesResponseBodyResourceShares) *ListResourceSharesResponseBody {
	s.ResourceShares = v
	return s
}

func (s *ListResourceSharesResponseBody) Validate() error {
	if s.ResourceShares != nil {
		for _, item := range s.ResourceShares {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceSharesResponseBodyResourceShares struct {
	// Indicates whether resources in the resource share can be shared with accounts outside the resource directory. Valid values:
	//
	// 	- false: Resources in the resource share can be shared only with accounts in the resource directory.
	//
	// 	- true: Resources in the resource share can be shared with both accounts in the resource directory and accounts outside the resource directory.
	//
	// example:
	//
	// false
	AllowExternalTargets *bool `json:"AllowExternalTargets,omitempty" xml:"AllowExternalTargets,omitempty"`
	// The time when the resource share was created.
	//
	// example:
	//
	// 2020-12-03T02:20:31.292Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekz5nlvlak****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the resource share.
	//
	// example:
	//
	// rs-PqysnzIj****
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The name of the resource share.
	//
	// example:
	//
	// test
	ResourceShareName *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	// The owner of the resource share.
	//
	// example:
	//
	// 151266687691****
	ResourceShareOwner *string `json:"ResourceShareOwner,omitempty" xml:"ResourceShareOwner,omitempty"`
	// The status of the resource share. Valid values:
	//
	// 	- Active
	//
	// 	- Pending
	//
	// 	- Deleting
	//
	// 	- Deleted
	//
	// >  The system automatically deletes the records of resource shares in the Deleted state within 48 hours to 96 hours after you delete the resource shares.
	//
	// example:
	//
	// Active
	ResourceShareStatus *string `json:"ResourceShareStatus,omitempty" xml:"ResourceShareStatus,omitempty"`
	// The tags.
	Tags []*ListResourceSharesResponseBodyResourceSharesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The time when the resource share was updated.
	//
	// example:
	//
	// 2020-12-03T08:01:43.638Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListResourceSharesResponseBodyResourceShares) String() string {
	return dara.Prettify(s)
}

func (s ListResourceSharesResponseBodyResourceShares) GoString() string {
	return s.String()
}

func (s *ListResourceSharesResponseBodyResourceShares) GetAllowExternalTargets() *bool {
	return s.AllowExternalTargets
}

func (s *ListResourceSharesResponseBodyResourceShares) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListResourceSharesResponseBodyResourceShares) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListResourceSharesResponseBodyResourceShares) GetResourceShareId() *string {
	return s.ResourceShareId
}

func (s *ListResourceSharesResponseBodyResourceShares) GetResourceShareName() *string {
	return s.ResourceShareName
}

func (s *ListResourceSharesResponseBodyResourceShares) GetResourceShareOwner() *string {
	return s.ResourceShareOwner
}

func (s *ListResourceSharesResponseBodyResourceShares) GetResourceShareStatus() *string {
	return s.ResourceShareStatus
}

func (s *ListResourceSharesResponseBodyResourceShares) GetTags() []*ListResourceSharesResponseBodyResourceSharesTags {
	return s.Tags
}

func (s *ListResourceSharesResponseBodyResourceShares) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListResourceSharesResponseBodyResourceShares) SetAllowExternalTargets(v bool) *ListResourceSharesResponseBodyResourceShares {
	s.AllowExternalTargets = &v
	return s
}

func (s *ListResourceSharesResponseBodyResourceShares) SetCreateTime(v string) *ListResourceSharesResponseBodyResourceShares {
	s.CreateTime = &v
	return s
}

func (s *ListResourceSharesResponseBodyResourceShares) SetResourceGroupId(v string) *ListResourceSharesResponseBodyResourceShares {
	s.ResourceGroupId = &v
	return s
}

func (s *ListResourceSharesResponseBodyResourceShares) SetResourceShareId(v string) *ListResourceSharesResponseBodyResourceShares {
	s.ResourceShareId = &v
	return s
}

func (s *ListResourceSharesResponseBodyResourceShares) SetResourceShareName(v string) *ListResourceSharesResponseBodyResourceShares {
	s.ResourceShareName = &v
	return s
}

func (s *ListResourceSharesResponseBodyResourceShares) SetResourceShareOwner(v string) *ListResourceSharesResponseBodyResourceShares {
	s.ResourceShareOwner = &v
	return s
}

func (s *ListResourceSharesResponseBodyResourceShares) SetResourceShareStatus(v string) *ListResourceSharesResponseBodyResourceShares {
	s.ResourceShareStatus = &v
	return s
}

func (s *ListResourceSharesResponseBodyResourceShares) SetTags(v []*ListResourceSharesResponseBodyResourceSharesTags) *ListResourceSharesResponseBodyResourceShares {
	s.Tags = v
	return s
}

func (s *ListResourceSharesResponseBodyResourceShares) SetUpdateTime(v string) *ListResourceSharesResponseBodyResourceShares {
	s.UpdateTime = &v
	return s
}

func (s *ListResourceSharesResponseBodyResourceShares) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceSharesResponseBodyResourceSharesTags struct {
	// The tag key.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListResourceSharesResponseBodyResourceSharesTags) String() string {
	return dara.Prettify(s)
}

func (s ListResourceSharesResponseBodyResourceSharesTags) GoString() string {
	return s.String()
}

func (s *ListResourceSharesResponseBodyResourceSharesTags) GetKey() *string {
	return s.Key
}

func (s *ListResourceSharesResponseBodyResourceSharesTags) GetValue() *string {
	return s.Value
}

func (s *ListResourceSharesResponseBodyResourceSharesTags) SetKey(v string) *ListResourceSharesResponseBodyResourceSharesTags {
	s.Key = &v
	return s
}

func (s *ListResourceSharesResponseBodyResourceSharesTags) SetValue(v string) *ListResourceSharesResponseBodyResourceSharesTags {
	s.Value = &v
	return s
}

func (s *ListResourceSharesResponseBodyResourceSharesTags) Validate() error {
	return dara.Validate(s)
}
