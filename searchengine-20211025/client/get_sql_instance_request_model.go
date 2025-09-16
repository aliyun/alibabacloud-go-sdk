// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersion(v int64) *GetSqlInstanceRequest
	GetVersion() *int64
}

type GetSqlInstanceRequest struct {
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetSqlInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSqlInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetSqlInstanceRequest) GetVersion() *int64 {
	return s.Version
}

func (s *GetSqlInstanceRequest) SetVersion(v int64) *GetSqlInstanceRequest {
	s.Version = &v
	return s
}

func (s *GetSqlInstanceRequest) Validate() error {
	return dara.Validate(s)
}
