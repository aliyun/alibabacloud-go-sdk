// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetVodDomainConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchSetVodDomainConfigsResponseBody
	GetRequestId() *string
}

type BatchSetVodDomainConfigsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-****-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchSetVodDomainConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchSetVodDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSetVodDomainConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchSetVodDomainConfigsResponseBody) SetRequestId(v string) *BatchSetVodDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchSetVodDomainConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}
