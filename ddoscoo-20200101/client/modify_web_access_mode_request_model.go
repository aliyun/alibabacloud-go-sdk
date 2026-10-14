// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebAccessModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessMode(v int32) *ModifyWebAccessModeRequest
	GetAccessMode() *int32
	SetDomain(v string) *ModifyWebAccessModeRequest
	GetDomain() *string
}

type ModifyWebAccessModeRequest struct {
	// The access mode for website services. Valid values:
	//
	// - **0**: A record
	//
	// - **1**: Anti-DDoS Pro mode
	//
	// - **2**: Back-to-origin mode
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	AccessMode *int32 `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	// The domain name of the website service.
	//
	// > The domain name must have forwarding rules configured. You can call [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s ModifyWebAccessModeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebAccessModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebAccessModeRequest) GetAccessMode() *int32 {
	return s.AccessMode
}

func (s *ModifyWebAccessModeRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyWebAccessModeRequest) SetAccessMode(v int32) *ModifyWebAccessModeRequest {
	s.AccessMode = &v
	return s
}

func (s *ModifyWebAccessModeRequest) SetDomain(v string) *ModifyWebAccessModeRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebAccessModeRequest) Validate() error {
	return dara.Validate(s)
}
