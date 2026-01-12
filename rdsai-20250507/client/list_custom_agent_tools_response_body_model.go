// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAgentToolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListCustomAgentToolsResponseBodyData) *ListCustomAgentToolsResponseBody
	GetData() []*ListCustomAgentToolsResponseBodyData
	SetRequestId(v string) *ListCustomAgentToolsResponseBody
	GetRequestId() *string
}

type ListCustomAgentToolsResponseBody struct {
	// The returned data.
	Data []*ListCustomAgentToolsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 32DEFB4A-861F-5D1D-ADD5-918E4FD7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCustomAgentToolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentToolsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomAgentToolsResponseBody) GetData() []*ListCustomAgentToolsResponseBodyData {
	return s.Data
}

func (s *ListCustomAgentToolsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomAgentToolsResponseBody) SetData(v []*ListCustomAgentToolsResponseBodyData) *ListCustomAgentToolsResponseBody {
	s.Data = v
	return s
}

func (s *ListCustomAgentToolsResponseBody) SetRequestId(v string) *ListCustomAgentToolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomAgentToolsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomAgentToolsResponseBodyData struct {
	// The description in English.
	//
	// example:
	//
	// Queries RDS instances.
	En *string `json:"En,omitempty" xml:"En,omitempty"`
	// The description in Japanese.
	Ja *string `json:"Ja,omitempty" xml:"Ja,omitempty"`
	// The tool name.
	//
	// example:
	//
	// describe_db_instances
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description in Traditional Chinese.
	Tc *string `json:"Tc,omitempty" xml:"Tc,omitempty"`
	// The read/write type of the tool.
	//
	// example:
	//
	// get
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The description in Simplified Chinese.
	Zh *string `json:"Zh,omitempty" xml:"Zh,omitempty"`
}

func (s ListCustomAgentToolsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentToolsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCustomAgentToolsResponseBodyData) GetEn() *string {
	return s.En
}

func (s *ListCustomAgentToolsResponseBodyData) GetJa() *string {
	return s.Ja
}

func (s *ListCustomAgentToolsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListCustomAgentToolsResponseBodyData) GetTc() *string {
	return s.Tc
}

func (s *ListCustomAgentToolsResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListCustomAgentToolsResponseBodyData) GetZh() *string {
	return s.Zh
}

func (s *ListCustomAgentToolsResponseBodyData) SetEn(v string) *ListCustomAgentToolsResponseBodyData {
	s.En = &v
	return s
}

func (s *ListCustomAgentToolsResponseBodyData) SetJa(v string) *ListCustomAgentToolsResponseBodyData {
	s.Ja = &v
	return s
}

func (s *ListCustomAgentToolsResponseBodyData) SetName(v string) *ListCustomAgentToolsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListCustomAgentToolsResponseBodyData) SetTc(v string) *ListCustomAgentToolsResponseBodyData {
	s.Tc = &v
	return s
}

func (s *ListCustomAgentToolsResponseBodyData) SetType(v string) *ListCustomAgentToolsResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListCustomAgentToolsResponseBodyData) SetZh(v string) *ListCustomAgentToolsResponseBodyData {
	s.Zh = &v
	return s
}

func (s *ListCustomAgentToolsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
