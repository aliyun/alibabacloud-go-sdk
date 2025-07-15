// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbolishApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *AbolishApiRequest
	GetApiId() *string
	SetGroupId(v string) *AbolishApiRequest
	GetGroupId() *string
	SetSecurityToken(v string) *AbolishApiRequest
	GetSecurityToken() *string
	SetStageName(v string) *AbolishApiRequest
	GetStageName() *string
}

type AbolishApiRequest struct {
	// The ID of the specified API.
	//
	// This parameter is required.
	//
	// example:
	//
	// d6f679aeb3be4b91b3688e887ca1fe16
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The ID of the API group.
	//
	// example:
	//
	// 123
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The name of the runtime environment. Valid values:
	//
	// 	- **RELEASE**
	//
	// 	- **TEST**
	//
	// This parameter is required.
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s AbolishApiRequest) String() string {
	return dara.Prettify(s)
}

func (s AbolishApiRequest) GoString() string {
	return s.String()
}

func (s *AbolishApiRequest) GetApiId() *string {
	return s.ApiId
}

func (s *AbolishApiRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *AbolishApiRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AbolishApiRequest) GetStageName() *string {
	return s.StageName
}

func (s *AbolishApiRequest) SetApiId(v string) *AbolishApiRequest {
	s.ApiId = &v
	return s
}

func (s *AbolishApiRequest) SetGroupId(v string) *AbolishApiRequest {
	s.GroupId = &v
	return s
}

func (s *AbolishApiRequest) SetSecurityToken(v string) *AbolishApiRequest {
	s.SecurityToken = &v
	return s
}

func (s *AbolishApiRequest) SetStageName(v string) *AbolishApiRequest {
	s.StageName = &v
	return s
}

func (s *AbolishApiRequest) Validate() error {
	return dara.Validate(s)
}
