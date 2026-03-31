// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRemediationTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListRemediationTemplatesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListRemediationTemplatesResponseBody
	GetPageSize() *int64
	SetRemediationTemplates(v []*ListRemediationTemplatesResponseBodyRemediationTemplates) *ListRemediationTemplatesResponseBody
	GetRemediationTemplates() []*ListRemediationTemplatesResponseBodyRemediationTemplates
	SetRequestId(v string) *ListRemediationTemplatesResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListRemediationTemplatesResponseBody
	GetTotalCount() *string
}

type ListRemediationTemplatesResponseBody struct {
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The queried remediation templates.
	RemediationTemplates []*ListRemediationTemplatesResponseBodyRemediationTemplates `json:"RemediationTemplates,omitempty" xml:"RemediationTemplates,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// FC2C4750-7024-499C-A69F-763543D1CBE3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of remediation templates.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRemediationTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRemediationTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRemediationTemplatesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListRemediationTemplatesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRemediationTemplatesResponseBody) GetRemediationTemplates() []*ListRemediationTemplatesResponseBodyRemediationTemplates {
	return s.RemediationTemplates
}

func (s *ListRemediationTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRemediationTemplatesResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListRemediationTemplatesResponseBody) SetPageNumber(v int64) *ListRemediationTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRemediationTemplatesResponseBody) SetPageSize(v int64) *ListRemediationTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRemediationTemplatesResponseBody) SetRemediationTemplates(v []*ListRemediationTemplatesResponseBodyRemediationTemplates) *ListRemediationTemplatesResponseBody {
	s.RemediationTemplates = v
	return s
}

func (s *ListRemediationTemplatesResponseBody) SetRequestId(v string) *ListRemediationTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRemediationTemplatesResponseBody) SetTotalCount(v string) *ListRemediationTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRemediationTemplatesResponseBody) Validate() error {
	if s.RemediationTemplates != nil {
		for _, item := range s.RemediationTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRemediationTemplatesResponseBodyRemediationTemplates struct {
	// The type of the remediation template. Valid value: OOS, which indicates Operation Orchestration Service.
	//
	// example:
	//
	// OOS
	RemediationType *string `json:"RemediationType,omitempty" xml:"RemediationType,omitempty"`
	// The definition of the remediation template.
	//
	// example:
	//
	// {
	//
	//     "Parameters": {
	//
	//         "bucketName": {
	//
	//             "Default": "{resourceId}",
	//
	//             "Description": {
	//
	//                 "zh-cn": "[Required]OSS Bucket Name.",
	//
	//                 "en": "[Required]OSS Bucket Name."
	//
	//             },
	//
	//             "Type": "String"
	//
	//         },
	//
	//         "regionId": {
	//
	//             "AssociationProperty": "RegionId",
	//
	//             "Default": "{regionId}",
	//
	//             "Description": {
	//
	//                 "zh-cn": "[Required]The id of target region.",
	//
	//                 "en": "[Required]The id of target region."
	//
	//             },
	//
	//             "Type": "String"
	//
	//         },
	//
	//         "permissionName": {
	//
	//             "AllowValues": "[\\"public-read-write\\",\\"public-read\\",\\"private\\"]",
	//
	//             "Default": "private",
	//
	//             "Description": {
	//
	//                 "zh-cn": "[Required]ACL Permission Name.",
	//
	//                 "en": "[Required]ACL Permission Name."
	//
	//             },
	//
	//             "Type": "String"
	//
	//         }
	//
	//     }
	//
	// }
	TemplateDefinition *string `json:"TemplateDefinition,omitempty" xml:"TemplateDefinition,omitempty"`
	// The description of the remediation template.
	//
	// example:
	//
	// Configure encryption rules for OSSBucket through the PutBucketEncryption interface. Be aware of the risks and exercise caution.
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	// The ID of the remediation template.
	//
	// example:
	//
	// ACS-OSS-PutBucketAcl
	TemplateIdentifier *string `json:"TemplateIdentifier,omitempty" xml:"TemplateIdentifier,omitempty"`
	// The name of the remediation template.
	//
	// example:
	//
	// Set the ACL of an OSS bucket to private
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListRemediationTemplatesResponseBodyRemediationTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListRemediationTemplatesResponseBodyRemediationTemplates) GoString() string {
	return s.String()
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) GetRemediationType() *string {
	return s.RemediationType
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) GetTemplateDefinition() *string {
	return s.TemplateDefinition
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) GetTemplateDescription() *string {
	return s.TemplateDescription
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) GetTemplateIdentifier() *string {
	return s.TemplateIdentifier
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) SetRemediationType(v string) *ListRemediationTemplatesResponseBodyRemediationTemplates {
	s.RemediationType = &v
	return s
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) SetTemplateDefinition(v string) *ListRemediationTemplatesResponseBodyRemediationTemplates {
	s.TemplateDefinition = &v
	return s
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) SetTemplateDescription(v string) *ListRemediationTemplatesResponseBodyRemediationTemplates {
	s.TemplateDescription = &v
	return s
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) SetTemplateIdentifier(v string) *ListRemediationTemplatesResponseBodyRemediationTemplates {
	s.TemplateIdentifier = &v
	return s
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) SetTemplateName(v string) *ListRemediationTemplatesResponseBodyRemediationTemplates {
	s.TemplateName = &v
	return s
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) Validate() error {
	return dara.Validate(s)
}
