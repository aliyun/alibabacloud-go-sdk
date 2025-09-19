// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetInfoPublishRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListAssetInfoPublishRequest
	GetName() *string
	SetUuidList(v []*string) *ListAssetInfoPublishRequest
	GetUuidList() []*string
}

type ListAssetInfoPublishRequest struct {
	// An extended parameter. This parameter is temporarily unavailable.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The UUIDs of the servers that you want to query.
	//
	// This parameter is required.
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s ListAssetInfoPublishRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAssetInfoPublishRequest) GoString() string {
	return s.String()
}

func (s *ListAssetInfoPublishRequest) GetName() *string {
	return s.Name
}

func (s *ListAssetInfoPublishRequest) GetUuidList() []*string {
	return s.UuidList
}

func (s *ListAssetInfoPublishRequest) SetName(v string) *ListAssetInfoPublishRequest {
	s.Name = &v
	return s
}

func (s *ListAssetInfoPublishRequest) SetUuidList(v []*string) *ListAssetInfoPublishRequest {
	s.UuidList = v
	return s
}

func (s *ListAssetInfoPublishRequest) Validate() error {
	return dara.Validate(s)
}
