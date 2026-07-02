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
	// The user configuration information.
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
	if s.TenantConfigModel != nil {
		if err := s.TenantConfigModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTenantConfigResponseBodyTenantConfigModel struct {
	// Indicates whether resource expiration reminders are enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Not enabled.
	//
	// example:
	//
	// true
	AppInstanceGroupExpireRemind *bool `json:"AppInstanceGroupExpireRemind,omitempty" xml:"AppInstanceGroupExpireRemind,omitempty"`
	// example:
	//
	// None
	MultiSessionSupportType      *string   `json:"MultiSessionSupportType,omitempty" xml:"MultiSessionSupportType,omitempty"`
	MultiSessionSupportedRegions []*string `json:"MultiSessionSupportedRegions,omitempty" xml:"MultiSessionSupportedRegions,omitempty" type:"Repeated"`
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

func (s *ListTenantConfigResponseBodyTenantConfigModel) GetMultiSessionSupportType() *string {
	return s.MultiSessionSupportType
}

func (s *ListTenantConfigResponseBodyTenantConfigModel) GetMultiSessionSupportedRegions() []*string {
	return s.MultiSessionSupportedRegions
}

func (s *ListTenantConfigResponseBodyTenantConfigModel) SetAppInstanceGroupExpireRemind(v bool) *ListTenantConfigResponseBodyTenantConfigModel {
	s.AppInstanceGroupExpireRemind = &v
	return s
}

func (s *ListTenantConfigResponseBodyTenantConfigModel) SetMultiSessionSupportType(v string) *ListTenantConfigResponseBodyTenantConfigModel {
	s.MultiSessionSupportType = &v
	return s
}

func (s *ListTenantConfigResponseBodyTenantConfigModel) SetMultiSessionSupportedRegions(v []*string) *ListTenantConfigResponseBodyTenantConfigModel {
	s.MultiSessionSupportedRegions = v
	return s
}

func (s *ListTenantConfigResponseBodyTenantConfigModel) Validate() error {
	return dara.Validate(s)
}
