// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateServiceInstanceNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ValidateServiceInstanceNameResponseBody
	GetRequestId() *string
}

type ValidateServiceInstanceNameResponseBody struct {
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ValidateServiceInstanceNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateServiceInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateServiceInstanceNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateServiceInstanceNameResponseBody) SetRequestId(v string) *ValidateServiceInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateServiceInstanceNameResponseBody) Validate() error {
	return dara.Validate(s)
}
