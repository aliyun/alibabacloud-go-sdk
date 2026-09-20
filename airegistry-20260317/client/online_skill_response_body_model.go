// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OnlineSkillResponseBody
	GetRequestId() *string
}

type OnlineSkillResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D9E87E66-9EF0-5C10-A5E6-924020A0C9B7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnlineSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnlineSkillResponseBody) GoString() string {
	return s.String()
}

func (s *OnlineSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnlineSkillResponseBody) SetRequestId(v string) *OnlineSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnlineSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
