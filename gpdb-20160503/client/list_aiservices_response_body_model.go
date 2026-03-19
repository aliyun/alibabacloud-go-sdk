// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v string) *ListAIServicesResponseBody
	GetPageNumber() *string
	SetPageRecordCount(v string) *ListAIServicesResponseBody
	GetPageRecordCount() *string
	SetRequestId(v string) *ListAIServicesResponseBody
	GetRequestId() *string
	SetServices(v []*ListAIServicesResponseBodyServices) *ListAIServicesResponseBody
	GetServices() []*ListAIServicesResponseBodyServices
	SetTotalRecordCount(v string) *ListAIServicesResponseBody
	GetTotalRecordCount() *string
}

type ListAIServicesResponseBody struct {
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageRecordCount *string `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Services  []*ListAIServicesResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalRecordCount *string `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListAIServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAIServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAIServicesResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListAIServicesResponseBody) GetPageRecordCount() *string {
	return s.PageRecordCount
}

func (s *ListAIServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAIServicesResponseBody) GetServices() []*ListAIServicesResponseBodyServices {
	return s.Services
}

func (s *ListAIServicesResponseBody) GetTotalRecordCount() *string {
	return s.TotalRecordCount
}

func (s *ListAIServicesResponseBody) SetPageNumber(v string) *ListAIServicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAIServicesResponseBody) SetPageRecordCount(v string) *ListAIServicesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *ListAIServicesResponseBody) SetRequestId(v string) *ListAIServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAIServicesResponseBody) SetServices(v []*ListAIServicesResponseBodyServices) *ListAIServicesResponseBody {
	s.Services = v
	return s
}

func (s *ListAIServicesResponseBody) SetTotalRecordCount(v string) *ListAIServicesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *ListAIServicesResponseBody) Validate() error {
	if s.Services != nil {
		for _, item := range s.Services {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAIServicesResponseBodyServices struct {
	// example:
	//
	// 2026-03-01T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// dramatest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 192.168.0.1/api-docs
	PrivateApiDevUrl *string `json:"PrivateApiDevUrl,omitempty" xml:"PrivateApiDevUrl,omitempty"`
	// example:
	//
	// 192.168.0.1
	PrivateWorkbenchUrl *string `json:"PrivateWorkbenchUrl,omitempty" xml:"PrivateWorkbenchUrl,omitempty"`
	// example:
	//
	// 8.8.8.8/api-docs
	PublicApiDevUrl *string `json:"PublicApiDevUrl,omitempty" xml:"PublicApiDevUrl,omitempty"`
	// example:
	//
	// 8.8.8.8
	PublicWorkbenchUrl *string `json:"PublicWorkbenchUrl,omitempty" xml:"PublicWorkbenchUrl,omitempty"`
	// example:
	//
	// 127.0.0.1
	SecurityIpList *string `json:"SecurityIpList,omitempty" xml:"SecurityIpList,omitempty"`
	// example:
	//
	// dramauser
	ServiceAccount *string `json:"ServiceAccount,omitempty" xml:"ServiceAccount,omitempty"`
	// example:
	//
	// drama-123456
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAIServicesResponseBodyServices) String() string {
	return dara.Prettify(s)
}

func (s ListAIServicesResponseBodyServices) GoString() string {
	return s.String()
}

func (s *ListAIServicesResponseBodyServices) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAIServicesResponseBodyServices) GetDescription() *string {
	return s.Description
}

func (s *ListAIServicesResponseBodyServices) GetPrivateApiDevUrl() *string {
	return s.PrivateApiDevUrl
}

func (s *ListAIServicesResponseBodyServices) GetPrivateWorkbenchUrl() *string {
	return s.PrivateWorkbenchUrl
}

func (s *ListAIServicesResponseBodyServices) GetPublicApiDevUrl() *string {
	return s.PublicApiDevUrl
}

func (s *ListAIServicesResponseBodyServices) GetPublicWorkbenchUrl() *string {
	return s.PublicWorkbenchUrl
}

func (s *ListAIServicesResponseBodyServices) GetSecurityIpList() *string {
	return s.SecurityIpList
}

func (s *ListAIServicesResponseBodyServices) GetServiceAccount() *string {
	return s.ServiceAccount
}

func (s *ListAIServicesResponseBodyServices) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListAIServicesResponseBodyServices) GetStatus() *string {
	return s.Status
}

func (s *ListAIServicesResponseBodyServices) SetCreateTime(v string) *ListAIServicesResponseBodyServices {
	s.CreateTime = &v
	return s
}

func (s *ListAIServicesResponseBodyServices) SetDescription(v string) *ListAIServicesResponseBodyServices {
	s.Description = &v
	return s
}

func (s *ListAIServicesResponseBodyServices) SetPrivateApiDevUrl(v string) *ListAIServicesResponseBodyServices {
	s.PrivateApiDevUrl = &v
	return s
}

func (s *ListAIServicesResponseBodyServices) SetPrivateWorkbenchUrl(v string) *ListAIServicesResponseBodyServices {
	s.PrivateWorkbenchUrl = &v
	return s
}

func (s *ListAIServicesResponseBodyServices) SetPublicApiDevUrl(v string) *ListAIServicesResponseBodyServices {
	s.PublicApiDevUrl = &v
	return s
}

func (s *ListAIServicesResponseBodyServices) SetPublicWorkbenchUrl(v string) *ListAIServicesResponseBodyServices {
	s.PublicWorkbenchUrl = &v
	return s
}

func (s *ListAIServicesResponseBodyServices) SetSecurityIpList(v string) *ListAIServicesResponseBodyServices {
	s.SecurityIpList = &v
	return s
}

func (s *ListAIServicesResponseBodyServices) SetServiceAccount(v string) *ListAIServicesResponseBodyServices {
	s.ServiceAccount = &v
	return s
}

func (s *ListAIServicesResponseBodyServices) SetServiceId(v string) *ListAIServicesResponseBodyServices {
	s.ServiceId = &v
	return s
}

func (s *ListAIServicesResponseBodyServices) SetStatus(v string) *ListAIServicesResponseBodyServices {
	s.Status = &v
	return s
}

func (s *ListAIServicesResponseBodyServices) Validate() error {
	return dara.Validate(s)
}
