// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeListenerAccessControlAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessControlStatus(v string) *DescribeListenerAccessControlAttributeResponseBody
	GetAccessControlStatus() *string
	SetRequestId(v string) *DescribeListenerAccessControlAttributeResponseBody
	GetRequestId() *string
	SetSourceItems(v string) *DescribeListenerAccessControlAttributeResponseBody
	GetSourceItems() *string
}

type DescribeListenerAccessControlAttributeResponseBody struct {
	// Indicates whether the whitelist is enabled. Valid values:
	//
	// 	- **open_white_list**: the whitelist is enabled.
	//
	// 	- **close**: the whitelist is disabled.
	//
	// example:
	//
	// open_white_list
	AccessControlStatus *string `json:"AccessControlStatus,omitempty" xml:"AccessControlStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried ACLs.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceItems *string `json:"SourceItems,omitempty" xml:"SourceItems,omitempty"`
}

func (s DescribeListenerAccessControlAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeListenerAccessControlAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeListenerAccessControlAttributeResponseBody) GetAccessControlStatus() *string {
	return s.AccessControlStatus
}

func (s *DescribeListenerAccessControlAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeListenerAccessControlAttributeResponseBody) GetSourceItems() *string {
	return s.SourceItems
}

func (s *DescribeListenerAccessControlAttributeResponseBody) SetAccessControlStatus(v string) *DescribeListenerAccessControlAttributeResponseBody {
	s.AccessControlStatus = &v
	return s
}

func (s *DescribeListenerAccessControlAttributeResponseBody) SetRequestId(v string) *DescribeListenerAccessControlAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeListenerAccessControlAttributeResponseBody) SetSourceItems(v string) *DescribeListenerAccessControlAttributeResponseBody {
	s.SourceItems = &v
	return s
}

func (s *DescribeListenerAccessControlAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
