// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQuotaTemplateServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyQuotaTemplateServiceStatusResponseBody
	GetRequestId() *string
	SetTemplateServiceStatus(v *ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) *ModifyQuotaTemplateServiceStatusResponseBody
	GetTemplateServiceStatus() *ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus
}

type ModifyQuotaTemplateServiceStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D47B3A10-CDAC-5412-B2EE-EC9A3DBE9053
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the quota template.
	TemplateServiceStatus *ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus `json:"TemplateServiceStatus,omitempty" xml:"TemplateServiceStatus,omitempty" type:"Struct"`
}

func (s ModifyQuotaTemplateServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyQuotaTemplateServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyQuotaTemplateServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyQuotaTemplateServiceStatusResponseBody) GetTemplateServiceStatus() *ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus {
	return s.TemplateServiceStatus
}

func (s *ModifyQuotaTemplateServiceStatusResponseBody) SetRequestId(v string) *ModifyQuotaTemplateServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyQuotaTemplateServiceStatusResponseBody) SetTemplateServiceStatus(v *ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) *ModifyQuotaTemplateServiceStatusResponseBody {
	s.TemplateServiceStatus = v
	return s
}

func (s *ModifyQuotaTemplateServiceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus struct {
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-pG****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the quota template. Valid values:
	//
	// 	- \\-1: The quota template is disabled.
	//
	// 	- 1: The quota template is enabled.
	//
	// example:
	//
	// 1
	ServiceStatus *int32 `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) String() string {
	return dara.Prettify(s)
}

func (s ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) GoString() string {
	return s.String()
}

func (s *ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) GetServiceStatus() *int32 {
	return s.ServiceStatus
}

func (s *ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) SetResourceDirectoryId(v string) *ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) SetServiceStatus(v int32) *ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus {
	s.ServiceStatus = &v
	return s
}

func (s *ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) Validate() error {
	return dara.Validate(s)
}
