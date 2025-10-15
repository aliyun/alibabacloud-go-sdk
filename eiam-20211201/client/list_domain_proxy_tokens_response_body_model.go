// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainProxyTokensResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainProxyTokens(v []*ListDomainProxyTokensResponseBodyDomainProxyTokens) *ListDomainProxyTokensResponseBody
	GetDomainProxyTokens() []*ListDomainProxyTokensResponseBodyDomainProxyTokens
	SetRequestId(v string) *ListDomainProxyTokensResponseBody
	GetRequestId() *string
}

type ListDomainProxyTokensResponseBody struct {
	// The proxy tokens of the domain name.
	DomainProxyTokens []*ListDomainProxyTokensResponseBodyDomainProxyTokens `json:"DomainProxyTokens,omitempty" xml:"DomainProxyTokens,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDomainProxyTokensResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDomainProxyTokensResponseBody) GoString() string {
	return s.String()
}

func (s *ListDomainProxyTokensResponseBody) GetDomainProxyTokens() []*ListDomainProxyTokensResponseBodyDomainProxyTokens {
	return s.DomainProxyTokens
}

func (s *ListDomainProxyTokensResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDomainProxyTokensResponseBody) SetDomainProxyTokens(v []*ListDomainProxyTokensResponseBodyDomainProxyTokens) *ListDomainProxyTokensResponseBody {
	s.DomainProxyTokens = v
	return s
}

func (s *ListDomainProxyTokensResponseBody) SetRequestId(v string) *ListDomainProxyTokensResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDomainProxyTokensResponseBody) Validate() error {
	if s.DomainProxyTokens != nil {
		for _, item := range s.DomainProxyTokens {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDomainProxyTokensResponseBodyDomainProxyTokens struct {
	// The time when the proxy token of the domain name was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1649830226000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// dm_examplexxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The proxy token of the domain.
	//
	// example:
	//
	// PTxxxxxxxx
	DomainProxyToken *string `json:"DomainProxyToken,omitempty" xml:"DomainProxyToken,omitempty"`
	// The ID of the proxy token of the domain.
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

func (s ListDomainProxyTokensResponseBodyDomainProxyTokens) String() string {
	return dara.Prettify(s)
}

func (s ListDomainProxyTokensResponseBodyDomainProxyTokens) GoString() string {
	return s.String()
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) GetDomainId() *string {
	return s.DomainId
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) GetDomainProxyToken() *string {
	return s.DomainProxyToken
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) GetDomainProxyTokenId() *string {
	return s.DomainProxyTokenId
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) GetLastUsedTime() *int64 {
	return s.LastUsedTime
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) GetStatus() *string {
	return s.Status
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) SetCreateTime(v int64) *ListDomainProxyTokensResponseBodyDomainProxyTokens {
	s.CreateTime = &v
	return s
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) SetDomainId(v string) *ListDomainProxyTokensResponseBodyDomainProxyTokens {
	s.DomainId = &v
	return s
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) SetDomainProxyToken(v string) *ListDomainProxyTokensResponseBodyDomainProxyTokens {
	s.DomainProxyToken = &v
	return s
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) SetDomainProxyTokenId(v string) *ListDomainProxyTokensResponseBodyDomainProxyTokens {
	s.DomainProxyTokenId = &v
	return s
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) SetInstanceId(v string) *ListDomainProxyTokensResponseBodyDomainProxyTokens {
	s.InstanceId = &v
	return s
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) SetLastUsedTime(v int64) *ListDomainProxyTokensResponseBodyDomainProxyTokens {
	s.LastUsedTime = &v
	return s
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) SetStatus(v string) *ListDomainProxyTokensResponseBodyDomainProxyTokens {
	s.Status = &v
	return s
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) SetUpdateTime(v int64) *ListDomainProxyTokensResponseBodyDomainProxyTokens {
	s.UpdateTime = &v
	return s
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) Validate() error {
	return dara.Validate(s)
}
