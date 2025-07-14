// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmartqQueryAbilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCubeId(v string) *SmartqQueryAbilityRequest
	GetCubeId() *string
	SetUserId(v string) *SmartqQueryAbilityRequest
	GetUserId() *string
	SetUserQuestion(v string) *SmartqQueryAbilityRequest
	GetUserQuestion() *string
}

type SmartqQueryAbilityRequest struct {
	// Dataset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// User ID.
	//
	// 	Notice: If this field is not filled, the data will be queried by default as the organization owner.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// Question text.
	//
	// This parameter is required.
	//
	// example:
	//
	// This year\\"s sales data
	UserQuestion *string `json:"UserQuestion,omitempty" xml:"UserQuestion,omitempty"`
}

func (s SmartqQueryAbilityRequest) String() string {
	return dara.Prettify(s)
}

func (s SmartqQueryAbilityRequest) GoString() string {
	return s.String()
}

func (s *SmartqQueryAbilityRequest) GetCubeId() *string {
	return s.CubeId
}

func (s *SmartqQueryAbilityRequest) GetUserId() *string {
	return s.UserId
}

func (s *SmartqQueryAbilityRequest) GetUserQuestion() *string {
	return s.UserQuestion
}

func (s *SmartqQueryAbilityRequest) SetCubeId(v string) *SmartqQueryAbilityRequest {
	s.CubeId = &v
	return s
}

func (s *SmartqQueryAbilityRequest) SetUserId(v string) *SmartqQueryAbilityRequest {
	s.UserId = &v
	return s
}

func (s *SmartqQueryAbilityRequest) SetUserQuestion(v string) *SmartqQueryAbilityRequest {
	s.UserQuestion = &v
	return s
}

func (s *SmartqQueryAbilityRequest) Validate() error {
	return dara.Validate(s)
}
