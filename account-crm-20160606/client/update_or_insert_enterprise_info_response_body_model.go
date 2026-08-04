// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOrInsertEnterpriseInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateOrInsertEnterpriseInfoResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateOrInsertEnterpriseInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateOrInsertEnterpriseInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateOrInsertEnterpriseInfoResponseBody
	GetSuccess() *bool
}

type UpdateOrInsertEnterpriseInfoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateOrInsertEnterpriseInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOrInsertEnterpriseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOrInsertEnterpriseInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateOrInsertEnterpriseInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateOrInsertEnterpriseInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOrInsertEnterpriseInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateOrInsertEnterpriseInfoResponseBody) SetCode(v string) *UpdateOrInsertEnterpriseInfoResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoResponseBody) SetMessage(v string) *UpdateOrInsertEnterpriseInfoResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoResponseBody) SetRequestId(v string) *UpdateOrInsertEnterpriseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoResponseBody) SetSuccess(v bool) *UpdateOrInsertEnterpriseInfoResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
