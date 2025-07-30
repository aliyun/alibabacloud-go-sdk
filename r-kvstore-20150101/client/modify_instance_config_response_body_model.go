// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceConfigResponseBody
	GetRequestId() *string
}

type ModifyInstanceConfigResponseBody struct {
	// 请求ID。
	//
	// example:
	//
	// 8D0C0AFC-E9CD-47A4-8395-5C31BF9B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceConfigResponseBody) SetRequestId(v string) *ModifyInstanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
