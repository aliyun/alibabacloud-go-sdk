// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateConnectionToOtherZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MigrateConnectionToOtherZoneResponseBody
	GetCode() *string
	SetMessage(v string) *MigrateConnectionToOtherZoneResponseBody
	GetMessage() *string
	SetRequestId(v string) *MigrateConnectionToOtherZoneResponseBody
	GetRequestId() *string
}

type MigrateConnectionToOtherZoneResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidParam
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// Invalid Parameter.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 65BDA532-28AF-4122-AA39-B382721EEE64
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MigrateConnectionToOtherZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MigrateConnectionToOtherZoneResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateConnectionToOtherZoneResponseBody) GetCode() *string {
	return s.Code
}

func (s *MigrateConnectionToOtherZoneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MigrateConnectionToOtherZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MigrateConnectionToOtherZoneResponseBody) SetCode(v string) *MigrateConnectionToOtherZoneResponseBody {
	s.Code = &v
	return s
}

func (s *MigrateConnectionToOtherZoneResponseBody) SetMessage(v string) *MigrateConnectionToOtherZoneResponseBody {
	s.Message = &v
	return s
}

func (s *MigrateConnectionToOtherZoneResponseBody) SetRequestId(v string) *MigrateConnectionToOtherZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateConnectionToOtherZoneResponseBody) Validate() error {
	return dara.Validate(s)
}
