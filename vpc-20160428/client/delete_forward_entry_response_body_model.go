// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteForwardEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteForwardEntryResponseBody
	GetRequestId() *string
}

type DeleteForwardEntryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteForwardEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteForwardEntryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteForwardEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteForwardEntryResponseBody) SetRequestId(v string) *DeleteForwardEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteForwardEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
