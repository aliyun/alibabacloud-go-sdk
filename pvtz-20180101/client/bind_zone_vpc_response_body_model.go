// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindZoneVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindZoneVpcResponseBody
	GetRequestId() *string
}

type BindZoneVpcResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 46973D4C-E3E4-4ABA-9190-9A9DE406C7E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindZoneVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindZoneVpcResponseBody) GoString() string {
	return s.String()
}

func (s *BindZoneVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindZoneVpcResponseBody) SetRequestId(v string) *BindZoneVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindZoneVpcResponseBody) Validate() error {
	return dara.Validate(s)
}
