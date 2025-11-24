// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGuestClusterNamespacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNsLabels(v map[string]interface{}) *DescribeGuestClusterNamespacesResponseBody
	GetNsLabels() map[string]interface{}
	SetNsList(v []*string) *DescribeGuestClusterNamespacesResponseBody
	GetNsList() []*string
	SetRequestId(v string) *DescribeGuestClusterNamespacesResponseBody
	GetRequestId() *string
}

type DescribeGuestClusterNamespacesResponseBody struct {
	// The labels of the namespaces. Labels are returned only when `ShowNsLabels` is set to `true`.
	//
	// example:
	//
	// {"default":{"istio-injection":"enabled"}}
	NsLabels map[string]interface{} `json:"NsLabels,omitempty" xml:"NsLabels,omitempty"`
	// The names of the namespaces.
	NsList []*string `json:"NsList,omitempty" xml:"NsList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 31d3a0f0-07ed-4f6e-9004-1804498c****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGuestClusterNamespacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGuestClusterNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGuestClusterNamespacesResponseBody) GetNsLabels() map[string]interface{} {
	return s.NsLabels
}

func (s *DescribeGuestClusterNamespacesResponseBody) GetNsList() []*string {
	return s.NsList
}

func (s *DescribeGuestClusterNamespacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGuestClusterNamespacesResponseBody) SetNsLabels(v map[string]interface{}) *DescribeGuestClusterNamespacesResponseBody {
	s.NsLabels = v
	return s
}

func (s *DescribeGuestClusterNamespacesResponseBody) SetNsList(v []*string) *DescribeGuestClusterNamespacesResponseBody {
	s.NsList = v
	return s
}

func (s *DescribeGuestClusterNamespacesResponseBody) SetRequestId(v string) *DescribeGuestClusterNamespacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGuestClusterNamespacesResponseBody) Validate() error {
	return dara.Validate(s)
}
