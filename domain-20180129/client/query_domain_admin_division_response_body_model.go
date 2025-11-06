// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainAdminDivisionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdminDivisions(v *QueryDomainAdminDivisionResponseBodyAdminDivisions) *QueryDomainAdminDivisionResponseBody
	GetAdminDivisions() *QueryDomainAdminDivisionResponseBodyAdminDivisions
	SetRequestId(v string) *QueryDomainAdminDivisionResponseBody
	GetRequestId() *string
}

type QueryDomainAdminDivisionResponseBody struct {
	AdminDivisions *QueryDomainAdminDivisionResponseBodyAdminDivisions `json:"AdminDivisions,omitempty" xml:"AdminDivisions,omitempty" type:"Struct"`
	// example:
	//
	// 4EA05A10-D4BC-47EA-AD9E-370A46BB4FB9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDomainAdminDivisionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainAdminDivisionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDomainAdminDivisionResponseBody) GetAdminDivisions() *QueryDomainAdminDivisionResponseBodyAdminDivisions {
	return s.AdminDivisions
}

func (s *QueryDomainAdminDivisionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDomainAdminDivisionResponseBody) SetAdminDivisions(v *QueryDomainAdminDivisionResponseBodyAdminDivisions) *QueryDomainAdminDivisionResponseBody {
	s.AdminDivisions = v
	return s
}

func (s *QueryDomainAdminDivisionResponseBody) SetRequestId(v string) *QueryDomainAdminDivisionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDomainAdminDivisionResponseBody) Validate() error {
	if s.AdminDivisions != nil {
		if err := s.AdminDivisions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDomainAdminDivisionResponseBodyAdminDivisions struct {
	AdminDivision []*QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivision `json:"AdminDivision,omitempty" xml:"AdminDivision,omitempty" type:"Repeated"`
}

func (s QueryDomainAdminDivisionResponseBodyAdminDivisions) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainAdminDivisionResponseBodyAdminDivisions) GoString() string {
	return s.String()
}

func (s *QueryDomainAdminDivisionResponseBodyAdminDivisions) GetAdminDivision() []*QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivision {
	return s.AdminDivision
}

func (s *QueryDomainAdminDivisionResponseBodyAdminDivisions) SetAdminDivision(v []*QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivision) *QueryDomainAdminDivisionResponseBodyAdminDivisions {
	s.AdminDivision = v
	return s
}

func (s *QueryDomainAdminDivisionResponseBodyAdminDivisions) Validate() error {
	if s.AdminDivision != nil {
		for _, item := range s.AdminDivision {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivision struct {
	Children     *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildren `json:"Children,omitempty" xml:"Children,omitempty" type:"Struct"`
	DivisionName *string                                                                  `json:"DivisionName,omitempty" xml:"DivisionName,omitempty"`
}

func (s QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivision) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivision) GoString() string {
	return s.String()
}

func (s *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivision) GetChildren() *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildren {
	return s.Children
}

func (s *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivision) GetDivisionName() *string {
	return s.DivisionName
}

func (s *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivision) SetChildren(v *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildren) *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivision {
	s.Children = v
	return s
}

func (s *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivision) SetDivisionName(v string) *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivision {
	s.DivisionName = &v
	return s
}

func (s *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivision) Validate() error {
	if s.Children != nil {
		if err := s.Children.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildren struct {
	Children []*QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildrenChildren `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
}

func (s QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildren) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildren) GoString() string {
	return s.String()
}

func (s *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildren) GetChildren() []*QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildrenChildren {
	return s.Children
}

func (s *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildren) SetChildren(v []*QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildrenChildren) *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildren {
	s.Children = v
	return s
}

func (s *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildren) Validate() error {
	if s.Children != nil {
		for _, item := range s.Children {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildrenChildren struct {
	ChildDivisionName *string `json:"ChildDivisionName,omitempty" xml:"ChildDivisionName,omitempty"`
}

func (s QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildrenChildren) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildrenChildren) GoString() string {
	return s.String()
}

func (s *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildrenChildren) GetChildDivisionName() *string {
	return s.ChildDivisionName
}

func (s *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildrenChildren) SetChildDivisionName(v string) *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildrenChildren {
	s.ChildDivisionName = &v
	return s
}

func (s *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildrenChildren) Validate() error {
	return dara.Validate(s)
}
