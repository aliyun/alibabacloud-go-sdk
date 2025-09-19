// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGroupPropertyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyGroupPropertyResponseBody
	GetRequestId() *string
}

type ModifyGroupPropertyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyGroupPropertyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyGroupPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGroupPropertyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyGroupPropertyResponseBody) SetRequestId(v string) *ModifyGroupPropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGroupPropertyResponseBody) Validate() error {
	return dara.Validate(s)
}
