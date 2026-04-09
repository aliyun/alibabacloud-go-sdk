// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateParameterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateParameterResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateParameterResponseBody
	GetRequestId() *string
}

type CreateParameterResponseBody struct {
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateParameterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateParameterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateParameterResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateParameterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateParameterResponseBody) SetId(v int64) *CreateParameterResponseBody {
	s.Id = &v
	return s
}

func (s *CreateParameterResponseBody) SetRequestId(v string) *CreateParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateParameterResponseBody) Validate() error {
	return dara.Validate(s)
}
