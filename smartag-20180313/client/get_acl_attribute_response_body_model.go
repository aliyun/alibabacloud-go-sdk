// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAclAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *GetAclAttributeResponseBody
	GetAclId() *string
	SetAclName(v string) *GetAclAttributeResponseBody
	GetAclName() *string
	SetErrorConfigSmartAGCount(v int32) *GetAclAttributeResponseBody
	GetErrorConfigSmartAGCount() *int32
	SetRequestId(v string) *GetAclAttributeResponseBody
	GetRequestId() *string
}

type GetAclAttributeResponseBody struct {
	// The ID of the ACL.
	//
	// example:
	//
	// acl-xhwhyuo43l0n*****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The name of the ACL.
	//
	// example:
	//
	// dpi_test
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// The number of SAG devices that are associated with the ACL who has DPI configuration errors.
	//
	// You can call the [ListDpiConfigError](https://help.aliyun.com/document_detail/197566.html) operation to query exception details and SAG device information.
	//
	// example:
	//
	// 0
	ErrorConfigSmartAGCount *int32 `json:"ErrorConfigSmartAGCount,omitempty" xml:"ErrorConfigSmartAGCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5D2013F0-85AB-4332-9094-8023A598C2C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAclAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAclAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAclAttributeResponseBody) GetAclId() *string {
	return s.AclId
}

func (s *GetAclAttributeResponseBody) GetAclName() *string {
	return s.AclName
}

func (s *GetAclAttributeResponseBody) GetErrorConfigSmartAGCount() *int32 {
	return s.ErrorConfigSmartAGCount
}

func (s *GetAclAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAclAttributeResponseBody) SetAclId(v string) *GetAclAttributeResponseBody {
	s.AclId = &v
	return s
}

func (s *GetAclAttributeResponseBody) SetAclName(v string) *GetAclAttributeResponseBody {
	s.AclName = &v
	return s
}

func (s *GetAclAttributeResponseBody) SetErrorConfigSmartAGCount(v int32) *GetAclAttributeResponseBody {
	s.ErrorConfigSmartAGCount = &v
	return s
}

func (s *GetAclAttributeResponseBody) SetRequestId(v string) *GetAclAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAclAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
