// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyParametersResponseBody
	GetRequestId() *string
}

type ModifyParametersResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 36923CC2-DDAB-4B48-A144-DA92C1E19537
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyParametersResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyParametersResponseBody) SetRequestId(v string) *ModifyParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyParametersResponseBody) Validate() error {
	return dara.Validate(s)
}
