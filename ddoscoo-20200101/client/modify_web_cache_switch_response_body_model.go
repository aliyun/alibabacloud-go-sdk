// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebCacheSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWebCacheSwitchResponseBody
	GetRequestId() *string
}

type ModifyWebCacheSwitchResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6623EA1F-30FB-5BC8-BEC9-74D55F6F08F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebCacheSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebCacheSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebCacheSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWebCacheSwitchResponseBody) SetRequestId(v string) *ModifyWebCacheSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWebCacheSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
