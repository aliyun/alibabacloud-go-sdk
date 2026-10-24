// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityEventMarkMissListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v []*int64) *DeleteSecurityEventMarkMissListRequest
	GetIds() []*int64
}

type DeleteSecurityEventMarkMissListRequest struct {
	// The list of alert whitelist rule IDs.
	//
	// This parameter is required.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
}

func (s DeleteSecurityEventMarkMissListRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityEventMarkMissListRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityEventMarkMissListRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *DeleteSecurityEventMarkMissListRequest) SetIds(v []*int64) *DeleteSecurityEventMarkMissListRequest {
	s.Ids = v
	return s
}

func (s *DeleteSecurityEventMarkMissListRequest) Validate() error {
	return dara.Validate(s)
}
