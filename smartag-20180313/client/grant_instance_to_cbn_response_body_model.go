// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantInstanceToCbnResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GrantInstanceToCbnResponseBody
	GetRequestId() *string
}

type GrantInstanceToCbnResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 22840034-ADE9-41D8-A5DC-A7CF435CEE75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantInstanceToCbnResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantInstanceToCbnResponseBody) GoString() string {
	return s.String()
}

func (s *GrantInstanceToCbnResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantInstanceToCbnResponseBody) SetRequestId(v string) *GrantInstanceToCbnResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantInstanceToCbnResponseBody) Validate() error {
	return dara.Validate(s)
}
