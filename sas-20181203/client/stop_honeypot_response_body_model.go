// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopHoneypotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopHoneypotResponseBody
	GetCode() *string
	SetData(v *StopHoneypotResponseBodyData) *StopHoneypotResponseBody
	GetData() *StopHoneypotResponseBodyData
	SetHttpStatusCode(v int32) *StopHoneypotResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StopHoneypotResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopHoneypotResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopHoneypotResponseBody
	GetSuccess() *bool
}

type StopHoneypotResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *StopHoneypotResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D65AADFC-1D20-5A6A-8F6A-9FA53C*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopHoneypotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopHoneypotResponseBody) GoString() string {
	return s.String()
}

func (s *StopHoneypotResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopHoneypotResponseBody) GetData() *StopHoneypotResponseBodyData {
	return s.Data
}

func (s *StopHoneypotResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StopHoneypotResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopHoneypotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopHoneypotResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopHoneypotResponseBody) SetCode(v string) *StopHoneypotResponseBody {
	s.Code = &v
	return s
}

func (s *StopHoneypotResponseBody) SetData(v *StopHoneypotResponseBodyData) *StopHoneypotResponseBody {
	s.Data = v
	return s
}

func (s *StopHoneypotResponseBody) SetHttpStatusCode(v int32) *StopHoneypotResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopHoneypotResponseBody) SetMessage(v string) *StopHoneypotResponseBody {
	s.Message = &v
	return s
}

func (s *StopHoneypotResponseBody) SetRequestId(v string) *StopHoneypotResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopHoneypotResponseBody) SetSuccess(v bool) *StopHoneypotResponseBody {
	s.Success = &v
	return s
}

func (s *StopHoneypotResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StopHoneypotResponseBodyData struct {
	// The name of the management node to which the honeypot belongs.
	//
	// example:
	//
	// managerNoden****
	ControlNodeName *string `json:"ControlNodeName,omitempty" xml:"ControlNodeName,omitempty"`
	// The ID of the honeypot.
	//
	// example:
	//
	// 9bf8cd373112263d4bc102fc5dba9d9f812ee05d4d35c487d330d52e937f****
	HoneypotId *string `json:"HoneypotId,omitempty" xml:"HoneypotId,omitempty"`
	// The display name of the image.
	//
	// example:
	//
	// RuoYi
	HoneypotImageDisplayName *string `json:"HoneypotImageDisplayName,omitempty" xml:"HoneypotImageDisplayName,omitempty"`
	// The name of the image that is used for the honeypot.
	//
	// example:
	//
	// metabase
	HoneypotImageName *string `json:"HoneypotImageName,omitempty" xml:"HoneypotImageName,omitempty"`
	// The name of the honeypot.
	//
	// example:
	//
	// hyl-phpmya****
	HoneypotName *string `json:"HoneypotName,omitempty" xml:"HoneypotName,omitempty"`
	// The ID of the management node.
	//
	// example:
	//
	// a882e590-b87b-45a6-87b9-d0a3e5a0****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The ID of the honeypot custom parameter.
	//
	// example:
	//
	// 868a7579-00b5-4a74-999d-8bd3f411****
	PresetId *string `json:"PresetId,omitempty" xml:"PresetId,omitempty"`
	// The statuses of the honeypots.
	State []*string `json:"State,omitempty" xml:"State,omitempty" type:"Repeated"`
}

func (s StopHoneypotResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StopHoneypotResponseBodyData) GoString() string {
	return s.String()
}

func (s *StopHoneypotResponseBodyData) GetControlNodeName() *string {
	return s.ControlNodeName
}

func (s *StopHoneypotResponseBodyData) GetHoneypotId() *string {
	return s.HoneypotId
}

func (s *StopHoneypotResponseBodyData) GetHoneypotImageDisplayName() *string {
	return s.HoneypotImageDisplayName
}

func (s *StopHoneypotResponseBodyData) GetHoneypotImageName() *string {
	return s.HoneypotImageName
}

func (s *StopHoneypotResponseBodyData) GetHoneypotName() *string {
	return s.HoneypotName
}

func (s *StopHoneypotResponseBodyData) GetNodeId() *string {
	return s.NodeId
}

func (s *StopHoneypotResponseBodyData) GetPresetId() *string {
	return s.PresetId
}

func (s *StopHoneypotResponseBodyData) GetState() []*string {
	return s.State
}

func (s *StopHoneypotResponseBodyData) SetControlNodeName(v string) *StopHoneypotResponseBodyData {
	s.ControlNodeName = &v
	return s
}

func (s *StopHoneypotResponseBodyData) SetHoneypotId(v string) *StopHoneypotResponseBodyData {
	s.HoneypotId = &v
	return s
}

func (s *StopHoneypotResponseBodyData) SetHoneypotImageDisplayName(v string) *StopHoneypotResponseBodyData {
	s.HoneypotImageDisplayName = &v
	return s
}

func (s *StopHoneypotResponseBodyData) SetHoneypotImageName(v string) *StopHoneypotResponseBodyData {
	s.HoneypotImageName = &v
	return s
}

func (s *StopHoneypotResponseBodyData) SetHoneypotName(v string) *StopHoneypotResponseBodyData {
	s.HoneypotName = &v
	return s
}

func (s *StopHoneypotResponseBodyData) SetNodeId(v string) *StopHoneypotResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *StopHoneypotResponseBodyData) SetPresetId(v string) *StopHoneypotResponseBodyData {
	s.PresetId = &v
	return s
}

func (s *StopHoneypotResponseBodyData) SetState(v []*string) *StopHoneypotResponseBodyData {
	s.State = v
	return s
}

func (s *StopHoneypotResponseBodyData) Validate() error {
	return dara.Validate(s)
}
