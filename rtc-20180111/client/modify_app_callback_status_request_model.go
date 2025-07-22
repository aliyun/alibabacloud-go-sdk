// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppCallbackStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyAppCallbackStatusRequest
	GetAppId() *string
}

type ModifyAppCallbackStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 223***JQb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ModifyAppCallbackStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppCallbackStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppCallbackStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyAppCallbackStatusRequest) SetAppId(v string) *ModifyAppCallbackStatusRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppCallbackStatusRequest) Validate() error {
	return dara.Validate(s)
}
