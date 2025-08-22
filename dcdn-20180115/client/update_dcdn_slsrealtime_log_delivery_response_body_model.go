// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDcdnSLSRealtimeLogDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContent) *UpdateDcdnSLSRealtimeLogDeliveryResponseBody
	GetContent() *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContent
	SetRequestId(v string) *UpdateDcdnSLSRealtimeLogDeliveryResponseBody
	GetRequestId() *string
}

type UpdateDcdnSLSRealtimeLogDeliveryResponseBody struct {
	// The configuration results of the domain name.
	Content *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// F32C57AA-7BF8-49AE-A2CC-9F42390F5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDcdnSLSRealtimeLogDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDcdnSLSRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBody) GetContent() *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	return s.Content
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBody) SetContent(v *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContent) *UpdateDcdnSLSRealtimeLogDeliveryResponseBody {
	s.Content = v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBody) SetRequestId(v string) *UpdateDcdnSLSRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContent struct {
	Domains []*UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContent) GetDomains() []*UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains {
	return s.Domains
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContent) SetDomains(v []*UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains) *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	s.Domains = v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContent) Validate() error {
	return dara.Validate(s)
}

type UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains struct {
	// The description of the returned result.
	//
	// example:
	//
	// created
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// cn
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Indicates whether the real-time log delivery project was successfully updated. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains) String() string {
	return dara.Prettify(s)
}

func (s UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains) GoString() string {
	return s.String()
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains) GetDesc() *string {
	return s.Desc
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains) GetRegion() *string {
	return s.Region
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains) GetStatus() *string {
	return s.Status
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains) SetDesc(v string) *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains {
	s.Desc = &v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains) SetDomainName(v string) *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains {
	s.DomainName = &v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains) SetRegion(v string) *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains {
	s.Region = &v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains) SetStatus(v string) *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains {
	s.Status = &v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains) Validate() error {
	return dara.Validate(s)
}
