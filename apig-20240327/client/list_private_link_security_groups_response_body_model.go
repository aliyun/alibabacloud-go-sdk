// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateLinkSecurityGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPrivateLinkSecurityGroupsResponseBody
	GetCode() *string
	SetData(v *ListPrivateLinkSecurityGroupsResponseBodyData) *ListPrivateLinkSecurityGroupsResponseBody
	GetData() *ListPrivateLinkSecurityGroupsResponseBodyData
	SetMessage(v string) *ListPrivateLinkSecurityGroupsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPrivateLinkSecurityGroupsResponseBody
	GetRequestId() *string
}

type ListPrivateLinkSecurityGroupsResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *ListPrivateLinkSecurityGroupsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 479AE38F-A574-52F7-87EA-E91199999F9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListPrivateLinkSecurityGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateLinkSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrivateLinkSecurityGroupsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPrivateLinkSecurityGroupsResponseBody) GetData() *ListPrivateLinkSecurityGroupsResponseBodyData {
	return s.Data
}

func (s *ListPrivateLinkSecurityGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPrivateLinkSecurityGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrivateLinkSecurityGroupsResponseBody) SetCode(v string) *ListPrivateLinkSecurityGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListPrivateLinkSecurityGroupsResponseBody) SetData(v *ListPrivateLinkSecurityGroupsResponseBodyData) *ListPrivateLinkSecurityGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListPrivateLinkSecurityGroupsResponseBody) SetMessage(v string) *ListPrivateLinkSecurityGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListPrivateLinkSecurityGroupsResponseBody) SetRequestId(v string) *ListPrivateLinkSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrivateLinkSecurityGroupsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPrivateLinkSecurityGroupsResponseBodyData struct {
	// The source information.
	Items []*ListPrivateLinkSecurityGroupsResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of results returned for this query.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for the next page.
	//
	// example:
	//
	// None
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListPrivateLinkSecurityGroupsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateLinkSecurityGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPrivateLinkSecurityGroupsResponseBodyData) GetItems() []*ListPrivateLinkSecurityGroupsResponseBodyDataItems {
	return s.Items
}

func (s *ListPrivateLinkSecurityGroupsResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPrivateLinkSecurityGroupsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPrivateLinkSecurityGroupsResponseBodyData) SetItems(v []*ListPrivateLinkSecurityGroupsResponseBodyDataItems) *ListPrivateLinkSecurityGroupsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListPrivateLinkSecurityGroupsResponseBodyData) SetMaxResults(v int32) *ListPrivateLinkSecurityGroupsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListPrivateLinkSecurityGroupsResponseBodyData) SetNextToken(v string) *ListPrivateLinkSecurityGroupsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListPrivateLinkSecurityGroupsResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPrivateLinkSecurityGroupsResponseBodyDataItems struct {
	// The security group ID.
	//
	// example:
	//
	// sg-bp1ftp5sm9os***
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// The security group name.
	//
	// example:
	//
	// Product Center Cluster Security Group
	SecurityGroupName *string `json:"securityGroupName,omitempty" xml:"securityGroupName,omitempty"`
	// Specifies whether the resource is created by a service account. Valid values:
	//
	// - 0: The resource is not created by a service account.
	//
	// - 1: The resource is created by a service account.
	ServiceManaged *bool `json:"serviceManaged,omitempty" xml:"serviceManaged,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-uf61resqa9am***
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListPrivateLinkSecurityGroupsResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateLinkSecurityGroupsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListPrivateLinkSecurityGroupsResponseBodyDataItems) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListPrivateLinkSecurityGroupsResponseBodyDataItems) GetSecurityGroupName() *string {
	return s.SecurityGroupName
}

func (s *ListPrivateLinkSecurityGroupsResponseBodyDataItems) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *ListPrivateLinkSecurityGroupsResponseBodyDataItems) GetVpcId() *string {
	return s.VpcId
}

func (s *ListPrivateLinkSecurityGroupsResponseBodyDataItems) SetSecurityGroupId(v string) *ListPrivateLinkSecurityGroupsResponseBodyDataItems {
	s.SecurityGroupId = &v
	return s
}

func (s *ListPrivateLinkSecurityGroupsResponseBodyDataItems) SetSecurityGroupName(v string) *ListPrivateLinkSecurityGroupsResponseBodyDataItems {
	s.SecurityGroupName = &v
	return s
}

func (s *ListPrivateLinkSecurityGroupsResponseBodyDataItems) SetServiceManaged(v bool) *ListPrivateLinkSecurityGroupsResponseBodyDataItems {
	s.ServiceManaged = &v
	return s
}

func (s *ListPrivateLinkSecurityGroupsResponseBodyDataItems) SetVpcId(v string) *ListPrivateLinkSecurityGroupsResponseBodyDataItems {
	s.VpcId = &v
	return s
}

func (s *ListPrivateLinkSecurityGroupsResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
