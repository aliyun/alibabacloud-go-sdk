// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesForExpressConnectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTagResourcesForExpressConnectResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTagResourcesForExpressConnectResponseBody
	GetRequestId() *string
	SetTagResources(v *ListTagResourcesForExpressConnectResponseBodyTagResources) *ListTagResourcesForExpressConnectResponseBody
	GetTagResources() *ListTagResourcesForExpressConnectResponseBodyTagResources
}

type ListTagResourcesForExpressConnectResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value is returned for **NextToken**, the value can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags that are added to the resource.
	TagResources *ListTagResourcesForExpressConnectResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesForExpressConnectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesForExpressConnectResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesForExpressConnectResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagResourcesForExpressConnectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagResourcesForExpressConnectResponseBody) GetTagResources() *ListTagResourcesForExpressConnectResponseBodyTagResources {
	return s.TagResources
}

func (s *ListTagResourcesForExpressConnectResponseBody) SetNextToken(v string) *ListTagResourcesForExpressConnectResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesForExpressConnectResponseBody) SetRequestId(v string) *ListTagResourcesForExpressConnectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesForExpressConnectResponseBody) SetTagResources(v *ListTagResourcesForExpressConnectResponseBodyTagResources) *ListTagResourcesForExpressConnectResponseBody {
	s.TagResources = v
	return s
}

func (s *ListTagResourcesForExpressConnectResponseBody) Validate() error {
	if s.TagResources != nil {
		if err := s.TagResources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTagResourcesForExpressConnectResponseBodyTagResources struct {
	TagResource []*ListTagResourcesForExpressConnectResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesForExpressConnectResponseBodyTagResources) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesForExpressConnectResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesForExpressConnectResponseBodyTagResources) GetTagResource() []*ListTagResourcesForExpressConnectResponseBodyTagResourcesTagResource {
	return s.TagResource
}

func (s *ListTagResourcesForExpressConnectResponseBodyTagResources) SetTagResource(v []*ListTagResourcesForExpressConnectResponseBodyTagResourcesTagResource) *ListTagResourcesForExpressConnectResponseBodyTagResources {
	s.TagResource = v
	return s
}

func (s *ListTagResourcesForExpressConnectResponseBodyTagResources) Validate() error {
	if s.TagResource != nil {
		for _, item := range s.TagResource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTagResourcesForExpressConnectResponseBodyTagResourcesTagResource struct {
	// The resource ID.
	//
	// example:
	//
	// pc-bp16qjewdsunr41m1****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- **PHYSICALCONNECTION**: Express Connect circuit.
	//
	// 	- **VIRTUALBORDERROUTER**: VBR.
	//
	// 	- **ROUTERINTERFACE**: router interface.
	//
	// example:
	//
	// PHYSICALCONNECTION
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the tag that is added to the resource.
	//
	// example:
	//
	// FinanceDept
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag that is added to the resource.
	//
	// example:
	//
	// FinanceJoshua
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesForExpressConnectResponseBodyTagResourcesTagResource) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesForExpressConnectResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesForExpressConnectResponseBodyTagResourcesTagResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTagResourcesForExpressConnectResponseBodyTagResourcesTagResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesForExpressConnectResponseBodyTagResourcesTagResource) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTagResourcesForExpressConnectResponseBodyTagResourcesTagResource) GetTagValue() *string {
	return s.TagValue
}

func (s *ListTagResourcesForExpressConnectResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesForExpressConnectResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesForExpressConnectResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesForExpressConnectResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesForExpressConnectResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesForExpressConnectResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesForExpressConnectResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesForExpressConnectResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesForExpressConnectResponseBodyTagResourcesTagResource) Validate() error {
	return dara.Validate(s)
}
