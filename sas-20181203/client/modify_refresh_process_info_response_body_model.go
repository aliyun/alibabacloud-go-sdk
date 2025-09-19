// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRefreshProcessInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRefreshProcessInfoResponseBody
	GetRequestId() *string
}

type ModifyRefreshProcessInfoResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 814FCBBC-3A02-5555-8D05-F8D9FD62A295
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRefreshProcessInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRefreshProcessInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRefreshProcessInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRefreshProcessInfoResponseBody) SetRequestId(v string) *ModifyRefreshProcessInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRefreshProcessInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
