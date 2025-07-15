// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOpsItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteOpsItemsResponseBody
	GetRequestId() *string
}

type DeleteOpsItemsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DBA6E6C8-F75D-41DE-AFF5-1FA03F551CA3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOpsItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpsItemsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOpsItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOpsItemsResponseBody) SetRequestId(v string) *DeleteOpsItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOpsItemsResponseBody) Validate() error {
	return dara.Validate(s)
}
