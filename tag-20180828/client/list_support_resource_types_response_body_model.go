// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupportResourceTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListSupportResourceTypesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSupportResourceTypesResponseBody
	GetRequestId() *string
	SetSupportResourceTypes(v []*ListSupportResourceTypesResponseBodySupportResourceTypes) *ListSupportResourceTypesResponseBody
	GetSupportResourceTypes() []*ListSupportResourceTypesResponseBodySupportResourceTypes
}

type ListSupportResourceTypesResponseBody struct {
	// Indicates whether the next query is required.
	//
	// 	- If the value of this parameter is empty, all results are returned, and the next query is not required.
	//
	// 	- If the value of this parameter is not empty, the next query is required, and the value is the token used to start the next query.
	//
	// example:
	//
	// AAAAAYws9fJ0Ur4MGm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ABC71772-F3A1-59CA-B811-4A5B0E0B72F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The supported resource types.
	SupportResourceTypes []*ListSupportResourceTypesResponseBodySupportResourceTypes `json:"SupportResourceTypes,omitempty" xml:"SupportResourceTypes,omitempty" type:"Repeated"`
}

func (s ListSupportResourceTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSupportResourceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSupportResourceTypesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSupportResourceTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSupportResourceTypesResponseBody) GetSupportResourceTypes() []*ListSupportResourceTypesResponseBodySupportResourceTypes {
	return s.SupportResourceTypes
}

func (s *ListSupportResourceTypesResponseBody) SetNextToken(v string) *ListSupportResourceTypesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSupportResourceTypesResponseBody) SetRequestId(v string) *ListSupportResourceTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSupportResourceTypesResponseBody) SetSupportResourceTypes(v []*ListSupportResourceTypesResponseBodySupportResourceTypes) *ListSupportResourceTypesResponseBody {
	s.SupportResourceTypes = v
	return s
}

func (s *ListSupportResourceTypesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSupportResourceTypesResponseBodySupportResourceTypes struct {
	// The resource ARN template.
	//
	// example:
	//
	// acs:ecs:*:*:instance/${ResourceId}
	ArnTemplate *string `json:"ArnTemplate,omitempty" xml:"ArnTemplate,omitempty"`
	// The service code.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The resource type.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The supported tag-related capability items.
	//
	// >  This parameter is returned only if the `ShowItems` parameter is set to `true`.
	SupportItems []*ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems `json:"SupportItems,omitempty" xml:"SupportItems,omitempty" type:"Repeated"`
}

func (s ListSupportResourceTypesResponseBodySupportResourceTypes) String() string {
	return dara.Prettify(s)
}

func (s ListSupportResourceTypesResponseBodySupportResourceTypes) GoString() string {
	return s.String()
}

func (s *ListSupportResourceTypesResponseBodySupportResourceTypes) GetArnTemplate() *string {
	return s.ArnTemplate
}

func (s *ListSupportResourceTypesResponseBodySupportResourceTypes) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListSupportResourceTypesResponseBodySupportResourceTypes) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListSupportResourceTypesResponseBodySupportResourceTypes) GetSupportItems() []*ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems {
	return s.SupportItems
}

func (s *ListSupportResourceTypesResponseBodySupportResourceTypes) SetArnTemplate(v string) *ListSupportResourceTypesResponseBodySupportResourceTypes {
	s.ArnTemplate = &v
	return s
}

func (s *ListSupportResourceTypesResponseBodySupportResourceTypes) SetProductCode(v string) *ListSupportResourceTypesResponseBodySupportResourceTypes {
	s.ProductCode = &v
	return s
}

func (s *ListSupportResourceTypesResponseBodySupportResourceTypes) SetResourceType(v string) *ListSupportResourceTypesResponseBodySupportResourceTypes {
	s.ResourceType = &v
	return s
}

func (s *ListSupportResourceTypesResponseBodySupportResourceTypes) SetSupportItems(v []*ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems) *ListSupportResourceTypesResponseBodySupportResourceTypes {
	s.SupportItems = v
	return s
}

func (s *ListSupportResourceTypesResponseBodySupportResourceTypes) Validate() error {
	return dara.Validate(s)
}

type ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems struct {
	// Indicates whether the tag-related capability item is supported. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Support *bool `json:"Support,omitempty" xml:"Support,omitempty"`
	// The code of the tag-related capability item.
	//
	// example:
	//
	// TAG_CONSOLE_SUPPORT
	SupportCode *string `json:"SupportCode,omitempty" xml:"SupportCode,omitempty"`
	// The details of the support for the tag-related capability item.
	SupportDetails []map[string]*string `json:"SupportDetails,omitempty" xml:"SupportDetails,omitempty" type:"Repeated"`
}

func (s ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems) String() string {
	return dara.Prettify(s)
}

func (s ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems) GoString() string {
	return s.String()
}

func (s *ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems) GetSupport() *bool {
	return s.Support
}

func (s *ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems) GetSupportCode() *string {
	return s.SupportCode
}

func (s *ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems) GetSupportDetails() []map[string]*string {
	return s.SupportDetails
}

func (s *ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems) SetSupport(v bool) *ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems {
	s.Support = &v
	return s
}

func (s *ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems) SetSupportCode(v string) *ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems {
	s.SupportCode = &v
	return s
}

func (s *ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems) SetSupportDetails(v []map[string]*string) *ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems {
	s.SupportDetails = v
	return s
}

func (s *ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems) Validate() error {
	return dara.Validate(s)
}
