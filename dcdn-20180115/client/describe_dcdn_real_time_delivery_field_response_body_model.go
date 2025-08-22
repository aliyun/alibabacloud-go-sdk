// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnRealTimeDeliveryFieldResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *DescribeDcdnRealTimeDeliveryFieldResponseBodyContent) *DescribeDcdnRealTimeDeliveryFieldResponseBody
	GetContent() *DescribeDcdnRealTimeDeliveryFieldResponseBodyContent
	SetRequestId(v string) *DescribeDcdnRealTimeDeliveryFieldResponseBody
	GetRequestId() *string
}

type DescribeDcdnRealTimeDeliveryFieldResponseBody struct {
	// The returned results.
	Content *DescribeDcdnRealTimeDeliveryFieldResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 3EACD23C-F49F-4BF7-B9AD-C2CD3BA888C4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnRealTimeDeliveryFieldResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRealTimeDeliveryFieldResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponseBody) GetContent() *DescribeDcdnRealTimeDeliveryFieldResponseBodyContent {
	return s.Content
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponseBody) SetContent(v *DescribeDcdnRealTimeDeliveryFieldResponseBodyContent) *DescribeDcdnRealTimeDeliveryFieldResponseBody {
	s.Content = v
	return s
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponseBody) SetRequestId(v string) *DescribeDcdnRealTimeDeliveryFieldResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnRealTimeDeliveryFieldResponseBodyContent struct {
	Fields []*DescribeDcdnRealTimeDeliveryFieldResponseBodyContentFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
}

func (s DescribeDcdnRealTimeDeliveryFieldResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRealTimeDeliveryFieldResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponseBodyContent) GetFields() []*DescribeDcdnRealTimeDeliveryFieldResponseBodyContentFields {
	return s.Fields
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponseBodyContent) SetFields(v []*DescribeDcdnRealTimeDeliveryFieldResponseBodyContentFields) *DescribeDcdnRealTimeDeliveryFieldResponseBodyContent {
	s.Fields = v
	return s
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponseBodyContent) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnRealTimeDeliveryFieldResponseBodyContentFields struct {
	// The description of the field.
	//
	// example:
	//
	// Access time
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the field. For more information about fields in real-time log entries, see [Fields in a real-time log](https://help.aliyun.com/document_detail/324199.html).
	//
	// example:
	//
	// unixtime
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
}

func (s DescribeDcdnRealTimeDeliveryFieldResponseBodyContentFields) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRealTimeDeliveryFieldResponseBodyContentFields) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponseBodyContentFields) GetDescription() *string {
	return s.Description
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponseBodyContentFields) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponseBodyContentFields) SetDescription(v string) *DescribeDcdnRealTimeDeliveryFieldResponseBodyContentFields {
	s.Description = &v
	return s
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponseBodyContentFields) SetFieldName(v string) *DescribeDcdnRealTimeDeliveryFieldResponseBodyContentFields {
	s.FieldName = &v
	return s
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponseBodyContentFields) Validate() error {
	return dara.Validate(s)
}
