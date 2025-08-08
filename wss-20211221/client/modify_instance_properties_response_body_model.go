// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstancePropertiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstancePropertiesResponseBody
	GetRequestId() *string
}

type ModifyInstancePropertiesResponseBody struct {
	// example:
	//
	// 833C4D2C-09C7-5CE6-8159-06758B964970
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstancePropertiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstancePropertiesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstancePropertiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstancePropertiesResponseBody) SetRequestId(v string) *ModifyInstancePropertiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstancePropertiesResponseBody) Validate() error {
	return dara.Validate(s)
}
