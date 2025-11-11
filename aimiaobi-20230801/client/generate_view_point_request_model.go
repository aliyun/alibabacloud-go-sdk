// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateViewPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GenerateViewPointRequest
	GetAgentKey() *string
	SetReferenceData(v *GenerateViewPointRequestReferenceData) *GenerateViewPointRequest
	GetReferenceData() *GenerateViewPointRequestReferenceData
}

type GenerateViewPointRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c160c841c8e54295bf2f441432785944_p_efm
	AgentKey      *string                                `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	ReferenceData *GenerateViewPointRequestReferenceData `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty" type:"Struct"`
}

func (s GenerateViewPointRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateViewPointRequest) GoString() string {
	return s.String()
}

func (s *GenerateViewPointRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GenerateViewPointRequest) GetReferenceData() *GenerateViewPointRequestReferenceData {
	return s.ReferenceData
}

func (s *GenerateViewPointRequest) SetAgentKey(v string) *GenerateViewPointRequest {
	s.AgentKey = &v
	return s
}

func (s *GenerateViewPointRequest) SetReferenceData(v *GenerateViewPointRequestReferenceData) *GenerateViewPointRequest {
	s.ReferenceData = v
	return s
}

func (s *GenerateViewPointRequest) Validate() error {
	if s.ReferenceData != nil {
		if err := s.ReferenceData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateViewPointRequestReferenceData struct {
	MiniDoc []*string `json:"MiniDoc,omitempty" xml:"MiniDoc,omitempty" type:"Repeated"`
}

func (s GenerateViewPointRequestReferenceData) String() string {
	return dara.Prettify(s)
}

func (s GenerateViewPointRequestReferenceData) GoString() string {
	return s.String()
}

func (s *GenerateViewPointRequestReferenceData) GetMiniDoc() []*string {
	return s.MiniDoc
}

func (s *GenerateViewPointRequestReferenceData) SetMiniDoc(v []*string) *GenerateViewPointRequestReferenceData {
	s.MiniDoc = v
	return s
}

func (s *GenerateViewPointRequestReferenceData) Validate() error {
	return dara.Validate(s)
}
