// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLaunchServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *LaunchServiceResponseBody
	GetRequestId() *string
	SetServiceLaunchResultType(v string) *LaunchServiceResponseBody
	GetServiceLaunchResultType() *string
	SetVersion(v string) *LaunchServiceResponseBody
	GetVersion() *string
}

type LaunchServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The mode of the service online. Valid Type
	//
	// - PublishNewVersion: Launch new version
	//
	// - PublishOfflineVersion:  The offline version is online again.
	//
	// - UpdateLatestVersion: Update the latest version online
	//
	// example:
	//
	// PublishNewVersion
	ServiceLaunchResultType *string `json:"ServiceLaunchResultType,omitempty" xml:"ServiceLaunchResultType,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s LaunchServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LaunchServiceResponseBody) GoString() string {
	return s.String()
}

func (s *LaunchServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LaunchServiceResponseBody) GetServiceLaunchResultType() *string {
	return s.ServiceLaunchResultType
}

func (s *LaunchServiceResponseBody) GetVersion() *string {
	return s.Version
}

func (s *LaunchServiceResponseBody) SetRequestId(v string) *LaunchServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *LaunchServiceResponseBody) SetServiceLaunchResultType(v string) *LaunchServiceResponseBody {
	s.ServiceLaunchResultType = &v
	return s
}

func (s *LaunchServiceResponseBody) SetVersion(v string) *LaunchServiceResponseBody {
	s.Version = &v
	return s
}

func (s *LaunchServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
