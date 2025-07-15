// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppSecurityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *DescribeAppSecurityResponseBody
	GetAppCode() *string
	SetAppKey(v string) *DescribeAppSecurityResponseBody
	GetAppKey() *string
	SetAppSecret(v string) *DescribeAppSecurityResponseBody
	GetAppSecret() *string
	SetCreatedTime(v string) *DescribeAppSecurityResponseBody
	GetCreatedTime() *string
	SetModifiedTime(v string) *DescribeAppSecurityResponseBody
	GetModifiedTime() *string
	SetRequestId(v string) *DescribeAppSecurityResponseBody
	GetRequestId() *string
}

type DescribeAppSecurityResponseBody struct {
	// The AppCode of the app.
	//
	// example:
	//
	// 3aaf905a0a1f4f0eabc6d891dfa08afc
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// The key of the app.
	//
	// example:
	//
	// 60030986
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The password of the app.
	//
	// example:
	//
	// c1ae97aaa7e45f21d10824bc44678fee
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	// The creation time (UTC) of the key, which is the same as the app creation time.
	//
	// example:
	//
	// 2016-07-31T04:10:19Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The modification time (UTC) of the key.
	//
	// example:
	//
	// 2016-07-31T04:10:19Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ015
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAppSecurityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppSecurityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppSecurityResponseBody) GetAppCode() *string {
	return s.AppCode
}

func (s *DescribeAppSecurityResponseBody) GetAppKey() *string {
	return s.AppKey
}

func (s *DescribeAppSecurityResponseBody) GetAppSecret() *string {
	return s.AppSecret
}

func (s *DescribeAppSecurityResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeAppSecurityResponseBody) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeAppSecurityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppSecurityResponseBody) SetAppCode(v string) *DescribeAppSecurityResponseBody {
	s.AppCode = &v
	return s
}

func (s *DescribeAppSecurityResponseBody) SetAppKey(v string) *DescribeAppSecurityResponseBody {
	s.AppKey = &v
	return s
}

func (s *DescribeAppSecurityResponseBody) SetAppSecret(v string) *DescribeAppSecurityResponseBody {
	s.AppSecret = &v
	return s
}

func (s *DescribeAppSecurityResponseBody) SetCreatedTime(v string) *DescribeAppSecurityResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeAppSecurityResponseBody) SetModifiedTime(v string) *DescribeAppSecurityResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeAppSecurityResponseBody) SetRequestId(v string) *DescribeAppSecurityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppSecurityResponseBody) Validate() error {
	return dara.Validate(s)
}
