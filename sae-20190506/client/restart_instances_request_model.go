// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RestartInstancesRequest
	GetAppId() *string
	SetInstanceIds(v string) *RestartInstancesRequest
	GetInstanceIds() *string
}

type RestartInstancesRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1daa7236-3844-4f36-b39a-605b0cc0****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the instance to be restarted. Separate multiple instance IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// mysae-1daa7236-3844-4f36-b39a-605b0cc0caa6-*****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s RestartInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartInstancesRequest) GoString() string {
	return s.String()
}

func (s *RestartInstancesRequest) GetAppId() *string {
	return s.AppId
}

func (s *RestartInstancesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *RestartInstancesRequest) SetAppId(v string) *RestartInstancesRequest {
	s.AppId = &v
	return s
}

func (s *RestartInstancesRequest) SetInstanceIds(v string) *RestartInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *RestartInstancesRequest) Validate() error {
	return dara.Validate(s)
}
