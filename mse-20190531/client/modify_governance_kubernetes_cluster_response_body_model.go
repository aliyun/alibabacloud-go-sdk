// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGovernanceKubernetesClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ModifyGovernanceKubernetesClusterResponseBody
	GetCode() *int32
	SetData(v bool) *ModifyGovernanceKubernetesClusterResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *ModifyGovernanceKubernetesClusterResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyGovernanceKubernetesClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyGovernanceKubernetesClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyGovernanceKubernetesClusterResponseBody
	GetSuccess() *bool
}

type ModifyGovernanceKubernetesClusterResponseBody struct {
	// The response code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The deletion result.
	//
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F9849E97-2455-54B0-A3B4-3F6E4E9FFEB1
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

func (s ModifyGovernanceKubernetesClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyGovernanceKubernetesClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGovernanceKubernetesClusterResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ModifyGovernanceKubernetesClusterResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModifyGovernanceKubernetesClusterResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyGovernanceKubernetesClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyGovernanceKubernetesClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyGovernanceKubernetesClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyGovernanceKubernetesClusterResponseBody) SetCode(v int32) *ModifyGovernanceKubernetesClusterResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyGovernanceKubernetesClusterResponseBody) SetData(v bool) *ModifyGovernanceKubernetesClusterResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyGovernanceKubernetesClusterResponseBody) SetHttpStatusCode(v int32) *ModifyGovernanceKubernetesClusterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyGovernanceKubernetesClusterResponseBody) SetMessage(v string) *ModifyGovernanceKubernetesClusterResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyGovernanceKubernetesClusterResponseBody) SetRequestId(v string) *ModifyGovernanceKubernetesClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGovernanceKubernetesClusterResponseBody) SetSuccess(v bool) *ModifyGovernanceKubernetesClusterResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyGovernanceKubernetesClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
