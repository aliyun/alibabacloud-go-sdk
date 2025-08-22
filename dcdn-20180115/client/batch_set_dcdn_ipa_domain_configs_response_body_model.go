// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetDcdnIpaDomainConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchSetDcdnIpaDomainConfigsResponseBody
	GetRequestId() *string
}

type BatchSetDcdnIpaDomainConfigsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchSetDcdnIpaDomainConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchSetDcdnIpaDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSetDcdnIpaDomainConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchSetDcdnIpaDomainConfigsResponseBody) SetRequestId(v string) *BatchSetDcdnIpaDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchSetDcdnIpaDomainConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}
