// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListTenantConfigResponseBody
	GetRequestId() *string
	SetTenantConfigModel(v *ListTenantConfigResponseBodyTenantConfigModel) *ListTenantConfigResponseBody
	GetTenantConfigModel() *ListTenantConfigResponseBodyTenantConfigModel
}

type ListTenantConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The user configurations.
	TenantConfigModel *ListTenantConfigResponseBodyTenantConfigModel `json:"TenantConfigModel,omitempty" xml:"TenantConfigModel,omitempty" type:"Struct"`
}

func (s ListTenantConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTenantConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListTenantConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTenantConfigResponseBody) GetTenantConfigModel() *ListTenantConfigResponseBodyTenantConfigModel {
	return s.TenantConfigModel
}

func (s *ListTenantConfigResponseBody) SetRequestId(v string) *ListTenantConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTenantConfigResponseBody) SetTenantConfigModel(v *ListTenantConfigResponseBodyTenantConfigModel) *ListTenantConfigResponseBody {
	s.TenantConfigModel = v
	return s
}

func (s *ListTenantConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTenantConfigResponseBodyTenantConfigModel struct {
	// Indicates whether the resource expiration reminder feature is enabled.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AppInstanceGroupExpireRemind *bool `json:"AppInstanceGroupExpireRemind,omitempty" xml:"AppInstanceGroupExpireRemind,omitempty"`
}

func (s ListTenantConfigResponseBodyTenantConfigModel) String() string {
	return dara.Prettify(s)
}

func (s ListTenantConfigResponseBodyTenantConfigModel) GoString() string {
	return s.String()
}

func (s *ListTenantConfigResponseBodyTenantConfigModel) GetAppInstanceGroupExpireRemind() *bool {
	return s.AppInstanceGroupExpireRemind
}

func (s *ListTenantConfigResponseBodyTenantConfigModel) SetAppInstanceGroupExpireRemind(v bool) *ListTenantConfigResponseBodyTenantConfigModel {
	s.AppInstanceGroupExpireRemind = &v
	return s
}

func (s *ListTenantConfigResponseBodyTenantConfigModel) Validate() error {
	return dara.Validate(s)
}
