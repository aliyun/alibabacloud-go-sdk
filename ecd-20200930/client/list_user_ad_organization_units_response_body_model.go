// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserAdOrganizationUnitsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListUserAdOrganizationUnitsResponseBody
	GetNextToken() *string
	SetOUNames(v []*ListUserAdOrganizationUnitsResponseBodyOUNames) *ListUserAdOrganizationUnitsResponseBody
	GetOUNames() []*ListUserAdOrganizationUnitsResponseBodyOUNames
	SetRequestId(v string) *ListUserAdOrganizationUnitsResponseBody
	GetRequestId() *string
}

type ListUserAdOrganizationUnitsResponseBody struct {
	// A pagination token.
	//
	// example:
	//
	// CAAAAA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The OUs of the AD domain.
	OUNames []*ListUserAdOrganizationUnitsResponseBodyOUNames `json:"OUNames,omitempty" xml:"OUNames,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUserAdOrganizationUnitsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserAdOrganizationUnitsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserAdOrganizationUnitsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserAdOrganizationUnitsResponseBody) GetOUNames() []*ListUserAdOrganizationUnitsResponseBodyOUNames {
	return s.OUNames
}

func (s *ListUserAdOrganizationUnitsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserAdOrganizationUnitsResponseBody) SetNextToken(v string) *ListUserAdOrganizationUnitsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserAdOrganizationUnitsResponseBody) SetOUNames(v []*ListUserAdOrganizationUnitsResponseBodyOUNames) *ListUserAdOrganizationUnitsResponseBody {
	s.OUNames = v
	return s
}

func (s *ListUserAdOrganizationUnitsResponseBody) SetRequestId(v string) *ListUserAdOrganizationUnitsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserAdOrganizationUnitsResponseBody) Validate() error {
	if s.OUNames != nil {
		for _, item := range s.OUNames {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserAdOrganizationUnitsResponseBodyOUNames struct {
	// The name of the OU.
	//
	// example:
	//
	// wuying_computers
	DisplayOUName *string `json:"DisplayOUName,omitempty" xml:"DisplayOUName,omitempty"`
	// The canonical name (CNAME) of the OU in the AD domain controller.
	//
	// example:
	//
	// example.com/wuying_computers
	OUName *string `json:"OUName,omitempty" xml:"OUName,omitempty"`
	// The enterprise AD office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-485361****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
}

func (s ListUserAdOrganizationUnitsResponseBodyOUNames) String() string {
	return dara.Prettify(s)
}

func (s ListUserAdOrganizationUnitsResponseBodyOUNames) GoString() string {
	return s.String()
}

func (s *ListUserAdOrganizationUnitsResponseBodyOUNames) GetDisplayOUName() *string {
	return s.DisplayOUName
}

func (s *ListUserAdOrganizationUnitsResponseBodyOUNames) GetOUName() *string {
	return s.OUName
}

func (s *ListUserAdOrganizationUnitsResponseBodyOUNames) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ListUserAdOrganizationUnitsResponseBodyOUNames) SetDisplayOUName(v string) *ListUserAdOrganizationUnitsResponseBodyOUNames {
	s.DisplayOUName = &v
	return s
}

func (s *ListUserAdOrganizationUnitsResponseBodyOUNames) SetOUName(v string) *ListUserAdOrganizationUnitsResponseBodyOUNames {
	s.OUName = &v
	return s
}

func (s *ListUserAdOrganizationUnitsResponseBodyOUNames) SetOfficeSiteId(v string) *ListUserAdOrganizationUnitsResponseBodyOUNames {
	s.OfficeSiteId = &v
	return s
}

func (s *ListUserAdOrganizationUnitsResponseBodyOUNames) Validate() error {
	return dara.Validate(s)
}
