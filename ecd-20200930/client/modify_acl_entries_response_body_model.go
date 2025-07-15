// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAclEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAclEntriesResponseBody
	GetRequestId() *string
}

type ModifyAclEntriesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 062B1439-709A-580E-85DF-CE97A1560565
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAclEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAclEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAclEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAclEntriesResponseBody) SetRequestId(v string) *ModifyAclEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAclEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}
