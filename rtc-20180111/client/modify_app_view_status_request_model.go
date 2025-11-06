// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppViewStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyAppViewStatusRequest
	GetAppId() *string
}

type ModifyAppViewStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ModifyAppViewStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppViewStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppViewStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyAppViewStatusRequest) SetAppId(v string) *ModifyAppViewStatusRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppViewStatusRequest) Validate() error {
	return dara.Validate(s)
}
