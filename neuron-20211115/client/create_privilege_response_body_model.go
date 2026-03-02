// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivilegeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*Privilege) *CreatePrivilegeResponseBody
	GetList() []*Privilege
	SetRequestId(v string) *CreatePrivilegeResponseBody
	GetRequestId() *string
}

type CreatePrivilegeResponseBody struct {
	List      []*Privilege `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	RequestId *string      `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreatePrivilegeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivilegeResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrivilegeResponseBody) GetList() []*Privilege {
	return s.List
}

func (s *CreatePrivilegeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePrivilegeResponseBody) SetList(v []*Privilege) *CreatePrivilegeResponseBody {
	s.List = v
	return s
}

func (s *CreatePrivilegeResponseBody) SetRequestId(v string) *CreatePrivilegeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePrivilegeResponseBody) Validate() error {
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
