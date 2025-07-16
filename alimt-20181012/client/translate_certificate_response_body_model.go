// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *TranslateCertificateResponseBodyData) *TranslateCertificateResponseBody
	GetData() *TranslateCertificateResponseBodyData
	SetRequestId(v string) *TranslateCertificateResponseBody
	GetRequestId() *string
}

type TranslateCertificateResponseBody struct {
	Data *TranslateCertificateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 6640DE48-201C-4110-AE87-FB8FA34412B9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TranslateCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TranslateCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateCertificateResponseBody) GetData() *TranslateCertificateResponseBodyData {
	return s.Data
}

func (s *TranslateCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TranslateCertificateResponseBody) SetData(v *TranslateCertificateResponseBodyData) *TranslateCertificateResponseBody {
	s.Data = v
	return s
}

func (s *TranslateCertificateResponseBody) SetRequestId(v string) *TranslateCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *TranslateCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

type TranslateCertificateResponseBodyData struct {
	TranslatedValues []*TranslateCertificateResponseBodyDataTranslatedValues `json:"TranslatedValues,omitempty" xml:"TranslatedValues,omitempty" type:"Repeated"`
}

func (s TranslateCertificateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TranslateCertificateResponseBodyData) GoString() string {
	return s.String()
}

func (s *TranslateCertificateResponseBodyData) GetTranslatedValues() []*TranslateCertificateResponseBodyDataTranslatedValues {
	return s.TranslatedValues
}

func (s *TranslateCertificateResponseBodyData) SetTranslatedValues(v []*TranslateCertificateResponseBodyDataTranslatedValues) *TranslateCertificateResponseBodyData {
	s.TranslatedValues = v
	return s
}

func (s *TranslateCertificateResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type TranslateCertificateResponseBodyDataTranslatedValues struct {
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// name
	KeyTranslation *string `json:"KeyTranslation,omitempty" xml:"KeyTranslation,omitempty"`
	Value          *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// Solemn
	ValueTranslation *string `json:"ValueTranslation,omitempty" xml:"ValueTranslation,omitempty"`
}

func (s TranslateCertificateResponseBodyDataTranslatedValues) String() string {
	return dara.Prettify(s)
}

func (s TranslateCertificateResponseBodyDataTranslatedValues) GoString() string {
	return s.String()
}

func (s *TranslateCertificateResponseBodyDataTranslatedValues) GetKey() *string {
	return s.Key
}

func (s *TranslateCertificateResponseBodyDataTranslatedValues) GetKeyTranslation() *string {
	return s.KeyTranslation
}

func (s *TranslateCertificateResponseBodyDataTranslatedValues) GetValue() *string {
	return s.Value
}

func (s *TranslateCertificateResponseBodyDataTranslatedValues) GetValueTranslation() *string {
	return s.ValueTranslation
}

func (s *TranslateCertificateResponseBodyDataTranslatedValues) SetKey(v string) *TranslateCertificateResponseBodyDataTranslatedValues {
	s.Key = &v
	return s
}

func (s *TranslateCertificateResponseBodyDataTranslatedValues) SetKeyTranslation(v string) *TranslateCertificateResponseBodyDataTranslatedValues {
	s.KeyTranslation = &v
	return s
}

func (s *TranslateCertificateResponseBodyDataTranslatedValues) SetValue(v string) *TranslateCertificateResponseBodyDataTranslatedValues {
	s.Value = &v
	return s
}

func (s *TranslateCertificateResponseBodyDataTranslatedValues) SetValueTranslation(v string) *TranslateCertificateResponseBodyDataTranslatedValues {
	s.ValueTranslation = &v
	return s
}

func (s *TranslateCertificateResponseBodyDataTranslatedValues) Validate() error {
	return dara.Validate(s)
}
