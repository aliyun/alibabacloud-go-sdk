// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceExtensionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*ListInstanceExtensionsResponseBodyItems) *ListInstanceExtensionsResponseBody
	GetItems() []*ListInstanceExtensionsResponseBodyItems
	SetPageNumber(v int32) *ListInstanceExtensionsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *ListInstanceExtensionsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *ListInstanceExtensionsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *ListInstanceExtensionsResponseBody
	GetTotalRecordCount() *int32
}

type ListInstanceExtensionsResponseBody struct {
	// The queried extensions.
	Items []*ListInstanceExtensionsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7565770E-7C45-462D-BA4A-8A5396F2CAD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListInstanceExtensionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceExtensionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceExtensionsResponseBody) GetItems() []*ListInstanceExtensionsResponseBodyItems {
	return s.Items
}

func (s *ListInstanceExtensionsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstanceExtensionsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *ListInstanceExtensionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceExtensionsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *ListInstanceExtensionsResponseBody) SetItems(v []*ListInstanceExtensionsResponseBodyItems) *ListInstanceExtensionsResponseBody {
	s.Items = v
	return s
}

func (s *ListInstanceExtensionsResponseBody) SetPageNumber(v int32) *ListInstanceExtensionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceExtensionsResponseBody) SetPageRecordCount(v int32) *ListInstanceExtensionsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *ListInstanceExtensionsResponseBody) SetRequestId(v string) *ListInstanceExtensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceExtensionsResponseBody) SetTotalRecordCount(v int32) *ListInstanceExtensionsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *ListInstanceExtensionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInstanceExtensionsResponseBodyItems struct {
	// The current version.
	//
	// example:
	//
	// 1.0
	CurrentVersion *string `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// The description of the extension.
	//
	// example:
	//
	// citext usage
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The extension ID.
	//
	// example:
	//
	// 1
	ExtensionId *string `json:"ExtensionId,omitempty" xml:"ExtensionId,omitempty"`
	// The names of the databases in which the extension is installed.
	//
	// example:
	//
	// test
	InstalledDatabases *string `json:"InstalledDatabases,omitempty" xml:"InstalledDatabases,omitempty"`
	// Indicates whether an instance restart is required after you install the extension for the extension to take effect.
	//
	// example:
	//
	// false
	IsInstallNeedRestart *bool `json:"IsInstallNeedRestart,omitempty" xml:"IsInstallNeedRestart,omitempty"`
	// The latest version.
	//
	// example:
	//
	// 1.1
	LatestVersion *string `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	// The name of the extension.
	//
	// example:
	//
	// citext
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the extension.
	//
	// example:
	//
	// installed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstanceExtensionsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceExtensionsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListInstanceExtensionsResponseBodyItems) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *ListInstanceExtensionsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListInstanceExtensionsResponseBodyItems) GetExtensionId() *string {
	return s.ExtensionId
}

func (s *ListInstanceExtensionsResponseBodyItems) GetInstalledDatabases() *string {
	return s.InstalledDatabases
}

func (s *ListInstanceExtensionsResponseBodyItems) GetIsInstallNeedRestart() *bool {
	return s.IsInstallNeedRestart
}

func (s *ListInstanceExtensionsResponseBodyItems) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *ListInstanceExtensionsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListInstanceExtensionsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListInstanceExtensionsResponseBodyItems) SetCurrentVersion(v string) *ListInstanceExtensionsResponseBodyItems {
	s.CurrentVersion = &v
	return s
}

func (s *ListInstanceExtensionsResponseBodyItems) SetDescription(v string) *ListInstanceExtensionsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListInstanceExtensionsResponseBodyItems) SetExtensionId(v string) *ListInstanceExtensionsResponseBodyItems {
	s.ExtensionId = &v
	return s
}

func (s *ListInstanceExtensionsResponseBodyItems) SetInstalledDatabases(v string) *ListInstanceExtensionsResponseBodyItems {
	s.InstalledDatabases = &v
	return s
}

func (s *ListInstanceExtensionsResponseBodyItems) SetIsInstallNeedRestart(v bool) *ListInstanceExtensionsResponseBodyItems {
	s.IsInstallNeedRestart = &v
	return s
}

func (s *ListInstanceExtensionsResponseBodyItems) SetLatestVersion(v string) *ListInstanceExtensionsResponseBodyItems {
	s.LatestVersion = &v
	return s
}

func (s *ListInstanceExtensionsResponseBodyItems) SetName(v string) *ListInstanceExtensionsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListInstanceExtensionsResponseBodyItems) SetStatus(v string) *ListInstanceExtensionsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListInstanceExtensionsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
