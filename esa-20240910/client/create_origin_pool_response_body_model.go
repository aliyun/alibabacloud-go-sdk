// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOriginPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateOriginPoolResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateOriginPoolResponseBody
	GetRequestId() *string
}

type CreateOriginPoolResponseBody struct {
	// The ID of the newly created origin address pool.
	//
	// example:
	//
	// 103852052519****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-A198-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOriginPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOriginPoolResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOriginPoolResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateOriginPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOriginPoolResponseBody) SetId(v int64) *CreateOriginPoolResponseBody {
	s.Id = &v
	return s
}

func (s *CreateOriginPoolResponseBody) SetRequestId(v string) *CreateOriginPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOriginPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
