// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomFieldResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFieldId(v string) *CreateCustomFieldResponseBody
	GetFieldId() *string
	SetRequestId(v string) *CreateCustomFieldResponseBody
	GetRequestId() *string
}

type CreateCustomFieldResponseBody struct {
	// example:
	//
	// ufd_001
	FieldId *string `json:"FieldId,omitempty" xml:"FieldId,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomFieldResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomFieldResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomFieldResponseBody) GetFieldId() *string {
	return s.FieldId
}

func (s *CreateCustomFieldResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomFieldResponseBody) SetFieldId(v string) *CreateCustomFieldResponseBody {
	s.FieldId = &v
	return s
}

func (s *CreateCustomFieldResponseBody) SetRequestId(v string) *CreateCustomFieldResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomFieldResponseBody) Validate() error {
	return dara.Validate(s)
}
