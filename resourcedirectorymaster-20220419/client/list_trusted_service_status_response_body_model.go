// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrustedServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnabledServicePrincipals(v *ListTrustedServiceStatusResponseBodyEnabledServicePrincipals) *ListTrustedServiceStatusResponseBody
	GetEnabledServicePrincipals() *ListTrustedServiceStatusResponseBodyEnabledServicePrincipals
	SetPageNumber(v int32) *ListTrustedServiceStatusResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTrustedServiceStatusResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTrustedServiceStatusResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTrustedServiceStatusResponseBody
	GetTotalCount() *int32
}

type ListTrustedServiceStatusResponseBody struct {
	// The information about the trusted services that are enabled.
	EnabledServicePrincipals *ListTrustedServiceStatusResponseBodyEnabledServicePrincipals `json:"EnabledServicePrincipals,omitempty" xml:"EnabledServicePrincipals,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CD76D376-2517-4924-92C5-DBC52262F93A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTrustedServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTrustedServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrustedServiceStatusResponseBody) GetEnabledServicePrincipals() *ListTrustedServiceStatusResponseBodyEnabledServicePrincipals {
	return s.EnabledServicePrincipals
}

func (s *ListTrustedServiceStatusResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTrustedServiceStatusResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTrustedServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTrustedServiceStatusResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTrustedServiceStatusResponseBody) SetEnabledServicePrincipals(v *ListTrustedServiceStatusResponseBodyEnabledServicePrincipals) *ListTrustedServiceStatusResponseBody {
	s.EnabledServicePrincipals = v
	return s
}

func (s *ListTrustedServiceStatusResponseBody) SetPageNumber(v int32) *ListTrustedServiceStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTrustedServiceStatusResponseBody) SetPageSize(v int32) *ListTrustedServiceStatusResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTrustedServiceStatusResponseBody) SetRequestId(v string) *ListTrustedServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrustedServiceStatusResponseBody) SetTotalCount(v int32) *ListTrustedServiceStatusResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTrustedServiceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTrustedServiceStatusResponseBodyEnabledServicePrincipals struct {
	EnabledServicePrincipal []*ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal `json:"EnabledServicePrincipal,omitempty" xml:"EnabledServicePrincipal,omitempty" type:"Repeated"`
}

func (s ListTrustedServiceStatusResponseBodyEnabledServicePrincipals) String() string {
	return dara.Prettify(s)
}

func (s ListTrustedServiceStatusResponseBodyEnabledServicePrincipals) GoString() string {
	return s.String()
}

func (s *ListTrustedServiceStatusResponseBodyEnabledServicePrincipals) GetEnabledServicePrincipal() []*ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal {
	return s.EnabledServicePrincipal
}

func (s *ListTrustedServiceStatusResponseBodyEnabledServicePrincipals) SetEnabledServicePrincipal(v []*ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal) *ListTrustedServiceStatusResponseBodyEnabledServicePrincipals {
	s.EnabledServicePrincipal = v
	return s
}

func (s *ListTrustedServiceStatusResponseBodyEnabledServicePrincipals) Validate() error {
	return dara.Validate(s)
}

type ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal struct {
	// The time when the trusted service was enabled.
	//
	// example:
	//
	// 2019-02-18T15:32:10.473Z
	EnableTime *string `json:"EnableTime,omitempty" xml:"EnableTime,omitempty"`
	// The identifier of the trusted service.
	//
	// example:
	//
	// config.aliyuncs.com
	ServicePrincipal *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty"`
}

func (s ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal) String() string {
	return dara.Prettify(s)
}

func (s ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal) GoString() string {
	return s.String()
}

func (s *ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal) GetEnableTime() *string {
	return s.EnableTime
}

func (s *ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal) GetServicePrincipal() *string {
	return s.ServicePrincipal
}

func (s *ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal) SetEnableTime(v string) *ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal {
	s.EnableTime = &v
	return s
}

func (s *ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal) SetServicePrincipal(v string) *ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal {
	s.ServicePrincipal = &v
	return s
}

func (s *ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal) Validate() error {
	return dara.Validate(s)
}
