// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ModifyWebLockStatusRequest
	GetLang() *string
	SetSourceIp(v string) *ModifyWebLockStatusRequest
	GetSourceIp() *string
	SetStatus(v string) *ModifyWebLockStatusRequest
	GetStatus() *string
	SetUuid(v string) *ModifyWebLockStatusRequest
	GetUuid() *string
}

type ModifyWebLockStatusRequest struct {
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
	// 125.71.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// Specifies whether to enable or disable web tamper proofing for the specified server. Valid values:
	//
	// 	- **on**: enables web tamper proofing
	//
	// 	- **off**: disables web tamper proofing
	//
	// > After you disable web tamper proofing for the specified server, one quota is released.
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UUID of the server for which you want to enable or disable web tamper proofing. You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// inet-1234567****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ModifyWebLockStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebLockStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyWebLockStatusRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyWebLockStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyWebLockStatusRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ModifyWebLockStatusRequest) SetLang(v string) *ModifyWebLockStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyWebLockStatusRequest) SetSourceIp(v string) *ModifyWebLockStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyWebLockStatusRequest) SetStatus(v string) *ModifyWebLockStatusRequest {
	s.Status = &v
	return s
}

func (s *ModifyWebLockStatusRequest) SetUuid(v string) *ModifyWebLockStatusRequest {
	s.Uuid = &v
	return s
}

func (s *ModifyWebLockStatusRequest) Validate() error {
	return dara.Validate(s)
}
