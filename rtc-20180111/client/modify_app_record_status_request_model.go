// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppRecordStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyAppRecordStatusRequest
	GetAppId() *string
	SetClientToken(v string) *ModifyAppRecordStatusRequest
	GetClientToken() *string
}

type ModifyAppRecordStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 53200b81-b761-4c10-842a-a0726xxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ModifyAppRecordStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppRecordStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppRecordStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyAppRecordStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyAppRecordStatusRequest) SetAppId(v string) *ModifyAppRecordStatusRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppRecordStatusRequest) SetClientToken(v string) *ModifyAppRecordStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyAppRecordStatusRequest) Validate() error {
	return dara.Validate(s)
}
