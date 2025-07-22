// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppAgentFunctionStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyAppAgentFunctionStatusRequest
	GetAppId() *string
}

type ModifyAppAgentFunctionStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ModifyAppAgentFunctionStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppAgentFunctionStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppAgentFunctionStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyAppAgentFunctionStatusRequest) SetAppId(v string) *ModifyAppAgentFunctionStatusRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppAgentFunctionStatusRequest) Validate() error {
	return dara.Validate(s)
}
