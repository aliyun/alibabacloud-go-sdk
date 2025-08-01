// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKubernetesSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetKubernetesSourceResponseBody
	GetCode() *int32
	SetData(v []*GetKubernetesSourceResponseBodyData) *GetKubernetesSourceResponseBody
	GetData() []*GetKubernetesSourceResponseBodyData
	SetHttpStatusCode(v int32) *GetKubernetesSourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetKubernetesSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetKubernetesSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetKubernetesSourceResponseBody
	GetSuccess() *bool
}

type GetKubernetesSourceResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data []*GetKubernetesSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1085A66C-DEF1-58EE-A0A4-31E00C9FC0D5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetKubernetesSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKubernetesSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetKubernetesSourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetKubernetesSourceResponseBody) GetData() []*GetKubernetesSourceResponseBodyData {
	return s.Data
}

func (s *GetKubernetesSourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetKubernetesSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetKubernetesSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKubernetesSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetKubernetesSourceResponseBody) SetCode(v int32) *GetKubernetesSourceResponseBody {
	s.Code = &v
	return s
}

func (s *GetKubernetesSourceResponseBody) SetData(v []*GetKubernetesSourceResponseBodyData) *GetKubernetesSourceResponseBody {
	s.Data = v
	return s
}

func (s *GetKubernetesSourceResponseBody) SetHttpStatusCode(v int32) *GetKubernetesSourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetKubernetesSourceResponseBody) SetMessage(v string) *GetKubernetesSourceResponseBody {
	s.Message = &v
	return s
}

func (s *GetKubernetesSourceResponseBody) SetRequestId(v string) *GetKubernetesSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKubernetesSourceResponseBody) SetSuccess(v bool) *GetKubernetesSourceResponseBody {
	s.Success = &v
	return s
}

func (s *GetKubernetesSourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetKubernetesSourceResponseBodyData struct {
	// The ID of the ACK cluster.
	//
	// example:
	//
	// cbc1efca895a64af097ff00b26f3f****
	Cluster *string `json:"Cluster,omitempty" xml:"Cluster,omitempty"`
	// The name of the ACK cluster.
	//
	// example:
	//
	// k8s-demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetKubernetesSourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetKubernetesSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetKubernetesSourceResponseBodyData) GetCluster() *string {
	return s.Cluster
}

func (s *GetKubernetesSourceResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetKubernetesSourceResponseBodyData) SetCluster(v string) *GetKubernetesSourceResponseBodyData {
	s.Cluster = &v
	return s
}

func (s *GetKubernetesSourceResponseBodyData) SetName(v string) *GetKubernetesSourceResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetKubernetesSourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
