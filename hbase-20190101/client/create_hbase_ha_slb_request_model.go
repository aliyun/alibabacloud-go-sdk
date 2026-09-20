// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHbaseHaSlbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBdsId(v string) *CreateHbaseHaSlbRequest
	GetBdsId() *string
	SetClientToken(v string) *CreateHbaseHaSlbRequest
	GetClientToken() *string
	SetHaId(v string) *CreateHbaseHaSlbRequest
	GetHaId() *string
	SetHaTypes(v string) *CreateHbaseHaSlbRequest
	GetHaTypes() *string
	SetHbaseType(v string) *CreateHbaseHaSlbRequest
	GetHbaseType() *string
}

type CreateHbaseHaSlbRequest struct {
	// The ID of the BDS cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// bds-t4n3496whj23****
	BdsId *string `json:"BdsId,omitempty" xml:"BdsId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The value cannot exceed 64 printable ASCII characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The high-availability ID in the BDS active-active management.
	//
	// This parameter is required.
	//
	// example:
	//
	// ha-v21tmnxjwh2yu****
	HaId *string `json:"HaId,omitempty" xml:"HaId,omitempty"`
	// The high-availability type. Valid values:
	//
	// - thrift
	//
	// - phoenix.
	//
	// This parameter is required.
	//
	// example:
	//
	// thrift
	HaTypes *string `json:"HaTypes,omitempty" xml:"HaTypes,omitempty"`
	// Specifies whether the high-availability type is on the primary or secondary instance. Valid values:
	//
	// - Active: The high-availability type is on the primary instance.
	//
	// - Standby: The high-availability type is on the secondary instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// Active
	HbaseType *string `json:"HbaseType,omitempty" xml:"HbaseType,omitempty"`
}

func (s CreateHbaseHaSlbRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHbaseHaSlbRequest) GoString() string {
	return s.String()
}

func (s *CreateHbaseHaSlbRequest) GetBdsId() *string {
	return s.BdsId
}

func (s *CreateHbaseHaSlbRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateHbaseHaSlbRequest) GetHaId() *string {
	return s.HaId
}

func (s *CreateHbaseHaSlbRequest) GetHaTypes() *string {
	return s.HaTypes
}

func (s *CreateHbaseHaSlbRequest) GetHbaseType() *string {
	return s.HbaseType
}

func (s *CreateHbaseHaSlbRequest) SetBdsId(v string) *CreateHbaseHaSlbRequest {
	s.BdsId = &v
	return s
}

func (s *CreateHbaseHaSlbRequest) SetClientToken(v string) *CreateHbaseHaSlbRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateHbaseHaSlbRequest) SetHaId(v string) *CreateHbaseHaSlbRequest {
	s.HaId = &v
	return s
}

func (s *CreateHbaseHaSlbRequest) SetHaTypes(v string) *CreateHbaseHaSlbRequest {
	s.HaTypes = &v
	return s
}

func (s *CreateHbaseHaSlbRequest) SetHbaseType(v string) *CreateHbaseHaSlbRequest {
	s.HbaseType = &v
	return s
}

func (s *CreateHbaseHaSlbRequest) Validate() error {
	return dara.Validate(s)
}
