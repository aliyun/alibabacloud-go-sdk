// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaTemplateServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetQuotaTemplateServiceStatusResponseBody
	GetRequestId() *string
	SetTemplateServiceStatus(v *GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) *GetQuotaTemplateServiceStatusResponseBody
	GetTemplateServiceStatus() *GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus
}

type GetQuotaTemplateServiceStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D47B3A10-CDAC-5412-B2EE-EC9A3DBE9053
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the quota template.
	TemplateServiceStatus *GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus `json:"TemplateServiceStatus,omitempty" xml:"TemplateServiceStatus,omitempty" type:"Struct"`
}

func (s GetQuotaTemplateServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaTemplateServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaTemplateServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQuotaTemplateServiceStatusResponseBody) GetTemplateServiceStatus() *GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus {
	return s.TemplateServiceStatus
}

func (s *GetQuotaTemplateServiceStatusResponseBody) SetRequestId(v string) *GetQuotaTemplateServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQuotaTemplateServiceStatusResponseBody) SetTemplateServiceStatus(v *GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) *GetQuotaTemplateServiceStatusResponseBody {
	s.TemplateServiceStatus = v
	return s
}

func (s *GetQuotaTemplateServiceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus struct {
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
	// 	- 0: No quota template is specified.
	//
	// example:
	//
	// 1
	ServiceStatus *int32 `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) GoString() string {
	return s.String()
}

func (s *GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) GetServiceStatus() *int32 {
	return s.ServiceStatus
}

func (s *GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) SetResourceDirectoryId(v string) *GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus {
	s.ResourceDirectoryId = &v
	return s
}

func (s *GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) SetServiceStatus(v int32) *GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus {
	s.ServiceStatus = &v
	return s
}

func (s *GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) Validate() error {
	return dara.Validate(s)
}
