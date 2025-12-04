// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDestinationCidrBlockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetDestinationCidrBlockRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetDestinationCidrBlockRequest
	GetRegionId() *string
}

type GetDestinationCidrBlockRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-xxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Region ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetDestinationCidrBlockRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDestinationCidrBlockRequest) GoString() string {
	return s.String()
}

func (s *GetDestinationCidrBlockRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDestinationCidrBlockRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDestinationCidrBlockRequest) SetInstanceId(v string) *GetDestinationCidrBlockRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDestinationCidrBlockRequest) SetRegionId(v string) *GetDestinationCidrBlockRequest {
	s.RegionId = &v
	return s
}

func (s *GetDestinationCidrBlockRequest) Validate() error {
	return dara.Validate(s)
}
