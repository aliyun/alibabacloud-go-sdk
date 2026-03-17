// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDnatEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDnatEntryResponseBody
	GetRequestId() *string
}

type DeleteDnatEntryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 68CE10C0-2EFF-4B82-9907-10AB7E2B0A6C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDnatEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDnatEntryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDnatEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDnatEntryResponseBody) SetRequestId(v string) *DeleteDnatEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDnatEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
