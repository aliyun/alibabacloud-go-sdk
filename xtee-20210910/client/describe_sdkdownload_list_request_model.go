// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSDKDownloadListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSDKDownloadListRequest
	GetLang() *string
	SetDeviceType(v string) *DescribeSDKDownloadListRequest
	GetDeviceType() *string
	SetListType(v string) *DescribeSDKDownloadListRequest
	GetListType() *string
	SetRegId(v string) *DescribeSDKDownloadListRequest
	GetRegId() *string
}

type DescribeSDKDownloadListRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Device type.
	//
	// example:
	//
	// ANDROID
	DeviceType *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	// Download type
	//
	// example:
	//
	// OLD
	ListType *string `json:"listType,omitempty" xml:"listType,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeSDKDownloadListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDKDownloadListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSDKDownloadListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSDKDownloadListRequest) GetDeviceType() *string {
	return s.DeviceType
}

func (s *DescribeSDKDownloadListRequest) GetListType() *string {
	return s.ListType
}

func (s *DescribeSDKDownloadListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSDKDownloadListRequest) SetLang(v string) *DescribeSDKDownloadListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSDKDownloadListRequest) SetDeviceType(v string) *DescribeSDKDownloadListRequest {
	s.DeviceType = &v
	return s
}

func (s *DescribeSDKDownloadListRequest) SetListType(v string) *DescribeSDKDownloadListRequest {
	s.ListType = &v
	return s
}

func (s *DescribeSDKDownloadListRequest) SetRegId(v string) *DescribeSDKDownloadListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSDKDownloadListRequest) Validate() error {
	return dara.Validate(s)
}
