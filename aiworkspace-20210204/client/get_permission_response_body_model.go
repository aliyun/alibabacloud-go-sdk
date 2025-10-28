// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPermissionCode(v string) *GetPermissionResponseBody
	GetPermissionCode() *string
	SetPermissionRules(v []*GetPermissionResponseBodyPermissionRules) *GetPermissionResponseBody
	GetPermissionRules() []*GetPermissionResponseBodyPermissionRules
	SetRequestId(v string) *GetPermissionResponseBody
	GetRequestId() *string
}

type GetPermissionResponseBody struct {
	// The permission name, which is unique in a region. For more information about permissions, see [Appendix: Roles and permissions](https://help.aliyun.com/document_detail/2840449.html).
	//
	// example:
	//
	// PaiDLC:ListJobs
	PermissionCode *string `json:"PermissionCode,omitempty" xml:"PermissionCode,omitempty"`
	// The permission rules.
	PermissionRules []*GetPermissionResponseBodyPermissionRules `json:"PermissionRules,omitempty" xml:"PermissionRules,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *GetPermissionResponseBody) GetPermissionCode() *string {
	return s.PermissionCode
}

func (s *GetPermissionResponseBody) GetPermissionRules() []*GetPermissionResponseBodyPermissionRules {
	return s.PermissionRules
}

func (s *GetPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPermissionResponseBody) SetPermissionCode(v string) *GetPermissionResponseBody {
	s.PermissionCode = &v
	return s
}

func (s *GetPermissionResponseBody) SetPermissionRules(v []*GetPermissionResponseBodyPermissionRules) *GetPermissionResponseBody {
	s.PermissionRules = v
	return s
}

func (s *GetPermissionResponseBody) SetRequestId(v string) *GetPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPermissionResponseBody) Validate() error {
	if s.PermissionRules != nil {
		for _, item := range s.PermissionRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPermissionResponseBodyPermissionRules struct {
	// The accessibility. Valid values:
	//
	// 	- PUBLIC: All members can access the workspace.
	//
	// 	- PRIVATE: Only the creator can access the workspace.
	//
	// 	- ANY: All users can access the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The access type. If you set Accessibility to PUBLIC, all users can access the workspace. This parameter is invalid. If you set Accessibility to PRIVATE, the value of this parameter can be:
	//
	// 	- PRIVATE: Only the creator can access the workspace.
	//
	// 	- ANY: All users can access the workspace.
	//
	// example:
	//
	// CREATOR
	EntityAccessType *string `json:"EntityAccessType,omitempty" xml:"EntityAccessType,omitempty"`
}

func (s GetPermissionResponseBodyPermissionRules) String() string {
	return dara.Prettify(s)
}

func (s GetPermissionResponseBodyPermissionRules) GoString() string {
	return s.String()
}

func (s *GetPermissionResponseBodyPermissionRules) GetAccessibility() *string {
	return s.Accessibility
}

func (s *GetPermissionResponseBodyPermissionRules) GetEntityAccessType() *string {
	return s.EntityAccessType
}

func (s *GetPermissionResponseBodyPermissionRules) SetAccessibility(v string) *GetPermissionResponseBodyPermissionRules {
	s.Accessibility = &v
	return s
}

func (s *GetPermissionResponseBodyPermissionRules) SetEntityAccessType(v string) *GetPermissionResponseBodyPermissionRules {
	s.EntityAccessType = &v
	return s
}

func (s *GetPermissionResponseBodyPermissionRules) Validate() error {
	return dara.Validate(s)
}
