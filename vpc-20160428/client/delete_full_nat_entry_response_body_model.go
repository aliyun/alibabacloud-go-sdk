// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFullNatEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFullNatEntryResponseBody
	GetRequestId() *string
}

type DeleteFullNatEntryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2BCC426F-A9F2-3F03-99D2-1E0D647236DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFullNatEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFullNatEntryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFullNatEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFullNatEntryResponseBody) SetRequestId(v string) *DeleteFullNatEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFullNatEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
