// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateAvailableZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MigrateAvailableZoneResponseBody
	GetRequestId() *string
}

type MigrateAvailableZoneResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0FDDC511-7252-4A4A-ADDA-5CB1BF63****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MigrateAvailableZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MigrateAvailableZoneResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateAvailableZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MigrateAvailableZoneResponseBody) SetRequestId(v string) *MigrateAvailableZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateAvailableZoneResponseBody) Validate() error {
	return dara.Validate(s)
}
