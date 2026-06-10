// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillHubConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSkillHubConfigResponseBody
	GetRequestId() *string
}

type CreateSkillHubConfigResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 3F976EF8-C10A-57DC-917C-BB7BEB508FFB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSkillHubConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillHubConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSkillHubConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSkillHubConfigResponseBody) SetRequestId(v string) *CreateSkillHubConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSkillHubConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
