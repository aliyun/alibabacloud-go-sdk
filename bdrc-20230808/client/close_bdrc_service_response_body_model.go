// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseBdrcServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CloseBdrcServiceResponseBody
	GetRequestId() *string
}

type CloseBdrcServiceResponseBody struct {
	// The unique identifier of the request.
	//
	// example:
	//
	// 663D8898-E0B5-5964-BF28-A191CE6A1825
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloseBdrcServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseBdrcServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CloseBdrcServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseBdrcServiceResponseBody) SetRequestId(v string) *CloseBdrcServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseBdrcServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
