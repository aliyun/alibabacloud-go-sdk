// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartScalingGroupZoneShiftRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartScalingGroupZoneShiftRequest
	GetClientToken() *string
	SetExpiryTime(v string) *StartScalingGroupZoneShiftRequest
	GetExpiryTime() *string
	SetImpairedZoneIds(v []*string) *StartScalingGroupZoneShiftRequest
	GetImpairedZoneIds() []*string
	SetRegionId(v string) *StartScalingGroupZoneShiftRequest
	GetRegionId() *string
	SetScalingGroupId(v string) *StartScalingGroupZoneShiftRequest
	GetScalingGroupId() *string
}

type StartScalingGroupZoneShiftRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see Ensure idempotence.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The expiration time of the zone transfer. After the time expires, the system automatically ends the transfer and restores the related zones. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mmZ format. The time must be in UTC. The value must be later than the current time and cannot be more than 3 days after the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-01-01T12:05Z
	ExpiryTime *string `json:"ExpiryTime,omitempty" xml:"ExpiryTime,omitempty"`
	// The IDs of the impaired zones. Valid values of N: 1 to 5.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-g
	ImpairedZoneIds []*string `json:"ImpairedZoneIds,omitempty" xml:"ImpairedZoneIds,omitempty" type:"Repeated"`
	// The region ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp14wlu85wrpchm0****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s StartScalingGroupZoneShiftRequest) String() string {
	return dara.Prettify(s)
}

func (s StartScalingGroupZoneShiftRequest) GoString() string {
	return s.String()
}

func (s *StartScalingGroupZoneShiftRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartScalingGroupZoneShiftRequest) GetExpiryTime() *string {
	return s.ExpiryTime
}

func (s *StartScalingGroupZoneShiftRequest) GetImpairedZoneIds() []*string {
	return s.ImpairedZoneIds
}

func (s *StartScalingGroupZoneShiftRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartScalingGroupZoneShiftRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *StartScalingGroupZoneShiftRequest) SetClientToken(v string) *StartScalingGroupZoneShiftRequest {
	s.ClientToken = &v
	return s
}

func (s *StartScalingGroupZoneShiftRequest) SetExpiryTime(v string) *StartScalingGroupZoneShiftRequest {
	s.ExpiryTime = &v
	return s
}

func (s *StartScalingGroupZoneShiftRequest) SetImpairedZoneIds(v []*string) *StartScalingGroupZoneShiftRequest {
	s.ImpairedZoneIds = v
	return s
}

func (s *StartScalingGroupZoneShiftRequest) SetRegionId(v string) *StartScalingGroupZoneShiftRequest {
	s.RegionId = &v
	return s
}

func (s *StartScalingGroupZoneShiftRequest) SetScalingGroupId(v string) *StartScalingGroupZoneShiftRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *StartScalingGroupZoneShiftRequest) Validate() error {
	return dara.Validate(s)
}
