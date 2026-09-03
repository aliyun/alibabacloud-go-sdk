// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKBSyncLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteKBSyncLinkResponseBody
	GetRequestId() *string
}

type DeleteKBSyncLinkResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 69A85BAF-1089-4CDF-A82F-0A140F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteKBSyncLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteKBSyncLinkResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKBSyncLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteKBSyncLinkResponseBody) SetRequestId(v string) *DeleteKBSyncLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKBSyncLinkResponseBody) Validate() error {
	return dara.Validate(s)
}
