// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRouteMapAssociationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListRouteMapAssociationsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListRouteMapAssociationsResponseBody
	GetCode() *string
	SetDynamicCode(v string) *ListRouteMapAssociationsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListRouteMapAssociationsResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ListRouteMapAssociationsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListRouteMapAssociationsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListRouteMapAssociationsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListRouteMapAssociationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListRouteMapAssociationsResponseBody
	GetRequestId() *string
	SetRouteMapAssociationList(v []*ListRouteMapAssociationsResponseBodyRouteMapAssociationList) *ListRouteMapAssociationsResponseBody
	GetRouteMapAssociationList() []*ListRouteMapAssociationsResponseBodyRouteMapAssociationList
	SetSuccess(v bool) *ListRouteMapAssociationsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListRouteMapAssociationsResponseBody
	GetTotalCount() *int32
}

type ListRouteMapAssociationsResponseBody struct {
	// The details of the access denied error.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code that indicates whether the call is successful. A value of 200 indicates success. Other values indicate failure. For more information about the causes of failure, see error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace `%s` in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the **DtsJobId*	- request parameter is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The message returned for the operation.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The token that is used to start the next query. Valid values:
	//
	// - If **NextToken*	- is empty, no next query is to be sent.
	//
	// - If **NextToken*	- has a return value, the value indicates the token used to start the next query.
	//
	// example:
	//
	// e0a2dbeb69a8beeeb8194e92b702df3fd3e7bfe6ce7bfc16e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 05130E79-588D-5C40-A718-C4863A59****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of route map associations.
	RouteMapAssociationList []*ListRouteMapAssociationsResponseBodyRouteMapAssociationList `json:"RouteMapAssociationList,omitempty" xml:"RouteMapAssociationList,omitempty" type:"Repeated"`
	// Indicates whether the call is successful. Valid values:
	//
	// - **True**: The call is successful.
	//
	// - **False**: The call fails.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of route map associations.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRouteMapAssociationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRouteMapAssociationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRouteMapAssociationsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListRouteMapAssociationsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRouteMapAssociationsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListRouteMapAssociationsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListRouteMapAssociationsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListRouteMapAssociationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRouteMapAssociationsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRouteMapAssociationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRouteMapAssociationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRouteMapAssociationsResponseBody) GetRouteMapAssociationList() []*ListRouteMapAssociationsResponseBodyRouteMapAssociationList {
	return s.RouteMapAssociationList
}

func (s *ListRouteMapAssociationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRouteMapAssociationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRouteMapAssociationsResponseBody) SetAccessDeniedDetail(v string) *ListRouteMapAssociationsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListRouteMapAssociationsResponseBody) SetCode(v string) *ListRouteMapAssociationsResponseBody {
	s.Code = &v
	return s
}

func (s *ListRouteMapAssociationsResponseBody) SetDynamicCode(v string) *ListRouteMapAssociationsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListRouteMapAssociationsResponseBody) SetDynamicMessage(v string) *ListRouteMapAssociationsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListRouteMapAssociationsResponseBody) SetHttpStatusCode(v int32) *ListRouteMapAssociationsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRouteMapAssociationsResponseBody) SetMaxResults(v int32) *ListRouteMapAssociationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRouteMapAssociationsResponseBody) SetMessage(v string) *ListRouteMapAssociationsResponseBody {
	s.Message = &v
	return s
}

func (s *ListRouteMapAssociationsResponseBody) SetNextToken(v string) *ListRouteMapAssociationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRouteMapAssociationsResponseBody) SetRequestId(v string) *ListRouteMapAssociationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRouteMapAssociationsResponseBody) SetRouteMapAssociationList(v []*ListRouteMapAssociationsResponseBodyRouteMapAssociationList) *ListRouteMapAssociationsResponseBody {
	s.RouteMapAssociationList = v
	return s
}

func (s *ListRouteMapAssociationsResponseBody) SetSuccess(v bool) *ListRouteMapAssociationsResponseBody {
	s.Success = &v
	return s
}

func (s *ListRouteMapAssociationsResponseBody) SetTotalCount(v int32) *ListRouteMapAssociationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRouteMapAssociationsResponseBody) Validate() error {
	if s.RouteMapAssociationList != nil {
		for _, item := range s.RouteMapAssociationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRouteMapAssociationsResponseBodyRouteMapAssociationList struct {
	// The ID of the route map association.
	//
	// example:
	//
	// ecr-routemap-assoc-qdoeneegffq4yv****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The direction in which the routing policy is applied. Valid values:
	//
	// - **Inbound**: the inbound region gateway direction, indicating the direction in which routing is transmitted to the Express Connect Router (ECR).
	//
	// - **Outbound**: the outbound region gateway direction, indicating the direction in which routing is transmitted from the Express Connect Router (ECR).
	//
	// example:
	//
	// Inbound
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The instance ID of the Express Connect Router (ECR).
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The time when the route map association was created.
	//
	// The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2026-03-16T03:34:09Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the route map association was updated.
	//
	// The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2026-10-29T03:27:24Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The priority of the route map. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 6
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The list of region IDs where the route map is distributed and takes effect.
	RegionIdList []*string `json:"RegionIdList,omitempty" xml:"RegionIdList,omitempty" type:"Repeated"`
	// The ID of the route map.
	//
	// example:
	//
	// rm-6m58b17417tg5a****
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	// The status of the route map association. Valid values:
	//
	// - **CREATING**: The association is being created.
	//
	// - **ACTIVE**: The association is in a normal state, and no operations are in progress.
	//
	// - **DELETING**: The association is being deleted.
	//
	// - **UPDATING**: The association is being updated.
	//
	// - **DELETED**: The association is deleted.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The version number of the route map. The version number increments when the route map in the CEN is changed.
	//
	// example:
	//
	// 3
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListRouteMapAssociationsResponseBodyRouteMapAssociationList) String() string {
	return dara.Prettify(s)
}

func (s ListRouteMapAssociationsResponseBodyRouteMapAssociationList) GoString() string {
	return s.String()
}

func (s *ListRouteMapAssociationsResponseBodyRouteMapAssociationList) GetAssociationId() *string {
	return s.AssociationId
}

func (s *ListRouteMapAssociationsResponseBodyRouteMapAssociationList) GetDirection() *string {
	return s.Direction
}

func (s *ListRouteMapAssociationsResponseBodyRouteMapAssociationList) GetEcrId() *string {
	return s.EcrId
}

func (s *ListRouteMapAssociationsResponseBodyRouteMapAssociationList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListRouteMapAssociationsResponseBodyRouteMapAssociationList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListRouteMapAssociationsResponseBodyRouteMapAssociationList) GetPriority() *int64 {
	return s.Priority
}

func (s *ListRouteMapAssociationsResponseBodyRouteMapAssociationList) GetRegionIdList() []*string {
	return s.RegionIdList
}

func (s *ListRouteMapAssociationsResponseBodyRouteMapAssociationList) GetRouteMapId() *string {
	return s.RouteMapId
}

func (s *ListRouteMapAssociationsResponseBodyRouteMapAssociationList) GetStatus() *string {
	return s.Status
}

func (s *ListRouteMapAssociationsResponseBodyRouteMapAssociationList) GetVersion() *int64 {
	return s.Version
}

func (s *ListRouteMapAssociationsResponseBodyRouteMapAssociationList) SetAssociationId(v string) *ListRouteMapAssociationsResponseBodyRouteMapAssociationList {
	s.AssociationId = &v
	return s
}

func (s *ListRouteMapAssociationsResponseBodyRouteMapAssociationList) SetDirection(v string) *ListRouteMapAssociationsResponseBodyRouteMapAssociationList {
	s.Direction = &v
	return s
}

func (s *ListRouteMapAssociationsResponseBodyRouteMapAssociationList) SetEcrId(v string) *ListRouteMapAssociationsResponseBodyRouteMapAssociationList {
	s.EcrId = &v
	return s
}

func (s *ListRouteMapAssociationsResponseBodyRouteMapAssociationList) SetGmtCreate(v string) *ListRouteMapAssociationsResponseBodyRouteMapAssociationList {
	s.GmtCreate = &v
	return s
}

func (s *ListRouteMapAssociationsResponseBodyRouteMapAssociationList) SetGmtModified(v string) *ListRouteMapAssociationsResponseBodyRouteMapAssociationList {
	s.GmtModified = &v
	return s
}

func (s *ListRouteMapAssociationsResponseBodyRouteMapAssociationList) SetPriority(v int64) *ListRouteMapAssociationsResponseBodyRouteMapAssociationList {
	s.Priority = &v
	return s
}

func (s *ListRouteMapAssociationsResponseBodyRouteMapAssociationList) SetRegionIdList(v []*string) *ListRouteMapAssociationsResponseBodyRouteMapAssociationList {
	s.RegionIdList = v
	return s
}

func (s *ListRouteMapAssociationsResponseBodyRouteMapAssociationList) SetRouteMapId(v string) *ListRouteMapAssociationsResponseBodyRouteMapAssociationList {
	s.RouteMapId = &v
	return s
}

func (s *ListRouteMapAssociationsResponseBodyRouteMapAssociationList) SetStatus(v string) *ListRouteMapAssociationsResponseBodyRouteMapAssociationList {
	s.Status = &v
	return s
}

func (s *ListRouteMapAssociationsResponseBodyRouteMapAssociationList) SetVersion(v int64) *ListRouteMapAssociationsResponseBodyRouteMapAssociationList {
	s.Version = &v
	return s
}

func (s *ListRouteMapAssociationsResponseBodyRouteMapAssociationList) Validate() error {
	return dara.Validate(s)
}
