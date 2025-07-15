// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySnatEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySnatEntryResponseBody
	GetRequestId() *string
}

type ModifySnatEntryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2315DEB7-5E92-423A-91F7-4C1EC9AD97C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySnatEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySnatEntryResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySnatEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySnatEntryResponseBody) SetRequestId(v string) *ModifySnatEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySnatEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
