// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSearchConditionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSearchConditionResponseBody
	GetRequestId() *string
}

type DeleteSearchConditionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2983C540-E51F-582A-B510-732C27CD914C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSearchConditionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSearchConditionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSearchConditionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSearchConditionResponseBody) SetRequestId(v string) *DeleteSearchConditionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSearchConditionResponseBody) Validate() error {
	return dara.Validate(s)
}
