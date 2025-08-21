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
	SetResult(v bool) *MigrateToOtherZoneResponseBody
	GetResult() *bool
}

type MigrateToOtherZoneResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return results:
	//
	// 	- true: migration succeeded
	//
	// 	- false: The migration fails
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
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

func (s *MigrateToOtherZoneResponseBody) GetResult() *bool {
	return s.Result
}

func (s *MigrateToOtherZoneResponseBody) SetRequestId(v string) *MigrateToOtherZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateToOtherZoneResponseBody) SetResult(v bool) *MigrateToOtherZoneResponseBody {
	s.Result = &v
	return s
}

func (s *MigrateToOtherZoneResponseBody) Validate() error {
	return dara.Validate(s)
}
