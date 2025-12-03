// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndAllCoordinationByOwnerRequest interface {
  dara.Model
  String() string
  GoString() string
  SetLoginRegionId(v string) *EndAllCoordinationByOwnerRequest
  GetLoginRegionId() *string 
  SetLoginToken(v string) *EndAllCoordinationByOwnerRequest
  GetLoginToken() *string 
  SetResourceId(v string) *EndAllCoordinationByOwnerRequest
  GetResourceId() *string 
  SetResourceRegionId(v string) *EndAllCoordinationByOwnerRequest
  GetResourceRegionId() *string 
  SetResourceType(v string) *EndAllCoordinationByOwnerRequest
  GetResourceType() *string 
  SetSessionId(v string) *EndAllCoordinationByOwnerRequest
  GetSessionId() *string 
}

type EndAllCoordinationByOwnerRequest struct {
  // example:
  // 
  // cn-hangzhou
  LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
  // example:
  // 
  // v2c4e2ef03d62******
  LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
  // example:
  // 
  // ecd-68a7ddrt0******
  ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
  // example:
  // 
  // cn-hangzhou
  ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
  // example:
  // 
  // CloudDesktop
  ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
  // example:
  // 
  // 09e2b2e6-3181******
  SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s EndAllCoordinationByOwnerRequest) String() string {
  return dara.Prettify(s)
}

func (s EndAllCoordinationByOwnerRequest) GoString() string {
  return s.String()
}

func (s *EndAllCoordinationByOwnerRequest) GetLoginRegionId() *string  {
  return s.LoginRegionId
}

func (s *EndAllCoordinationByOwnerRequest) GetLoginToken() *string  {
  return s.LoginToken
}

func (s *EndAllCoordinationByOwnerRequest) GetResourceId() *string  {
  return s.ResourceId
}

func (s *EndAllCoordinationByOwnerRequest) GetResourceRegionId() *string  {
  return s.ResourceRegionId
}

func (s *EndAllCoordinationByOwnerRequest) GetResourceType() *string  {
  return s.ResourceType
}

func (s *EndAllCoordinationByOwnerRequest) GetSessionId() *string  {
  return s.SessionId
}

func (s *EndAllCoordinationByOwnerRequest) SetLoginRegionId(v string) *EndAllCoordinationByOwnerRequest {
  s.LoginRegionId = &v
  return s
}

func (s *EndAllCoordinationByOwnerRequest) SetLoginToken(v string) *EndAllCoordinationByOwnerRequest {
  s.LoginToken = &v
  return s
}

func (s *EndAllCoordinationByOwnerRequest) SetResourceId(v string) *EndAllCoordinationByOwnerRequest {
  s.ResourceId = &v
  return s
}

func (s *EndAllCoordinationByOwnerRequest) SetResourceRegionId(v string) *EndAllCoordinationByOwnerRequest {
  s.ResourceRegionId = &v
  return s
}

func (s *EndAllCoordinationByOwnerRequest) SetResourceType(v string) *EndAllCoordinationByOwnerRequest {
  s.ResourceType = &v
  return s
}

func (s *EndAllCoordinationByOwnerRequest) SetSessionId(v string) *EndAllCoordinationByOwnerRequest {
  s.SessionId = &v
  return s
}

func (s *EndAllCoordinationByOwnerRequest) Validate() error {
  return dara.Validate(s)
}

