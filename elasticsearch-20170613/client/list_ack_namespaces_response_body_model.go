// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAckNamespacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAckNamespacesResponseBody
	GetRequestId() *string
	SetResult(v []*ListAckNamespacesResponseBodyResult) *ListAckNamespacesResponseBody
	GetResult() []*ListAckNamespacesResponseBodyResult
}

type ListAckNamespacesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 95789100-A329-473B-9D14-9E0B7DB4BD5A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result []*ListAckNamespacesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListAckNamespacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAckNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAckNamespacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAckNamespacesResponseBody) GetResult() []*ListAckNamespacesResponseBodyResult {
	return s.Result
}

func (s *ListAckNamespacesResponseBody) SetRequestId(v string) *ListAckNamespacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAckNamespacesResponseBody) SetResult(v []*ListAckNamespacesResponseBodyResult) *ListAckNamespacesResponseBody {
	s.Result = v
	return s
}

func (s *ListAckNamespacesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAckNamespacesResponseBodyResult struct {
	// The namespace of the cluster.
	//
	// example:
	//
	// logging
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The status of the namespace.
	//
	// example:
	//
	// Active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListAckNamespacesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListAckNamespacesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAckNamespacesResponseBodyResult) GetNamespace() *string {
	return s.Namespace
}

func (s *ListAckNamespacesResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListAckNamespacesResponseBodyResult) SetNamespace(v string) *ListAckNamespacesResponseBodyResult {
	s.Namespace = &v
	return s
}

func (s *ListAckNamespacesResponseBodyResult) SetStatus(v string) *ListAckNamespacesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListAckNamespacesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
