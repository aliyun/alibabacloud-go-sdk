// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExtensionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstalledExtensions(v []*DescribeExtensionsResponseBodyInstalledExtensions) *DescribeExtensionsResponseBody
	GetInstalledExtensions() []*DescribeExtensionsResponseBodyInstalledExtensions
	SetOverview(v string) *DescribeExtensionsResponseBody
	GetOverview() *string
	SetRequestId(v string) *DescribeExtensionsResponseBody
	GetRequestId() *string
	SetUninstalledExtensions(v []*DescribeExtensionsResponseBodyUninstalledExtensions) *DescribeExtensionsResponseBody
	GetUninstalledExtensions() []*DescribeExtensionsResponseBodyUninstalledExtensions
}

type DescribeExtensionsResponseBody struct {
	InstalledExtensions []*DescribeExtensionsResponseBodyInstalledExtensions `json:"InstalledExtensions,omitempty" xml:"InstalledExtensions,omitempty" type:"Repeated"`
	Overview            *string                                              `json:"Overview,omitempty" xml:"Overview,omitempty"`
	// Id of the request
	RequestId             *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UninstalledExtensions []*DescribeExtensionsResponseBodyUninstalledExtensions `json:"UninstalledExtensions,omitempty" xml:"UninstalledExtensions,omitempty" type:"Repeated"`
}

func (s DescribeExtensionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExtensionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExtensionsResponseBody) GetInstalledExtensions() []*DescribeExtensionsResponseBodyInstalledExtensions {
	return s.InstalledExtensions
}

func (s *DescribeExtensionsResponseBody) GetOverview() *string {
	return s.Overview
}

func (s *DescribeExtensionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExtensionsResponseBody) GetUninstalledExtensions() []*DescribeExtensionsResponseBodyUninstalledExtensions {
	return s.UninstalledExtensions
}

func (s *DescribeExtensionsResponseBody) SetInstalledExtensions(v []*DescribeExtensionsResponseBodyInstalledExtensions) *DescribeExtensionsResponseBody {
	s.InstalledExtensions = v
	return s
}

func (s *DescribeExtensionsResponseBody) SetOverview(v string) *DescribeExtensionsResponseBody {
	s.Overview = &v
	return s
}

func (s *DescribeExtensionsResponseBody) SetRequestId(v string) *DescribeExtensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExtensionsResponseBody) SetUninstalledExtensions(v []*DescribeExtensionsResponseBodyUninstalledExtensions) *DescribeExtensionsResponseBody {
	s.UninstalledExtensions = v
	return s
}

func (s *DescribeExtensionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeExtensionsResponseBodyInstalledExtensions struct {
	// example:
	//
	// geography_space, self_develop
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// OK
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 7.7
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// example:
	//
	// 7.7
	InstalledVersion *string `json:"InstalledVersion,omitempty" xml:"InstalledVersion,omitempty"`
	// example:
	//
	// jueming
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// alton
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// ganos_networking
	Requires *string `json:"Requires,omitempty" xml:"Requires,omitempty"`
	// example:
	//
	// true
	Restart *string `json:"Restart,omitempty" xml:"Restart,omitempty"`
}

func (s DescribeExtensionsResponseBodyInstalledExtensions) String() string {
	return dara.Prettify(s)
}

func (s DescribeExtensionsResponseBodyInstalledExtensions) GoString() string {
	return s.String()
}

func (s *DescribeExtensionsResponseBodyInstalledExtensions) GetCategory() *string {
	return s.Category
}

func (s *DescribeExtensionsResponseBodyInstalledExtensions) GetComment() *string {
	return s.Comment
}

func (s *DescribeExtensionsResponseBodyInstalledExtensions) GetDefaultVersion() *string {
	return s.DefaultVersion
}

func (s *DescribeExtensionsResponseBodyInstalledExtensions) GetInstalledVersion() *string {
	return s.InstalledVersion
}

func (s *DescribeExtensionsResponseBodyInstalledExtensions) GetName() *string {
	return s.Name
}

func (s *DescribeExtensionsResponseBodyInstalledExtensions) GetOwner() *string {
	return s.Owner
}

func (s *DescribeExtensionsResponseBodyInstalledExtensions) GetPriority() *string {
	return s.Priority
}

func (s *DescribeExtensionsResponseBodyInstalledExtensions) GetRequires() *string {
	return s.Requires
}

func (s *DescribeExtensionsResponseBodyInstalledExtensions) GetRestart() *string {
	return s.Restart
}

func (s *DescribeExtensionsResponseBodyInstalledExtensions) SetCategory(v string) *DescribeExtensionsResponseBodyInstalledExtensions {
	s.Category = &v
	return s
}

func (s *DescribeExtensionsResponseBodyInstalledExtensions) SetComment(v string) *DescribeExtensionsResponseBodyInstalledExtensions {
	s.Comment = &v
	return s
}

func (s *DescribeExtensionsResponseBodyInstalledExtensions) SetDefaultVersion(v string) *DescribeExtensionsResponseBodyInstalledExtensions {
	s.DefaultVersion = &v
	return s
}

func (s *DescribeExtensionsResponseBodyInstalledExtensions) SetInstalledVersion(v string) *DescribeExtensionsResponseBodyInstalledExtensions {
	s.InstalledVersion = &v
	return s
}

func (s *DescribeExtensionsResponseBodyInstalledExtensions) SetName(v string) *DescribeExtensionsResponseBodyInstalledExtensions {
	s.Name = &v
	return s
}

func (s *DescribeExtensionsResponseBodyInstalledExtensions) SetOwner(v string) *DescribeExtensionsResponseBodyInstalledExtensions {
	s.Owner = &v
	return s
}

func (s *DescribeExtensionsResponseBodyInstalledExtensions) SetPriority(v string) *DescribeExtensionsResponseBodyInstalledExtensions {
	s.Priority = &v
	return s
}

func (s *DescribeExtensionsResponseBodyInstalledExtensions) SetRequires(v string) *DescribeExtensionsResponseBodyInstalledExtensions {
	s.Requires = &v
	return s
}

func (s *DescribeExtensionsResponseBodyInstalledExtensions) SetRestart(v string) *DescribeExtensionsResponseBodyInstalledExtensions {
	s.Restart = &v
	return s
}

func (s *DescribeExtensionsResponseBodyInstalledExtensions) Validate() error {
	return dara.Validate(s)
}

type DescribeExtensionsResponseBodyUninstalledExtensions struct {
	// example:
	//
	// geography_space, self_develop
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// OK
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 7.7
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// example:
	//
	// 7.7
	InstalledVersion *string `json:"InstalledVersion,omitempty" xml:"InstalledVersion,omitempty"`
	// example:
	//
	// jueming
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// alton
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// ganos_networking
	Requires *string `json:"Requires,omitempty" xml:"Requires,omitempty"`
	// example:
	//
	// true
	Restart *string `json:"Restart,omitempty" xml:"Restart,omitempty"`
}

func (s DescribeExtensionsResponseBodyUninstalledExtensions) String() string {
	return dara.Prettify(s)
}

func (s DescribeExtensionsResponseBodyUninstalledExtensions) GoString() string {
	return s.String()
}

func (s *DescribeExtensionsResponseBodyUninstalledExtensions) GetCategory() *string {
	return s.Category
}

func (s *DescribeExtensionsResponseBodyUninstalledExtensions) GetComment() *string {
	return s.Comment
}

func (s *DescribeExtensionsResponseBodyUninstalledExtensions) GetDefaultVersion() *string {
	return s.DefaultVersion
}

func (s *DescribeExtensionsResponseBodyUninstalledExtensions) GetInstalledVersion() *string {
	return s.InstalledVersion
}

func (s *DescribeExtensionsResponseBodyUninstalledExtensions) GetName() *string {
	return s.Name
}

func (s *DescribeExtensionsResponseBodyUninstalledExtensions) GetOwner() *string {
	return s.Owner
}

func (s *DescribeExtensionsResponseBodyUninstalledExtensions) GetPriority() *string {
	return s.Priority
}

func (s *DescribeExtensionsResponseBodyUninstalledExtensions) GetRequires() *string {
	return s.Requires
}

func (s *DescribeExtensionsResponseBodyUninstalledExtensions) GetRestart() *string {
	return s.Restart
}

func (s *DescribeExtensionsResponseBodyUninstalledExtensions) SetCategory(v string) *DescribeExtensionsResponseBodyUninstalledExtensions {
	s.Category = &v
	return s
}

func (s *DescribeExtensionsResponseBodyUninstalledExtensions) SetComment(v string) *DescribeExtensionsResponseBodyUninstalledExtensions {
	s.Comment = &v
	return s
}

func (s *DescribeExtensionsResponseBodyUninstalledExtensions) SetDefaultVersion(v string) *DescribeExtensionsResponseBodyUninstalledExtensions {
	s.DefaultVersion = &v
	return s
}

func (s *DescribeExtensionsResponseBodyUninstalledExtensions) SetInstalledVersion(v string) *DescribeExtensionsResponseBodyUninstalledExtensions {
	s.InstalledVersion = &v
	return s
}

func (s *DescribeExtensionsResponseBodyUninstalledExtensions) SetName(v string) *DescribeExtensionsResponseBodyUninstalledExtensions {
	s.Name = &v
	return s
}

func (s *DescribeExtensionsResponseBodyUninstalledExtensions) SetOwner(v string) *DescribeExtensionsResponseBodyUninstalledExtensions {
	s.Owner = &v
	return s
}

func (s *DescribeExtensionsResponseBodyUninstalledExtensions) SetPriority(v string) *DescribeExtensionsResponseBodyUninstalledExtensions {
	s.Priority = &v
	return s
}

func (s *DescribeExtensionsResponseBodyUninstalledExtensions) SetRequires(v string) *DescribeExtensionsResponseBodyUninstalledExtensions {
	s.Requires = &v
	return s
}

func (s *DescribeExtensionsResponseBodyUninstalledExtensions) SetRestart(v string) *DescribeExtensionsResponseBodyUninstalledExtensions {
	s.Restart = &v
	return s
}

func (s *DescribeExtensionsResponseBodyUninstalledExtensions) Validate() error {
	return dara.Validate(s)
}
