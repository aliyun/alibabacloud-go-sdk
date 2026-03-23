// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostgresExtensionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstalledExtensions(v []*DescribePostgresExtensionsResponseBodyInstalledExtensions) *DescribePostgresExtensionsResponseBody
	GetInstalledExtensions() []*DescribePostgresExtensionsResponseBodyInstalledExtensions
	SetOverview(v map[string]interface{}) *DescribePostgresExtensionsResponseBody
	GetOverview() map[string]interface{}
	SetRequestId(v string) *DescribePostgresExtensionsResponseBody
	GetRequestId() *string
	SetUninstalledExtensions(v []*DescribePostgresExtensionsResponseBodyUninstalledExtensions) *DescribePostgresExtensionsResponseBody
	GetUninstalledExtensions() []*DescribePostgresExtensionsResponseBodyUninstalledExtensions
}

type DescribePostgresExtensionsResponseBody struct {
	InstalledExtensions   []*DescribePostgresExtensionsResponseBodyInstalledExtensions   `json:"InstalledExtensions,omitempty" xml:"InstalledExtensions,omitempty" type:"Repeated"`
	Overview              map[string]interface{}                                         `json:"Overview,omitempty" xml:"Overview,omitempty"`
	RequestId             *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UninstalledExtensions []*DescribePostgresExtensionsResponseBodyUninstalledExtensions `json:"UninstalledExtensions,omitempty" xml:"UninstalledExtensions,omitempty" type:"Repeated"`
}

func (s DescribePostgresExtensionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePostgresExtensionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePostgresExtensionsResponseBody) GetInstalledExtensions() []*DescribePostgresExtensionsResponseBodyInstalledExtensions {
	return s.InstalledExtensions
}

func (s *DescribePostgresExtensionsResponseBody) GetOverview() map[string]interface{} {
	return s.Overview
}

func (s *DescribePostgresExtensionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePostgresExtensionsResponseBody) GetUninstalledExtensions() []*DescribePostgresExtensionsResponseBodyUninstalledExtensions {
	return s.UninstalledExtensions
}

func (s *DescribePostgresExtensionsResponseBody) SetInstalledExtensions(v []*DescribePostgresExtensionsResponseBodyInstalledExtensions) *DescribePostgresExtensionsResponseBody {
	s.InstalledExtensions = v
	return s
}

func (s *DescribePostgresExtensionsResponseBody) SetOverview(v map[string]interface{}) *DescribePostgresExtensionsResponseBody {
	s.Overview = v
	return s
}

func (s *DescribePostgresExtensionsResponseBody) SetRequestId(v string) *DescribePostgresExtensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePostgresExtensionsResponseBody) SetUninstalledExtensions(v []*DescribePostgresExtensionsResponseBodyUninstalledExtensions) *DescribePostgresExtensionsResponseBody {
	s.UninstalledExtensions = v
	return s
}

func (s *DescribePostgresExtensionsResponseBody) Validate() error {
	if s.InstalledExtensions != nil {
		for _, item := range s.InstalledExtensions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UninstalledExtensions != nil {
		for _, item := range s.UninstalledExtensions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePostgresExtensionsResponseBodyInstalledExtensions struct {
	Category         *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Comment          *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	DefaultVersion   *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	InstalledVersion *string `json:"InstalledVersion,omitempty" xml:"InstalledVersion,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Owner            *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Priority         *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Requires         *string `json:"Requires,omitempty" xml:"Requires,omitempty"`
	Uid              *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DescribePostgresExtensionsResponseBodyInstalledExtensions) String() string {
	return dara.Prettify(s)
}

func (s DescribePostgresExtensionsResponseBodyInstalledExtensions) GoString() string {
	return s.String()
}

func (s *DescribePostgresExtensionsResponseBodyInstalledExtensions) GetCategory() *string {
	return s.Category
}

func (s *DescribePostgresExtensionsResponseBodyInstalledExtensions) GetComment() *string {
	return s.Comment
}

func (s *DescribePostgresExtensionsResponseBodyInstalledExtensions) GetDefaultVersion() *string {
	return s.DefaultVersion
}

func (s *DescribePostgresExtensionsResponseBodyInstalledExtensions) GetInstalledVersion() *string {
	return s.InstalledVersion
}

func (s *DescribePostgresExtensionsResponseBodyInstalledExtensions) GetName() *string {
	return s.Name
}

func (s *DescribePostgresExtensionsResponseBodyInstalledExtensions) GetOwner() *string {
	return s.Owner
}

func (s *DescribePostgresExtensionsResponseBodyInstalledExtensions) GetPriority() *string {
	return s.Priority
}

func (s *DescribePostgresExtensionsResponseBodyInstalledExtensions) GetRequires() *string {
	return s.Requires
}

func (s *DescribePostgresExtensionsResponseBodyInstalledExtensions) GetUid() *string {
	return s.Uid
}

func (s *DescribePostgresExtensionsResponseBodyInstalledExtensions) SetCategory(v string) *DescribePostgresExtensionsResponseBodyInstalledExtensions {
	s.Category = &v
	return s
}

func (s *DescribePostgresExtensionsResponseBodyInstalledExtensions) SetComment(v string) *DescribePostgresExtensionsResponseBodyInstalledExtensions {
	s.Comment = &v
	return s
}

func (s *DescribePostgresExtensionsResponseBodyInstalledExtensions) SetDefaultVersion(v string) *DescribePostgresExtensionsResponseBodyInstalledExtensions {
	s.DefaultVersion = &v
	return s
}

func (s *DescribePostgresExtensionsResponseBodyInstalledExtensions) SetInstalledVersion(v string) *DescribePostgresExtensionsResponseBodyInstalledExtensions {
	s.InstalledVersion = &v
	return s
}

func (s *DescribePostgresExtensionsResponseBodyInstalledExtensions) SetName(v string) *DescribePostgresExtensionsResponseBodyInstalledExtensions {
	s.Name = &v
	return s
}

func (s *DescribePostgresExtensionsResponseBodyInstalledExtensions) SetOwner(v string) *DescribePostgresExtensionsResponseBodyInstalledExtensions {
	s.Owner = &v
	return s
}

func (s *DescribePostgresExtensionsResponseBodyInstalledExtensions) SetPriority(v string) *DescribePostgresExtensionsResponseBodyInstalledExtensions {
	s.Priority = &v
	return s
}

func (s *DescribePostgresExtensionsResponseBodyInstalledExtensions) SetRequires(v string) *DescribePostgresExtensionsResponseBodyInstalledExtensions {
	s.Requires = &v
	return s
}

func (s *DescribePostgresExtensionsResponseBodyInstalledExtensions) SetUid(v string) *DescribePostgresExtensionsResponseBodyInstalledExtensions {
	s.Uid = &v
	return s
}

func (s *DescribePostgresExtensionsResponseBodyInstalledExtensions) Validate() error {
	return dara.Validate(s)
}

type DescribePostgresExtensionsResponseBodyUninstalledExtensions struct {
	Category         *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Comment          *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	DefaultVersion   *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	InstalledVersion *string `json:"InstalledVersion,omitempty" xml:"InstalledVersion,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Owner            *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Priority         *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Requires         *string `json:"Requires,omitempty" xml:"Requires,omitempty"`
	Uid              *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DescribePostgresExtensionsResponseBodyUninstalledExtensions) String() string {
	return dara.Prettify(s)
}

func (s DescribePostgresExtensionsResponseBodyUninstalledExtensions) GoString() string {
	return s.String()
}

func (s *DescribePostgresExtensionsResponseBodyUninstalledExtensions) GetCategory() *string {
	return s.Category
}

func (s *DescribePostgresExtensionsResponseBodyUninstalledExtensions) GetComment() *string {
	return s.Comment
}

func (s *DescribePostgresExtensionsResponseBodyUninstalledExtensions) GetDefaultVersion() *string {
	return s.DefaultVersion
}

func (s *DescribePostgresExtensionsResponseBodyUninstalledExtensions) GetInstalledVersion() *string {
	return s.InstalledVersion
}

func (s *DescribePostgresExtensionsResponseBodyUninstalledExtensions) GetName() *string {
	return s.Name
}

func (s *DescribePostgresExtensionsResponseBodyUninstalledExtensions) GetOwner() *string {
	return s.Owner
}

func (s *DescribePostgresExtensionsResponseBodyUninstalledExtensions) GetPriority() *string {
	return s.Priority
}

func (s *DescribePostgresExtensionsResponseBodyUninstalledExtensions) GetRequires() *string {
	return s.Requires
}

func (s *DescribePostgresExtensionsResponseBodyUninstalledExtensions) GetUid() *string {
	return s.Uid
}

func (s *DescribePostgresExtensionsResponseBodyUninstalledExtensions) SetCategory(v string) *DescribePostgresExtensionsResponseBodyUninstalledExtensions {
	s.Category = &v
	return s
}

func (s *DescribePostgresExtensionsResponseBodyUninstalledExtensions) SetComment(v string) *DescribePostgresExtensionsResponseBodyUninstalledExtensions {
	s.Comment = &v
	return s
}

func (s *DescribePostgresExtensionsResponseBodyUninstalledExtensions) SetDefaultVersion(v string) *DescribePostgresExtensionsResponseBodyUninstalledExtensions {
	s.DefaultVersion = &v
	return s
}

func (s *DescribePostgresExtensionsResponseBodyUninstalledExtensions) SetInstalledVersion(v string) *DescribePostgresExtensionsResponseBodyUninstalledExtensions {
	s.InstalledVersion = &v
	return s
}

func (s *DescribePostgresExtensionsResponseBodyUninstalledExtensions) SetName(v string) *DescribePostgresExtensionsResponseBodyUninstalledExtensions {
	s.Name = &v
	return s
}

func (s *DescribePostgresExtensionsResponseBodyUninstalledExtensions) SetOwner(v string) *DescribePostgresExtensionsResponseBodyUninstalledExtensions {
	s.Owner = &v
	return s
}

func (s *DescribePostgresExtensionsResponseBodyUninstalledExtensions) SetPriority(v string) *DescribePostgresExtensionsResponseBodyUninstalledExtensions {
	s.Priority = &v
	return s
}

func (s *DescribePostgresExtensionsResponseBodyUninstalledExtensions) SetRequires(v string) *DescribePostgresExtensionsResponseBodyUninstalledExtensions {
	s.Requires = &v
	return s
}

func (s *DescribePostgresExtensionsResponseBodyUninstalledExtensions) SetUid(v string) *DescribePostgresExtensionsResponseBodyUninstalledExtensions {
	s.Uid = &v
	return s
}

func (s *DescribePostgresExtensionsResponseBodyUninstalledExtensions) Validate() error {
	return dara.Validate(s)
}
