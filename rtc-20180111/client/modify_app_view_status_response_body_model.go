// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppViewStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAppViewStatusResponseBody
	GetRequestId() *string
}

type ModifyAppViewStatusResponseBody struct {
	// example:
	//
	// 94D9A316-9750-5928-B18C-59DF182F6BF7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppViewStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppViewStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppViewStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppViewStatusResponseBody) SetRequestId(v string) *ModifyAppViewStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppViewStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
