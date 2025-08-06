// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStopVodDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchStopVodDomainResponseBody
	GetRequestId() *string
}

type BatchStopVodDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-****-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchStopVodDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchStopVodDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStopVodDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchStopVodDomainResponseBody) SetRequestId(v string) *BatchStopVodDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchStopVodDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
