// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAndroidInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceId(v string) *ModifyAndroidInstanceRequest
	GetAndroidInstanceId() *string
	SetDownBandwidthLimit(v int32) *ModifyAndroidInstanceRequest
	GetDownBandwidthLimit() *int32
	SetInstanceIds(v []*string) *ModifyAndroidInstanceRequest
	GetInstanceIds() []*string
	SetNewAndroidInstanceName(v string) *ModifyAndroidInstanceRequest
	GetNewAndroidInstanceName() *string
	SetUpBandwidthLimit(v int32) *ModifyAndroidInstanceRequest
	GetUpBandwidthLimit() *int32
}

type ModifyAndroidInstanceRequest struct {
	// The ID of the cloud phone instance.
	//
	// example:
	//
	// acp-8v5bjld0r7tkl****
	AndroidInstanceId  *string   `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	DownBandwidthLimit *int32    `json:"DownBandwidthLimit,omitempty" xml:"DownBandwidthLimit,omitempty"`
	InstanceIds        []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The new name of the cloud phone instance.
	//
	// >  The name can be up to 30 characters in length. It can contain letters, digits, colons (:), underscores (_), periods (.), or hyphens (-). It must start with letters but cannot start with http:// or https://.
	//
	// example:
	//
	// new_name
	NewAndroidInstanceName *string `json:"NewAndroidInstanceName,omitempty" xml:"NewAndroidInstanceName,omitempty"`
	UpBandwidthLimit       *int32  `json:"UpBandwidthLimit,omitempty" xml:"UpBandwidthLimit,omitempty"`
}

func (s ModifyAndroidInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAndroidInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyAndroidInstanceRequest) GetAndroidInstanceId() *string {
	return s.AndroidInstanceId
}

func (s *ModifyAndroidInstanceRequest) GetDownBandwidthLimit() *int32 {
	return s.DownBandwidthLimit
}

func (s *ModifyAndroidInstanceRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ModifyAndroidInstanceRequest) GetNewAndroidInstanceName() *string {
	return s.NewAndroidInstanceName
}

func (s *ModifyAndroidInstanceRequest) GetUpBandwidthLimit() *int32 {
	return s.UpBandwidthLimit
}

func (s *ModifyAndroidInstanceRequest) SetAndroidInstanceId(v string) *ModifyAndroidInstanceRequest {
	s.AndroidInstanceId = &v
	return s
}

func (s *ModifyAndroidInstanceRequest) SetDownBandwidthLimit(v int32) *ModifyAndroidInstanceRequest {
	s.DownBandwidthLimit = &v
	return s
}

func (s *ModifyAndroidInstanceRequest) SetInstanceIds(v []*string) *ModifyAndroidInstanceRequest {
	s.InstanceIds = v
	return s
}

func (s *ModifyAndroidInstanceRequest) SetNewAndroidInstanceName(v string) *ModifyAndroidInstanceRequest {
	s.NewAndroidInstanceName = &v
	return s
}

func (s *ModifyAndroidInstanceRequest) SetUpBandwidthLimit(v int32) *ModifyAndroidInstanceRequest {
	s.UpBandwidthLimit = &v
	return s
}

func (s *ModifyAndroidInstanceRequest) Validate() error {
	return dara.Validate(s)
}
