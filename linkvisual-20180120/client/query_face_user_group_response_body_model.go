// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryFaceUserGroupResponseBody
	GetCode() *string
	SetData(v *QueryFaceUserGroupResponseBodyData) *QueryFaceUserGroupResponseBody
	GetData() *QueryFaceUserGroupResponseBodyData
	SetErrorMessage(v string) *QueryFaceUserGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryFaceUserGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryFaceUserGroupResponseBody
	GetSuccess() *bool
}

type QueryFaceUserGroupResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *QueryFaceUserGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0A13279A-5640-45E7-87AA-83750541AD0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceUserGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryFaceUserGroupResponseBody) GetData() *QueryFaceUserGroupResponseBodyData {
	return s.Data
}

func (s *QueryFaceUserGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryFaceUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFaceUserGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryFaceUserGroupResponseBody) SetCode(v string) *QueryFaceUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceUserGroupResponseBody) SetData(v *QueryFaceUserGroupResponseBodyData) *QueryFaceUserGroupResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceUserGroupResponseBody) SetErrorMessage(v string) *QueryFaceUserGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceUserGroupResponseBody) SetRequestId(v string) *QueryFaceUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceUserGroupResponseBody) SetSuccess(v bool) *QueryFaceUserGroupResponseBody {
	s.Success = &v
	return s
}

func (s *QueryFaceUserGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryFaceUserGroupResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of user groups.
	//
	// example:
	//
	// 5
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The list of user groups.
	UserGroupList []*QueryFaceUserGroupResponseBodyDataUserGroupList `json:"UserGroupList,omitempty" xml:"UserGroupList,omitempty" type:"Repeated"`
}

func (s QueryFaceUserGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceUserGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceUserGroupResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryFaceUserGroupResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryFaceUserGroupResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *QueryFaceUserGroupResponseBodyData) GetUserGroupList() []*QueryFaceUserGroupResponseBodyDataUserGroupList {
	return s.UserGroupList
}

func (s *QueryFaceUserGroupResponseBodyData) SetPageNo(v int32) *QueryFaceUserGroupResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QueryFaceUserGroupResponseBodyData) SetPageSize(v int32) *QueryFaceUserGroupResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryFaceUserGroupResponseBodyData) SetTotal(v int32) *QueryFaceUserGroupResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryFaceUserGroupResponseBodyData) SetUserGroupList(v []*QueryFaceUserGroupResponseBodyDataUserGroupList) *QueryFaceUserGroupResponseBodyData {
	s.UserGroupList = v
	return s
}

func (s *QueryFaceUserGroupResponseBodyData) Validate() error {
	if s.UserGroupList != nil {
		for _, item := range s.UserGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryFaceUserGroupResponseBodyDataUserGroupList struct {
	// The update time.
	//
	// example:
	//
	// 2020-05-12 17:12:37
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the user group.
	//
	// example:
	//
	// TestUserGroupId
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The name of the user group.
	//
	// example:
	//
	// TestUserGroupName
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s QueryFaceUserGroupResponseBodyDataUserGroupList) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceUserGroupResponseBodyDataUserGroupList) GoString() string {
	return s.String()
}

func (s *QueryFaceUserGroupResponseBodyDataUserGroupList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *QueryFaceUserGroupResponseBodyDataUserGroupList) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *QueryFaceUserGroupResponseBodyDataUserGroupList) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *QueryFaceUserGroupResponseBodyDataUserGroupList) SetModifiedTime(v string) *QueryFaceUserGroupResponseBodyDataUserGroupList {
	s.ModifiedTime = &v
	return s
}

func (s *QueryFaceUserGroupResponseBodyDataUserGroupList) SetUserGroupId(v string) *QueryFaceUserGroupResponseBodyDataUserGroupList {
	s.UserGroupId = &v
	return s
}

func (s *QueryFaceUserGroupResponseBodyDataUserGroupList) SetUserGroupName(v string) *QueryFaceUserGroupResponseBodyDataUserGroupList {
	s.UserGroupName = &v
	return s
}

func (s *QueryFaceUserGroupResponseBodyDataUserGroupList) Validate() error {
	return dara.Validate(s)
}
