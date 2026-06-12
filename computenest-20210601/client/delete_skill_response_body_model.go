// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSkillResponseBody
	GetRequestId() *string
}

type DeleteSkillResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSkillResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSkillResponseBody) SetRequestId(v string) *DeleteSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
