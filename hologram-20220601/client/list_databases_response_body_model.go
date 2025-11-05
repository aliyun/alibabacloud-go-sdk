// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseList(v []*ListDatabasesResponseBodyDatabaseList) *ListDatabasesResponseBody
	GetDatabaseList() []*ListDatabasesResponseBodyDatabaseList
	SetErrorCode(v string) *ListDatabasesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDatabasesResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *ListDatabasesResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *ListDatabasesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ListDatabasesResponseBody
	GetSuccess() *string
}

type ListDatabasesResponseBody struct {
	DatabaseList []*ListDatabasesResponseBodyDatabaseList `json:"DatabaseList,omitempty" xml:"DatabaseList,omitempty" type:"Repeated"`
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// null
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0C4935F5-6217-569A-902F-931B2F3E28BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDatabasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBody) GetDatabaseList() []*ListDatabasesResponseBodyDatabaseList {
	return s.DatabaseList
}

func (s *ListDatabasesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDatabasesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDatabasesResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ListDatabasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatabasesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ListDatabasesResponseBody) SetDatabaseList(v []*ListDatabasesResponseBodyDatabaseList) *ListDatabasesResponseBody {
	s.DatabaseList = v
	return s
}

func (s *ListDatabasesResponseBody) SetErrorCode(v string) *ListDatabasesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDatabasesResponseBody) SetErrorMessage(v string) *ListDatabasesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDatabasesResponseBody) SetHttpStatusCode(v string) *ListDatabasesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDatabasesResponseBody) SetRequestId(v string) *ListDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatabasesResponseBody) SetSuccess(v string) *ListDatabasesResponseBody {
	s.Success = &v
	return s
}

func (s *ListDatabasesResponseBody) Validate() error {
	if s.DatabaseList != nil {
		for _, item := range s.DatabaseList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDatabasesResponseBodyDatabaseList struct {
	// example:
	//
	// false
	External *bool `json:"External,omitempty" xml:"External,omitempty"`
	// example:
	//
	// my_db
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// SPM
	PermissionModel *string `json:"PermissionModel,omitempty" xml:"PermissionModel,omitempty"`
	// example:
	//
	// developer
	Privilege *string `json:"Privilege,omitempty" xml:"Privilege,omitempty"`
}

func (s ListDatabasesResponseBodyDatabaseList) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesResponseBodyDatabaseList) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBodyDatabaseList) GetExternal() *bool {
	return s.External
}

func (s *ListDatabasesResponseBodyDatabaseList) GetName() *string {
	return s.Name
}

func (s *ListDatabasesResponseBodyDatabaseList) GetPermissionModel() *string {
	return s.PermissionModel
}

func (s *ListDatabasesResponseBodyDatabaseList) GetPrivilege() *string {
	return s.Privilege
}

func (s *ListDatabasesResponseBodyDatabaseList) SetExternal(v bool) *ListDatabasesResponseBodyDatabaseList {
	s.External = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseList) SetName(v string) *ListDatabasesResponseBodyDatabaseList {
	s.Name = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseList) SetPermissionModel(v string) *ListDatabasesResponseBodyDatabaseList {
	s.PermissionModel = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseList) SetPrivilege(v string) *ListDatabasesResponseBodyDatabaseList {
	s.Privilege = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseList) Validate() error {
	return dara.Validate(s)
}
