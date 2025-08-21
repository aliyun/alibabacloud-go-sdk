// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKeyPairsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteKeyPairsResponseBody
	GetRequestId() *string
}

type DeleteKeyPairsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteKeyPairsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteKeyPairsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKeyPairsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteKeyPairsResponseBody) SetRequestId(v string) *DeleteKeyPairsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKeyPairsResponseBody) Validate() error {
	return dara.Validate(s)
}
