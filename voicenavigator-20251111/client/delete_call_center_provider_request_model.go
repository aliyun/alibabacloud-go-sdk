// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCallCenterProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteCallCenterProviderRequest
	GetInstanceId() *string
	SetProviderId(v string) *DeleteCallCenterProviderRequest
	GetProviderId() *string
}

type DeleteCallCenterProviderRequest struct {
	// example:
	//
	// 36e9a4cd-a821-4f78-86fa-a9a4aefeea2e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// xxxxxxxxx
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
}

func (s DeleteCallCenterProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCallCenterProviderRequest) GoString() string {
	return s.String()
}

func (s *DeleteCallCenterProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteCallCenterProviderRequest) GetProviderId() *string {
	return s.ProviderId
}

func (s *DeleteCallCenterProviderRequest) SetInstanceId(v string) *DeleteCallCenterProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteCallCenterProviderRequest) SetProviderId(v string) *DeleteCallCenterProviderRequest {
	s.ProviderId = &v
	return s
}

func (s *DeleteCallCenterProviderRequest) Validate() error {
	return dara.Validate(s)
}
