// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPolicyAttachmentResponseBody
	GetCode() *string
	SetData(v *GetPolicyAttachmentResponseBodyData) *GetPolicyAttachmentResponseBody
	GetData() *GetPolicyAttachmentResponseBodyData
	SetMessage(v string) *GetPolicyAttachmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPolicyAttachmentResponseBody
	GetRequestId() *string
}

type GetPolicyAttachmentResponseBody struct {
	// Response code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response data.
	Data *GetPolicyAttachmentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 2C3B9A12-3868-5EB9-fBEA-F99E03DD1***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetPolicyAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolicyAttachmentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPolicyAttachmentResponseBody) GetData() *GetPolicyAttachmentResponseBodyData {
	return s.Data
}

func (s *GetPolicyAttachmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPolicyAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPolicyAttachmentResponseBody) SetCode(v string) *GetPolicyAttachmentResponseBody {
	s.Code = &v
	return s
}

func (s *GetPolicyAttachmentResponseBody) SetData(v *GetPolicyAttachmentResponseBodyData) *GetPolicyAttachmentResponseBody {
	s.Data = v
	return s
}

func (s *GetPolicyAttachmentResponseBody) SetMessage(v string) *GetPolicyAttachmentResponseBody {
	s.Message = &v
	return s
}

func (s *GetPolicyAttachmentResponseBody) SetRequestId(v string) *GetPolicyAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPolicyAttachmentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPolicyAttachmentResponseBodyData struct {
	// Attached Resource ID
	//
	// example:
	//
	// op-csbkd9llhtgqhqua***
	AttachResourceId *string `json:"attachResourceId,omitempty" xml:"attachResourceId,omitempty"`
	// Attached resource type, HttpApi, GatewayRoute, Operation, GatewayService, GatewayServicePort, Gateway, Domain
	//
	// example:
	//
	// Operation
	AttachResourceType *string `json:"attachResourceType,omitempty" xml:"attachResourceType,omitempty"`
	// Policy attachment configuration
	//
	// example:
	//
	// {"unitNum":1,"timeUnit":"s","enable":true}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// Environment ID
	//
	// example:
	//
	// env-cq7l5s5lhtgi6qa***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// Gateway Instance ID
	//
	// example:
	//
	// gw-cq2vundlhtg***
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// Policy Attachment ID
	//
	// example:
	//
	// pr-cqoojualhtgquuj***
	PolicyAttachmentId *string `json:"policyAttachmentId,omitempty" xml:"policyAttachmentId,omitempty"`
	// Policy ID
	//
	// example:
	//
	// p-cq7l5s5bblhtgi6qas***
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
}

func (s GetPolicyAttachmentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyAttachmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPolicyAttachmentResponseBodyData) GetAttachResourceId() *string {
	return s.AttachResourceId
}

func (s *GetPolicyAttachmentResponseBodyData) GetAttachResourceType() *string {
	return s.AttachResourceType
}

func (s *GetPolicyAttachmentResponseBodyData) GetConfig() *string {
	return s.Config
}

func (s *GetPolicyAttachmentResponseBodyData) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *GetPolicyAttachmentResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetPolicyAttachmentResponseBodyData) GetPolicyAttachmentId() *string {
	return s.PolicyAttachmentId
}

func (s *GetPolicyAttachmentResponseBodyData) GetPolicyId() *string {
	return s.PolicyId
}

func (s *GetPolicyAttachmentResponseBodyData) SetAttachResourceId(v string) *GetPolicyAttachmentResponseBodyData {
	s.AttachResourceId = &v
	return s
}

func (s *GetPolicyAttachmentResponseBodyData) SetAttachResourceType(v string) *GetPolicyAttachmentResponseBodyData {
	s.AttachResourceType = &v
	return s
}

func (s *GetPolicyAttachmentResponseBodyData) SetConfig(v string) *GetPolicyAttachmentResponseBodyData {
	s.Config = &v
	return s
}

func (s *GetPolicyAttachmentResponseBodyData) SetEnvironmentId(v string) *GetPolicyAttachmentResponseBodyData {
	s.EnvironmentId = &v
	return s
}

func (s *GetPolicyAttachmentResponseBodyData) SetGatewayId(v string) *GetPolicyAttachmentResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetPolicyAttachmentResponseBodyData) SetPolicyAttachmentId(v string) *GetPolicyAttachmentResponseBodyData {
	s.PolicyAttachmentId = &v
	return s
}

func (s *GetPolicyAttachmentResponseBodyData) SetPolicyId(v string) *GetPolicyAttachmentResponseBodyData {
	s.PolicyId = &v
	return s
}

func (s *GetPolicyAttachmentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
