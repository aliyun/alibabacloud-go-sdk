// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAppResponseBody
	GetRequestId() *string
}

type ModifyAppResponseBody struct {
	// example:
	//
	// 126D4DDD-05A5-49B1-B18C-39C4A929BFB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppResponseBody) SetRequestId(v string) *ModifyAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppResponseBody) Validate() error {
	return dara.Validate(s)
}
