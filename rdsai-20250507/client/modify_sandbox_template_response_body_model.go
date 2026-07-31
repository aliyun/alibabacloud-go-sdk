// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySandboxTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *ModifySandboxTemplateResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *ModifySandboxTemplateResponseBody
	GetRequestId() *string
}

type ModifySandboxTemplateResponseBody struct {
	// The instance ID of the AI application.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Id of the request
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySandboxTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySandboxTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySandboxTemplateResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifySandboxTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySandboxTemplateResponseBody) SetInstanceName(v string) *ModifySandboxTemplateResponseBody {
	s.InstanceName = &v
	return s
}

func (s *ModifySandboxTemplateResponseBody) SetRequestId(v string) *ModifySandboxTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySandboxTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
