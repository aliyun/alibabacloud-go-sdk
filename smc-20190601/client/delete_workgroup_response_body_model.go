// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkgroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWorkgroupResponseBody
	GetRequestId() *string
}

type DeleteWorkgroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 410E6073-66D0-45D3-AB3E-4DC3F5E4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWorkgroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkgroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkgroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkgroupResponseBody) SetRequestId(v string) *DeleteWorkgroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkgroupResponseBody) Validate() error {
	return dara.Validate(s)
}
