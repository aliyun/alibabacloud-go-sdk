// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLocalitySettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeLocalitySettingRequest
	GetAppId() *string
	SetNamespaceId(v string) *DescribeLocalitySettingRequest
	GetNamespaceId() *string
	SetRegion(v string) *DescribeLocalitySettingRequest
	GetRegion() *string
}

type DescribeLocalitySettingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dc190221-22b5-491c-a548-82f5fa1e3e26
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeLocalitySettingRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLocalitySettingRequest) GoString() string {
	return s.String()
}

func (s *DescribeLocalitySettingRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeLocalitySettingRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeLocalitySettingRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeLocalitySettingRequest) SetAppId(v string) *DescribeLocalitySettingRequest {
	s.AppId = &v
	return s
}

func (s *DescribeLocalitySettingRequest) SetNamespaceId(v string) *DescribeLocalitySettingRequest {
	s.NamespaceId = &v
	return s
}

func (s *DescribeLocalitySettingRequest) SetRegion(v string) *DescribeLocalitySettingRequest {
	s.Region = &v
	return s
}

func (s *DescribeLocalitySettingRequest) Validate() error {
	return dara.Validate(s)
}
