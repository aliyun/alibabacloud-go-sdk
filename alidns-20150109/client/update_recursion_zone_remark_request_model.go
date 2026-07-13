// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionZoneRemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateRecursionZoneRemarkRequest
	GetClientToken() *string
	SetRemark(v string) *UpdateRecursionZoneRemarkRequest
	GetRemark() *string
	SetZoneId(v string) *UpdateRecursionZoneRemarkRequest
	GetZoneId() *string
}

type UpdateRecursionZoneRemarkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You must generate a unique value for this parameter. The client token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 21079fa016944979537637959d09bc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The remarks.
	//
	// example:
	//
	// 备注
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The zone ID. This is the unique identifier of the zone.
	//
	// example:
	//
	// 173671468000011
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateRecursionZoneRemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionZoneRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecursionZoneRemarkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRecursionZoneRemarkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateRecursionZoneRemarkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpdateRecursionZoneRemarkRequest) SetClientToken(v string) *UpdateRecursionZoneRemarkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRecursionZoneRemarkRequest) SetRemark(v string) *UpdateRecursionZoneRemarkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateRecursionZoneRemarkRequest) SetZoneId(v string) *UpdateRecursionZoneRemarkRequest {
	s.ZoneId = &v
	return s
}

func (s *UpdateRecursionZoneRemarkRequest) Validate() error {
	return dara.Validate(s)
}
