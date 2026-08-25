// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSCIMSynchronizationStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *SetSCIMSynchronizationStatusRequest
	GetDirectoryId() *string
	SetSCIMSynchronizationStatus(v string) *SetSCIMSynchronizationStatusRequest
	GetSCIMSynchronizationStatus() *string
}

type SetSCIMSynchronizationStatusRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The status of SCIM synchronization. Valid values:
	//
	// 	- Enabled
	//
	// 	- Disabled
	//
	// example:
	//
	// Enabled
	SCIMSynchronizationStatus *string `json:"SCIMSynchronizationStatus,omitempty" xml:"SCIMSynchronizationStatus,omitempty"`
}

func (s SetSCIMSynchronizationStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSCIMSynchronizationStatusRequest) GoString() string {
	return s.String()
}

func (s *SetSCIMSynchronizationStatusRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *SetSCIMSynchronizationStatusRequest) GetSCIMSynchronizationStatus() *string {
	return s.SCIMSynchronizationStatus
}

func (s *SetSCIMSynchronizationStatusRequest) SetDirectoryId(v string) *SetSCIMSynchronizationStatusRequest {
	s.DirectoryId = &v
	return s
}

func (s *SetSCIMSynchronizationStatusRequest) SetSCIMSynchronizationStatus(v string) *SetSCIMSynchronizationStatusRequest {
	s.SCIMSynchronizationStatus = &v
	return s
}

func (s *SetSCIMSynchronizationStatusRequest) Validate() error {
	return dara.Validate(s)
}
