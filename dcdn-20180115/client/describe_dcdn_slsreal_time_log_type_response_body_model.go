// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnSLSRealTimeLogTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *DescribeDcdnSLSRealTimeLogTypeResponseBodyContent) *DescribeDcdnSLSRealTimeLogTypeResponseBody
	GetContent() *DescribeDcdnSLSRealTimeLogTypeResponseBodyContent
	SetRequestId(v string) *DescribeDcdnSLSRealTimeLogTypeResponseBody
	GetRequestId() *string
}

type DescribeDcdnSLSRealTimeLogTypeResponseBody struct {
	// The returned results.
	Content *DescribeDcdnSLSRealTimeLogTypeResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnSLSRealTimeLogTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSLSRealTimeLogTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSLSRealTimeLogTypeResponseBody) GetContent() *DescribeDcdnSLSRealTimeLogTypeResponseBodyContent {
	return s.Content
}

func (s *DescribeDcdnSLSRealTimeLogTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnSLSRealTimeLogTypeResponseBody) SetContent(v *DescribeDcdnSLSRealTimeLogTypeResponseBodyContent) *DescribeDcdnSLSRealTimeLogTypeResponseBody {
	s.Content = v
	return s
}

func (s *DescribeDcdnSLSRealTimeLogTypeResponseBody) SetRequestId(v string) *DescribeDcdnSLSRealTimeLogTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnSLSRealTimeLogTypeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnSLSRealTimeLogTypeResponseBodyContent struct {
	Business []*DescribeDcdnSLSRealTimeLogTypeResponseBodyContentBusiness `json:"Business,omitempty" xml:"Business,omitempty" type:"Repeated"`
}

func (s DescribeDcdnSLSRealTimeLogTypeResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSLSRealTimeLogTypeResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSLSRealTimeLogTypeResponseBodyContent) GetBusiness() []*DescribeDcdnSLSRealTimeLogTypeResponseBodyContentBusiness {
	return s.Business
}

func (s *DescribeDcdnSLSRealTimeLogTypeResponseBodyContent) SetBusiness(v []*DescribeDcdnSLSRealTimeLogTypeResponseBodyContentBusiness) *DescribeDcdnSLSRealTimeLogTypeResponseBodyContent {
	s.Business = v
	return s
}

func (s *DescribeDcdnSLSRealTimeLogTypeResponseBodyContent) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnSLSRealTimeLogTypeResponseBodyContentBusiness struct {
	// The type of real-time logs. Valid values:
	//
	// 	- **dcdn_log_access_l1**: access logs.
	//
	// 	- **dcdn_log_er**: EdgeRoutine logs
	//
	// 	- **dcdn_log_waf**: WAF interception logs
	//
	// example:
	//
	// dcdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The description of the real-time log type.
	//
	// example:
	//
	// product_U8JE
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
}

func (s DescribeDcdnSLSRealTimeLogTypeResponseBodyContentBusiness) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSLSRealTimeLogTypeResponseBodyContentBusiness) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSLSRealTimeLogTypeResponseBodyContentBusiness) GetBusinessType() *string {
	return s.BusinessType
}

func (s *DescribeDcdnSLSRealTimeLogTypeResponseBodyContentBusiness) GetDesc() *string {
	return s.Desc
}

func (s *DescribeDcdnSLSRealTimeLogTypeResponseBodyContentBusiness) SetBusinessType(v string) *DescribeDcdnSLSRealTimeLogTypeResponseBodyContentBusiness {
	s.BusinessType = &v
	return s
}

func (s *DescribeDcdnSLSRealTimeLogTypeResponseBodyContentBusiness) SetDesc(v string) *DescribeDcdnSLSRealTimeLogTypeResponseBodyContentBusiness {
	s.Desc = &v
	return s
}

func (s *DescribeDcdnSLSRealTimeLogTypeResponseBodyContentBusiness) Validate() error {
	return dara.Validate(s)
}
