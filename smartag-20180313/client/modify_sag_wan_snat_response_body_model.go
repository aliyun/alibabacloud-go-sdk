// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagWanSnatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySagWanSnatResponseBody
	GetRequestId() *string
}

type ModifySagWanSnatResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 96AF7326-B6DE-4188-8638-56A6164F62D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySagWanSnatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySagWanSnatResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySagWanSnatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySagWanSnatResponseBody) SetRequestId(v string) *ModifySagWanSnatResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySagWanSnatResponseBody) Validate() error {
	return dara.Validate(s)
}
