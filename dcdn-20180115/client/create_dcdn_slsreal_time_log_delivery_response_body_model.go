// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDcdnSLSRealTimeLogDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContent) *CreateDcdnSLSRealTimeLogDeliveryResponseBody
	GetContent() *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContent
	SetRequestId(v string) *CreateDcdnSLSRealTimeLogDeliveryResponseBody
	GetRequestId() *string
}

type CreateDcdnSLSRealTimeLogDeliveryResponseBody struct {
	// The configuration results of the domain name.
	Content *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// F32C57AA-7BF8-49AE-A2CC-9F42390F5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDcdnSLSRealTimeLogDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDcdnSLSRealTimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBody) GetContent() *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContent {
	return s.Content
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBody) SetContent(v *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContent) *CreateDcdnSLSRealTimeLogDeliveryResponseBody {
	s.Content = v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBody) SetRequestId(v string) *CreateDcdnSLSRealTimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateDcdnSLSRealTimeLogDeliveryResponseBodyContent struct {
	Domains []*CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s CreateDcdnSLSRealTimeLogDeliveryResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s CreateDcdnSLSRealTimeLogDeliveryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContent) GetDomains() []*CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains {
	return s.Domains
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContent) SetDomains(v []*CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains) *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContent {
	s.Domains = v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContent) Validate() error {
	return dara.Validate(s)
}

type CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains struct {
	// The description of the returned result.
	//
	// example:
	//
	// ok
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The domain name from which real-time logs were collected.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The region to which real-time logs were delivered.
	//
	// example:
	//
	// cn
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The status of real-time logs.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains) String() string {
	return dara.Prettify(s)
}

func (s CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains) GoString() string {
	return s.String()
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains) GetDesc() *string {
	return s.Desc
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains) GetRegion() *string {
	return s.Region
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains) GetStatus() *string {
	return s.Status
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains) SetDesc(v string) *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains {
	s.Desc = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains) SetDomainName(v string) *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains {
	s.DomainName = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains) SetRegion(v string) *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains {
	s.Region = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains) SetStatus(v string) *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains {
	s.Status = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains) Validate() error {
	return dara.Validate(s)
}
