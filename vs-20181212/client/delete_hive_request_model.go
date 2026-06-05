// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHiveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHiveId(v string) *DeleteHiveRequest
	GetHiveId() *string
}

type DeleteHiveRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// g-xxxx
	HiveId *string `json:"HiveId,omitempty" xml:"HiveId,omitempty"`
}

func (s DeleteHiveRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHiveRequest) GoString() string {
	return s.String()
}

func (s *DeleteHiveRequest) GetHiveId() *string {
	return s.HiveId
}

func (s *DeleteHiveRequest) SetHiveId(v string) *DeleteHiveRequest {
	s.HiveId = &v
	return s
}

func (s *DeleteHiveRequest) Validate() error {
	return dara.Validate(s)
}
