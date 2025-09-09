// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionRecordWeightEnableStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateRecursionRecordWeightEnableStatusRequest
	GetClientToken() *string
	SetEnableStatus(v string) *UpdateRecursionRecordWeightEnableStatusRequest
	GetEnableStatus() *string
	SetRequestSource(v string) *UpdateRecursionRecordWeightEnableStatusRequest
	GetRequestSource() *string
	SetRr(v string) *UpdateRecursionRecordWeightEnableStatusRequest
	GetRr() *string
	SetType(v string) *UpdateRecursionRecordWeightEnableStatusRequest
	GetType() *string
	SetZoneId(v string) *UpdateRecursionRecordWeightEnableStatusRequest
	GetZoneId() *string
}

type UpdateRecursionRecordWeightEnableStatusRequest struct {
	// example:
	//
	// 21079fa016944979537637959d09bc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// example:
	//
	// default
	RequestSource *string `json:"RequestSource,omitempty" xml:"RequestSource,omitempty"`
	// example:
	//
	// www
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 176432424234
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateRecursionRecordWeightEnableStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionRecordWeightEnableStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) GetRequestSource() *string {
	return s.RequestSource
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) GetRr() *string {
	return s.Rr
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) GetType() *string {
	return s.Type
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) SetClientToken(v string) *UpdateRecursionRecordWeightEnableStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) SetEnableStatus(v string) *UpdateRecursionRecordWeightEnableStatusRequest {
	s.EnableStatus = &v
	return s
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) SetRequestSource(v string) *UpdateRecursionRecordWeightEnableStatusRequest {
	s.RequestSource = &v
	return s
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) SetRr(v string) *UpdateRecursionRecordWeightEnableStatusRequest {
	s.Rr = &v
	return s
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) SetType(v string) *UpdateRecursionRecordWeightEnableStatusRequest {
	s.Type = &v
	return s
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) SetZoneId(v string) *UpdateRecursionRecordWeightEnableStatusRequest {
	s.ZoneId = &v
	return s
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) Validate() error {
	return dara.Validate(s)
}
