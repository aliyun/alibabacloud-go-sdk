// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUsersCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *GetUsersCountRequest
	GetBizType() *string
	SetBusinessChannel(v string) *GetUsersCountRequest
	GetBusinessChannel() *string
	SetSolutionId(v string) *GetUsersCountRequest
	GetSolutionId() *string
}

type GetUsersCountRequest struct {
	// > This parameter is not available.
	//
	// example:
	//
	// ENTERPRISE
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The business channel type.
	//
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// > This parameter is currently not available.
	//
	// example:
	//
	// co-0esnf80jab***
	SolutionId *string `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
}

func (s GetUsersCountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUsersCountRequest) GoString() string {
	return s.String()
}

func (s *GetUsersCountRequest) GetBizType() *string {
	return s.BizType
}

func (s *GetUsersCountRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *GetUsersCountRequest) GetSolutionId() *string {
	return s.SolutionId
}

func (s *GetUsersCountRequest) SetBizType(v string) *GetUsersCountRequest {
	s.BizType = &v
	return s
}

func (s *GetUsersCountRequest) SetBusinessChannel(v string) *GetUsersCountRequest {
	s.BusinessChannel = &v
	return s
}

func (s *GetUsersCountRequest) SetSolutionId(v string) *GetUsersCountRequest {
	s.SolutionId = &v
	return s
}

func (s *GetUsersCountRequest) Validate() error {
	return dara.Validate(s)
}
