// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationSupportedProvisionProtocolTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationSupportedProvisionProtocolType(v *ListApplicationSupportedProvisionProtocolTypesResponseBodyApplicationSupportedProvisionProtocolType) *ListApplicationSupportedProvisionProtocolTypesResponseBody
	GetApplicationSupportedProvisionProtocolType() *ListApplicationSupportedProvisionProtocolTypesResponseBodyApplicationSupportedProvisionProtocolType
	SetRequestId(v string) *ListApplicationSupportedProvisionProtocolTypesResponseBody
	GetRequestId() *string
}

type ListApplicationSupportedProvisionProtocolTypesResponseBody struct {
	ApplicationSupportedProvisionProtocolType *ListApplicationSupportedProvisionProtocolTypesResponseBodyApplicationSupportedProvisionProtocolType `json:"ApplicationSupportedProvisionProtocolType,omitempty" xml:"ApplicationSupportedProvisionProtocolType,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListApplicationSupportedProvisionProtocolTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationSupportedProvisionProtocolTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationSupportedProvisionProtocolTypesResponseBody) GetApplicationSupportedProvisionProtocolType() *ListApplicationSupportedProvisionProtocolTypesResponseBodyApplicationSupportedProvisionProtocolType {
	return s.ApplicationSupportedProvisionProtocolType
}

func (s *ListApplicationSupportedProvisionProtocolTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationSupportedProvisionProtocolTypesResponseBody) SetApplicationSupportedProvisionProtocolType(v *ListApplicationSupportedProvisionProtocolTypesResponseBodyApplicationSupportedProvisionProtocolType) *ListApplicationSupportedProvisionProtocolTypesResponseBody {
	s.ApplicationSupportedProvisionProtocolType = v
	return s
}

func (s *ListApplicationSupportedProvisionProtocolTypesResponseBody) SetRequestId(v string) *ListApplicationSupportedProvisionProtocolTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationSupportedProvisionProtocolTypesResponseBody) Validate() error {
	if s.ApplicationSupportedProvisionProtocolType != nil {
		if err := s.ApplicationSupportedProvisionProtocolType.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApplicationSupportedProvisionProtocolTypesResponseBodyApplicationSupportedProvisionProtocolType struct {
	// 账户同步支持类型
	ProvisionProtocolType []*string `json:"ProvisionProtocolType,omitempty" xml:"ProvisionProtocolType,omitempty" type:"Repeated"`
}

func (s ListApplicationSupportedProvisionProtocolTypesResponseBodyApplicationSupportedProvisionProtocolType) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationSupportedProvisionProtocolTypesResponseBodyApplicationSupportedProvisionProtocolType) GoString() string {
	return s.String()
}

func (s *ListApplicationSupportedProvisionProtocolTypesResponseBodyApplicationSupportedProvisionProtocolType) GetProvisionProtocolType() []*string {
	return s.ProvisionProtocolType
}

func (s *ListApplicationSupportedProvisionProtocolTypesResponseBodyApplicationSupportedProvisionProtocolType) SetProvisionProtocolType(v []*string) *ListApplicationSupportedProvisionProtocolTypesResponseBodyApplicationSupportedProvisionProtocolType {
	s.ProvisionProtocolType = v
	return s
}

func (s *ListApplicationSupportedProvisionProtocolTypesResponseBodyApplicationSupportedProvisionProtocolType) Validate() error {
	return dara.Validate(s)
}
