// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodRealtimeLogDeliveryDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *ListVodRealtimeLogDeliveryDomainsResponseBodyContent) *ListVodRealtimeLogDeliveryDomainsResponseBody
	GetContent() *ListVodRealtimeLogDeliveryDomainsResponseBodyContent
	SetRequestId(v string) *ListVodRealtimeLogDeliveryDomainsResponseBody
	GetRequestId() *string
}

type ListVodRealtimeLogDeliveryDomainsResponseBody struct {
	Content   *ListVodRealtimeLogDeliveryDomainsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVodRealtimeLogDeliveryDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryDomainsResponseBody) GetContent() *ListVodRealtimeLogDeliveryDomainsResponseBodyContent {
	return s.Content
}

func (s *ListVodRealtimeLogDeliveryDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVodRealtimeLogDeliveryDomainsResponseBody) SetContent(v *ListVodRealtimeLogDeliveryDomainsResponseBodyContent) *ListVodRealtimeLogDeliveryDomainsResponseBody {
	s.Content = v
	return s
}

func (s *ListVodRealtimeLogDeliveryDomainsResponseBody) SetRequestId(v string) *ListVodRealtimeLogDeliveryDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListVodRealtimeLogDeliveryDomainsResponseBodyContent struct {
	Domains []*ListVodRealtimeLogDeliveryDomainsResponseBodyContentDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s ListVodRealtimeLogDeliveryDomainsResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryDomainsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryDomainsResponseBodyContent) GetDomains() []*ListVodRealtimeLogDeliveryDomainsResponseBodyContentDomains {
	return s.Domains
}

func (s *ListVodRealtimeLogDeliveryDomainsResponseBodyContent) SetDomains(v []*ListVodRealtimeLogDeliveryDomainsResponseBodyContentDomains) *ListVodRealtimeLogDeliveryDomainsResponseBodyContent {
	s.Domains = v
	return s
}

func (s *ListVodRealtimeLogDeliveryDomainsResponseBodyContent) Validate() error {
	return dara.Validate(s)
}

type ListVodRealtimeLogDeliveryDomainsResponseBodyContentDomains struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListVodRealtimeLogDeliveryDomainsResponseBodyContentDomains) String() string {
	return dara.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryDomainsResponseBodyContentDomains) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryDomainsResponseBodyContentDomains) GetDomainName() *string {
	return s.DomainName
}

func (s *ListVodRealtimeLogDeliveryDomainsResponseBodyContentDomains) GetStatus() *string {
	return s.Status
}

func (s *ListVodRealtimeLogDeliveryDomainsResponseBodyContentDomains) SetDomainName(v string) *ListVodRealtimeLogDeliveryDomainsResponseBodyContentDomains {
	s.DomainName = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryDomainsResponseBodyContentDomains) SetStatus(v string) *ListVodRealtimeLogDeliveryDomainsResponseBodyContentDomains {
	s.Status = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryDomainsResponseBodyContentDomains) Validate() error {
	return dara.Validate(s)
}
