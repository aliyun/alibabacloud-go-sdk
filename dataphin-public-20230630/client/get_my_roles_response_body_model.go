// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMyRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMyRolesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetMyRolesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetMyRolesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMyRolesResponseBody
	GetRequestId() *string
	SetRoleList(v []*GetMyRolesResponseBodyRoleList) *GetMyRolesResponseBody
	GetRoleList() []*GetMyRolesResponseBodyRoleList
	SetSuccess(v bool) *GetMyRolesResponseBody
	GetSuccess() *bool
}

type GetMyRolesResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RoleList  []*GetMyRolesResponseBodyRoleList `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMyRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMyRolesResponseBody) GoString() string {
	return s.String()
}

func (s *GetMyRolesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMyRolesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMyRolesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMyRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMyRolesResponseBody) GetRoleList() []*GetMyRolesResponseBodyRoleList {
	return s.RoleList
}

func (s *GetMyRolesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMyRolesResponseBody) SetCode(v string) *GetMyRolesResponseBody {
	s.Code = &v
	return s
}

func (s *GetMyRolesResponseBody) SetHttpStatusCode(v int32) *GetMyRolesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMyRolesResponseBody) SetMessage(v string) *GetMyRolesResponseBody {
	s.Message = &v
	return s
}

func (s *GetMyRolesResponseBody) SetRequestId(v string) *GetMyRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMyRolesResponseBody) SetRoleList(v []*GetMyRolesResponseBodyRoleList) *GetMyRolesResponseBody {
	s.RoleList = v
	return s
}

func (s *GetMyRolesResponseBody) SetSuccess(v bool) *GetMyRolesResponseBody {
	s.Success = &v
	return s
}

func (s *GetMyRolesResponseBody) Validate() error {
	if s.RoleList != nil {
		for _, item := range s.RoleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMyRolesResponseBodyRoleList struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 300047957
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// dataphinAdmin
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetMyRolesResponseBodyRoleList) String() string {
	return dara.Prettify(s)
}

func (s GetMyRolesResponseBodyRoleList) GoString() string {
	return s.String()
}

func (s *GetMyRolesResponseBodyRoleList) GetDescription() *string {
	return s.Description
}

func (s *GetMyRolesResponseBodyRoleList) GetId() *int64 {
	return s.Id
}

func (s *GetMyRolesResponseBodyRoleList) GetName() *string {
	return s.Name
}

func (s *GetMyRolesResponseBodyRoleList) SetDescription(v string) *GetMyRolesResponseBodyRoleList {
	s.Description = &v
	return s
}

func (s *GetMyRolesResponseBodyRoleList) SetId(v int64) *GetMyRolesResponseBodyRoleList {
	s.Id = &v
	return s
}

func (s *GetMyRolesResponseBodyRoleList) SetName(v string) *GetMyRolesResponseBodyRoleList {
	s.Name = &v
	return s
}

func (s *GetMyRolesResponseBodyRoleList) Validate() error {
	return dara.Validate(s)
}
