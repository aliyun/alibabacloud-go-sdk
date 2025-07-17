// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDifyAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppUuid(v string) *DescribeDifyAttributeRequest
	GetAppUuid() *string
	SetClientToken(v string) *DescribeDifyAttributeRequest
	GetClientToken() *string
	SetDataRegion(v string) *DescribeDifyAttributeRequest
	GetDataRegion() *string
	SetWorkspaceId(v string) *DescribeDifyAttributeRequest
	GetWorkspaceId() *string
}

type DescribeDifyAttributeRequest struct {
	// example:
	//
	// 92748163-af62-4ca4-ad85-1****
	AppUuid *string `json:"AppUuid,omitempty" xml:"AppUuid,omitempty"`
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// cn-hangzhou
	DataRegion *string `json:"DataRegion,omitempty" xml:"DataRegion,omitempty"`
	// example:
	//
	// 339170706****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DescribeDifyAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDifyAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDifyAttributeRequest) GetAppUuid() *string {
	return s.AppUuid
}

func (s *DescribeDifyAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeDifyAttributeRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *DescribeDifyAttributeRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeDifyAttributeRequest) SetAppUuid(v string) *DescribeDifyAttributeRequest {
	s.AppUuid = &v
	return s
}

func (s *DescribeDifyAttributeRequest) SetClientToken(v string) *DescribeDifyAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDifyAttributeRequest) SetDataRegion(v string) *DescribeDifyAttributeRequest {
	s.DataRegion = &v
	return s
}

func (s *DescribeDifyAttributeRequest) SetWorkspaceId(v string) *DescribeDifyAttributeRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeDifyAttributeRequest) Validate() error {
	return dara.Validate(s)
}
