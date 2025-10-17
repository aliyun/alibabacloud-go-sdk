// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDirectoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDirectories(v []*DescribeDirectoriesResponseBodyDirectories) *DescribeDirectoriesResponseBody
	GetDirectories() []*DescribeDirectoriesResponseBodyDirectories
	SetRequestId(v string) *DescribeDirectoriesResponseBody
	GetRequestId() *string
}

type DescribeDirectoriesResponseBody struct {
	// The directories.
	Directories []*DescribeDirectoriesResponseBodyDirectories `json:"Directories,omitempty" xml:"Directories,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F369A091-002F-49C8-AD55-02A77629****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDirectoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBody) GetDirectories() []*DescribeDirectoriesResponseBodyDirectories {
	return s.Directories
}

func (s *DescribeDirectoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDirectoriesResponseBody) SetDirectories(v []*DescribeDirectoriesResponseBodyDirectories) *DescribeDirectoriesResponseBody {
	s.Directories = v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetRequestId(v string) *DescribeDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDirectoriesResponseBody) Validate() error {
	if s.Directories != nil {
		for _, item := range s.Directories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDirectoriesResponseBodyDirectories struct {
	// The connection method.
	//
	// Valid values:
	//
	// 	- VPC: End users connect to cloud computers over an enterprise virtual private cloud (VPC).
	//
	// 	- INTERNET: End users connect to cloud computers over the Internet.
	//
	// 	- ANY: End users connect to cloud computers over the Internet or an enterprise VPC.
	//
	// example:
	//
	// INTERNET
	DesktopAccessType *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// cn-hangzhou+dir-gx2x1dhsmu52rd****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The directory type.
	//
	// example:
	//
	// AD_CONNECTOR
	DirectoryType *string `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	// The provider ID.
	//
	// example:
	//
	// 26842
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	// The URL of the SSO service.
	//
	// example:
	//
	// https://eds-cn-shanghai-67726****
	SsoServiceUrl *string `json:"SsoServiceUrl,omitempty" xml:"SsoServiceUrl,omitempty"`
}

func (s DescribeDirectoriesResponseBodyDirectories) String() string {
	return dara.Prettify(s)
}

func (s DescribeDirectoriesResponseBodyDirectories) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetDesktopAccessType() *string {
	return s.DesktopAccessType
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetDirectoryType() *string {
	return s.DirectoryType
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetProviderId() *string {
	return s.ProviderId
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetSsoServiceUrl() *string {
	return s.SsoServiceUrl
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDesktopAccessType(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DesktopAccessType = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDirectoryId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDirectoryType(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DirectoryType = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetProviderId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.ProviderId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetSsoServiceUrl(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.SsoServiceUrl = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) Validate() error {
	return dara.Validate(s)
}
