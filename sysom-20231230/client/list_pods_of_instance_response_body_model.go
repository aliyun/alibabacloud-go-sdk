// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPodsOfInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListPodsOfInstanceResponseBody
	GetRequestId() *string
	SetCode(v string) *ListPodsOfInstanceResponseBody
	GetCode() *string
	SetData(v []*ListPodsOfInstanceResponseBodyData) *ListPodsOfInstanceResponseBody
	GetData() []*ListPodsOfInstanceResponseBodyData
	SetMessage(v string) *ListPodsOfInstanceResponseBody
	GetMessage() *string
	SetTotal(v int64) *ListPodsOfInstanceResponseBody
	GetTotal() *int64
}

type ListPodsOfInstanceResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// B149FD9C-ED5C-5765-B3AD-05AA4A4D64D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListPodsOfInstanceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// instance not exists
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 42
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListPodsOfInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPodsOfInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListPodsOfInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPodsOfInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPodsOfInstanceResponseBody) GetData() []*ListPodsOfInstanceResponseBodyData {
	return s.Data
}

func (s *ListPodsOfInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPodsOfInstanceResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListPodsOfInstanceResponseBody) SetRequestId(v string) *ListPodsOfInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPodsOfInstanceResponseBody) SetCode(v string) *ListPodsOfInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ListPodsOfInstanceResponseBody) SetData(v []*ListPodsOfInstanceResponseBodyData) *ListPodsOfInstanceResponseBody {
	s.Data = v
	return s
}

func (s *ListPodsOfInstanceResponseBody) SetMessage(v string) *ListPodsOfInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ListPodsOfInstanceResponseBody) SetTotal(v int64) *ListPodsOfInstanceResponseBody {
	s.Total = &v
	return s
}

func (s *ListPodsOfInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPodsOfInstanceResponseBodyData struct {
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// test-pod
	Pod *string `json:"pod,omitempty" xml:"pod,omitempty"`
}

func (s ListPodsOfInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPodsOfInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPodsOfInstanceResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *ListPodsOfInstanceResponseBodyData) GetPod() *string {
	return s.Pod
}

func (s *ListPodsOfInstanceResponseBodyData) SetNamespace(v string) *ListPodsOfInstanceResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *ListPodsOfInstanceResponseBodyData) SetPod(v string) *ListPodsOfInstanceResponseBodyData {
	s.Pod = &v
	return s
}

func (s *ListPodsOfInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
