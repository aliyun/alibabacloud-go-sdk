// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySearchParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySearchParametersResponseBody
	GetRequestId() *string
}

type ModifySearchParametersResponseBody struct {
	// The request ID. A successful response indicates that the parameter modification task has been accepted, but does not indicate that the parameters have taken effect.
	//
	// example:
	//
	// e3a24663-eee5-4009-b635-2e9e8d052e74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySearchParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySearchParametersResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySearchParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySearchParametersResponseBody) SetRequestId(v string) *ModifySearchParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySearchParametersResponseBody) Validate() error {
	return dara.Validate(s)
}
