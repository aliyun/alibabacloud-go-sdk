// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterPrimaryZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterPrimaryZoneResponseBody
	GetRequestId() *string
}

type ModifyDBClusterPrimaryZoneResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ED12C6FF-3261-4571-AB57-3570F6******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterPrimaryZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterPrimaryZoneResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterPrimaryZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterPrimaryZoneResponseBody) SetRequestId(v string) *ModifyDBClusterPrimaryZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneResponseBody) Validate() error {
	return dara.Validate(s)
}
