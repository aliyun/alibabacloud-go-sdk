// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartHoneypotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *StartHoneypotResponseBodyData) *StartHoneypotResponseBody
	GetData() *StartHoneypotResponseBodyData
	SetRequestId(v string) *StartHoneypotResponseBody
	GetRequestId() *string
}

type StartHoneypotResponseBody struct {
	// The information about the honeypot.
	Data *StartHoneypotResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 20456DD5-5CBF-5015-9173-12CA4246B***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartHoneypotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartHoneypotResponseBody) GoString() string {
	return s.String()
}

func (s *StartHoneypotResponseBody) GetData() *StartHoneypotResponseBodyData {
	return s.Data
}

func (s *StartHoneypotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartHoneypotResponseBody) SetData(v *StartHoneypotResponseBodyData) *StartHoneypotResponseBody {
	s.Data = v
	return s
}

func (s *StartHoneypotResponseBody) SetRequestId(v string) *StartHoneypotResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartHoneypotResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartHoneypotResponseBodyData struct {
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
	// 123
	HoneypotId *string `json:"HoneypotId,omitempty" xml:"HoneypotId,omitempty"`
	// The display name of the image.
	//
	// example:
	//
	// Webmin
	HoneypotImageDisplayName *string `json:"HoneypotImageDisplayName,omitempty" xml:"HoneypotImageDisplayName,omitempty"`
	// The name of the image that is used for the honeypot.
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
	// The statuses of the honeypots.
	State []*string `json:"State,omitempty" xml:"State,omitempty" type:"Repeated"`
}

func (s StartHoneypotResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StartHoneypotResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartHoneypotResponseBodyData) GetControlNodeName() *string {
	return s.ControlNodeName
}

func (s *StartHoneypotResponseBodyData) GetHoneypotId() *string {
	return s.HoneypotId
}

func (s *StartHoneypotResponseBodyData) GetHoneypotImageDisplayName() *string {
	return s.HoneypotImageDisplayName
}

func (s *StartHoneypotResponseBodyData) GetHoneypotImageName() *string {
	return s.HoneypotImageName
}

func (s *StartHoneypotResponseBodyData) GetHoneypotName() *string {
	return s.HoneypotName
}

func (s *StartHoneypotResponseBodyData) GetNodeId() *string {
	return s.NodeId
}

func (s *StartHoneypotResponseBodyData) GetPresetId() *string {
	return s.PresetId
}

func (s *StartHoneypotResponseBodyData) GetState() []*string {
	return s.State
}

func (s *StartHoneypotResponseBodyData) SetControlNodeName(v string) *StartHoneypotResponseBodyData {
	s.ControlNodeName = &v
	return s
}

func (s *StartHoneypotResponseBodyData) SetHoneypotId(v string) *StartHoneypotResponseBodyData {
	s.HoneypotId = &v
	return s
}

func (s *StartHoneypotResponseBodyData) SetHoneypotImageDisplayName(v string) *StartHoneypotResponseBodyData {
	s.HoneypotImageDisplayName = &v
	return s
}

func (s *StartHoneypotResponseBodyData) SetHoneypotImageName(v string) *StartHoneypotResponseBodyData {
	s.HoneypotImageName = &v
	return s
}

func (s *StartHoneypotResponseBodyData) SetHoneypotName(v string) *StartHoneypotResponseBodyData {
	s.HoneypotName = &v
	return s
}

func (s *StartHoneypotResponseBodyData) SetNodeId(v string) *StartHoneypotResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *StartHoneypotResponseBodyData) SetPresetId(v string) *StartHoneypotResponseBodyData {
	s.PresetId = &v
	return s
}

func (s *StartHoneypotResponseBodyData) SetState(v []*string) *StartHoneypotResponseBodyData {
	s.State = v
	return s
}

func (s *StartHoneypotResponseBodyData) Validate() error {
	return dara.Validate(s)
}
