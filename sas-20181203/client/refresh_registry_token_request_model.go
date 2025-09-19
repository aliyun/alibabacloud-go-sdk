// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshRegistryTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegistryId(v int64) *RefreshRegistryTokenRequest
	GetRegistryId() *int64
}

type RefreshRegistryTokenRequest struct {
	// The ID of the image repository.
	//
	// >  You can call the [PageImageRegistry](~~PageImageRegistry~~) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	RegistryId *int64 `json:"RegistryId,omitempty" xml:"RegistryId,omitempty"`
}

func (s RefreshRegistryTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshRegistryTokenRequest) GoString() string {
	return s.String()
}

func (s *RefreshRegistryTokenRequest) GetRegistryId() *int64 {
	return s.RegistryId
}

func (s *RefreshRegistryTokenRequest) SetRegistryId(v int64) *RefreshRegistryTokenRequest {
	s.RegistryId = &v
	return s
}

func (s *RefreshRegistryTokenRequest) Validate() error {
	return dara.Validate(s)
}
