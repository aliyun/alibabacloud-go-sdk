// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMapPkFromHidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *MapPkFromHidRequest
	GetAppName() *string
	SetBid(v string) *MapPkFromHidRequest
	GetBid() *string
	SetHid(v string) *MapPkFromHidRequest
	GetHid() *string
	SetMappingScenes(v string) *MapPkFromHidRequest
	GetMappingScenes() *string
}

type MapPkFromHidRequest struct {
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Bid     *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// This parameter is required.
	Hid *string `json:"Hid,omitempty" xml:"Hid,omitempty"`
	// This parameter is required.
	MappingScenes *string `json:"MappingScenes,omitempty" xml:"MappingScenes,omitempty"`
}

func (s MapPkFromHidRequest) String() string {
	return dara.Prettify(s)
}

func (s MapPkFromHidRequest) GoString() string {
	return s.String()
}

func (s *MapPkFromHidRequest) GetAppName() *string {
	return s.AppName
}

func (s *MapPkFromHidRequest) GetBid() *string {
	return s.Bid
}

func (s *MapPkFromHidRequest) GetHid() *string {
	return s.Hid
}

func (s *MapPkFromHidRequest) GetMappingScenes() *string {
	return s.MappingScenes
}

func (s *MapPkFromHidRequest) SetAppName(v string) *MapPkFromHidRequest {
	s.AppName = &v
	return s
}

func (s *MapPkFromHidRequest) SetBid(v string) *MapPkFromHidRequest {
	s.Bid = &v
	return s
}

func (s *MapPkFromHidRequest) SetHid(v string) *MapPkFromHidRequest {
	s.Hid = &v
	return s
}

func (s *MapPkFromHidRequest) SetMappingScenes(v string) *MapPkFromHidRequest {
	s.MappingScenes = &v
	return s
}

func (s *MapPkFromHidRequest) Validate() error {
	return dara.Validate(s)
}
