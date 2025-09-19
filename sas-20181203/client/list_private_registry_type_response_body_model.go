// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateRegistryTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegistryTypeInfos(v []*ListPrivateRegistryTypeResponseBodyRegistryTypeInfos) *ListPrivateRegistryTypeResponseBody
	GetRegistryTypeInfos() []*ListPrivateRegistryTypeResponseBodyRegistryTypeInfos
	SetRequestId(v string) *ListPrivateRegistryTypeResponseBody
	GetRequestId() *string
}

type ListPrivateRegistryTypeResponseBody struct {
	// An array that consists of image repository types.
	RegistryTypeInfos []*ListPrivateRegistryTypeResponseBodyRegistryTypeInfos `json:"RegistryTypeInfos,omitempty" xml:"RegistryTypeInfos,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 24A20733-10A0-4AF6-BE6B-E3322413****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrivateRegistryTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateRegistryTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrivateRegistryTypeResponseBody) GetRegistryTypeInfos() []*ListPrivateRegistryTypeResponseBodyRegistryTypeInfos {
	return s.RegistryTypeInfos
}

func (s *ListPrivateRegistryTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrivateRegistryTypeResponseBody) SetRegistryTypeInfos(v []*ListPrivateRegistryTypeResponseBodyRegistryTypeInfos) *ListPrivateRegistryTypeResponseBody {
	s.RegistryTypeInfos = v
	return s
}

func (s *ListPrivateRegistryTypeResponseBody) SetRequestId(v string) *ListPrivateRegistryTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrivateRegistryTypeResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPrivateRegistryTypeResponseBodyRegistryTypeInfos struct {
	// The number of image repositories.
	//
	// example:
	//
	// 2
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the image repository type. Valid values:
	//
	// 	- **acr**: Container Registry
	//
	// 	- **harbor**: Harbor
	//
	// 	- **quay**: Quay
	//
	// 	- **CI/CD**: Jenkins
	//
	// example:
	//
	// harbor
	RegistryType *string `json:"RegistryType,omitempty" xml:"RegistryType,omitempty"`
}

func (s ListPrivateRegistryTypeResponseBodyRegistryTypeInfos) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateRegistryTypeResponseBodyRegistryTypeInfos) GoString() string {
	return s.String()
}

func (s *ListPrivateRegistryTypeResponseBodyRegistryTypeInfos) GetCount() *int64 {
	return s.Count
}

func (s *ListPrivateRegistryTypeResponseBodyRegistryTypeInfos) GetRegistryType() *string {
	return s.RegistryType
}

func (s *ListPrivateRegistryTypeResponseBodyRegistryTypeInfos) SetCount(v int64) *ListPrivateRegistryTypeResponseBodyRegistryTypeInfos {
	s.Count = &v
	return s
}

func (s *ListPrivateRegistryTypeResponseBodyRegistryTypeInfos) SetRegistryType(v string) *ListPrivateRegistryTypeResponseBodyRegistryTypeInfos {
	s.RegistryType = &v
	return s
}

func (s *ListPrivateRegistryTypeResponseBodyRegistryTypeInfos) Validate() error {
	return dara.Validate(s)
}
