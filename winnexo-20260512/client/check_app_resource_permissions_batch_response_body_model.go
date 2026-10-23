// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAppResourcePermissionsBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckAppResourcePermissionsBatchResponseBody
	GetCode() *string
	SetMessage(v string) *CheckAppResourcePermissionsBatchResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckAppResourcePermissionsBatchResponseBody
	GetRequestId() *string
	SetResults(v []*CheckAppResourcePermissionsBatchResponseBodyResults) *CheckAppResourcePermissionsBatchResponseBody
	GetResults() []*CheckAppResourcePermissionsBatchResponseBodyResults
}

type CheckAppResourcePermissionsBatchResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error message. This field is empty if the request is successful.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The results, returned in the same order as the requested items.
	Results []*CheckAppResourcePermissionsBatchResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s CheckAppResourcePermissionsBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckAppResourcePermissionsBatchResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAppResourcePermissionsBatchResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckAppResourcePermissionsBatchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckAppResourcePermissionsBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckAppResourcePermissionsBatchResponseBody) GetResults() []*CheckAppResourcePermissionsBatchResponseBodyResults {
	return s.Results
}

func (s *CheckAppResourcePermissionsBatchResponseBody) SetCode(v string) *CheckAppResourcePermissionsBatchResponseBody {
	s.Code = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchResponseBody) SetMessage(v string) *CheckAppResourcePermissionsBatchResponseBody {
	s.Message = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchResponseBody) SetRequestId(v string) *CheckAppResourcePermissionsBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchResponseBody) SetResults(v []*CheckAppResourcePermissionsBatchResponseBodyResults) *CheckAppResourcePermissionsBatchResponseBody {
	s.Results = v
	return s
}

func (s *CheckAppResourcePermissionsBatchResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckAppResourcePermissionsBatchResponseBodyResults struct {
	// Indicates whether the user has the permission.
	//
	// example:
	//
	// true
	Allowed *bool `json:"allowed,omitempty" xml:"allowed,omitempty"`
	// The target permission.
	//
	// example:
	//
	// APP/VIEW
	Permission *string `json:"permission,omitempty" xml:"permission,omitempty"`
	// The dot-separated resource key whose first segment is the APP application key.
	//
	// example:
	//
	// app.demo.resource_1
	ResourceKey *string `json:"resourceKey,omitempty" xml:"resourceKey,omitempty"`
	// The resource type. Set the value to APP_RESOURCE.
	//
	// example:
	//
	// APP_RESOURCE
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s CheckAppResourcePermissionsBatchResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s CheckAppResourcePermissionsBatchResponseBodyResults) GoString() string {
	return s.String()
}

func (s *CheckAppResourcePermissionsBatchResponseBodyResults) GetAllowed() *bool {
	return s.Allowed
}

func (s *CheckAppResourcePermissionsBatchResponseBodyResults) GetPermission() *string {
	return s.Permission
}

func (s *CheckAppResourcePermissionsBatchResponseBodyResults) GetResourceKey() *string {
	return s.ResourceKey
}

func (s *CheckAppResourcePermissionsBatchResponseBodyResults) GetResourceType() *string {
	return s.ResourceType
}

func (s *CheckAppResourcePermissionsBatchResponseBodyResults) SetAllowed(v bool) *CheckAppResourcePermissionsBatchResponseBodyResults {
	s.Allowed = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchResponseBodyResults) SetPermission(v string) *CheckAppResourcePermissionsBatchResponseBodyResults {
	s.Permission = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchResponseBodyResults) SetResourceKey(v string) *CheckAppResourcePermissionsBatchResponseBodyResults {
	s.ResourceKey = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchResponseBodyResults) SetResourceType(v string) *CheckAppResourcePermissionsBatchResponseBodyResults {
	s.ResourceType = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
