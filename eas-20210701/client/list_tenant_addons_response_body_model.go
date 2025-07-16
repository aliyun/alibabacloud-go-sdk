// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantAddonsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddons(v []*ListTenantAddonsResponseBodyAddons) *ListTenantAddonsResponseBody
	GetAddons() []*ListTenantAddonsResponseBodyAddons
	SetRequestId(v string) *ListTenantAddonsResponseBody
	GetRequestId() *string
}

type ListTenantAddonsResponseBody struct {
	// The information about the plug-in.
	Addons []*ListTenantAddonsResponseBodyAddons `json:"Addons,omitempty" xml:"Addons,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTenantAddonsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTenantAddonsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTenantAddonsResponseBody) GetAddons() []*ListTenantAddonsResponseBodyAddons {
	return s.Addons
}

func (s *ListTenantAddonsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTenantAddonsResponseBody) SetAddons(v []*ListTenantAddonsResponseBodyAddons) *ListTenantAddonsResponseBody {
	s.Addons = v
	return s
}

func (s *ListTenantAddonsResponseBody) SetRequestId(v string) *ListTenantAddonsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTenantAddonsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTenantAddonsResponseBodyAddons struct {
	// The attributes of the plug-in.
	Attributes map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The name of the plug-in.
	//
	// example:
	//
	// prometheus_discovery
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListTenantAddonsResponseBodyAddons) String() string {
	return dara.Prettify(s)
}

func (s ListTenantAddonsResponseBodyAddons) GoString() string {
	return s.String()
}

func (s *ListTenantAddonsResponseBodyAddons) GetAttributes() map[string]*string {
	return s.Attributes
}

func (s *ListTenantAddonsResponseBodyAddons) GetName() *string {
	return s.Name
}

func (s *ListTenantAddonsResponseBodyAddons) SetAttributes(v map[string]*string) *ListTenantAddonsResponseBodyAddons {
	s.Attributes = v
	return s
}

func (s *ListTenantAddonsResponseBodyAddons) SetName(v string) *ListTenantAddonsResponseBodyAddons {
	s.Name = &v
	return s
}

func (s *ListTenantAddonsResponseBodyAddons) Validate() error {
	return dara.Validate(s)
}
