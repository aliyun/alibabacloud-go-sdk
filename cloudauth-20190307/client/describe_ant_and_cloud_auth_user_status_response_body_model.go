// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAntAndCloudAuthUserStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAntcloudauthEnabled(v bool) *DescribeAntAndCloudAuthUserStatusResponseBody
	GetAntcloudauthEnabled() *bool
	SetCloudauthEnabled(v bool) *DescribeAntAndCloudAuthUserStatusResponseBody
	GetCloudauthEnabled() *bool
	SetCloudauthstEnabled(v bool) *DescribeAntAndCloudAuthUserStatusResponseBody
	GetCloudauthstEnabled() *bool
	SetInforverifyEnabled(v bool) *DescribeAntAndCloudAuthUserStatusResponseBody
	GetInforverifyEnabled() *bool
	SetRequestId(v string) *DescribeAntAndCloudAuthUserStatusResponseBody
	GetRequestId() *string
}

type DescribeAntAndCloudAuthUserStatusResponseBody struct {
	// Indicates whether financial-grade real-person authentication is activated. Values:
	//
	// - **true**: Activated
	//
	// - **false**: Not activated
	//
	// example:
	//
	// true
	AntcloudauthEnabled *bool `json:"AntcloudauthEnabled,omitempty" xml:"AntcloudauthEnabled,omitempty"`
	// Indicates whether real-person authentication is activated. Values:
	//
	// - **true**: Activated
	//
	// - **false**: Not activated
	//
	// example:
	//
	// false
	CloudauthEnabled *bool `json:"CloudauthEnabled,omitempty" xml:"CloudauthEnabled,omitempty"`
	// Indicates whether the enhanced version of real-person authentication is activated. Values:
	//
	// - **true**: Activated
	//
	// - **false**: Not activated
	//
	// example:
	//
	// true
	CloudauthstEnabled *bool `json:"CloudauthstEnabled,omitempty" xml:"CloudauthstEnabled,omitempty"`
	// Indicates whether information verification is activated. Values:
	//
	// - **true**: Activated
	//
	// - **false**: Not activated
	//
	// example:
	//
	// false
	InforverifyEnabled *bool `json:"InforverifyEnabled,omitempty" xml:"InforverifyEnabled,omitempty"`
	// The ID of this request.
	//
	// example:
	//
	// 3FE07CCE-DF47-51C2-9D32-CD70ED62C91B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAntAndCloudAuthUserStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAntAndCloudAuthUserStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntAndCloudAuthUserStatusResponseBody) GetAntcloudauthEnabled() *bool {
	return s.AntcloudauthEnabled
}

func (s *DescribeAntAndCloudAuthUserStatusResponseBody) GetCloudauthEnabled() *bool {
	return s.CloudauthEnabled
}

func (s *DescribeAntAndCloudAuthUserStatusResponseBody) GetCloudauthstEnabled() *bool {
	return s.CloudauthstEnabled
}

func (s *DescribeAntAndCloudAuthUserStatusResponseBody) GetInforverifyEnabled() *bool {
	return s.InforverifyEnabled
}

func (s *DescribeAntAndCloudAuthUserStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAntAndCloudAuthUserStatusResponseBody) SetAntcloudauthEnabled(v bool) *DescribeAntAndCloudAuthUserStatusResponseBody {
	s.AntcloudauthEnabled = &v
	return s
}

func (s *DescribeAntAndCloudAuthUserStatusResponseBody) SetCloudauthEnabled(v bool) *DescribeAntAndCloudAuthUserStatusResponseBody {
	s.CloudauthEnabled = &v
	return s
}

func (s *DescribeAntAndCloudAuthUserStatusResponseBody) SetCloudauthstEnabled(v bool) *DescribeAntAndCloudAuthUserStatusResponseBody {
	s.CloudauthstEnabled = &v
	return s
}

func (s *DescribeAntAndCloudAuthUserStatusResponseBody) SetInforverifyEnabled(v bool) *DescribeAntAndCloudAuthUserStatusResponseBody {
	s.InforverifyEnabled = &v
	return s
}

func (s *DescribeAntAndCloudAuthUserStatusResponseBody) SetRequestId(v string) *DescribeAntAndCloudAuthUserStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntAndCloudAuthUserStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
