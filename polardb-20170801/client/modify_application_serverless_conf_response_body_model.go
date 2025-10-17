// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationServerlessConfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ModifyApplicationServerlessConfResponseBody
	GetApplicationId() *string
	SetRequestId(v string) *ModifyApplicationServerlessConfResponseBody
	GetRequestId() *string
}

type ModifyApplicationServerlessConfResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD86******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApplicationServerlessConfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationServerlessConfResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApplicationServerlessConfResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ModifyApplicationServerlessConfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApplicationServerlessConfResponseBody) SetApplicationId(v string) *ModifyApplicationServerlessConfResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *ModifyApplicationServerlessConfResponseBody) SetRequestId(v string) *ModifyApplicationServerlessConfResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApplicationServerlessConfResponseBody) Validate() error {
	return dara.Validate(s)
}
