// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyBastionAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMobile(v int64) *ApplyBastionAccountRequest
	GetMobile() *int64
	SetProjectId(v int64) *ApplyBastionAccountRequest
	GetProjectId() *int64
}

type ApplyBastionAccountRequest struct {
	// Mobile Phone Number
	//
	// example:
	//
	// ***
	Mobile *int64 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// Project ID
	//
	// example:
	//
	// 90912
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ApplyBastionAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyBastionAccountRequest) GoString() string {
	return s.String()
}

func (s *ApplyBastionAccountRequest) GetMobile() *int64 {
	return s.Mobile
}

func (s *ApplyBastionAccountRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ApplyBastionAccountRequest) SetMobile(v int64) *ApplyBastionAccountRequest {
	s.Mobile = &v
	return s
}

func (s *ApplyBastionAccountRequest) SetProjectId(v int64) *ApplyBastionAccountRequest {
	s.ProjectId = &v
	return s
}

func (s *ApplyBastionAccountRequest) Validate() error {
	return dara.Validate(s)
}
