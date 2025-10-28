// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportK8sClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ImportK8sClusterResponseBody
	GetCode() *int32
	SetData(v string) *ImportK8sClusterResponseBody
	GetData() *string
	SetMessage(v string) *ImportK8sClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportK8sClusterResponseBody
	GetRequestId() *string
}

type ImportK8sClusterResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the cluster in EDAS.
	//
	// example:
	//
	// cf96d49a-6be2-4b6d-****-75c7fb86****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// a5281053-08e4-47a5-b2ab-5c0323de7b5a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportK8sClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportK8sClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ImportK8sClusterResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ImportK8sClusterResponseBody) GetData() *string {
	return s.Data
}

func (s *ImportK8sClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportK8sClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportK8sClusterResponseBody) SetCode(v int32) *ImportK8sClusterResponseBody {
	s.Code = &v
	return s
}

func (s *ImportK8sClusterResponseBody) SetData(v string) *ImportK8sClusterResponseBody {
	s.Data = &v
	return s
}

func (s *ImportK8sClusterResponseBody) SetMessage(v string) *ImportK8sClusterResponseBody {
	s.Message = &v
	return s
}

func (s *ImportK8sClusterResponseBody) SetRequestId(v string) *ImportK8sClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportK8sClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
