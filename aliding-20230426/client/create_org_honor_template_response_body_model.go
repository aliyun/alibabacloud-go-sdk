// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrgHonorTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHonorId(v string) *CreateOrgHonorTemplateResponseBody
	GetHonorId() *string
	SetRequestId(v string) *CreateOrgHonorTemplateResponseBody
	GetRequestId() *string
}

type CreateOrgHonorTemplateResponseBody struct {
	HonorId   *string `json:"honorId,omitempty" xml:"honorId,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateOrgHonorTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrgHonorTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrgHonorTemplateResponseBody) GetHonorId() *string {
	return s.HonorId
}

func (s *CreateOrgHonorTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrgHonorTemplateResponseBody) SetHonorId(v string) *CreateOrgHonorTemplateResponseBody {
	s.HonorId = &v
	return s
}

func (s *CreateOrgHonorTemplateResponseBody) SetRequestId(v string) *CreateOrgHonorTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrgHonorTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
