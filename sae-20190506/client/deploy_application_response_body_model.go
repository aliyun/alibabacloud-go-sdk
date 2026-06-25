// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeployApplicationResponseBody
	GetCode() *string
	SetData(v *DeployApplicationResponseBodyData) *DeployApplicationResponseBody
	GetData() *DeployApplicationResponseBodyData
	SetErrorCode(v string) *DeployApplicationResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DeployApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeployApplicationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeployApplicationResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeployApplicationResponseBody
	GetTraceId() *string
}

type DeployApplicationResponseBody struct {
	// The API status or POP error code. Values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirection.
	//
	// - **4xx**: Request error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *DeployApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Values:
	//
	// - On success: This field is not returned.
	//
	// - On failure: This field is returned. For details, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// 空
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Additional information. Values:
	//
	// - On success, returns **success**.
	//
	// - On failure, returns a specific error code.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01CF26C7-00A3-4AA6-BA76-7E95F2A3***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the deployment succeeded. Values:
	//
	// - **true**: Deployment succeeded.
	//
	// - **false**: Deployment failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID for precise query of call information.
	//
	// example:
	//
	// ac1a0b2215622246421415014e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DeployApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeployApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeployApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeployApplicationResponseBody) GetData() *DeployApplicationResponseBodyData {
	return s.Data
}

func (s *DeployApplicationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeployApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeployApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployApplicationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeployApplicationResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeployApplicationResponseBody) SetCode(v string) *DeployApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *DeployApplicationResponseBody) SetData(v *DeployApplicationResponseBodyData) *DeployApplicationResponseBody {
	s.Data = v
	return s
}

func (s *DeployApplicationResponseBody) SetErrorCode(v string) *DeployApplicationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeployApplicationResponseBody) SetMessage(v string) *DeployApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *DeployApplicationResponseBody) SetRequestId(v string) *DeployApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployApplicationResponseBody) SetSuccess(v bool) *DeployApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *DeployApplicationResponseBody) SetTraceId(v string) *DeployApplicationResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeployApplicationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeployApplicationResponseBodyData struct {
	// The application ID.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The returned change order ID. Use it to query task execution status.
	//
	// example:
	//
	// 01db03d3-3ee9-48b3-b3d0-dfce2d88****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	// Whether RAM users need approval to deploy changes. Values:
	//
	// - **true**: Approval required.
	//
	// - **false**: No approval required.
	//
	// example:
	//
	// true
	IsNeedApproval *bool `json:"IsNeedApproval,omitempty" xml:"IsNeedApproval,omitempty"`
}

func (s DeployApplicationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeployApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeployApplicationResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *DeployApplicationResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *DeployApplicationResponseBodyData) GetIsNeedApproval() *bool {
	return s.IsNeedApproval
}

func (s *DeployApplicationResponseBodyData) SetAppId(v string) *DeployApplicationResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DeployApplicationResponseBodyData) SetChangeOrderId(v string) *DeployApplicationResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *DeployApplicationResponseBodyData) SetIsNeedApproval(v bool) *DeployApplicationResponseBodyData {
	s.IsNeedApproval = &v
	return s
}

func (s *DeployApplicationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
