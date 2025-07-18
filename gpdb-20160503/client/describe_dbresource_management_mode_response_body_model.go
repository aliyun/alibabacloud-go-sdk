// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBResourceManagementModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDBResourceManagementModeResponseBody
	GetRequestId() *string
	SetResourceManagementMode(v string) *DescribeDBResourceManagementModeResponseBody
	GetResourceManagementMode() *string
}

type DescribeDBResourceManagementModeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource management mode. Valid values:
	//
	// 	- resourceGroup: resource group management.
	//
	// 	- resourceQueue: resource queue management.
	//
	// example:
	//
	// resourceGroup
	ResourceManagementMode *string `json:"ResourceManagementMode,omitempty" xml:"ResourceManagementMode,omitempty"`
}

func (s DescribeDBResourceManagementModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourceManagementModeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceManagementModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBResourceManagementModeResponseBody) GetResourceManagementMode() *string {
	return s.ResourceManagementMode
}

func (s *DescribeDBResourceManagementModeResponseBody) SetRequestId(v string) *DescribeDBResourceManagementModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBResourceManagementModeResponseBody) SetResourceManagementMode(v string) *DescribeDBResourceManagementModeResponseBody {
	s.ResourceManagementMode = &v
	return s
}

func (s *DescribeDBResourceManagementModeResponseBody) Validate() error {
	return dara.Validate(s)
}
