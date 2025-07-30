// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManagerInfoByAuthCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrgId(v string) *GetManagerInfoByAuthCodeResponseBody
	GetOrgId() *string
	SetPhone(v string) *GetManagerInfoByAuthCodeResponseBody
	GetPhone() *string
	SetRequestId(v string) *GetManagerInfoByAuthCodeResponseBody
	GetRequestId() *string
	SetTeamName(v string) *GetManagerInfoByAuthCodeResponseBody
	GetTeamName() *string
	SetUserName(v string) *GetManagerInfoByAuthCodeResponseBody
	GetUserName() *string
	SetWaId(v int64) *GetManagerInfoByAuthCodeResponseBody
	GetWaId() *int64
}

type GetManagerInfoByAuthCodeResponseBody struct {
	// The organization ID.
	//
	// example:
	//
	// 12345678901234****
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The mobile number.
	//
	// example:
	//
	// 1301234****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The team name.
	//
	// example:
	//
	// devteam
	TeamName *string `json:"TeamName,omitempty" xml:"TeamName,omitempty"`
	// The tenant name.
	//
	// example:
	//
	// zhangsan
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The ID of the Elastic Desktop Service account.
	//
	// example:
	//
	// 12345678901234****
	WaId *int64 `json:"WaId,omitempty" xml:"WaId,omitempty"`
}

func (s GetManagerInfoByAuthCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetManagerInfoByAuthCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetManagerInfoByAuthCodeResponseBody) GetOrgId() *string {
	return s.OrgId
}

func (s *GetManagerInfoByAuthCodeResponseBody) GetPhone() *string {
	return s.Phone
}

func (s *GetManagerInfoByAuthCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetManagerInfoByAuthCodeResponseBody) GetTeamName() *string {
	return s.TeamName
}

func (s *GetManagerInfoByAuthCodeResponseBody) GetUserName() *string {
	return s.UserName
}

func (s *GetManagerInfoByAuthCodeResponseBody) GetWaId() *int64 {
	return s.WaId
}

func (s *GetManagerInfoByAuthCodeResponseBody) SetOrgId(v string) *GetManagerInfoByAuthCodeResponseBody {
	s.OrgId = &v
	return s
}

func (s *GetManagerInfoByAuthCodeResponseBody) SetPhone(v string) *GetManagerInfoByAuthCodeResponseBody {
	s.Phone = &v
	return s
}

func (s *GetManagerInfoByAuthCodeResponseBody) SetRequestId(v string) *GetManagerInfoByAuthCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetManagerInfoByAuthCodeResponseBody) SetTeamName(v string) *GetManagerInfoByAuthCodeResponseBody {
	s.TeamName = &v
	return s
}

func (s *GetManagerInfoByAuthCodeResponseBody) SetUserName(v string) *GetManagerInfoByAuthCodeResponseBody {
	s.UserName = &v
	return s
}

func (s *GetManagerInfoByAuthCodeResponseBody) SetWaId(v int64) *GetManagerInfoByAuthCodeResponseBody {
	s.WaId = &v
	return s
}

func (s *GetManagerInfoByAuthCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
