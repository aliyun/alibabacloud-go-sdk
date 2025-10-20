// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoginRedirectApplicationForBrandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBrandLoginRedirectApplication(v *GetLoginRedirectApplicationForBrandResponseBodyBrandLoginRedirectApplication) *GetLoginRedirectApplicationForBrandResponseBody
	GetBrandLoginRedirectApplication() *GetLoginRedirectApplicationForBrandResponseBodyBrandLoginRedirectApplication
	SetRequestId(v string) *GetLoginRedirectApplicationForBrandResponseBody
	GetRequestId() *string
}

type GetLoginRedirectApplicationForBrandResponseBody struct {
	BrandLoginRedirectApplication *GetLoginRedirectApplicationForBrandResponseBodyBrandLoginRedirectApplication `json:"BrandLoginRedirectApplication,omitempty" xml:"BrandLoginRedirectApplication,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLoginRedirectApplicationForBrandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLoginRedirectApplicationForBrandResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoginRedirectApplicationForBrandResponseBody) GetBrandLoginRedirectApplication() *GetLoginRedirectApplicationForBrandResponseBodyBrandLoginRedirectApplication {
	return s.BrandLoginRedirectApplication
}

func (s *GetLoginRedirectApplicationForBrandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLoginRedirectApplicationForBrandResponseBody) SetBrandLoginRedirectApplication(v *GetLoginRedirectApplicationForBrandResponseBodyBrandLoginRedirectApplication) *GetLoginRedirectApplicationForBrandResponseBody {
	s.BrandLoginRedirectApplication = v
	return s
}

func (s *GetLoginRedirectApplicationForBrandResponseBody) SetRequestId(v string) *GetLoginRedirectApplicationForBrandResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLoginRedirectApplicationForBrandResponseBody) Validate() error {
	if s.BrandLoginRedirectApplication != nil {
		if err := s.BrandLoginRedirectApplication.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLoginRedirectApplicationForBrandResponseBodyBrandLoginRedirectApplication struct {
	// 应用ID
	//
	// example:
	//
	// app_xxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 品牌ID
	//
	// example:
	//
	// brand_xxxx
	BrandId *string `json:"BrandId,omitempty" xml:"BrandId,omitempty"`
	// 实例ID
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetLoginRedirectApplicationForBrandResponseBodyBrandLoginRedirectApplication) String() string {
	return dara.Prettify(s)
}

func (s GetLoginRedirectApplicationForBrandResponseBodyBrandLoginRedirectApplication) GoString() string {
	return s.String()
}

func (s *GetLoginRedirectApplicationForBrandResponseBodyBrandLoginRedirectApplication) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetLoginRedirectApplicationForBrandResponseBodyBrandLoginRedirectApplication) GetBrandId() *string {
	return s.BrandId
}

func (s *GetLoginRedirectApplicationForBrandResponseBodyBrandLoginRedirectApplication) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLoginRedirectApplicationForBrandResponseBodyBrandLoginRedirectApplication) SetApplicationId(v string) *GetLoginRedirectApplicationForBrandResponseBodyBrandLoginRedirectApplication {
	s.ApplicationId = &v
	return s
}

func (s *GetLoginRedirectApplicationForBrandResponseBodyBrandLoginRedirectApplication) SetBrandId(v string) *GetLoginRedirectApplicationForBrandResponseBodyBrandLoginRedirectApplication {
	s.BrandId = &v
	return s
}

func (s *GetLoginRedirectApplicationForBrandResponseBodyBrandLoginRedirectApplication) SetInstanceId(v string) *GetLoginRedirectApplicationForBrandResponseBodyBrandLoginRedirectApplication {
	s.InstanceId = &v
	return s
}

func (s *GetLoginRedirectApplicationForBrandResponseBodyBrandLoginRedirectApplication) Validate() error {
	return dara.Validate(s)
}
