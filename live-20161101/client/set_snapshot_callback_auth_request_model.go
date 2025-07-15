// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSnapshotCallbackAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackAuthKey(v string) *SetSnapshotCallbackAuthRequest
	GetCallbackAuthKey() *string
	SetCallbackReqAuth(v string) *SetSnapshotCallbackAuthRequest
	GetCallbackReqAuth() *string
	SetDomainName(v string) *SetSnapshotCallbackAuthRequest
	GetDomainName() *string
	SetOwnerId(v int64) *SetSnapshotCallbackAuthRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetSnapshotCallbackAuthRequest
	GetRegionId() *string
}

type SetSnapshotCallbackAuthRequest struct {
	// The custom key that is used for callback authentication.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourkey
	CallbackAuthKey *string `json:"CallbackAuthKey,omitempty" xml:"CallbackAuthKey,omitempty"`
	// Specifies whether to enable callback authentication. Valid values:
	//
	// 	- **yes**: enables callback authentication.
	//
	// 	- **no**: disables callback authentication.
	//
	// This parameter is required.
	//
	// example:
	//
	// yes
	CallbackReqAuth *string `json:"CallbackReqAuth,omitempty" xml:"CallbackReqAuth,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetSnapshotCallbackAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSnapshotCallbackAuthRequest) GoString() string {
	return s.String()
}

func (s *SetSnapshotCallbackAuthRequest) GetCallbackAuthKey() *string {
	return s.CallbackAuthKey
}

func (s *SetSnapshotCallbackAuthRequest) GetCallbackReqAuth() *string {
	return s.CallbackReqAuth
}

func (s *SetSnapshotCallbackAuthRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetSnapshotCallbackAuthRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetSnapshotCallbackAuthRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetSnapshotCallbackAuthRequest) SetCallbackAuthKey(v string) *SetSnapshotCallbackAuthRequest {
	s.CallbackAuthKey = &v
	return s
}

func (s *SetSnapshotCallbackAuthRequest) SetCallbackReqAuth(v string) *SetSnapshotCallbackAuthRequest {
	s.CallbackReqAuth = &v
	return s
}

func (s *SetSnapshotCallbackAuthRequest) SetDomainName(v string) *SetSnapshotCallbackAuthRequest {
	s.DomainName = &v
	return s
}

func (s *SetSnapshotCallbackAuthRequest) SetOwnerId(v int64) *SetSnapshotCallbackAuthRequest {
	s.OwnerId = &v
	return s
}

func (s *SetSnapshotCallbackAuthRequest) SetRegionId(v string) *SetSnapshotCallbackAuthRequest {
	s.RegionId = &v
	return s
}

func (s *SetSnapshotCallbackAuthRequest) Validate() error {
	return dara.Validate(s)
}
