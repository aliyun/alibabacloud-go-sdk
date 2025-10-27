// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetHoneypotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ResetHoneypotResponseBodyData) *ResetHoneypotResponseBody
	GetData() *ResetHoneypotResponseBodyData
	SetRequestId(v string) *ResetHoneypotResponseBody
	GetRequestId() *string
}

type ResetHoneypotResponseBody struct {
	// The information about the honeypot.
	Data *ResetHoneypotResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A57C711B-AA15-55B2-8F61-4D09CEXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetHoneypotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetHoneypotResponseBody) GoString() string {
	return s.String()
}

func (s *ResetHoneypotResponseBody) GetData() *ResetHoneypotResponseBodyData {
	return s.Data
}

func (s *ResetHoneypotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetHoneypotResponseBody) SetData(v *ResetHoneypotResponseBodyData) *ResetHoneypotResponseBody {
	s.Data = v
	return s
}

func (s *ResetHoneypotResponseBody) SetRequestId(v string) *ResetHoneypotResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetHoneypotResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResetHoneypotResponseBodyData struct {
	// The name of the management node to which the honeypot belongs.
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
	// ruoyi
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
	// 868a7579-00b5-4a74-999d-8bd3f411****
	PresetId *string `json:"PresetId,omitempty" xml:"PresetId,omitempty"`
	// The statuses of the honeypots.
	State []*string `json:"State,omitempty" xml:"State,omitempty" type:"Repeated"`
}

func (s ResetHoneypotResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ResetHoneypotResponseBodyData) GoString() string {
	return s.String()
}

func (s *ResetHoneypotResponseBodyData) GetControlNodeName() *string {
	return s.ControlNodeName
}

func (s *ResetHoneypotResponseBodyData) GetHoneypotId() *string {
	return s.HoneypotId
}

func (s *ResetHoneypotResponseBodyData) GetHoneypotImageDisplayName() *string {
	return s.HoneypotImageDisplayName
}

func (s *ResetHoneypotResponseBodyData) GetHoneypotImageName() *string {
	return s.HoneypotImageName
}

func (s *ResetHoneypotResponseBodyData) GetHoneypotName() *string {
	return s.HoneypotName
}

func (s *ResetHoneypotResponseBodyData) GetNodeId() *string {
	return s.NodeId
}

func (s *ResetHoneypotResponseBodyData) GetPresetId() *string {
	return s.PresetId
}

func (s *ResetHoneypotResponseBodyData) GetState() []*string {
	return s.State
}

func (s *ResetHoneypotResponseBodyData) SetControlNodeName(v string) *ResetHoneypotResponseBodyData {
	s.ControlNodeName = &v
	return s
}

func (s *ResetHoneypotResponseBodyData) SetHoneypotId(v string) *ResetHoneypotResponseBodyData {
	s.HoneypotId = &v
	return s
}

func (s *ResetHoneypotResponseBodyData) SetHoneypotImageDisplayName(v string) *ResetHoneypotResponseBodyData {
	s.HoneypotImageDisplayName = &v
	return s
}

func (s *ResetHoneypotResponseBodyData) SetHoneypotImageName(v string) *ResetHoneypotResponseBodyData {
	s.HoneypotImageName = &v
	return s
}

func (s *ResetHoneypotResponseBodyData) SetHoneypotName(v string) *ResetHoneypotResponseBodyData {
	s.HoneypotName = &v
	return s
}

func (s *ResetHoneypotResponseBodyData) SetNodeId(v string) *ResetHoneypotResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *ResetHoneypotResponseBodyData) SetPresetId(v string) *ResetHoneypotResponseBodyData {
	s.PresetId = &v
	return s
}

func (s *ResetHoneypotResponseBodyData) SetState(v []*string) *ResetHoneypotResponseBodyData {
	s.State = v
	return s
}

func (s *ResetHoneypotResponseBodyData) Validate() error {
	return dara.Validate(s)
}
