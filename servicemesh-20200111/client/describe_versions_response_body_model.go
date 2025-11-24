// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVersionsResponseBody
	GetRequestId() *string
	SetVersionInfo(v []*DescribeVersionsResponseBodyVersionInfo) *DescribeVersionsResponseBody
	GetVersionInfo() []*DescribeVersionsResponseBodyVersionInfo
}

type DescribeVersionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The available ASM versions.
	VersionInfo []*DescribeVersionsResponseBodyVersionInfo `json:"VersionInfo,omitempty" xml:"VersionInfo,omitempty" type:"Repeated"`
}

func (s DescribeVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVersionsResponseBody) GetVersionInfo() []*DescribeVersionsResponseBodyVersionInfo {
	return s.VersionInfo
}

func (s *DescribeVersionsResponseBody) SetRequestId(v string) *DescribeVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVersionsResponseBody) SetVersionInfo(v []*DescribeVersionsResponseBodyVersionInfo) *DescribeVersionsResponseBody {
	s.VersionInfo = v
	return s
}

func (s *DescribeVersionsResponseBody) Validate() error {
	if s.VersionInfo != nil {
		for _, item := range s.VersionInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVersionsResponseBodyVersionInfo struct {
	// The edition of the ASM instance. Valid values:
	//
	// 	- `Default`: Standard Edition
	//
	// 	- `Pro`: Professional Edition that is commercially released
	//
	// example:
	//
	// Default
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The list of ASM versions available for the ASM instance of the current edition.
	Versions []*string `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s DescribeVersionsResponseBodyVersionInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeVersionsResponseBodyVersionInfo) GoString() string {
	return s.String()
}

func (s *DescribeVersionsResponseBodyVersionInfo) GetEdition() *string {
	return s.Edition
}

func (s *DescribeVersionsResponseBodyVersionInfo) GetVersions() []*string {
	return s.Versions
}

func (s *DescribeVersionsResponseBodyVersionInfo) SetEdition(v string) *DescribeVersionsResponseBodyVersionInfo {
	s.Edition = &v
	return s
}

func (s *DescribeVersionsResponseBodyVersionInfo) SetVersions(v []*string) *DescribeVersionsResponseBodyVersionInfo {
	s.Versions = v
	return s
}

func (s *DescribeVersionsResponseBodyVersionInfo) Validate() error {
	return dara.Validate(s)
}
