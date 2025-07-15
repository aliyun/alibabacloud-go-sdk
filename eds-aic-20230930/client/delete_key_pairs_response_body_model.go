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
	// The ID of the request.
	//
	// example:
	//
	// 5C5CEF0A-D6E1-58D3-8750-67DB4F82****
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
