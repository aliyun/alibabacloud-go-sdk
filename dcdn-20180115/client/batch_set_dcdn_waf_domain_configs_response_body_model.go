// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetDcdnWafDomainConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchSetDcdnWafDomainConfigsResponseBody
	GetRequestId() *string
}

type BatchSetDcdnWafDomainConfigsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-3C82-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchSetDcdnWafDomainConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchSetDcdnWafDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSetDcdnWafDomainConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchSetDcdnWafDomainConfigsResponseBody) SetRequestId(v string) *BatchSetDcdnWafDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchSetDcdnWafDomainConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}
