// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceAllUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryFaceAllUserGroupResponseBody
	GetCode() *string
	SetData(v *QueryFaceAllUserGroupResponseBodyData) *QueryFaceAllUserGroupResponseBody
	GetData() *QueryFaceAllUserGroupResponseBodyData
	SetErrorMessage(v string) *QueryFaceAllUserGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryFaceAllUserGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryFaceAllUserGroupResponseBody
	GetSuccess() *bool
}

type QueryFaceAllUserGroupResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *QueryFaceAllUserGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceAllUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceAllUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryFaceAllUserGroupResponseBody) GetData() *QueryFaceAllUserGroupResponseBodyData {
	return s.Data
}

func (s *QueryFaceAllUserGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryFaceAllUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFaceAllUserGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryFaceAllUserGroupResponseBody) SetCode(v string) *QueryFaceAllUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceAllUserGroupResponseBody) SetData(v *QueryFaceAllUserGroupResponseBodyData) *QueryFaceAllUserGroupResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceAllUserGroupResponseBody) SetErrorMessage(v string) *QueryFaceAllUserGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceAllUserGroupResponseBody) SetRequestId(v string) *QueryFaceAllUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceAllUserGroupResponseBody) SetSuccess(v bool) *QueryFaceAllUserGroupResponseBody {
	s.Success = &v
	return s
}

func (s *QueryFaceAllUserGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryFaceAllUserGroupResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The page size.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of user groups.
	//
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The list of user group information.
	UserGroupList []*QueryFaceAllUserGroupResponseBodyDataUserGroupList `json:"UserGroupList,omitempty" xml:"UserGroupList,omitempty" type:"Repeated"`
}

func (s QueryFaceAllUserGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceAllUserGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryFaceAllUserGroupResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryFaceAllUserGroupResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *QueryFaceAllUserGroupResponseBodyData) GetUserGroupList() []*QueryFaceAllUserGroupResponseBodyDataUserGroupList {
	return s.UserGroupList
}

func (s *QueryFaceAllUserGroupResponseBodyData) SetPageNo(v int32) *QueryFaceAllUserGroupResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QueryFaceAllUserGroupResponseBodyData) SetPageSize(v int32) *QueryFaceAllUserGroupResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryFaceAllUserGroupResponseBodyData) SetTotal(v int32) *QueryFaceAllUserGroupResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryFaceAllUserGroupResponseBodyData) SetUserGroupList(v []*QueryFaceAllUserGroupResponseBodyDataUserGroupList) *QueryFaceAllUserGroupResponseBodyData {
	s.UserGroupList = v
	return s
}

func (s *QueryFaceAllUserGroupResponseBodyData) Validate() error {
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

type QueryFaceAllUserGroupResponseBodyDataUserGroupList struct {
	// The last modified time.
	//
	// example:
	//
	// 2020-05-12 17:12:37
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The user group ID.
	//
	// example:
	//
	// dmjpxfdk
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The user group name.
	//
	// example:
	//
	// TestUserGroupName
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s QueryFaceAllUserGroupResponseBodyDataUserGroupList) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceAllUserGroupResponseBodyDataUserGroupList) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupResponseBodyDataUserGroupList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *QueryFaceAllUserGroupResponseBodyDataUserGroupList) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *QueryFaceAllUserGroupResponseBodyDataUserGroupList) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *QueryFaceAllUserGroupResponseBodyDataUserGroupList) SetModifiedTime(v string) *QueryFaceAllUserGroupResponseBodyDataUserGroupList {
	s.ModifiedTime = &v
	return s
}

func (s *QueryFaceAllUserGroupResponseBodyDataUserGroupList) SetUserGroupId(v string) *QueryFaceAllUserGroupResponseBodyDataUserGroupList {
	s.UserGroupId = &v
	return s
}

func (s *QueryFaceAllUserGroupResponseBodyDataUserGroupList) SetUserGroupName(v string) *QueryFaceAllUserGroupResponseBodyDataUserGroupList {
	s.UserGroupName = &v
	return s
}

func (s *QueryFaceAllUserGroupResponseBodyDataUserGroupList) Validate() error {
	return dara.Validate(s)
}
