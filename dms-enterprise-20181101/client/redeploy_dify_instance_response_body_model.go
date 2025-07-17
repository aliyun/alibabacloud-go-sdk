// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedeployDifyInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RedeployDifyInstanceResponseBody
	GetCode() *string
	SetData(v *RedeployDifyInstanceResponseBodyData) *RedeployDifyInstanceResponseBody
	GetData() *RedeployDifyInstanceResponseBodyData
	SetErrorCode(v string) *RedeployDifyInstanceResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *RedeployDifyInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RedeployDifyInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *RedeployDifyInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RedeployDifyInstanceResponseBody
	GetSuccess() *bool
}

type RedeployDifyInstanceResponseBody struct {
	Code           *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *RedeployDifyInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode      *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpStatusCode *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RedeployDifyInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RedeployDifyInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RedeployDifyInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *RedeployDifyInstanceResponseBody) GetData() *RedeployDifyInstanceResponseBodyData {
	return s.Data
}

func (s *RedeployDifyInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RedeployDifyInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RedeployDifyInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RedeployDifyInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RedeployDifyInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RedeployDifyInstanceResponseBody) SetCode(v string) *RedeployDifyInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *RedeployDifyInstanceResponseBody) SetData(v *RedeployDifyInstanceResponseBodyData) *RedeployDifyInstanceResponseBody {
	s.Data = v
	return s
}

func (s *RedeployDifyInstanceResponseBody) SetErrorCode(v string) *RedeployDifyInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RedeployDifyInstanceResponseBody) SetHttpStatusCode(v int32) *RedeployDifyInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RedeployDifyInstanceResponseBody) SetMessage(v string) *RedeployDifyInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *RedeployDifyInstanceResponseBody) SetRequestId(v string) *RedeployDifyInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RedeployDifyInstanceResponseBody) SetSuccess(v bool) *RedeployDifyInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *RedeployDifyInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type RedeployDifyInstanceResponseBodyData struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RedeployDifyInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RedeployDifyInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RedeployDifyInstanceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RedeployDifyInstanceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *RedeployDifyInstanceResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RedeployDifyInstanceResponseBodyData) SetInstanceId(v string) *RedeployDifyInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *RedeployDifyInstanceResponseBodyData) SetStatus(v string) *RedeployDifyInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *RedeployDifyInstanceResponseBodyData) SetWorkspaceId(v string) *RedeployDifyInstanceResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *RedeployDifyInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
