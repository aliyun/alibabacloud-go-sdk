// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateToOtherZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MigrateToOtherZoneResponseBody
	GetRequestId() *string
}

type MigrateToOtherZoneResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0FDDC511-7252-4A4A-ADDA-5CB1BF63873D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MigrateToOtherZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MigrateToOtherZoneResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateToOtherZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MigrateToOtherZoneResponseBody) SetRequestId(v string) *MigrateToOtherZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateToOtherZoneResponseBody) Validate() error {
	return dara.Validate(s)
}
