// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveWhiteListSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v []*int64) *RemoveWhiteListSettingRequest
	GetIds() []*int64
	SetServiceCode(v string) *RemoveWhiteListSettingRequest
	GetServiceCode() *string
}

type RemoveWhiteListSettingRequest struct {
	// The IDs of the whitelists to delete in batch.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// The ServiceCode of the ID Verification service. Set this parameter to **antcloudauth**.
	//
	// example:
	//
	// antcloudauth
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s RemoveWhiteListSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveWhiteListSettingRequest) GoString() string {
	return s.String()
}

func (s *RemoveWhiteListSettingRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *RemoveWhiteListSettingRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *RemoveWhiteListSettingRequest) SetIds(v []*int64) *RemoveWhiteListSettingRequest {
	s.Ids = v
	return s
}

func (s *RemoveWhiteListSettingRequest) SetServiceCode(v string) *RemoveWhiteListSettingRequest {
	s.ServiceCode = &v
	return s
}

func (s *RemoveWhiteListSettingRequest) Validate() error {
	return dara.Validate(s)
}
