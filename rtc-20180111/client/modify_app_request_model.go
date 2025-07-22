// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyAppRequest
	GetAppId() *string
	SetAppName(v string) *ModifyAppRequest
	GetAppName() *string
	SetOwnerId(v int64) *ModifyAppRequest
	GetOwnerId() *int64
}

type ModifyAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ioeh****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// defaultName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ModifyAppRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *ModifyAppRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAppRequest) SetAppId(v string) *ModifyAppRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppRequest) SetAppName(v string) *ModifyAppRequest {
	s.AppName = &v
	return s
}

func (s *ModifyAppRequest) SetOwnerId(v int64) *ModifyAppRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAppRequest) Validate() error {
	return dara.Validate(s)
}
