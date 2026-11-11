// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRightsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAiDesktop(v bool) *ListRightsRequest
	GetAiDesktop() *bool
	SetBusinessSceneCode(v string) *ListRightsRequest
	GetBusinessSceneCode() *string
	SetMaxResults(v int32) *ListRightsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListRightsRequest
	GetNextToken() *string
	SetOrderField(v string) *ListRightsRequest
	GetOrderField() *string
	SetOrderType(v string) *ListRightsRequest
	GetOrderType() *string
	SetOverseas(v int32) *ListRightsRequest
	GetOverseas() *int32
	SetRightsActiveStatus(v string) *ListRightsRequest
	GetRightsActiveStatus() *string
	SetRightsAssignedUser(v string) *ListRightsRequest
	GetRightsAssignedUser() *string
	SetRightsIds(v []*string) *ListRightsRequest
	GetRightsIds() []*string
	SetRightsPeriod(v string) *ListRightsRequest
	GetRightsPeriod() *string
	SetRightsResourceId(v string) *ListRightsRequest
	GetRightsResourceId() *string
	SetRightsResourceType(v string) *ListRightsRequest
	GetRightsResourceType() *string
	SetRightsResourceTypes(v []*string) *ListRightsRequest
	GetRightsResourceTypes() []*string
	SetRightsSpec(v string) *ListRightsRequest
	GetRightsSpec() *string
}

type ListRightsRequest struct {
	// Optional AI rights filter. Set to true to query only AI rights, or false to query only regular rights, including historical unmarked rights. If you do not specify this parameter, all rights are queried.
	//
	// example:
	//
	// true
	AiDesktop *bool `json:"AiDesktop,omitempty" xml:"AiDesktop,omitempty"`
	// The scenario code for the commercial edition.
	//
	// example:
	//
	// GENERAL_OFFICE
	BusinessSceneCode *string `json:"BusinessSceneCode,omitempty" xml:"BusinessSceneCode,omitempty"`
	// The maximum number of results per page. Specify a positive integer from 1 to 200.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Do not specify this parameter for the first query. For subsequent queries, pass the NextToken value returned from the previous page exactly as it was returned. The example is for illustration purposes only. Do not construct the token manually.
	//
	// example:
	//
	// b5f4e3d2c1a09876b5f4e3d2c1a09876
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The field by which to sort the results, such as gmtCreate, rightsId, or rightsExpireTime. If you do not specify this parameter, the default sorting order of the server is used.
	//
	// example:
	//
	// gmtCreate
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The sorting order. Valid values: asc and desc (lowercase). Use this parameter together with OrderField.
	//
	// example:
	//
	// desc
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// Specifies whether the rights apply to regions outside the Chinese mainland, including Hong Kong (China), Macao (China), and Taiwan (China). Valid values:
	//
	// - 0: The Chinese mainland.
	//
	// - 1: Regions outside the Chinese mainland, including Hong Kong (China), Macao (China), and Taiwan (China).
	//
	// example:
	//
	// 0
	Overseas *int32 `json:"Overseas,omitempty" xml:"Overseas,omitempty"`
	// The rights activation status. Use the returned value to determine whether the activation is complete.
	//
	// example:
	//
	// ACTED
	RightsActiveStatus *string `json:"RightsActiveStatus,omitempty" xml:"RightsActiveStatus,omitempty"`
	// The end user to whom the rights are assigned. For convenience accounts, use EndUserId.
	//
	// example:
	//
	// user-example
	RightsAssignedUser *string `json:"RightsAssignedUser,omitempty" xml:"RightsAssignedUser,omitempty"`
	// The list of rights IDs to query or activate. Specify the RightsId values returned by ListRights.
	//
	// example:
	//
	// ["New_beds-example_0001"]
	RightsIds []*string `json:"RightsIds,omitempty" xml:"RightsIds,omitempty" type:"Repeated"`
	// The rights period code.
	//
	// example:
	//
	// BUSINESS_WEEK_1
	RightsPeriod *string `json:"RightsPeriod,omitempty" xml:"RightsPeriod,omitempty"`
	// The ID of the resource associated with the rights. This value is returned after the resource is created.
	//
	// example:
	//
	// ecd-example
	RightsResourceId *string `json:"RightsResourceId,omitempty" xml:"RightsResourceId,omitempty"`
	// The rights resource type. Use Desktop for regular cloud computers. For other commodities, specify the actual rights type.
	//
	// example:
	//
	// Desktop
	RightsResourceType *string `json:"RightsResourceType,omitempty" xml:"RightsResourceType,omitempty"`
	// The list of rights resource types to filter.
	//
	// example:
	//
	// ["Desktop"]
	RightsResourceTypes []*string `json:"RightsResourceTypes,omitempty" xml:"RightsResourceTypes,omitempty" type:"Repeated"`
	// The rights specification code. Refer to the rights query results for the actual value.
	//
	// example:
	//
	// eds.business_edition_office.4c8g
	RightsSpec *string `json:"RightsSpec,omitempty" xml:"RightsSpec,omitempty"`
}

func (s ListRightsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRightsRequest) GoString() string {
	return s.String()
}

func (s *ListRightsRequest) GetAiDesktop() *bool {
	return s.AiDesktop
}

func (s *ListRightsRequest) GetBusinessSceneCode() *string {
	return s.BusinessSceneCode
}

func (s *ListRightsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRightsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRightsRequest) GetOrderField() *string {
	return s.OrderField
}

func (s *ListRightsRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListRightsRequest) GetOverseas() *int32 {
	return s.Overseas
}

func (s *ListRightsRequest) GetRightsActiveStatus() *string {
	return s.RightsActiveStatus
}

func (s *ListRightsRequest) GetRightsAssignedUser() *string {
	return s.RightsAssignedUser
}

func (s *ListRightsRequest) GetRightsIds() []*string {
	return s.RightsIds
}

func (s *ListRightsRequest) GetRightsPeriod() *string {
	return s.RightsPeriod
}

func (s *ListRightsRequest) GetRightsResourceId() *string {
	return s.RightsResourceId
}

func (s *ListRightsRequest) GetRightsResourceType() *string {
	return s.RightsResourceType
}

func (s *ListRightsRequest) GetRightsResourceTypes() []*string {
	return s.RightsResourceTypes
}

func (s *ListRightsRequest) GetRightsSpec() *string {
	return s.RightsSpec
}

func (s *ListRightsRequest) SetAiDesktop(v bool) *ListRightsRequest {
	s.AiDesktop = &v
	return s
}

func (s *ListRightsRequest) SetBusinessSceneCode(v string) *ListRightsRequest {
	s.BusinessSceneCode = &v
	return s
}

func (s *ListRightsRequest) SetMaxResults(v int32) *ListRightsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRightsRequest) SetNextToken(v string) *ListRightsRequest {
	s.NextToken = &v
	return s
}

func (s *ListRightsRequest) SetOrderField(v string) *ListRightsRequest {
	s.OrderField = &v
	return s
}

func (s *ListRightsRequest) SetOrderType(v string) *ListRightsRequest {
	s.OrderType = &v
	return s
}

func (s *ListRightsRequest) SetOverseas(v int32) *ListRightsRequest {
	s.Overseas = &v
	return s
}

func (s *ListRightsRequest) SetRightsActiveStatus(v string) *ListRightsRequest {
	s.RightsActiveStatus = &v
	return s
}

func (s *ListRightsRequest) SetRightsAssignedUser(v string) *ListRightsRequest {
	s.RightsAssignedUser = &v
	return s
}

func (s *ListRightsRequest) SetRightsIds(v []*string) *ListRightsRequest {
	s.RightsIds = v
	return s
}

func (s *ListRightsRequest) SetRightsPeriod(v string) *ListRightsRequest {
	s.RightsPeriod = &v
	return s
}

func (s *ListRightsRequest) SetRightsResourceId(v string) *ListRightsRequest {
	s.RightsResourceId = &v
	return s
}

func (s *ListRightsRequest) SetRightsResourceType(v string) *ListRightsRequest {
	s.RightsResourceType = &v
	return s
}

func (s *ListRightsRequest) SetRightsResourceTypes(v []*string) *ListRightsRequest {
	s.RightsResourceTypes = v
	return s
}

func (s *ListRightsRequest) SetRightsSpec(v string) *ListRightsRequest {
	s.RightsSpec = &v
	return s
}

func (s *ListRightsRequest) Validate() error {
	return dara.Validate(s)
}
