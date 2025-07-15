// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetShowListBackgroundResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetShowListBackgroundResponseBody
	GetRequestId() *string
}

type SetShowListBackgroundResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 987DA143-A39C-5B5D-AF5B-3B07944A0036
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetShowListBackgroundResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetShowListBackgroundResponseBody) GoString() string {
	return s.String()
}

func (s *SetShowListBackgroundResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetShowListBackgroundResponseBody) SetRequestId(v string) *SetShowListBackgroundResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetShowListBackgroundResponseBody) Validate() error {
	return dara.Validate(s)
}
