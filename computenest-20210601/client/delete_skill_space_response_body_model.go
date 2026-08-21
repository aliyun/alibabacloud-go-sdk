// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSkillSpaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSkillSpaceResponseBody
	GetRequestId() *string
}

type DeleteSkillSpaceResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 13FE89A5-C036-56BF-A0FF-A31C59819FD7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSkillSpaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSkillSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSkillSpaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSkillSpaceResponseBody) SetRequestId(v string) *DeleteSkillSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSkillSpaceResponseBody) Validate() error {
	return dara.Validate(s)
}
