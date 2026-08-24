// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirusScanScheduledStrategiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVirusScanScheduledStrategiesResponseBody
	GetRequestId() *string
}

type DeleteVirusScanScheduledStrategiesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVirusScanScheduledStrategiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirusScanScheduledStrategiesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVirusScanScheduledStrategiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVirusScanScheduledStrategiesResponseBody) SetRequestId(v string) *DeleteVirusScanScheduledStrategiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVirusScanScheduledStrategiesResponseBody) Validate() error {
	return dara.Validate(s)
}
