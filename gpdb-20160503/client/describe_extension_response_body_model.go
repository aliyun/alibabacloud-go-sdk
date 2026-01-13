// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExtensionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentVersion(v string) *DescribeExtensionResponseBody
	GetCurrentVersion() *string
	SetDescription(v string) *DescribeExtensionResponseBody
	GetDescription() *string
	SetExtensionId(v string) *DescribeExtensionResponseBody
	GetExtensionId() *string
	SetExtensionName(v string) *DescribeExtensionResponseBody
	GetExtensionName() *string
	SetIsInstallNeedRestart(v bool) *DescribeExtensionResponseBody
	GetIsInstallNeedRestart() *bool
	SetIsLatestVersion(v bool) *DescribeExtensionResponseBody
	GetIsLatestVersion() *bool
	SetLatestVersion(v string) *DescribeExtensionResponseBody
	GetLatestVersion() *string
	SetRequestId(v string) *DescribeExtensionResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeExtensionResponseBody
	GetStatus() *string
}

type DescribeExtensionResponseBody struct {
	// The current version.
	//
	// example:
	//
	// 2.1
	CurrentVersion *string `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// The description of the extension.
	//
	// example:
	//
	// zhparser
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The extension ID.
	//
	// example:
	//
	// 37
	ExtensionId *string `json:"ExtensionId,omitempty" xml:"ExtensionId,omitempty"`
	// The extension name.
	//
	// example:
	//
	// zhparser
	ExtensionName *string `json:"ExtensionName,omitempty" xml:"ExtensionName,omitempty"`
	// Indicates whether an instance restart is required after you install the extension.
	//
	// example:
	//
	// False
	IsInstallNeedRestart *bool `json:"IsInstallNeedRestart,omitempty" xml:"IsInstallNeedRestart,omitempty"`
	// Whether it is the latest version extension.
	//
	// example:
	//
	// True
	IsLatestVersion *bool `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	// The latest version.
	//
	// example:
	//
	// 2.1
	LatestVersion *string `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	// Request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the extension.
	//
	// example:
	//
	// installed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeExtensionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExtensionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExtensionResponseBody) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *DescribeExtensionResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeExtensionResponseBody) GetExtensionId() *string {
	return s.ExtensionId
}

func (s *DescribeExtensionResponseBody) GetExtensionName() *string {
	return s.ExtensionName
}

func (s *DescribeExtensionResponseBody) GetIsInstallNeedRestart() *bool {
	return s.IsInstallNeedRestart
}

func (s *DescribeExtensionResponseBody) GetIsLatestVersion() *bool {
	return s.IsLatestVersion
}

func (s *DescribeExtensionResponseBody) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *DescribeExtensionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExtensionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeExtensionResponseBody) SetCurrentVersion(v string) *DescribeExtensionResponseBody {
	s.CurrentVersion = &v
	return s
}

func (s *DescribeExtensionResponseBody) SetDescription(v string) *DescribeExtensionResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeExtensionResponseBody) SetExtensionId(v string) *DescribeExtensionResponseBody {
	s.ExtensionId = &v
	return s
}

func (s *DescribeExtensionResponseBody) SetExtensionName(v string) *DescribeExtensionResponseBody {
	s.ExtensionName = &v
	return s
}

func (s *DescribeExtensionResponseBody) SetIsInstallNeedRestart(v bool) *DescribeExtensionResponseBody {
	s.IsInstallNeedRestart = &v
	return s
}

func (s *DescribeExtensionResponseBody) SetIsLatestVersion(v bool) *DescribeExtensionResponseBody {
	s.IsLatestVersion = &v
	return s
}

func (s *DescribeExtensionResponseBody) SetLatestVersion(v string) *DescribeExtensionResponseBody {
	s.LatestVersion = &v
	return s
}

func (s *DescribeExtensionResponseBody) SetRequestId(v string) *DescribeExtensionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExtensionResponseBody) SetStatus(v string) *DescribeExtensionResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeExtensionResponseBody) Validate() error {
	return dara.Validate(s)
}
