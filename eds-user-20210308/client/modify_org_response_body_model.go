// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOrgResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyOrgResponseBody
	GetRequestId() *string
}

type ModifyOrgResponseBody struct {
	// example:
	//
	// 0296EDF8-3C8A-5128-8682-27B29C99****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOrgResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyOrgResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOrgResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyOrgResponseBody) SetRequestId(v string) *ModifyOrgResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyOrgResponseBody) Validate() error {
	return dara.Validate(s)
}
