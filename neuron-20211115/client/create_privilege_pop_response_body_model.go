// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivilegePopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*Privilege) *CreatePrivilegePopResponseBody
	GetList() []*Privilege
	SetRequestId(v string) *CreatePrivilegePopResponseBody
	GetRequestId() *string
}

type CreatePrivilegePopResponseBody struct {
	List []*Privilege `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// example:
	//
	// sdadawqewe
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreatePrivilegePopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivilegePopResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrivilegePopResponseBody) GetList() []*Privilege {
	return s.List
}

func (s *CreatePrivilegePopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePrivilegePopResponseBody) SetList(v []*Privilege) *CreatePrivilegePopResponseBody {
	s.List = v
	return s
}

func (s *CreatePrivilegePopResponseBody) SetRequestId(v string) *CreatePrivilegePopResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePrivilegePopResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
