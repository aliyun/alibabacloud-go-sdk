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
	EndUserId   *string   `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
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
