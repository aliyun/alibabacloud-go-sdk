// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisconnectAndroidInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndUserId(v string) *DisconnectAndroidInstanceRequest
	GetEndUserId() *string
	SetInstanceIds(v []*string) *DisconnectAndroidInstanceRequest
	GetInstanceIds() []*string
}

type DisconnectAndroidInstanceRequest struct {
	// <props="china">
	//
	// If you use the Cloud Phone Matrix Edition and the instance stream pattern is collaborative mode, you can specify `EndUserId` to disconnect a specific user and invalidate the corresponding ticket.
	//
	//
	//
	// <props="intl">
	//
	// This parameter is not publicly available.
	//
	// example:
	//
	// user1
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// A list of instance IDs. You can specify 1 to 100 IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DisconnectAndroidInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DisconnectAndroidInstanceRequest) GoString() string {
	return s.String()
}

func (s *DisconnectAndroidInstanceRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DisconnectAndroidInstanceRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DisconnectAndroidInstanceRequest) SetEndUserId(v string) *DisconnectAndroidInstanceRequest {
	s.EndUserId = &v
	return s
}

func (s *DisconnectAndroidInstanceRequest) SetInstanceIds(v []*string) *DisconnectAndroidInstanceRequest {
	s.InstanceIds = v
	return s
}

func (s *DisconnectAndroidInstanceRequest) Validate() error {
	return dara.Validate(s)
}
