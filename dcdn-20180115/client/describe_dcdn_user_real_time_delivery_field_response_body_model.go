// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserRealTimeDeliveryFieldResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContent) *DescribeDcdnUserRealTimeDeliveryFieldResponseBody
	GetContent() *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContent
	SetRequestId(v string) *DescribeDcdnUserRealTimeDeliveryFieldResponseBody
	GetRequestId() *string
}

type DescribeDcdnUserRealTimeDeliveryFieldResponseBody struct {
	// The data returned.
	Content *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 3EACD23C-F49F-4BF7-B9AD-C2CD3BA888C4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnUserRealTimeDeliveryFieldResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserRealTimeDeliveryFieldResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponseBody) GetContent() *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContent {
	return s.Content
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponseBody) SetContent(v *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContent) *DescribeDcdnUserRealTimeDeliveryFieldResponseBody {
	s.Content = v
	return s
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponseBody) SetRequestId(v string) *DescribeDcdnUserRealTimeDeliveryFieldResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContent struct {
	Fields []*DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContent) GetFields() []*DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields {
	return s.Fields
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContent) SetFields(v []*DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields) *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContent {
	s.Fields = v
	return s
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContent) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields struct {
	// The description of the field.
	//
	// example:
	//
	// The timestamp of the request
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the field.
	//
	// example:
	//
	// unixtime
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// Indicates whether the field was selected.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields) GetDescription() *string {
	return s.Description
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields) GetSelected() *bool {
	return s.Selected
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields) SetDescription(v string) *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields {
	s.Description = &v
	return s
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields) SetFieldName(v string) *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields {
	s.FieldName = &v
	return s
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields) SetSelected(v bool) *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields {
	s.Selected = &v
	return s
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields) Validate() error {
	return dara.Validate(s)
}
