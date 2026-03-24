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
	// The request ID.
	//
	// example:
	//
	// 9178CB86-285F-5679-A30A-3B3F007E4206
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the available HTTP and HTTPS port ranges for hybrid cloud access.
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
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHybridCloudUserResponseBodyUserInfo struct {
	// The available HTTP ports. The value is a string. If multiple ports are returned, they are separated by commas (,). Example: **port1,port2,port3**.
	//
	// example:
	//
	// 80,8080
	HttpPorts *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	// The available HTTPS ports. The value is a string. If multiple ports are returned, they are separated by commas (,). Example: **port1,port2,port3**.
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
