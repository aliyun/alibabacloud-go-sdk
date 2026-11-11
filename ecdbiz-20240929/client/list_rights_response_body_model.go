// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRightsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRightsResponseBody
	GetCode() *string
	SetData(v []*ListRightsResponseBodyData) *ListRightsResponseBody
	GetData() []*ListRightsResponseBodyData
	SetHttpStatusCode(v int32) *ListRightsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListRightsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListRightsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListRightsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListRightsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListRightsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListRightsResponseBody
	GetTotalCount() *int32
}

type ListRightsResponseBody struct {
	// The business processing result code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of rights on the current page.
	//
	// example:
	//
	// []
	Data []*ListRightsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The maximum number of results per page. Specify a positive integer from 1 to 200.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The description of the processing result. This is used for troubleshooting when a failure occurs.
	//
	// example:
	//
	// The request is processed successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination token for the next page. An empty value indicates that the current query has ended. If the value is not empty, pass it exactly as it is in the next query. The example is for illustration purposes only.
	//
	// example:
	//
	// b5f4e3d2c1a09876b5f4e3d2c1a09876
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The unique identifier of the request. This is used to troubleshoot call issues.
	//
	// example:
	//
	// 01A00000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was processed successfully. A successful activation request does not mean that the resource has been created.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of rights that meet the query conditions.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRightsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRightsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRightsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRightsResponseBody) GetData() []*ListRightsResponseBodyData {
	return s.Data
}

func (s *ListRightsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListRightsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRightsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRightsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRightsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRightsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRightsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRightsResponseBody) SetCode(v string) *ListRightsResponseBody {
	s.Code = &v
	return s
}

func (s *ListRightsResponseBody) SetData(v []*ListRightsResponseBodyData) *ListRightsResponseBody {
	s.Data = v
	return s
}

func (s *ListRightsResponseBody) SetHttpStatusCode(v int32) *ListRightsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRightsResponseBody) SetMaxResults(v int32) *ListRightsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRightsResponseBody) SetMessage(v string) *ListRightsResponseBody {
	s.Message = &v
	return s
}

func (s *ListRightsResponseBody) SetNextToken(v string) *ListRightsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRightsResponseBody) SetRequestId(v string) *ListRightsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRightsResponseBody) SetSuccess(v bool) *ListRightsResponseBody {
	s.Success = &v
	return s
}

func (s *ListRightsResponseBody) SetTotalCount(v int32) *ListRightsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRightsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRightsResponseBodyData struct {
	// The size of the additional data cloud disk. Unit: GiB.
	//
	// example:
	//
	// 0
	AddDataDiskSize *int32 `json:"AddDataDiskSize,omitempty" xml:"AddDataDiskSize,omitempty"`
	// The size of the additional system cloud disk. Unit: GiB.
	//
	// example:
	//
	// 0
	AddRootDiskSize *int32 `json:"AddRootDiskSize,omitempty" xml:"AddRootDiskSize,omitempty"`
	// Indicates whether the cloud computer is an AI cloud computer. Valid values:
	//
	// - true: AI cloud computer.
	//
	// - false: Regular cloud computer.
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
	// The compute plan code. Refer to the actual returned value of the rights.
	//
	// example:
	//
	// HOURS_120_PER_MONTH
	CalculatingCode *string `json:"CalculatingCode,omitempty" xml:"CalculatingCode,omitempty"`
	// The size of the data cloud disk. Unit: GiB.
	//
	// example:
	//
	// 80
	DataDiskSize *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// The time when the rights record was created. The time is in RFC 3339 UTC format and ends with Z.
	//
	// example:
	//
	// 2026-11-06T01:00:00Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The hibernation options for the rights. The actual policy is determined in combination with the compute plan.
	//
	// example:
	//
	// 0
	Hibernate *int32 `json:"Hibernate,omitempty" xml:"Hibernate,omitempty"`
	// The monthly credit information corresponding to the rights. This value may be empty if this configuration is not available.
	//
	// example:
	//
	// 2000
	MonthlyCredit *string `json:"MonthlyCredit,omitempty" xml:"MonthlyCredit,omitempty"`
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
	// The cause of the rights activation failure. This value may be empty if the activation is successful.
	//
	// example:
	//
	// The image is unavailable. Select an available image and try again
	RightsActiveFailureCause *string `json:"RightsActiveFailureCause,omitempty" xml:"RightsActiveFailureCause,omitempty"`
	// The rights activation status. Use the returned value to determine whether the activation is complete.
	//
	// example:
	//
	// ACTED
	RightsActiveStatus *string `json:"RightsActiveStatus,omitempty" xml:"RightsActiveStatus,omitempty"`
	// The time when the rights were activated. This value may be empty if the rights are not yet activated. The time is in RFC 3339 UTC format and ends with Z.
	//
	// example:
	//
	// 2026-11-06T01:00:00Z
	RightsActiveTime *string `json:"RightsActiveTime,omitempty" xml:"RightsActiveTime,omitempty"`
	// The end user to whom the rights are assigned. For convenience accounts, use EndUserId.
	//
	// example:
	//
	// user-example
	RightsAssignedUser *string `json:"RightsAssignedUser,omitempty" xml:"RightsAssignedUser,omitempty"`
	// The deadline for manual activation of the rights. This is different from RightsResourceExpireTime, which is the expiration time of the resource after activation. The time is in RFC 3339 UTC format and ends with Z.
	//
	// example:
	//
	// 2026-12-06T00:00:00Z
	RightsExpireTime *string `json:"RightsExpireTime,omitempty" xml:"RightsExpireTime,omitempty"`
	// The GPU specification. Refer to the returned value in the rights details.
	//
	// example:
	//
	// eds.graphics_flagship_pro.8c16g.2g
	RightsGpuSpec *string `json:"RightsGpuSpec,omitempty" xml:"RightsGpuSpec,omitempty"`
	// The rights ID. Use this value for activation. This is not the cloud computer resource ID.
	//
	// example:
	//
	// New_beds-example_0001
	RightsId *string `json:"RightsId,omitempty" xml:"RightsId,omitempty"`
	// The inventory specification corresponding to the rights.
	//
	// example:
	//
	// eds.business_edition_office.4c8g
	RightsInventorySpec *string `json:"RightsInventorySpec,omitempty" xml:"RightsInventorySpec,omitempty"`
	// The rights period code.
	//
	// example:
	//
	// BUSINESS_WEEK_1
	RightsPeriod *string `json:"RightsPeriod,omitempty" xml:"RightsPeriod,omitempty"`
	// The area to which the rights belong. This cannot replace the resource region in a new activation request.
	//
	// example:
	//
	// cn-east
	RightsResourceArea *string `json:"RightsResourceArea,omitempty" xml:"RightsResourceArea,omitempty"`
	// The actual expiration time of the resource. This is different from the deadline for rights activation. The time is in RFC 3339 UTC format and ends with Z.
	//
	// example:
	//
	// 2026-12-06T01:00:00Z
	RightsResourceExpireTime *string `json:"RightsResourceExpireTime,omitempty" xml:"RightsResourceExpireTime,omitempty"`
	// The ID of the resource associated with the rights. This value is returned after the resource is created.
	//
	// example:
	//
	// ecd-example
	RightsResourceId *string `json:"RightsResourceId,omitempty" xml:"RightsResourceId,omitempty"`
	// The region ID of the resource. This is configured separately from the POP request routing region.
	//
	// example:
	//
	// cn-hangzhou
	RightsResourceRegion *string `json:"RightsResourceRegion,omitempty" xml:"RightsResourceRegion,omitempty"`
	// The status of the associated resource. Refer to the returned value for the actual status.
	//
	// example:
	//
	// Running
	RightsResourceStatus *string `json:"RightsResourceStatus,omitempty" xml:"RightsResourceStatus,omitempty"`
	// The rights resource type. Use Desktop for regular cloud computers. For other commodities, specify the actual rights type.
	//
	// example:
	//
	// Desktop
	RightsResourceType *string `json:"RightsResourceType,omitempty" xml:"RightsResourceType,omitempty"`
	// The name of the rights source.
	//
	// example:
	//
	// Channel Service Provider Co., Ltd
	RightsSourceName *string `json:"RightsSourceName,omitempty" xml:"RightsSourceName,omitempty"`
	// The rights specification code. Refer to the rights query results for the actual value.
	//
	// example:
	//
	// eds.business_edition_office.4c8g
	RightsSpec *string `json:"RightsSpec,omitempty" xml:"RightsSpec,omitempty"`
	// The size of the system cloud disk. Unit: GiB.
	//
	// example:
	//
	// 80
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The list of value-added feature information associated with the rights.
	//
	// example:
	//
	// []
	ValuePackageInfos []*ListRightsResponseBodyDataValuePackageInfos `json:"ValuePackageInfos,omitempty" xml:"ValuePackageInfos,omitempty" type:"Repeated"`
}

func (s ListRightsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRightsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRightsResponseBodyData) GetAddDataDiskSize() *int32 {
	return s.AddDataDiskSize
}

func (s *ListRightsResponseBodyData) GetAddRootDiskSize() *int32 {
	return s.AddRootDiskSize
}

func (s *ListRightsResponseBodyData) GetAiDesktop() *bool {
	return s.AiDesktop
}

func (s *ListRightsResponseBodyData) GetBusinessSceneCode() *string {
	return s.BusinessSceneCode
}

func (s *ListRightsResponseBodyData) GetCalculatingCode() *string {
	return s.CalculatingCode
}

func (s *ListRightsResponseBodyData) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *ListRightsResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListRightsResponseBodyData) GetHibernate() *int32 {
	return s.Hibernate
}

func (s *ListRightsResponseBodyData) GetMonthlyCredit() *string {
	return s.MonthlyCredit
}

func (s *ListRightsResponseBodyData) GetOverseas() *int32 {
	return s.Overseas
}

func (s *ListRightsResponseBodyData) GetRightsActiveFailureCause() *string {
	return s.RightsActiveFailureCause
}

func (s *ListRightsResponseBodyData) GetRightsActiveStatus() *string {
	return s.RightsActiveStatus
}

func (s *ListRightsResponseBodyData) GetRightsActiveTime() *string {
	return s.RightsActiveTime
}

func (s *ListRightsResponseBodyData) GetRightsAssignedUser() *string {
	return s.RightsAssignedUser
}

func (s *ListRightsResponseBodyData) GetRightsExpireTime() *string {
	return s.RightsExpireTime
}

func (s *ListRightsResponseBodyData) GetRightsGpuSpec() *string {
	return s.RightsGpuSpec
}

func (s *ListRightsResponseBodyData) GetRightsId() *string {
	return s.RightsId
}

func (s *ListRightsResponseBodyData) GetRightsInventorySpec() *string {
	return s.RightsInventorySpec
}

func (s *ListRightsResponseBodyData) GetRightsPeriod() *string {
	return s.RightsPeriod
}

func (s *ListRightsResponseBodyData) GetRightsResourceArea() *string {
	return s.RightsResourceArea
}

func (s *ListRightsResponseBodyData) GetRightsResourceExpireTime() *string {
	return s.RightsResourceExpireTime
}

func (s *ListRightsResponseBodyData) GetRightsResourceId() *string {
	return s.RightsResourceId
}

func (s *ListRightsResponseBodyData) GetRightsResourceRegion() *string {
	return s.RightsResourceRegion
}

func (s *ListRightsResponseBodyData) GetRightsResourceStatus() *string {
	return s.RightsResourceStatus
}

func (s *ListRightsResponseBodyData) GetRightsResourceType() *string {
	return s.RightsResourceType
}

func (s *ListRightsResponseBodyData) GetRightsSourceName() *string {
	return s.RightsSourceName
}

func (s *ListRightsResponseBodyData) GetRightsSpec() *string {
	return s.RightsSpec
}

func (s *ListRightsResponseBodyData) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *ListRightsResponseBodyData) GetValuePackageInfos() []*ListRightsResponseBodyDataValuePackageInfos {
	return s.ValuePackageInfos
}

func (s *ListRightsResponseBodyData) SetAddDataDiskSize(v int32) *ListRightsResponseBodyData {
	s.AddDataDiskSize = &v
	return s
}

func (s *ListRightsResponseBodyData) SetAddRootDiskSize(v int32) *ListRightsResponseBodyData {
	s.AddRootDiskSize = &v
	return s
}

func (s *ListRightsResponseBodyData) SetAiDesktop(v bool) *ListRightsResponseBodyData {
	s.AiDesktop = &v
	return s
}

func (s *ListRightsResponseBodyData) SetBusinessSceneCode(v string) *ListRightsResponseBodyData {
	s.BusinessSceneCode = &v
	return s
}

func (s *ListRightsResponseBodyData) SetCalculatingCode(v string) *ListRightsResponseBodyData {
	s.CalculatingCode = &v
	return s
}

func (s *ListRightsResponseBodyData) SetDataDiskSize(v int32) *ListRightsResponseBodyData {
	s.DataDiskSize = &v
	return s
}

func (s *ListRightsResponseBodyData) SetGmtCreate(v string) *ListRightsResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListRightsResponseBodyData) SetHibernate(v int32) *ListRightsResponseBodyData {
	s.Hibernate = &v
	return s
}

func (s *ListRightsResponseBodyData) SetMonthlyCredit(v string) *ListRightsResponseBodyData {
	s.MonthlyCredit = &v
	return s
}

func (s *ListRightsResponseBodyData) SetOverseas(v int32) *ListRightsResponseBodyData {
	s.Overseas = &v
	return s
}

func (s *ListRightsResponseBodyData) SetRightsActiveFailureCause(v string) *ListRightsResponseBodyData {
	s.RightsActiveFailureCause = &v
	return s
}

func (s *ListRightsResponseBodyData) SetRightsActiveStatus(v string) *ListRightsResponseBodyData {
	s.RightsActiveStatus = &v
	return s
}

func (s *ListRightsResponseBodyData) SetRightsActiveTime(v string) *ListRightsResponseBodyData {
	s.RightsActiveTime = &v
	return s
}

func (s *ListRightsResponseBodyData) SetRightsAssignedUser(v string) *ListRightsResponseBodyData {
	s.RightsAssignedUser = &v
	return s
}

func (s *ListRightsResponseBodyData) SetRightsExpireTime(v string) *ListRightsResponseBodyData {
	s.RightsExpireTime = &v
	return s
}

func (s *ListRightsResponseBodyData) SetRightsGpuSpec(v string) *ListRightsResponseBodyData {
	s.RightsGpuSpec = &v
	return s
}

func (s *ListRightsResponseBodyData) SetRightsId(v string) *ListRightsResponseBodyData {
	s.RightsId = &v
	return s
}

func (s *ListRightsResponseBodyData) SetRightsInventorySpec(v string) *ListRightsResponseBodyData {
	s.RightsInventorySpec = &v
	return s
}

func (s *ListRightsResponseBodyData) SetRightsPeriod(v string) *ListRightsResponseBodyData {
	s.RightsPeriod = &v
	return s
}

func (s *ListRightsResponseBodyData) SetRightsResourceArea(v string) *ListRightsResponseBodyData {
	s.RightsResourceArea = &v
	return s
}

func (s *ListRightsResponseBodyData) SetRightsResourceExpireTime(v string) *ListRightsResponseBodyData {
	s.RightsResourceExpireTime = &v
	return s
}

func (s *ListRightsResponseBodyData) SetRightsResourceId(v string) *ListRightsResponseBodyData {
	s.RightsResourceId = &v
	return s
}

func (s *ListRightsResponseBodyData) SetRightsResourceRegion(v string) *ListRightsResponseBodyData {
	s.RightsResourceRegion = &v
	return s
}

func (s *ListRightsResponseBodyData) SetRightsResourceStatus(v string) *ListRightsResponseBodyData {
	s.RightsResourceStatus = &v
	return s
}

func (s *ListRightsResponseBodyData) SetRightsResourceType(v string) *ListRightsResponseBodyData {
	s.RightsResourceType = &v
	return s
}

func (s *ListRightsResponseBodyData) SetRightsSourceName(v string) *ListRightsResponseBodyData {
	s.RightsSourceName = &v
	return s
}

func (s *ListRightsResponseBodyData) SetRightsSpec(v string) *ListRightsResponseBodyData {
	s.RightsSpec = &v
	return s
}

func (s *ListRightsResponseBodyData) SetSystemDiskSize(v int32) *ListRightsResponseBodyData {
	s.SystemDiskSize = &v
	return s
}

func (s *ListRightsResponseBodyData) SetValuePackageInfos(v []*ListRightsResponseBodyDataValuePackageInfos) *ListRightsResponseBodyData {
	s.ValuePackageInfos = v
	return s
}

func (s *ListRightsResponseBodyData) Validate() error {
	if s.ValuePackageInfos != nil {
		for _, item := range s.ValuePackageInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRightsResponseBodyDataValuePackageInfos struct {
	// The content of the value-added feature.
	//
	// example:
	//
	// Basic monitoring, real-time monitoring, session monitoring, and application control
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// The value-added feature code.
	//
	// example:
	//
	// Monitor
	ValuePackageCode *string `json:"ValuePackageCode,omitempty" xml:"ValuePackageCode,omitempty"`
}

func (s ListRightsResponseBodyDataValuePackageInfos) String() string {
	return dara.Prettify(s)
}

func (s ListRightsResponseBodyDataValuePackageInfos) GoString() string {
	return s.String()
}

func (s *ListRightsResponseBodyDataValuePackageInfos) GetFunction() *string {
	return s.Function
}

func (s *ListRightsResponseBodyDataValuePackageInfos) GetValuePackageCode() *string {
	return s.ValuePackageCode
}

func (s *ListRightsResponseBodyDataValuePackageInfos) SetFunction(v string) *ListRightsResponseBodyDataValuePackageInfos {
	s.Function = &v
	return s
}

func (s *ListRightsResponseBodyDataValuePackageInfos) SetValuePackageCode(v string) *ListRightsResponseBodyDataValuePackageInfos {
	s.ValuePackageCode = &v
	return s
}

func (s *ListRightsResponseBodyDataValuePackageInfos) Validate() error {
	return dara.Validate(s)
}
