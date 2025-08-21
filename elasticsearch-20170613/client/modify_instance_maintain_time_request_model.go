// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceMaintainTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyInstanceMaintainTimeRequest
	GetClientToken() *string
	SetBody(v string) *ModifyInstanceMaintainTimeRequest
	GetBody() *string
}

type ModifyInstanceMaintainTimeRequest struct {
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// {     "openMaintainTime": true,     "maintainStartTime": "03:00Z",     "maintainEndTime": "04:00Z" }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceMaintainTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceMaintainTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMaintainTimeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyInstanceMaintainTimeRequest) GetBody() *string {
	return s.Body
}

func (s *ModifyInstanceMaintainTimeRequest) SetClientToken(v string) *ModifyInstanceMaintainTimeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) SetBody(v string) *ModifyInstanceMaintainTimeRequest {
	s.Body = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) Validate() error {
	return dara.Validate(s)
}
