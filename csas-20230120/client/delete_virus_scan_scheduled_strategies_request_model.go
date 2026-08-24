// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirusScanScheduledStrategiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStrategyIds(v []*string) *DeleteVirusScanScheduledStrategiesRequest
	GetStrategyIds() []*string
}

type DeleteVirusScanScheduledStrategiesRequest struct {
	// The IDs of the virus scheduled scan policies to delete. The collection must contain at least 1 and at most 100 IDs. Duplicate IDs are not allowed.
	//
	// This parameter is required.
	StrategyIds []*string `json:"StrategyIds,omitempty" xml:"StrategyIds,omitempty" type:"Repeated"`
}

func (s DeleteVirusScanScheduledStrategiesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirusScanScheduledStrategiesRequest) GoString() string {
	return s.String()
}

func (s *DeleteVirusScanScheduledStrategiesRequest) GetStrategyIds() []*string {
	return s.StrategyIds
}

func (s *DeleteVirusScanScheduledStrategiesRequest) SetStrategyIds(v []*string) *DeleteVirusScanScheduledStrategiesRequest {
	s.StrategyIds = v
	return s
}

func (s *DeleteVirusScanScheduledStrategiesRequest) Validate() error {
	return dara.Validate(s)
}
