// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefaultLevelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDefaultLevelResponseBody
	GetRequestId() *string
}

type ModifyDefaultLevelResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 208B016D-4CB9-4A85-96A5-0B8ED1EBF271
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefaultLevelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefaultLevelResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefaultLevelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDefaultLevelResponseBody) SetRequestId(v string) *ModifyDefaultLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDefaultLevelResponseBody) Validate() error {
	return dara.Validate(s)
}
