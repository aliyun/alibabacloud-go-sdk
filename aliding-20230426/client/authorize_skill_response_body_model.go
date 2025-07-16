// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *AuthorizeSkillResponseBody
	GetData() interface{}
	SetRequestId(v string) *AuthorizeSkillResponseBody
	GetRequestId() *string
}

type AuthorizeSkillResponseBody struct {
	// example:
	//
	// {}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// A348BA5D-FFD4-57E4-9450-23A14D72F331
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeSkillResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeSkillResponseBody) GetData() interface{} {
	return s.Data
}

func (s *AuthorizeSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeSkillResponseBody) SetData(v interface{}) *AuthorizeSkillResponseBody {
	s.Data = v
	return s
}

func (s *AuthorizeSkillResponseBody) SetRequestId(v string) *AuthorizeSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
