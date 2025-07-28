// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeHybridCloudUserResponseBody
	GetRequestId() *string
	SetUserInfo(v *DescribeHybridCloudUserResponseBodyUserInfo) *DescribeHybridCloudUserResponseBody
	GetUserInfo() *DescribeHybridCloudUserResponseBodyUserInfo
}

type DescribeHybridCloudUserResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9178CB86-285F-5679-A30A-3B3F007E4206
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the ports that can be used by a hybrid cloud cluster.
	UserInfo *DescribeHybridCloudUserResponseBodyUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s DescribeHybridCloudUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudUserResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridCloudUserResponseBody) GetUserInfo() *DescribeHybridCloudUserResponseBodyUserInfo {
	return s.UserInfo
}

func (s *DescribeHybridCloudUserResponseBody) SetRequestId(v string) *DescribeHybridCloudUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridCloudUserResponseBody) SetUserInfo(v *DescribeHybridCloudUserResponseBodyUserInfo) *DescribeHybridCloudUserResponseBody {
	s.UserInfo = v
	return s
}

func (s *DescribeHybridCloudUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridCloudUserResponseBodyUserInfo struct {
	// The HTTP ports. The value is a string. If multiple ports are returned, the value is in the **port1,port2,port3*	- format.
	//
	// example:
	//
	// 80,8080
	HttpPorts *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	// The HTTPS ports. The value is a string. If multiple ports are returned, the value is in the **port1,port2,port3*	- format.
	//
	// example:
	//
	// 8443,443
	HttpsPorts *string `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
}

func (s DescribeHybridCloudUserResponseBodyUserInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudUserResponseBodyUserInfo) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudUserResponseBodyUserInfo) GetHttpPorts() *string {
	return s.HttpPorts
}

func (s *DescribeHybridCloudUserResponseBodyUserInfo) GetHttpsPorts() *string {
	return s.HttpsPorts
}

func (s *DescribeHybridCloudUserResponseBodyUserInfo) SetHttpPorts(v string) *DescribeHybridCloudUserResponseBodyUserInfo {
	s.HttpPorts = &v
	return s
}

func (s *DescribeHybridCloudUserResponseBodyUserInfo) SetHttpsPorts(v string) *DescribeHybridCloudUserResponseBodyUserInfo {
	s.HttpsPorts = &v
	return s
}

func (s *DescribeHybridCloudUserResponseBodyUserInfo) Validate() error {
	return dara.Validate(s)
}
