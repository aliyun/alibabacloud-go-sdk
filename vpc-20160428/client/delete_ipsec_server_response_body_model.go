// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpsecServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIpsecServerResponseBody
	GetRequestId() *string
}

type DeleteIpsecServerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpsecServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpsecServerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpsecServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIpsecServerResponseBody) SetRequestId(v string) *DeleteIpsecServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIpsecServerResponseBody) Validate() error {
	return dara.Validate(s)
}
