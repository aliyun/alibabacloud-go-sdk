// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeNamespaceResponseBody
	GetDBInstanceId() *string
	SetMessage(v string) *DescribeNamespaceResponseBody
	GetMessage() *string
	SetNamespace(v string) *DescribeNamespaceResponseBody
	GetNamespace() *string
	SetNamespaceInfo(v map[string]*string) *DescribeNamespaceResponseBody
	GetNamespaceInfo() map[string]*string
	SetRegionId(v string) *DescribeNamespaceResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeNamespaceResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeNamespaceResponseBody
	GetStatus() *string
}

type DescribeNamespaceResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The queried namespace.
	NamespaceInfo map[string]*string `json:"NamespaceInfo,omitempty" xml:"NamespaceInfo,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeNamespaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeNamespaceResponseBody) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeNamespaceResponseBody) GetNamespaceInfo() map[string]*string {
	return s.NamespaceInfo
}

func (s *DescribeNamespaceResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNamespaceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeNamespaceResponseBody) SetDBInstanceId(v string) *DescribeNamespaceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeNamespaceResponseBody) SetMessage(v string) *DescribeNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeNamespaceResponseBody) SetNamespace(v string) *DescribeNamespaceResponseBody {
	s.Namespace = &v
	return s
}

func (s *DescribeNamespaceResponseBody) SetNamespaceInfo(v map[string]*string) *DescribeNamespaceResponseBody {
	s.NamespaceInfo = v
	return s
}

func (s *DescribeNamespaceResponseBody) SetRegionId(v string) *DescribeNamespaceResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeNamespaceResponseBody) SetRequestId(v string) *DescribeNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNamespaceResponseBody) SetStatus(v string) *DescribeNamespaceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}
