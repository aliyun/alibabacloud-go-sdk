// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPotentialFailZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsPlanId(v bool) *GetPotentialFailZonesRequest
	GetIsPlanId() *bool
	SetObjectId(v string) *GetPotentialFailZonesRequest
	GetObjectId() *string
}

type GetPotentialFailZonesRequest struct {
	// Specifies whether the value of this parameter is the ID of a disaster recovery set.
	//
	// example:
	//
	// true
	IsPlanId *bool `json:"IsPlanId,omitempty" xml:"IsPlanId,omitempty"`
	// If you set IsPlanId to false, specify the ID of a disaster recovery application. If you set IsPlanId to true, specify the ID of a disaster recovery set.
	//
	// example:
	//
	// FS3ATPTOSC4SE1GG
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
}

func (s GetPotentialFailZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPotentialFailZonesRequest) GoString() string {
	return s.String()
}

func (s *GetPotentialFailZonesRequest) GetIsPlanId() *bool {
	return s.IsPlanId
}

func (s *GetPotentialFailZonesRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *GetPotentialFailZonesRequest) SetIsPlanId(v bool) *GetPotentialFailZonesRequest {
	s.IsPlanId = &v
	return s
}

func (s *GetPotentialFailZonesRequest) SetObjectId(v string) *GetPotentialFailZonesRequest {
	s.ObjectId = &v
	return s
}

func (s *GetPotentialFailZonesRequest) Validate() error {
	return dara.Validate(s)
}
