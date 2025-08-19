// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTargetAttachmentsForControlPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTargetAttachmentsForControlPolicyResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTargetAttachmentsForControlPolicyResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTargetAttachmentsForControlPolicyResponseBody
	GetRequestId() *string
	SetTargetAttachments(v *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments) *ListTargetAttachmentsForControlPolicyResponseBody
	GetTargetAttachments() *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments
	SetTotalCount(v int32) *ListTargetAttachmentsForControlPolicyResponseBody
	GetTotalCount() *int32
}

type ListTargetAttachmentsForControlPolicyResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B32BD3D6-1089-41F3-8E70-E0079BC7D760
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the objects to which the access control policy is attached.
	TargetAttachments *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments `json:"TargetAttachments,omitempty" xml:"TargetAttachments,omitempty" type:"Struct"`
	// The total number of objects to which the access control policy is attached.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTargetAttachmentsForControlPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTargetAttachmentsForControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListTargetAttachmentsForControlPolicyResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTargetAttachmentsForControlPolicyResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTargetAttachmentsForControlPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTargetAttachmentsForControlPolicyResponseBody) GetTargetAttachments() *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments {
	return s.TargetAttachments
}

func (s *ListTargetAttachmentsForControlPolicyResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTargetAttachmentsForControlPolicyResponseBody) SetPageNumber(v int32) *ListTargetAttachmentsForControlPolicyResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBody) SetPageSize(v int32) *ListTargetAttachmentsForControlPolicyResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBody) SetRequestId(v string) *ListTargetAttachmentsForControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBody) SetTargetAttachments(v *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments) *ListTargetAttachmentsForControlPolicyResponseBody {
	s.TargetAttachments = v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBody) SetTotalCount(v int32) *ListTargetAttachmentsForControlPolicyResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments struct {
	TargetAttachment []*ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment `json:"TargetAttachment,omitempty" xml:"TargetAttachment,omitempty" type:"Repeated"`
}

func (s ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments) String() string {
	return dara.Prettify(s)
}

func (s ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments) GoString() string {
	return s.String()
}

func (s *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments) GetTargetAttachment() []*ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment {
	return s.TargetAttachment
}

func (s *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments) SetTargetAttachment(v []*ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments {
	s.TargetAttachment = v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments) Validate() error {
	return dara.Validate(s)
}

type ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment struct {
	// The time when the access control policy was attached to the object.
	//
	// example:
	//
	// 2021-03-19T02:56:24Z
	AttachDate *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	// The ID of the object.
	//
	// example:
	//
	// fd-ZDNPiT****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the object.
	//
	// example:
	//
	// Dev_Department
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The type of the object. Valid values:
	//
	// 	- Root: Root folder
	//
	// 	- Folder: subfolder of the Root folder
	//
	// 	- Account: member
	//
	// example:
	//
	// Folder
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) String() string {
	return dara.Prettify(s)
}

func (s ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) GoString() string {
	return s.String()
}

func (s *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) GetAttachDate() *string {
	return s.AttachDate
}

func (s *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) GetTargetId() *string {
	return s.TargetId
}

func (s *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) GetTargetName() *string {
	return s.TargetName
}

func (s *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) GetTargetType() *string {
	return s.TargetType
}

func (s *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) SetAttachDate(v string) *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment {
	s.AttachDate = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) SetTargetId(v string) *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment {
	s.TargetId = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) SetTargetName(v string) *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment {
	s.TargetName = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) SetTargetType(v string) *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment {
	s.TargetType = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) Validate() error {
	return dara.Validate(s)
}
