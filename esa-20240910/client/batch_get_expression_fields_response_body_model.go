// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetExpressionFieldsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFields(v []*BatchGetExpressionFieldsResponseBodyFields) *BatchGetExpressionFieldsResponseBody
	GetFields() []*BatchGetExpressionFieldsResponseBodyFields
	SetRequestId(v string) *BatchGetExpressionFieldsResponseBody
	GetRequestId() *string
}

type BatchGetExpressionFieldsResponseBody struct {
	// List of match fields.
	Fields []*BatchGetExpressionFieldsResponseBodyFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchGetExpressionFieldsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchGetExpressionFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetExpressionFieldsResponseBody) GetFields() []*BatchGetExpressionFieldsResponseBodyFields {
	return s.Fields
}

func (s *BatchGetExpressionFieldsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchGetExpressionFieldsResponseBody) SetFields(v []*BatchGetExpressionFieldsResponseBodyFields) *BatchGetExpressionFieldsResponseBody {
	s.Fields = v
	return s
}

func (s *BatchGetExpressionFieldsResponseBody) SetRequestId(v string) *BatchGetExpressionFieldsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetExpressionFieldsResponseBody) Validate() error {
	return dara.Validate(s)
}

type BatchGetExpressionFieldsResponseBodyFields struct {
	// List of match fields for a single expression.
	Fields []*string `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// ID of the expression, corresponding to the ID in the input parameters.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s BatchGetExpressionFieldsResponseBodyFields) String() string {
	return dara.Prettify(s)
}

func (s BatchGetExpressionFieldsResponseBodyFields) GoString() string {
	return s.String()
}

func (s *BatchGetExpressionFieldsResponseBodyFields) GetFields() []*string {
	return s.Fields
}

func (s *BatchGetExpressionFieldsResponseBodyFields) GetId() *string {
	return s.Id
}

func (s *BatchGetExpressionFieldsResponseBodyFields) SetFields(v []*string) *BatchGetExpressionFieldsResponseBodyFields {
	s.Fields = v
	return s
}

func (s *BatchGetExpressionFieldsResponseBodyFields) SetId(v string) *BatchGetExpressionFieldsResponseBodyFields {
	s.Id = &v
	return s
}

func (s *BatchGetExpressionFieldsResponseBodyFields) Validate() error {
	return dara.Validate(s)
}
