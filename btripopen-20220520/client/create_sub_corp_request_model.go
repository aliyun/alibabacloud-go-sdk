// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubCorpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOuterCorpId(v string) *CreateSubCorpRequest
	GetOuterCorpId() *string
	SetOuterCorpName(v string) *CreateSubCorpRequest
	GetOuterCorpName() *string
	SetUserId(v string) *CreateSubCorpRequest
	GetUserId() *string
}

type CreateSubCorpRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// corp123
	OuterCorpId *string `json:"outer_corp_id,omitempty" xml:"outer_corp_id,omitempty"`
	// This parameter is required.
	OuterCorpName *string `json:"outer_corp_name,omitempty" xml:"outer_corp_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CreateSubCorpRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSubCorpRequest) GoString() string {
	return s.String()
}

func (s *CreateSubCorpRequest) GetOuterCorpId() *string {
	return s.OuterCorpId
}

func (s *CreateSubCorpRequest) GetOuterCorpName() *string {
	return s.OuterCorpName
}

func (s *CreateSubCorpRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateSubCorpRequest) SetOuterCorpId(v string) *CreateSubCorpRequest {
	s.OuterCorpId = &v
	return s
}

func (s *CreateSubCorpRequest) SetOuterCorpName(v string) *CreateSubCorpRequest {
	s.OuterCorpName = &v
	return s
}

func (s *CreateSubCorpRequest) SetUserId(v string) *CreateSubCorpRequest {
	s.UserId = &v
	return s
}

func (s *CreateSubCorpRequest) Validate() error {
	return dara.Validate(s)
}
