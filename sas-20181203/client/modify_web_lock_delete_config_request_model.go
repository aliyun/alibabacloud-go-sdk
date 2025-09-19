// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockDeleteConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int32) *ModifyWebLockDeleteConfigRequest
	GetId() *int32
	SetLang(v string) *ModifyWebLockDeleteConfigRequest
	GetLang() *string
	SetSourceIp(v string) *ModifyWebLockDeleteConfigRequest
	GetSourceIp() *string
	SetUuid(v string) *ModifyWebLockDeleteConfigRequest
	GetUuid() *string
}

type ModifyWebLockDeleteConfigRequest struct {
	// The ID of the protected directory that you want to delete.
	//
	// > You can call the [DescribeWebLockConfigList](~~DescribeWebLockConfigList~~) operation to query the IDs of protected directories.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 1.2.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The UUID of the server for which you want to delete the protected directory.
	//
	// > You can call the [DescribeWebLockConfigList](~~DescribeWebLockConfigList~~) operation to query the UUIDs of servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7f7fe9a2-55de-4b9d-a37a-0d981d36****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ModifyWebLockDeleteConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockDeleteConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebLockDeleteConfigRequest) GetId() *int32 {
	return s.Id
}

func (s *ModifyWebLockDeleteConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyWebLockDeleteConfigRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyWebLockDeleteConfigRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ModifyWebLockDeleteConfigRequest) SetId(v int32) *ModifyWebLockDeleteConfigRequest {
	s.Id = &v
	return s
}

func (s *ModifyWebLockDeleteConfigRequest) SetLang(v string) *ModifyWebLockDeleteConfigRequest {
	s.Lang = &v
	return s
}

func (s *ModifyWebLockDeleteConfigRequest) SetSourceIp(v string) *ModifyWebLockDeleteConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyWebLockDeleteConfigRequest) SetUuid(v string) *ModifyWebLockDeleteConfigRequest {
	s.Uuid = &v
	return s
}

func (s *ModifyWebLockDeleteConfigRequest) Validate() error {
	return dara.Validate(s)
}
