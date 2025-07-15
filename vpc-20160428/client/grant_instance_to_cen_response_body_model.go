// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantInstanceToCenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GrantInstanceToCenResponseBody
	GetRequestId() *string
}

type GrantInstanceToCenResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantInstanceToCenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantInstanceToCenResponseBody) GoString() string {
	return s.String()
}

func (s *GrantInstanceToCenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantInstanceToCenResponseBody) SetRequestId(v string) *GrantInstanceToCenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantInstanceToCenResponseBody) Validate() error {
	return dara.Validate(s)
}
