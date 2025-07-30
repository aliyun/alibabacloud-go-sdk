// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainDomainProxyTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainProxyToken(v *ObtainDomainProxyTokenResponseBodyDomainProxyToken) *ObtainDomainProxyTokenResponseBody
	GetDomainProxyToken() *ObtainDomainProxyTokenResponseBodyDomainProxyToken
	SetRequestId(v string) *ObtainDomainProxyTokenResponseBody
	GetRequestId() *string
}

type ObtainDomainProxyTokenResponseBody struct {
	// The information about the proxy token.
	DomainProxyToken *ObtainDomainProxyTokenResponseBodyDomainProxyToken `json:"DomainProxyToken,omitempty" xml:"DomainProxyToken,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ObtainDomainProxyTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ObtainDomainProxyTokenResponseBody) GoString() string {
	return s.String()
}

func (s *ObtainDomainProxyTokenResponseBody) GetDomainProxyToken() *ObtainDomainProxyTokenResponseBodyDomainProxyToken {
	return s.DomainProxyToken
}

func (s *ObtainDomainProxyTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ObtainDomainProxyTokenResponseBody) SetDomainProxyToken(v *ObtainDomainProxyTokenResponseBodyDomainProxyToken) *ObtainDomainProxyTokenResponseBody {
	s.DomainProxyToken = v
	return s
}

func (s *ObtainDomainProxyTokenResponseBody) SetRequestId(v string) *ObtainDomainProxyTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *ObtainDomainProxyTokenResponseBody) Validate() error {
	return dara.Validate(s)
}

type ObtainDomainProxyTokenResponseBodyDomainProxyToken struct {
	// The time when the proxy token of the domain name was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1649830226000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the domain name.
	//
	// example:
	//
	// dm_examplexxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The proxy token of the domain name.
	//
	// example:
	//
	// PTxxxxxxxx
	DomainProxyToken *string `json:"DomainProxyToken,omitempty" xml:"DomainProxyToken,omitempty"`
	// The ID of the proxy token of the domain name.
	//
	// example:
	//
	// pt_examplexxxx
	DomainProxyTokenId *string `json:"DomainProxyTokenId,omitempty" xml:"DomainProxyTokenId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the proxy token of the domain name was last used. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1649830226000
	LastUsedTime *int64 `json:"LastUsedTime,omitempty" xml:"LastUsedTime,omitempty"`
	// The state of the proxy token. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the proxy token of the domain name was last updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1649830226000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ObtainDomainProxyTokenResponseBodyDomainProxyToken) String() string {
	return dara.Prettify(s)
}

func (s ObtainDomainProxyTokenResponseBodyDomainProxyToken) GoString() string {
	return s.String()
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) GetDomainId() *string {
	return s.DomainId
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) GetDomainProxyToken() *string {
	return s.DomainProxyToken
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) GetDomainProxyTokenId() *string {
	return s.DomainProxyTokenId
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) GetLastUsedTime() *int64 {
	return s.LastUsedTime
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) GetStatus() *string {
	return s.Status
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) SetCreateTime(v int64) *ObtainDomainProxyTokenResponseBodyDomainProxyToken {
	s.CreateTime = &v
	return s
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) SetDomainId(v string) *ObtainDomainProxyTokenResponseBodyDomainProxyToken {
	s.DomainId = &v
	return s
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) SetDomainProxyToken(v string) *ObtainDomainProxyTokenResponseBodyDomainProxyToken {
	s.DomainProxyToken = &v
	return s
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) SetDomainProxyTokenId(v string) *ObtainDomainProxyTokenResponseBodyDomainProxyToken {
	s.DomainProxyTokenId = &v
	return s
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) SetInstanceId(v string) *ObtainDomainProxyTokenResponseBodyDomainProxyToken {
	s.InstanceId = &v
	return s
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) SetLastUsedTime(v int64) *ObtainDomainProxyTokenResponseBodyDomainProxyToken {
	s.LastUsedTime = &v
	return s
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) SetStatus(v string) *ObtainDomainProxyTokenResponseBodyDomainProxyToken {
	s.Status = &v
	return s
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) SetUpdateTime(v int64) *ObtainDomainProxyTokenResponseBodyDomainProxyToken {
	s.UpdateTime = &v
	return s
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) Validate() error {
	return dara.Validate(s)
}
