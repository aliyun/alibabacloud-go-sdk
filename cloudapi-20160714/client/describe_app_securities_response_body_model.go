// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppSecuritiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppSecuritys(v *DescribeAppSecuritiesResponseBodyAppSecuritys) *DescribeAppSecuritiesResponseBody
	GetAppSecuritys() *DescribeAppSecuritiesResponseBodyAppSecuritys
	SetRequestId(v string) *DescribeAppSecuritiesResponseBody
	GetRequestId() *string
}

type DescribeAppSecuritiesResponseBody struct {
	// The associated security policy information.
	AppSecuritys *DescribeAppSecuritiesResponseBodyAppSecuritys `json:"AppSecuritys,omitempty" xml:"AppSecuritys,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// EF924FE4-2EDD-4CD3-89EC-34E4708574E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAppSecuritiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppSecuritiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppSecuritiesResponseBody) GetAppSecuritys() *DescribeAppSecuritiesResponseBodyAppSecuritys {
	return s.AppSecuritys
}

func (s *DescribeAppSecuritiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppSecuritiesResponseBody) SetAppSecuritys(v *DescribeAppSecuritiesResponseBodyAppSecuritys) *DescribeAppSecuritiesResponseBody {
	s.AppSecuritys = v
	return s
}

func (s *DescribeAppSecuritiesResponseBody) SetRequestId(v string) *DescribeAppSecuritiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppSecuritiesResponseBody) Validate() error {
	if s.AppSecuritys != nil {
		if err := s.AppSecuritys.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAppSecuritiesResponseBodyAppSecuritys struct {
	AppSecurity []*DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity `json:"AppSecurity,omitempty" xml:"AppSecurity,omitempty" type:"Repeated"`
}

func (s DescribeAppSecuritiesResponseBodyAppSecuritys) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppSecuritiesResponseBodyAppSecuritys) GoString() string {
	return s.String()
}

func (s *DescribeAppSecuritiesResponseBodyAppSecuritys) GetAppSecurity() []*DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity {
	return s.AppSecurity
}

func (s *DescribeAppSecuritiesResponseBodyAppSecuritys) SetAppSecurity(v []*DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity) *DescribeAppSecuritiesResponseBodyAppSecuritys {
	s.AppSecurity = v
	return s
}

func (s *DescribeAppSecuritiesResponseBodyAppSecuritys) Validate() error {
	if s.AppSecurity != nil {
		for _, item := range s.AppSecurity {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity struct {
	// The application AppCode.
	//
	// example:
	//
	// d2350ecd62c44cbfbe35a7f182e35105
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// The application AppKey.
	//
	// example:
	//
	// 34379343
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The application AppSecret.
	//
	// example:
	//
	// ea5291a7aff343769eb3139a2f6de8c9
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	// The time when the AppKey was created.
	//
	// example:
	//
	// 2021-09-14T18:50:59
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The time when the AppSecret was last modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-08-14T18:03:00+08:00
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
}

func (s DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity) GoString() string {
	return s.String()
}

func (s *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity) GetAppCode() *string {
	return s.AppCode
}

func (s *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity) GetAppKey() *string {
	return s.AppKey
}

func (s *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity) GetAppSecret() *string {
	return s.AppSecret
}

func (s *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity) SetAppCode(v string) *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity {
	s.AppCode = &v
	return s
}

func (s *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity) SetAppKey(v string) *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity {
	s.AppKey = &v
	return s
}

func (s *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity) SetAppSecret(v string) *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity {
	s.AppSecret = &v
	return s
}

func (s *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity) SetCreatedTime(v string) *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity {
	s.CreatedTime = &v
	return s
}

func (s *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity) SetModifiedTime(v string) *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity) Validate() error {
	return dara.Validate(s)
}
