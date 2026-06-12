// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelServiceUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CancelServiceUsageRequest
	GetClientToken() *string
	SetNeedDelete(v bool) *CancelServiceUsageRequest
	GetNeedDelete() *bool
	SetRegionId(v string) *CancelServiceUsageRequest
	GetRegionId() *string
	SetServiceId(v string) *CancelServiceUsageRequest
	GetServiceId() *string
}

type CancelServiceUsageRequest struct {
	// A client-generated token that ensures the idempotence of the request. The token must be unique for each request. It can contain only ASCII characters and must be no more than 64 characters long.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to delete the request.
	//
	// > If you delete the request, you must submit a new one.
	//
	// example:
	//
	// true
	NeedDelete *bool `json:"NeedDelete,omitempty" xml:"NeedDelete,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-d6fc5f949a9246xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s CancelServiceUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelServiceUsageRequest) GoString() string {
	return s.String()
}

func (s *CancelServiceUsageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CancelServiceUsageRequest) GetNeedDelete() *bool {
	return s.NeedDelete
}

func (s *CancelServiceUsageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelServiceUsageRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *CancelServiceUsageRequest) SetClientToken(v string) *CancelServiceUsageRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelServiceUsageRequest) SetNeedDelete(v bool) *CancelServiceUsageRequest {
	s.NeedDelete = &v
	return s
}

func (s *CancelServiceUsageRequest) SetRegionId(v string) *CancelServiceUsageRequest {
	s.RegionId = &v
	return s
}

func (s *CancelServiceUsageRequest) SetServiceId(v string) *CancelServiceUsageRequest {
	s.ServiceId = &v
	return s
}

func (s *CancelServiceUsageRequest) Validate() error {
	return dara.Validate(s)
}
