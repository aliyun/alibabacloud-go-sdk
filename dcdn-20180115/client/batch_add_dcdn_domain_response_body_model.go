// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddDcdnDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchAddDcdnDomainResponseBody
	GetRequestId() *string
}

type BatchAddDcdnDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchAddDcdnDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchAddDcdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAddDcdnDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchAddDcdnDomainResponseBody) SetRequestId(v string) *BatchAddDcdnDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchAddDcdnDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
