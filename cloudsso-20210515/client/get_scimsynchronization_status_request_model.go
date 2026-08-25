// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSCIMSynchronizationStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetSCIMSynchronizationStatusRequest
	GetDirectoryId() *string
}

type GetSCIMSynchronizationStatusRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetSCIMSynchronizationStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSCIMSynchronizationStatusRequest) GoString() string {
	return s.String()
}

func (s *GetSCIMSynchronizationStatusRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetSCIMSynchronizationStatusRequest) SetDirectoryId(v string) *GetSCIMSynchronizationStatusRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetSCIMSynchronizationStatusRequest) Validate() error {
	return dara.Validate(s)
}
