// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolarFsInstanceDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyPolarFsInstanceDescriptionResponseBody
	GetRequestId() *string
}

type ModifyPolarFsInstanceDescriptionResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 24A1990B-4F6E-482B-B8CB-75C612******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPolarFsInstanceDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolarFsInstanceDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPolarFsInstanceDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPolarFsInstanceDescriptionResponseBody) SetRequestId(v string) *ModifyPolarFsInstanceDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPolarFsInstanceDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
