// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamespacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListNamespacesResponseBody
	GetCount() *int32
	SetDBInstanceId(v string) *ListNamespacesResponseBody
	GetDBInstanceId() *string
	SetMessage(v string) *ListNamespacesResponseBody
	GetMessage() *string
	SetNamespaces(v *ListNamespacesResponseBodyNamespaces) *ListNamespacesResponseBody
	GetNamespaces() *ListNamespacesResponseBodyNamespaces
	SetRegionId(v string) *ListNamespacesResponseBody
	GetRegionId() *string
	SetRequestId(v string) *ListNamespacesResponseBody
	GetRequestId() *string
	SetStatus(v string) *ListNamespacesResponseBody
	GetStatus() *string
}

type ListNamespacesResponseBody struct {
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
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
	// The queried namespaces.
	Namespaces *ListNamespacesResponseBodyNamespaces `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Struct"`
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

func (s ListNamespacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListNamespacesResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListNamespacesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListNamespacesResponseBody) GetNamespaces() *ListNamespacesResponseBodyNamespaces {
	return s.Namespaces
}

func (s *ListNamespacesResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNamespacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNamespacesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListNamespacesResponseBody) SetCount(v int32) *ListNamespacesResponseBody {
	s.Count = &v
	return s
}

func (s *ListNamespacesResponseBody) SetDBInstanceId(v string) *ListNamespacesResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ListNamespacesResponseBody) SetMessage(v string) *ListNamespacesResponseBody {
	s.Message = &v
	return s
}

func (s *ListNamespacesResponseBody) SetNamespaces(v *ListNamespacesResponseBodyNamespaces) *ListNamespacesResponseBody {
	s.Namespaces = v
	return s
}

func (s *ListNamespacesResponseBody) SetRegionId(v string) *ListNamespacesResponseBody {
	s.RegionId = &v
	return s
}

func (s *ListNamespacesResponseBody) SetRequestId(v string) *ListNamespacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNamespacesResponseBody) SetStatus(v string) *ListNamespacesResponseBody {
	s.Status = &v
	return s
}

func (s *ListNamespacesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListNamespacesResponseBodyNamespaces struct {
	Namespace []*string `json:"Namespace,omitempty" xml:"Namespace,omitempty" type:"Repeated"`
}

func (s ListNamespacesResponseBodyNamespaces) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesResponseBodyNamespaces) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBodyNamespaces) GetNamespace() []*string {
	return s.Namespace
}

func (s *ListNamespacesResponseBodyNamespaces) SetNamespace(v []*string) *ListNamespacesResponseBodyNamespaces {
	s.Namespace = v
	return s
}

func (s *ListNamespacesResponseBodyNamespaces) Validate() error {
	return dara.Validate(s)
}
