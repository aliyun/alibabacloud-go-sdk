// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CloseBackupRequest
	GetClusterId() *string
}

type CloseBackupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s CloseBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseBackupRequest) GoString() string {
	return s.String()
}

func (s *CloseBackupRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CloseBackupRequest) SetClusterId(v string) *CloseBackupRequest {
	s.ClusterId = &v
	return s
}

func (s *CloseBackupRequest) Validate() error {
	return dara.Validate(s)
}
