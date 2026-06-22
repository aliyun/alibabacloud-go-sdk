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
	// The language type for requests and responses. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 125.71.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The protection status of the server. Valid values:
	//
	// - **on**: Enables protection.
	//
	// - **off**: Shuts down protection.
	//
	// > After you shut down web tamper-proofing for the server, a tamper-proofing authorization quota is released.
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UUID of the server for which you want to modify the brute-force attacks prevention status.
	//
	// You can invoke the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to obtain the UUID of the server.
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
