// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStopCdnDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchStopCdnDomainResponseBody
	GetRequestId() *string
}

type BatchStopCdnDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 324AEFFF-308C-4DA7-8CD3-01B277B98F28
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchStopCdnDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchStopCdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStopCdnDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchStopCdnDomainResponseBody) SetRequestId(v string) *BatchStopCdnDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchStopCdnDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
