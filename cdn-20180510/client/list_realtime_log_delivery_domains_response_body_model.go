// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRealtimeLogDeliveryDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *ListRealtimeLogDeliveryDomainsResponseBodyContent) *ListRealtimeLogDeliveryDomainsResponseBody
	GetContent() *ListRealtimeLogDeliveryDomainsResponseBodyContent
	SetRequestId(v string) *ListRealtimeLogDeliveryDomainsResponseBody
	GetRequestId() *string
}

type ListRealtimeLogDeliveryDomainsResponseBody struct {
	// The information about the accelerated domain names.
	Content *ListRealtimeLogDeliveryDomainsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 95D5B69F-8AEC-419B-8F3A-612B35032B0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRealtimeLogDeliveryDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeLogDeliveryDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRealtimeLogDeliveryDomainsResponseBody) GetContent() *ListRealtimeLogDeliveryDomainsResponseBodyContent {
	return s.Content
}

func (s *ListRealtimeLogDeliveryDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRealtimeLogDeliveryDomainsResponseBody) SetContent(v *ListRealtimeLogDeliveryDomainsResponseBodyContent) *ListRealtimeLogDeliveryDomainsResponseBody {
	s.Content = v
	return s
}

func (s *ListRealtimeLogDeliveryDomainsResponseBody) SetRequestId(v string) *ListRealtimeLogDeliveryDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRealtimeLogDeliveryDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRealtimeLogDeliveryDomainsResponseBodyContent struct {
	Domains []*ListRealtimeLogDeliveryDomainsResponseBodyContentDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s ListRealtimeLogDeliveryDomainsResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeLogDeliveryDomainsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListRealtimeLogDeliveryDomainsResponseBodyContent) GetDomains() []*ListRealtimeLogDeliveryDomainsResponseBodyContentDomains {
	return s.Domains
}

func (s *ListRealtimeLogDeliveryDomainsResponseBodyContent) SetDomains(v []*ListRealtimeLogDeliveryDomainsResponseBodyContentDomains) *ListRealtimeLogDeliveryDomainsResponseBodyContent {
	s.Domains = v
	return s
}

func (s *ListRealtimeLogDeliveryDomainsResponseBodyContent) Validate() error {
	return dara.Validate(s)
}

type ListRealtimeLogDeliveryDomainsResponseBodyContentDomains struct {
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The status. Valid values:
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

func (s ListRealtimeLogDeliveryDomainsResponseBodyContentDomains) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeLogDeliveryDomainsResponseBodyContentDomains) GoString() string {
	return s.String()
}

func (s *ListRealtimeLogDeliveryDomainsResponseBodyContentDomains) GetDomainName() *string {
	return s.DomainName
}

func (s *ListRealtimeLogDeliveryDomainsResponseBodyContentDomains) GetStatus() *string {
	return s.Status
}

func (s *ListRealtimeLogDeliveryDomainsResponseBodyContentDomains) SetDomainName(v string) *ListRealtimeLogDeliveryDomainsResponseBodyContentDomains {
	s.DomainName = &v
	return s
}

func (s *ListRealtimeLogDeliveryDomainsResponseBodyContentDomains) SetStatus(v string) *ListRealtimeLogDeliveryDomainsResponseBodyContentDomains {
	s.Status = &v
	return s
}

func (s *ListRealtimeLogDeliveryDomainsResponseBodyContentDomains) Validate() error {
	return dara.Validate(s)
}
