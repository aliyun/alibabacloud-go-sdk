// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartParentPlatformResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *StartParentPlatformResponseBody
	GetId() *string
	SetRequestId(v string) *StartParentPlatformResponseBody
	GetRequestId() *string
}

type StartParentPlatformResponseBody struct {
	// example:
	//
	// 359*****374-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartParentPlatformResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartParentPlatformResponseBody) GoString() string {
	return s.String()
}

func (s *StartParentPlatformResponseBody) GetId() *string {
	return s.Id
}

func (s *StartParentPlatformResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartParentPlatformResponseBody) SetId(v string) *StartParentPlatformResponseBody {
	s.Id = &v
	return s
}

func (s *StartParentPlatformResponseBody) SetRequestId(v string) *StartParentPlatformResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartParentPlatformResponseBody) Validate() error {
	return dara.Validate(s)
}
