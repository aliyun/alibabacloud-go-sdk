// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRulesResponseBody
	GetRequestId() *string
}

type DeleteRulesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9DEC9C28-AB05-4DDF-9A78-6B08EC9CE18C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRulesResponseBody) SetRequestId(v string) *DeleteRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
