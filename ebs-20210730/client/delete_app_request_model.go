// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteAppRequest
	GetAppId() *string
	SetClientToken(v string) *DeleteAppRequest
	GetClientToken() *string
	SetOwner(v string) *DeleteAppRequest
	GetOwner() *string
	SetRegionId(v string) *DeleteAppRequest
	GetRegionId() *string
}

type DeleteAppRequest struct {
	// example:
	//
	// app-xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// anchashi
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteAppRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteAppRequest) GetOwner() *string {
	return s.Owner
}

func (s *DeleteAppRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAppRequest) SetAppId(v string) *DeleteAppRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppRequest) SetClientToken(v string) *DeleteAppRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAppRequest) SetOwner(v string) *DeleteAppRequest {
	s.Owner = &v
	return s
}

func (s *DeleteAppRequest) SetRegionId(v string) *DeleteAppRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAppRequest) Validate() error {
	return dara.Validate(s)
}
