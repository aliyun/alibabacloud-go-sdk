// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagManagementPortResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySagManagementPortResponseBody
	GetRequestId() *string
}

type ModifySagManagementPortResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6877D55B-08F7-4DA3-916B-32A6FD402E06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySagManagementPortResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySagManagementPortResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySagManagementPortResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySagManagementPortResponseBody) SetRequestId(v string) *ModifySagManagementPortResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySagManagementPortResponseBody) Validate() error {
	return dara.Validate(s)
}
