// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveRealtimeLogDeliveryDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *ListLiveRealtimeLogDeliveryDomainsResponseBodyContent) *ListLiveRealtimeLogDeliveryDomainsResponseBody
	GetContent() *ListLiveRealtimeLogDeliveryDomainsResponseBodyContent
	SetRequestId(v string) *ListLiveRealtimeLogDeliveryDomainsResponseBody
	GetRequestId() *string
}

type ListLiveRealtimeLogDeliveryDomainsResponseBody struct {
	// The domain names.
	Content *ListLiveRealtimeLogDeliveryDomainsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 95D5B69F-8AEC-419B-8F3A-612B35032B0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLiveRealtimeLogDeliveryDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponseBody) GetContent() *ListLiveRealtimeLogDeliveryDomainsResponseBodyContent {
	return s.Content
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponseBody) SetContent(v *ListLiveRealtimeLogDeliveryDomainsResponseBodyContent) *ListLiveRealtimeLogDeliveryDomainsResponseBody {
	s.Content = v
	return s
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponseBody) SetRequestId(v string) *ListLiveRealtimeLogDeliveryDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLiveRealtimeLogDeliveryDomainsResponseBodyContent struct {
	Domains []*ListLiveRealtimeLogDeliveryDomainsResponseBodyContentDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s ListLiveRealtimeLogDeliveryDomainsResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryDomainsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponseBodyContent) GetDomains() []*ListLiveRealtimeLogDeliveryDomainsResponseBodyContentDomains {
	return s.Domains
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponseBodyContent) SetDomains(v []*ListLiveRealtimeLogDeliveryDomainsResponseBodyContentDomains) *ListLiveRealtimeLogDeliveryDomainsResponseBodyContent {
	s.Domains = v
	return s
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponseBodyContent) Validate() error {
	return dara.Validate(s)
}

type ListLiveRealtimeLogDeliveryDomainsResponseBodyContentDomains struct {
	// The streaming domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The status of real-time log delivery. Valid values:
	//
	// 	- **online**: enabled
	//
	// 	- **offline**: disabled
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLiveRealtimeLogDeliveryDomainsResponseBodyContentDomains) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryDomainsResponseBodyContentDomains) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponseBodyContentDomains) GetDomainName() *string {
	return s.DomainName
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponseBodyContentDomains) GetStatus() *string {
	return s.Status
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponseBodyContentDomains) SetDomainName(v string) *ListLiveRealtimeLogDeliveryDomainsResponseBodyContentDomains {
	s.DomainName = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponseBodyContentDomains) SetStatus(v string) *ListLiveRealtimeLogDeliveryDomainsResponseBodyContentDomains {
	s.Status = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponseBodyContentDomains) Validate() error {
	return dara.Validate(s)
}
