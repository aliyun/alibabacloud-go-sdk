// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDBClusterResponseBody
	GetRequestId() *string
}

type DeleteDBClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 05321590-BB65-4720-8CB6-8218E041CDD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBClusterResponseBody) SetRequestId(v string) *DeleteDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
