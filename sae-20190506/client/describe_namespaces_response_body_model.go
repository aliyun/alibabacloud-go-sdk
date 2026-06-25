// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNamespacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeNamespacesResponseBody
	GetCode() *string
	SetData(v *DescribeNamespacesResponseBodyData) *DescribeNamespacesResponseBody
	GetData() *DescribeNamespacesResponseBodyData
	SetErrorCode(v string) *DescribeNamespacesResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeNamespacesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeNamespacesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeNamespacesResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeNamespacesResponseBody
	GetTraceId() *string
}

type DescribeNamespacesResponseBody struct {
	// The HTTP status code.
	//
	// - **2xx**: The request was successful.
	//
	// - **3xx**: The request was redirected.
	//
	// - **4xx**: The request was invalid.
	//
	// - **5xx**: A server-side error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the namespaces.
	Data *DescribeNamespacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - This parameter is returned only if the request fails.
	//
	// - For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned by the system.
	//
	// - Returns **success*	- if the request is successful.
	//
	// - Returns a specific error message if the request fails.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID that can be used to query the details of a call.
	//
	// example:
	//
	// 0a981dd515966966104121683d****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeNamespacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeNamespacesResponseBody) GetData() *DescribeNamespacesResponseBodyData {
	return s.Data
}

func (s *DescribeNamespacesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeNamespacesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeNamespacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNamespacesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeNamespacesResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeNamespacesResponseBody) SetCode(v string) *DescribeNamespacesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeNamespacesResponseBody) SetData(v *DescribeNamespacesResponseBodyData) *DescribeNamespacesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeNamespacesResponseBody) SetErrorCode(v string) *DescribeNamespacesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeNamespacesResponseBody) SetMessage(v string) *DescribeNamespacesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeNamespacesResponseBody) SetRequestId(v string) *DescribeNamespacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNamespacesResponseBody) SetSuccess(v bool) *DescribeNamespacesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeNamespacesResponseBody) SetTraceId(v string) *DescribeNamespacesResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeNamespacesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNamespacesResponseBodyData struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The list of namespaces.
	Namespaces []*DescribeNamespacesResponseBodyDataNamespaces `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of namespaces.
	//
	// example:
	//
	// 100
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribeNamespacesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespacesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeNamespacesResponseBodyData) GetNamespaces() []*DescribeNamespacesResponseBodyDataNamespaces {
	return s.Namespaces
}

func (s *DescribeNamespacesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNamespacesResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *DescribeNamespacesResponseBodyData) SetCurrentPage(v int32) *DescribeNamespacesResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *DescribeNamespacesResponseBodyData) SetNamespaces(v []*DescribeNamespacesResponseBodyDataNamespaces) *DescribeNamespacesResponseBodyData {
	s.Namespaces = v
	return s
}

func (s *DescribeNamespacesResponseBodyData) SetPageSize(v int32) *DescribeNamespacesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeNamespacesResponseBodyData) SetTotalSize(v int32) *DescribeNamespacesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *DescribeNamespacesResponseBodyData) Validate() error {
	if s.Namespaces != nil {
		for _, item := range s.Namespaces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNamespacesResponseBodyDataNamespaces struct {
	// The access key ID for Application Configuration Management (ACM), used to manage data in an ACM namespace. For more information, see [Differences between an Alibaba Cloud access key and an ACM-specific access key](https://help.aliyun.com/document_detail/68941.html).
	//
	// example:
	//
	// b34dbe3315c64f9f99b58ea447ec****
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The service endpoint.
	//
	// example:
	//
	// addr-bj-internal.edas.aliyun.com
	AddressServerHost *string `json:"AddressServerHost,omitempty" xml:"AddressServerHost,omitempty"`
	// The short ID of the namespace.
	//
	// example:
	//
	// test
	NameSpaceShortId *string `json:"NameSpaceShortId,omitempty" xml:"NameSpaceShortId,omitempty"`
	// The description of the namespace.
	//
	// example:
	//
	// desc
	NamespaceDescription *string `json:"NamespaceDescription,omitempty" xml:"NamespaceDescription,omitempty"`
	// The namespace ID. The default namespace cannot be queried, modified, or deleted.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// name
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The ID of the region. For example, \\"cn-beijing\\" indicates China (Beijing).
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The secret access key for Application Configuration Management (ACM), used to manage data in an ACM namespace. For more information, see [Differences between an Alibaba Cloud access key and an ACM-specific access key](https://help.aliyun.com/document_detail/68941.html).
	//
	// example:
	//
	// G/w6sseK7+nb3S6HBmANDBMD****
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 838cad95-973f-48fe-830b-2a8546d7****
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeNamespacesResponseBodyDataNamespaces) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespacesResponseBodyDataNamespaces) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) GetAccessKey() *string {
	return s.AccessKey
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) GetAddressServerHost() *string {
	return s.AddressServerHost
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) GetNameSpaceShortId() *string {
	return s.NameSpaceShortId
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) GetNamespaceDescription() *string {
	return s.NamespaceDescription
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) GetSecretKey() *string {
	return s.SecretKey
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) SetAccessKey(v string) *DescribeNamespacesResponseBodyDataNamespaces {
	s.AccessKey = &v
	return s
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) SetAddressServerHost(v string) *DescribeNamespacesResponseBodyDataNamespaces {
	s.AddressServerHost = &v
	return s
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) SetNameSpaceShortId(v string) *DescribeNamespacesResponseBodyDataNamespaces {
	s.NameSpaceShortId = &v
	return s
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) SetNamespaceDescription(v string) *DescribeNamespacesResponseBodyDataNamespaces {
	s.NamespaceDescription = &v
	return s
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) SetNamespaceId(v string) *DescribeNamespacesResponseBodyDataNamespaces {
	s.NamespaceId = &v
	return s
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) SetNamespaceName(v string) *DescribeNamespacesResponseBodyDataNamespaces {
	s.NamespaceName = &v
	return s
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) SetRegionId(v string) *DescribeNamespacesResponseBodyDataNamespaces {
	s.RegionId = &v
	return s
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) SetSecretKey(v string) *DescribeNamespacesResponseBodyDataNamespaces {
	s.SecretKey = &v
	return s
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) SetTenantId(v string) *DescribeNamespacesResponseBodyDataNamespaces {
	s.TenantId = &v
	return s
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) Validate() error {
	return dara.Validate(s)
}
