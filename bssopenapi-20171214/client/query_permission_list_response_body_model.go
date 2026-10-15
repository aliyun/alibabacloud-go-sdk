// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPermissionListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryPermissionListResponseBody
	GetCode() *string
	SetData(v *QueryPermissionListResponseBodyData) *QueryPermissionListResponseBody
	GetData() *QueryPermissionListResponseBodyData
	SetMessage(v string) *QueryPermissionListResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryPermissionListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryPermissionListResponseBody
	GetSuccess() *bool
}

type QueryPermissionListResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryPermissionListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F6E29451-A3CD-4705-806D-0112D08F5C49
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryPermissionListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPermissionListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPermissionListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryPermissionListResponseBody) GetData() *QueryPermissionListResponseBodyData {
	return s.Data
}

func (s *QueryPermissionListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryPermissionListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPermissionListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryPermissionListResponseBody) SetCode(v string) *QueryPermissionListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPermissionListResponseBody) SetData(v *QueryPermissionListResponseBodyData) *QueryPermissionListResponseBody {
	s.Data = v
	return s
}

func (s *QueryPermissionListResponseBody) SetMessage(v string) *QueryPermissionListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPermissionListResponseBody) SetRequestId(v string) *QueryPermissionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPermissionListResponseBody) SetSuccess(v bool) *QueryPermissionListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryPermissionListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPermissionListResponseBodyData struct {
	// The time when the relationship expires. If this parameter is empty, the relationship has not expired.
	//
	// example:
	//
	// 2021-03-06T15:12Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1990699401005016
	MasterId *int64 `json:"MasterId,omitempty" xml:"MasterId,omitempty"`
	// The ID of the member account.
	//
	// example:
	//
	// 1851253838840762
	MemberId *int64 `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The list of permissions.
	PermissionList []*QueryPermissionListResponseBodyDataPermissionList `json:"PermissionList,omitempty" xml:"PermissionList,omitempty" type:"Repeated"`
	// The type of the relationship. Valid values: FinancialManagement and FinancialTrusteeship.
	//
	// example:
	//
	// FinancialManagement
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// The time when the relationship was established. The time is in UTC and follows the ISO 8601 standard. The format is YYYY-MM-DDThh:mm:ssZ. For example, 2016-05-23T12:00:00Z is 20:00:00 on May 23, 2016 in UTC+8.
	//
	// example:
	//
	// 2021-03-02T15:12Z
	SetupTime *string `json:"SetupTime,omitempty" xml:"SetupTime,omitempty"`
	// The time when the relationship takes effect. The time is in UTC and follows the ISO 8601 standard. The format is YYYY-MM-DDThh:mm:ssZ. For example, 2016-05-23T12:00:00Z is 20:00:00 on May 23, 2016 in UTC+8.
	//
	// example:
	//
	// 2021-03-02T15:12Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the relationship. The value is an enumeration member of RelationshipStatusEnum.
	//
	// example:
	//
	// RELATED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s QueryPermissionListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryPermissionListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPermissionListResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryPermissionListResponseBodyData) GetMasterId() *int64 {
	return s.MasterId
}

func (s *QueryPermissionListResponseBodyData) GetMemberId() *int64 {
	return s.MemberId
}

func (s *QueryPermissionListResponseBodyData) GetPermissionList() []*QueryPermissionListResponseBodyDataPermissionList {
	return s.PermissionList
}

func (s *QueryPermissionListResponseBodyData) GetRelationType() *string {
	return s.RelationType
}

func (s *QueryPermissionListResponseBodyData) GetSetupTime() *string {
	return s.SetupTime
}

func (s *QueryPermissionListResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryPermissionListResponseBodyData) GetState() *string {
	return s.State
}

func (s *QueryPermissionListResponseBodyData) SetEndTime(v string) *QueryPermissionListResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *QueryPermissionListResponseBodyData) SetMasterId(v int64) *QueryPermissionListResponseBodyData {
	s.MasterId = &v
	return s
}

func (s *QueryPermissionListResponseBodyData) SetMemberId(v int64) *QueryPermissionListResponseBodyData {
	s.MemberId = &v
	return s
}

func (s *QueryPermissionListResponseBodyData) SetPermissionList(v []*QueryPermissionListResponseBodyDataPermissionList) *QueryPermissionListResponseBodyData {
	s.PermissionList = v
	return s
}

func (s *QueryPermissionListResponseBodyData) SetRelationType(v string) *QueryPermissionListResponseBodyData {
	s.RelationType = &v
	return s
}

func (s *QueryPermissionListResponseBodyData) SetSetupTime(v string) *QueryPermissionListResponseBodyData {
	s.SetupTime = &v
	return s
}

func (s *QueryPermissionListResponseBodyData) SetStartTime(v string) *QueryPermissionListResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *QueryPermissionListResponseBodyData) SetState(v string) *QueryPermissionListResponseBodyData {
	s.State = &v
	return s
}

func (s *QueryPermissionListResponseBodyData) Validate() error {
	if s.PermissionList != nil {
		for _, item := range s.PermissionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryPermissionListResponseBodyDataPermissionList struct {
	// The time when the permission expires. If this parameter is empty, the permission is still valid. The time is in UTC and follows the ISO 8601 standard. The format is YYYY-MM-DDThh:mm:ssZ. For example, 2016-05-23T12:00:00Z is 20:00:00 on May 23, 2016 in UTC+8.
	//
	// example:
	//
	// 2021-03-05T15:12Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The permission code.
	//
	// example:
	//
	// allow_synchronize_finance_identity
	PermissionCode *string `json:"PermissionCode,omitempty" xml:"PermissionCode,omitempty"`
	// The permission name.
	//
	// example:
	//
	// 主账号 子账号共享主账号信控身份
	PermissionName *string `json:"PermissionName,omitempty" xml:"PermissionName,omitempty"`
	// The time when the permission takes effect. The time is in UTC and follows the ISO 8601 standard. The format is YYYY-MM-DDThh:mm:ssZ. For example, 2016-05-23T12:00:00Z is 20:00:00 on May 23, 2016 in UTC+8.
	//
	// example:
	//
	// 2021-03-02T15:12Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryPermissionListResponseBodyDataPermissionList) String() string {
	return dara.Prettify(s)
}

func (s QueryPermissionListResponseBodyDataPermissionList) GoString() string {
	return s.String()
}

func (s *QueryPermissionListResponseBodyDataPermissionList) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryPermissionListResponseBodyDataPermissionList) GetPermissionCode() *string {
	return s.PermissionCode
}

func (s *QueryPermissionListResponseBodyDataPermissionList) GetPermissionName() *string {
	return s.PermissionName
}

func (s *QueryPermissionListResponseBodyDataPermissionList) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryPermissionListResponseBodyDataPermissionList) SetEndTime(v string) *QueryPermissionListResponseBodyDataPermissionList {
	s.EndTime = &v
	return s
}

func (s *QueryPermissionListResponseBodyDataPermissionList) SetPermissionCode(v string) *QueryPermissionListResponseBodyDataPermissionList {
	s.PermissionCode = &v
	return s
}

func (s *QueryPermissionListResponseBodyDataPermissionList) SetPermissionName(v string) *QueryPermissionListResponseBodyDataPermissionList {
	s.PermissionName = &v
	return s
}

func (s *QueryPermissionListResponseBodyDataPermissionList) SetStartTime(v string) *QueryPermissionListResponseBodyDataPermissionList {
	s.StartTime = &v
	return s
}

func (s *QueryPermissionListResponseBodyDataPermissionList) Validate() error {
	return dara.Validate(s)
}
