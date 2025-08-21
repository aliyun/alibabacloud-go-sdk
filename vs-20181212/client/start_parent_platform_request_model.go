// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartParentPlatformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *StartParentPlatformRequest
	GetId() *string
	SetOwnerId(v int64) *StartParentPlatformRequest
	GetOwnerId() *int64
}

type StartParentPlatformRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 359*****374-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s StartParentPlatformRequest) String() string {
	return dara.Prettify(s)
}

func (s StartParentPlatformRequest) GoString() string {
	return s.String()
}

func (s *StartParentPlatformRequest) GetId() *string {
	return s.Id
}

func (s *StartParentPlatformRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartParentPlatformRequest) SetId(v string) *StartParentPlatformRequest {
	s.Id = &v
	return s
}

func (s *StartParentPlatformRequest) SetOwnerId(v int64) *StartParentPlatformRequest {
	s.OwnerId = &v
	return s
}

func (s *StartParentPlatformRequest) Validate() error {
	return dara.Validate(s)
}
