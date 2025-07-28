// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApisecEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyApisecEventsResponseBody
	GetRequestId() *string
}

type ModifyApisecEventsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D7861F61-5B61-****-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApisecEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApisecEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApisecEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApisecEventsResponseBody) SetRequestId(v string) *ModifyApisecEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApisecEventsResponseBody) Validate() error {
	return dara.Validate(s)
}
