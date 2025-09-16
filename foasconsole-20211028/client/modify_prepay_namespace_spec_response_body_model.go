// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPrepayNamespaceSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyPrepayNamespaceSpecResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyPrepayNamespaceSpecResponseBody
	GetSuccess() *bool
}

type ModifyPrepayNamespaceSpecResponseBody struct {
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyPrepayNamespaceSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayNamespaceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPrepayNamespaceSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPrepayNamespaceSpecResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyPrepayNamespaceSpecResponseBody) SetRequestId(v string) *ModifyPrepayNamespaceSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecResponseBody) SetSuccess(v bool) *ModifyPrepayNamespaceSpecResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
