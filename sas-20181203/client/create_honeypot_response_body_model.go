// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHoneypotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateHoneypotResponseBody
	GetCode() *string
	SetData(v *CreateHoneypotResponseBodyData) *CreateHoneypotResponseBody
	GetData() *CreateHoneypotResponseBodyData
	SetHttpStatusCode(v int32) *CreateHoneypotResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateHoneypotResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateHoneypotResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateHoneypotResponseBody
	GetSuccess() *bool
}

type CreateHoneypotResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the honeypot.
	Data *CreateHoneypotResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 29874225-EAAC-5415-8501-32DD20FD29F6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateHoneypotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHoneypotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHoneypotResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateHoneypotResponseBody) GetData() *CreateHoneypotResponseBodyData {
	return s.Data
}

func (s *CreateHoneypotResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateHoneypotResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateHoneypotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHoneypotResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateHoneypotResponseBody) SetCode(v string) *CreateHoneypotResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHoneypotResponseBody) SetData(v *CreateHoneypotResponseBodyData) *CreateHoneypotResponseBody {
	s.Data = v
	return s
}

func (s *CreateHoneypotResponseBody) SetHttpStatusCode(v int32) *CreateHoneypotResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateHoneypotResponseBody) SetMessage(v string) *CreateHoneypotResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHoneypotResponseBody) SetRequestId(v string) *CreateHoneypotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHoneypotResponseBody) SetSuccess(v bool) *CreateHoneypotResponseBody {
	s.Success = &v
	return s
}

func (s *CreateHoneypotResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateHoneypotResponseBodyData struct {
	// The name of the management node.
	//
	// example:
	//
	// managerNodename
	ControlNodeName *string `json:"ControlNodeName,omitempty" xml:"ControlNodeName,omitempty"`
	// The ID of the honeypot.
	//
	// example:
	//
	// 9bf8cd373112263d4bc102fc5dba9d9f812ee05d4d35c487d330d52e937f****
	HoneypotId *string `json:"HoneypotId,omitempty" xml:"HoneypotId,omitempty"`
	// The display name of the honeypot image.
	//
	// example:
	//
	// HoneyDisplayName
	HoneypotImageDisplayName *string `json:"HoneypotImageDisplayName,omitempty" xml:"HoneypotImageDisplayName,omitempty"`
	// The name of the honeypot image.
	//
	// example:
	//
	// tcp_proxy
	HoneypotImageName *string `json:"HoneypotImageName,omitempty" xml:"HoneypotImageName,omitempty"`
	// The custom name of the honeypot.
	//
	// example:
	//
	// ruoyi
	HoneypotName *string `json:"HoneypotName,omitempty" xml:"HoneypotName,omitempty"`
	// The ID of the management node.
	//
	// example:
	//
	// a882e590-b87b-45a6-87b9-d0a3e5a0****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The ID of the custom configuration for the honeypot.
	//
	// example:
	//
	// ddh3731641137fe4b72b245346a2721d4b6tdgg3731641137fe4b72b245346a2721***
	PresetId *string `json:"PresetId,omitempty" xml:"PresetId,omitempty"`
	// An array that consists of the status information about the honeypot.
	State []*string `json:"State,omitempty" xml:"State,omitempty" type:"Repeated"`
}

func (s CreateHoneypotResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateHoneypotResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateHoneypotResponseBodyData) GetControlNodeName() *string {
	return s.ControlNodeName
}

func (s *CreateHoneypotResponseBodyData) GetHoneypotId() *string {
	return s.HoneypotId
}

func (s *CreateHoneypotResponseBodyData) GetHoneypotImageDisplayName() *string {
	return s.HoneypotImageDisplayName
}

func (s *CreateHoneypotResponseBodyData) GetHoneypotImageName() *string {
	return s.HoneypotImageName
}

func (s *CreateHoneypotResponseBodyData) GetHoneypotName() *string {
	return s.HoneypotName
}

func (s *CreateHoneypotResponseBodyData) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateHoneypotResponseBodyData) GetPresetId() *string {
	return s.PresetId
}

func (s *CreateHoneypotResponseBodyData) GetState() []*string {
	return s.State
}

func (s *CreateHoneypotResponseBodyData) SetControlNodeName(v string) *CreateHoneypotResponseBodyData {
	s.ControlNodeName = &v
	return s
}

func (s *CreateHoneypotResponseBodyData) SetHoneypotId(v string) *CreateHoneypotResponseBodyData {
	s.HoneypotId = &v
	return s
}

func (s *CreateHoneypotResponseBodyData) SetHoneypotImageDisplayName(v string) *CreateHoneypotResponseBodyData {
	s.HoneypotImageDisplayName = &v
	return s
}

func (s *CreateHoneypotResponseBodyData) SetHoneypotImageName(v string) *CreateHoneypotResponseBodyData {
	s.HoneypotImageName = &v
	return s
}

func (s *CreateHoneypotResponseBodyData) SetHoneypotName(v string) *CreateHoneypotResponseBodyData {
	s.HoneypotName = &v
	return s
}

func (s *CreateHoneypotResponseBodyData) SetNodeId(v string) *CreateHoneypotResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *CreateHoneypotResponseBodyData) SetPresetId(v string) *CreateHoneypotResponseBodyData {
	s.PresetId = &v
	return s
}

func (s *CreateHoneypotResponseBodyData) SetState(v []*string) *CreateHoneypotResponseBodyData {
	s.State = v
	return s
}

func (s *CreateHoneypotResponseBodyData) Validate() error {
	return dara.Validate(s)
}
