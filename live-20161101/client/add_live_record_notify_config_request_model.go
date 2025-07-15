// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveRecordNotifyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *AddLiveRecordNotifyConfigRequest
	GetDomainName() *string
	SetNeedStatusNotify(v bool) *AddLiveRecordNotifyConfigRequest
	GetNeedStatusNotify() *bool
	SetNotifyUrl(v string) *AddLiveRecordNotifyConfigRequest
	GetNotifyUrl() *string
	SetOnDemandUrl(v string) *AddLiveRecordNotifyConfigRequest
	GetOnDemandUrl() *string
	SetOwnerId(v int64) *AddLiveRecordNotifyConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *AddLiveRecordNotifyConfigRequest
	GetSecurityToken() *string
}

type AddLiveRecordNotifyConfigRequest struct {
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Specifies whether to enable callbacks for recording status. Valid values:
	//
	// 	- true: enables callbacks for recording status. If you set this parameter to **true**, an example of recording status callback is returned.
	//
	// 	- false (default): disables callbacks for recording status.
	//
	// example:
	//
	// false
	NeedStatusNotify *bool `json:"NeedStatusNotify,omitempty" xml:"NeedStatusNotify,omitempty"`
	// The callback URL that is used to receive notifications about recording events and status.
	//
	// >  The URL must start with `http://` or `https://`. For more information, see [Callbacks for live stream recording](https://help.aliyun.com/document_detail/55016.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// http://demo.aliyundoc.com/examplecallback.action
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	// The callback URL for on-demand recordings.
	//
	// >  The URL must start with `http://` or `https://`. For more information, see [On-demand recording](https://help.aliyun.com/document_detail/85910.html).
	//
	// example:
	//
	// http://learn.aliyundoc.com/ondemandcallback.action
	OnDemandUrl   *string `json:"OnDemandUrl,omitempty" xml:"OnDemandUrl,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s AddLiveRecordNotifyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLiveRecordNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLiveRecordNotifyConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddLiveRecordNotifyConfigRequest) GetNeedStatusNotify() *bool {
	return s.NeedStatusNotify
}

func (s *AddLiveRecordNotifyConfigRequest) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *AddLiveRecordNotifyConfigRequest) GetOnDemandUrl() *string {
	return s.OnDemandUrl
}

func (s *AddLiveRecordNotifyConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLiveRecordNotifyConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AddLiveRecordNotifyConfigRequest) SetDomainName(v string) *AddLiveRecordNotifyConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddLiveRecordNotifyConfigRequest) SetNeedStatusNotify(v bool) *AddLiveRecordNotifyConfigRequest {
	s.NeedStatusNotify = &v
	return s
}

func (s *AddLiveRecordNotifyConfigRequest) SetNotifyUrl(v string) *AddLiveRecordNotifyConfigRequest {
	s.NotifyUrl = &v
	return s
}

func (s *AddLiveRecordNotifyConfigRequest) SetOnDemandUrl(v string) *AddLiveRecordNotifyConfigRequest {
	s.OnDemandUrl = &v
	return s
}

func (s *AddLiveRecordNotifyConfigRequest) SetOwnerId(v int64) *AddLiveRecordNotifyConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLiveRecordNotifyConfigRequest) SetSecurityToken(v string) *AddLiveRecordNotifyConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddLiveRecordNotifyConfigRequest) Validate() error {
	return dara.Validate(s)
}
