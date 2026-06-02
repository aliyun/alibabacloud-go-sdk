// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallerAliUid(v string) *DeleteClientsRequest
	GetCallerAliUid() *string
	SetInManage(v bool) *DeleteClientsRequest
	GetInManage() *bool
	SetUuids(v []*string) *DeleteClientsRequest
	GetUuids() []*string
}

type DeleteClientsRequest struct {
	// aliuid
	//
	// example:
	//
	// ***
	CallerAliUid *string `json:"CallerAliUid,omitempty" xml:"CallerAliUid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// True
	InManage *bool `json:"InManage,omitempty" xml:"InManage,omitempty"`
	// This parameter is required.
	Uuids []*string `json:"Uuids,omitempty" xml:"Uuids,omitempty" type:"Repeated"`
}

func (s DeleteClientsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientsRequest) GoString() string {
	return s.String()
}

func (s *DeleteClientsRequest) GetCallerAliUid() *string {
	return s.CallerAliUid
}

func (s *DeleteClientsRequest) GetInManage() *bool {
	return s.InManage
}

func (s *DeleteClientsRequest) GetUuids() []*string {
	return s.Uuids
}

func (s *DeleteClientsRequest) SetCallerAliUid(v string) *DeleteClientsRequest {
	s.CallerAliUid = &v
	return s
}

func (s *DeleteClientsRequest) SetInManage(v bool) *DeleteClientsRequest {
	s.InManage = &v
	return s
}

func (s *DeleteClientsRequest) SetUuids(v []*string) *DeleteClientsRequest {
	s.Uuids = v
	return s
}

func (s *DeleteClientsRequest) Validate() error {
	return dara.Validate(s)
}
