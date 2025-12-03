// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *OpenBackupRequest
	GetClusterId() *string
}

type OpenBackupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s OpenBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenBackupRequest) GoString() string {
	return s.String()
}

func (s *OpenBackupRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OpenBackupRequest) SetClusterId(v string) *OpenBackupRequest {
	s.ClusterId = &v
	return s
}

func (s *OpenBackupRequest) Validate() error {
	return dara.Validate(s)
}
