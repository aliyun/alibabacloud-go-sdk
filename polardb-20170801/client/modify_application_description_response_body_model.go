// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ModifyApplicationDescriptionResponseBody
	GetApplicationId() *string
	SetRequestId(v string) *ModifyApplicationDescriptionResponseBody
	GetRequestId() *string
}

type ModifyApplicationDescriptionResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApplicationDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApplicationDescriptionResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ModifyApplicationDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApplicationDescriptionResponseBody) SetApplicationId(v string) *ModifyApplicationDescriptionResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *ModifyApplicationDescriptionResponseBody) SetRequestId(v string) *ModifyApplicationDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApplicationDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
