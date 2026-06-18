// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateQuotaResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateQuotaResponseBody
	GetMessage() *string
	SetQuota(v *Quota) *UpdateQuotaResponseBody
	GetQuota() *Quota
	SetRequestId(v string) *UpdateQuotaResponseBody
	GetRequestId() *string
}

type UpdateQuotaResponseBody struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	Quota     *Quota  `json:"quota,omitempty" xml:"quota,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQuotaResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateQuotaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateQuotaResponseBody) GetQuota() *Quota {
	return s.Quota
}

func (s *UpdateQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateQuotaResponseBody) SetCode(v string) *UpdateQuotaResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateQuotaResponseBody) SetMessage(v string) *UpdateQuotaResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateQuotaResponseBody) SetQuota(v *Quota) *UpdateQuotaResponseBody {
	s.Quota = v
	return s
}

func (s *UpdateQuotaResponseBody) SetRequestId(v string) *UpdateQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateQuotaResponseBody) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}
