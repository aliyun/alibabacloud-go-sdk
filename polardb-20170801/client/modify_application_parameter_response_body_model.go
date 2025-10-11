// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationParameterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ModifyApplicationParameterResponseBody
	GetApplicationId() *string
	SetRequestId(v string) *ModifyApplicationParameterResponseBody
	GetRequestId() *string
}

type ModifyApplicationParameterResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApplicationParameterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationParameterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApplicationParameterResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ModifyApplicationParameterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApplicationParameterResponseBody) SetApplicationId(v string) *ModifyApplicationParameterResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *ModifyApplicationParameterResponseBody) SetRequestId(v string) *ModifyApplicationParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApplicationParameterResponseBody) Validate() error {
	return dara.Validate(s)
}
