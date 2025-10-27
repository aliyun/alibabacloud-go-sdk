// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHoneypotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateHoneypotResponseBody
	GetCode() *string
	SetData(v *UpdateHoneypotResponseBodyData) *UpdateHoneypotResponseBody
	GetData() *UpdateHoneypotResponseBodyData
	SetHttpStatusCode(v int32) *UpdateHoneypotResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateHoneypotResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateHoneypotResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateHoneypotResponseBody
	GetSuccess() *bool
}

type UpdateHoneypotResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the honeypot.
	Data *UpdateHoneypotResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3B323ADD-6CF8-51F6-9047-2F0A4E3F5EFD
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

func (s UpdateHoneypotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHoneypotResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHoneypotResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateHoneypotResponseBody) GetData() *UpdateHoneypotResponseBodyData {
	return s.Data
}

func (s *UpdateHoneypotResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateHoneypotResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateHoneypotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHoneypotResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateHoneypotResponseBody) SetCode(v string) *UpdateHoneypotResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHoneypotResponseBody) SetData(v *UpdateHoneypotResponseBodyData) *UpdateHoneypotResponseBody {
	s.Data = v
	return s
}

func (s *UpdateHoneypotResponseBody) SetHttpStatusCode(v int32) *UpdateHoneypotResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateHoneypotResponseBody) SetMessage(v string) *UpdateHoneypotResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHoneypotResponseBody) SetRequestId(v string) *UpdateHoneypotResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHoneypotResponseBody) SetSuccess(v bool) *UpdateHoneypotResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateHoneypotResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateHoneypotResponseBodyData struct {
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
	// Webmin
	HoneypotImageDisplayName *string `json:"HoneypotImageDisplayName,omitempty" xml:"HoneypotImageDisplayName,omitempty"`
	// The name of the honeypot image.
	//
	// example:
	//
	// webpage
	HoneypotImageName *string `json:"HoneypotImageName,omitempty" xml:"HoneypotImageName,omitempty"`
	// The custom name of the honeypot.
	//
	// example:
	//
	// hyl-phpmyadmin
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
	// 868a7579-00b5-4a74-999d-8bd3f411e8a3
	PresetId *string `json:"PresetId,omitempty" xml:"PresetId,omitempty"`
	// An array that consists of the status information about the honeypot.
	State []*string `json:"State,omitempty" xml:"State,omitempty" type:"Repeated"`
}

func (s UpdateHoneypotResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateHoneypotResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateHoneypotResponseBodyData) GetHoneypotId() *string {
	return s.HoneypotId
}

func (s *UpdateHoneypotResponseBodyData) GetHoneypotImageDisplayName() *string {
	return s.HoneypotImageDisplayName
}

func (s *UpdateHoneypotResponseBodyData) GetHoneypotImageName() *string {
	return s.HoneypotImageName
}

func (s *UpdateHoneypotResponseBodyData) GetHoneypotName() *string {
	return s.HoneypotName
}

func (s *UpdateHoneypotResponseBodyData) GetNodeId() *string {
	return s.NodeId
}

func (s *UpdateHoneypotResponseBodyData) GetPresetId() *string {
	return s.PresetId
}

func (s *UpdateHoneypotResponseBodyData) GetState() []*string {
	return s.State
}

func (s *UpdateHoneypotResponseBodyData) SetHoneypotId(v string) *UpdateHoneypotResponseBodyData {
	s.HoneypotId = &v
	return s
}

func (s *UpdateHoneypotResponseBodyData) SetHoneypotImageDisplayName(v string) *UpdateHoneypotResponseBodyData {
	s.HoneypotImageDisplayName = &v
	return s
}

func (s *UpdateHoneypotResponseBodyData) SetHoneypotImageName(v string) *UpdateHoneypotResponseBodyData {
	s.HoneypotImageName = &v
	return s
}

func (s *UpdateHoneypotResponseBodyData) SetHoneypotName(v string) *UpdateHoneypotResponseBodyData {
	s.HoneypotName = &v
	return s
}

func (s *UpdateHoneypotResponseBodyData) SetNodeId(v string) *UpdateHoneypotResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *UpdateHoneypotResponseBodyData) SetPresetId(v string) *UpdateHoneypotResponseBodyData {
	s.PresetId = &v
	return s
}

func (s *UpdateHoneypotResponseBodyData) SetState(v []*string) *UpdateHoneypotResponseBodyData {
	s.State = v
	return s
}

func (s *UpdateHoneypotResponseBodyData) Validate() error {
	return dara.Validate(s)
}
