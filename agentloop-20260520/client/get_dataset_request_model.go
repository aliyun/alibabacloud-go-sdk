// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersion(v string) *GetDatasetRequest
	GetVersion() *string
}

type GetDatasetRequest struct {
	// The version number.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetRequest) GoString() string {
	return s.String()
}

func (s *GetDatasetRequest) GetVersion() *string {
	return s.Version
}

func (s *GetDatasetRequest) SetVersion(v string) *GetDatasetRequest {
	s.Version = &v
	return s
}

func (s *GetDatasetRequest) Validate() error {
	return dara.Validate(s)
}
