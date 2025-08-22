// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDcdnWafGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateDcdnWafGroupResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateDcdnWafGroupResponseBody
	GetRequestId() *string
}

type CreateDcdnWafGroupResponseBody struct {
	// The ID of the created WAF rule group.
	//
	// example:
	//
	// 30000166
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 675F4820-400E-5929-8B03-2C031A5D5391
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDcdnWafGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDcdnWafGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDcdnWafGroupResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateDcdnWafGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDcdnWafGroupResponseBody) SetId(v int64) *CreateDcdnWafGroupResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDcdnWafGroupResponseBody) SetRequestId(v string) *CreateDcdnWafGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDcdnWafGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
