// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateParentPlatformResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateParentPlatformResponseBody
	GetId() *string
	SetRequestId(v string) *CreateParentPlatformResponseBody
	GetRequestId() *string
}

type CreateParentPlatformResponseBody struct {
	// example:
	//
	// 359*****374-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateParentPlatformResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateParentPlatformResponseBody) GoString() string {
	return s.String()
}

func (s *CreateParentPlatformResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateParentPlatformResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateParentPlatformResponseBody) SetId(v string) *CreateParentPlatformResponseBody {
	s.Id = &v
	return s
}

func (s *CreateParentPlatformResponseBody) SetRequestId(v string) *CreateParentPlatformResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateParentPlatformResponseBody) Validate() error {
	return dara.Validate(s)
}
