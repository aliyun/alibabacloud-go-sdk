// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIgnoreCheckItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *IgnoreCheckItemsResponseBody
	GetRequestId() *string
}

type IgnoreCheckItemsResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 11EBEC99-B4B5-542E-8C17-B87B624C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IgnoreCheckItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IgnoreCheckItemsResponseBody) GoString() string {
	return s.String()
}

func (s *IgnoreCheckItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IgnoreCheckItemsResponseBody) SetRequestId(v string) *IgnoreCheckItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *IgnoreCheckItemsResponseBody) Validate() error {
	return dara.Validate(s)
}
