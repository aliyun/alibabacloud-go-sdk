// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveWhiteListSettingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdsShrink(v string) *RemoveWhiteListSettingShrinkRequest
	GetIdsShrink() *string
	SetServiceCode(v string) *RemoveWhiteListSettingShrinkRequest
	GetServiceCode() *string
}

type RemoveWhiteListSettingShrinkRequest struct {
	// IDs of the whitelist to be deleted in bulk.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// ServiceCode for the real person cloud product, only value: **antcloudauth**.
	//
	// example:
	//
	// antcloudauth
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s RemoveWhiteListSettingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveWhiteListSettingShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveWhiteListSettingShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *RemoveWhiteListSettingShrinkRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *RemoveWhiteListSettingShrinkRequest) SetIdsShrink(v string) *RemoveWhiteListSettingShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *RemoveWhiteListSettingShrinkRequest) SetServiceCode(v string) *RemoveWhiteListSettingShrinkRequest {
	s.ServiceCode = &v
	return s
}

func (s *RemoveWhiteListSettingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
