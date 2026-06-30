// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSandboxTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *DeleteSandboxTemplateResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *DeleteSandboxTemplateResponseBody
	GetRequestId() *string
}

type DeleteSandboxTemplateResponseBody struct {
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSandboxTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSandboxTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSandboxTemplateResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DeleteSandboxTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSandboxTemplateResponseBody) SetInstanceName(v string) *DeleteSandboxTemplateResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DeleteSandboxTemplateResponseBody) SetRequestId(v string) *DeleteSandboxTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSandboxTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
