// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSandboxTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *CreateSandboxTemplateResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *CreateSandboxTemplateResponseBody
	GetRequestId() *string
}

type CreateSandboxTemplateResponseBody struct {
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSandboxTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSandboxTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSandboxTemplateResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateSandboxTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSandboxTemplateResponseBody) SetInstanceName(v string) *CreateSandboxTemplateResponseBody {
	s.InstanceName = &v
	return s
}

func (s *CreateSandboxTemplateResponseBody) SetRequestId(v string) *CreateSandboxTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSandboxTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
