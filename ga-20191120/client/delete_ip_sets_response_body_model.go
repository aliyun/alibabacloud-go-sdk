// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpSetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIpSetsResponseBody
	GetRequestId() *string
}

type DeleteIpSetsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B49B60F6-F6C8-49E5-B06B-E33E3A469A92
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpSetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpSetsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpSetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIpSetsResponseBody) SetRequestId(v string) *DeleteIpSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIpSetsResponseBody) Validate() error {
	return dara.Validate(s)
}
