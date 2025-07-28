// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseResourceXffResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDefenseResourceXffResponseBody
	GetRequestId() *string
}

type ModifyDefenseResourceXffResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6C094583-9B3F-5BD8-8748-DC638E****BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefenseResourceXffResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseResourceXffResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefenseResourceXffResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDefenseResourceXffResponseBody) SetRequestId(v string) *ModifyDefenseResourceXffResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDefenseResourceXffResponseBody) Validate() error {
	return dara.Validate(s)
}
