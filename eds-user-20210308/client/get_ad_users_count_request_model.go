// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdUsersCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *GetAdUsersCountRequest
	GetBizType() *string
	SetBusinessChannel(v string) *GetAdUsersCountRequest
	GetBusinessChannel() *string
	SetSolutionId(v string) *GetAdUsersCountRequest
	GetSolutionId() *string
}

type GetAdUsersCountRequest struct {
	// example:
	//
	// ENTERPRISE
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// example:
	//
	// co-0esnf80jab***
	SolutionId *string `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
}

func (s GetAdUsersCountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAdUsersCountRequest) GoString() string {
	return s.String()
}

func (s *GetAdUsersCountRequest) GetBizType() *string {
	return s.BizType
}

func (s *GetAdUsersCountRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *GetAdUsersCountRequest) GetSolutionId() *string {
	return s.SolutionId
}

func (s *GetAdUsersCountRequest) SetBizType(v string) *GetAdUsersCountRequest {
	s.BizType = &v
	return s
}

func (s *GetAdUsersCountRequest) SetBusinessChannel(v string) *GetAdUsersCountRequest {
	s.BusinessChannel = &v
	return s
}

func (s *GetAdUsersCountRequest) SetSolutionId(v string) *GetAdUsersCountRequest {
	s.SolutionId = &v
	return s
}

func (s *GetAdUsersCountRequest) Validate() error {
	return dara.Validate(s)
}
