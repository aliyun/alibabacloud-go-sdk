// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateCdnUrlAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputUrl(v string) *ValidateCdnUrlAuthRequest
	GetInputUrl() *string
	SetOwnerId(v string) *ValidateCdnUrlAuthRequest
	GetOwnerId() *string
	SetResourceRealOwnerId(v int64) *ValidateCdnUrlAuthRequest
	GetResourceRealOwnerId() *int64
	SetType(v string) *ValidateCdnUrlAuthRequest
	GetType() *string
}

type ValidateCdnUrlAuthRequest struct {
	// This parameter is required.
	InputUrl            *string `json:"InputUrl,omitempty" xml:"InputUrl,omitempty"`
	OwnerId             *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceRealOwnerId *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ValidateCdnUrlAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateCdnUrlAuthRequest) GoString() string {
	return s.String()
}

func (s *ValidateCdnUrlAuthRequest) GetInputUrl() *string {
	return s.InputUrl
}

func (s *ValidateCdnUrlAuthRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ValidateCdnUrlAuthRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *ValidateCdnUrlAuthRequest) GetType() *string {
	return s.Type
}

func (s *ValidateCdnUrlAuthRequest) SetInputUrl(v string) *ValidateCdnUrlAuthRequest {
	s.InputUrl = &v
	return s
}

func (s *ValidateCdnUrlAuthRequest) SetOwnerId(v string) *ValidateCdnUrlAuthRequest {
	s.OwnerId = &v
	return s
}

func (s *ValidateCdnUrlAuthRequest) SetResourceRealOwnerId(v int64) *ValidateCdnUrlAuthRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *ValidateCdnUrlAuthRequest) SetType(v string) *ValidateCdnUrlAuthRequest {
	s.Type = &v
	return s
}

func (s *ValidateCdnUrlAuthRequest) Validate() error {
	return dara.Validate(s)
}
