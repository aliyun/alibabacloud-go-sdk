// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotwordLibraryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreationTime(v string) *GetHotwordLibraryResponseBody
	GetCreationTime() *string
	SetDescription(v string) *GetHotwordLibraryResponseBody
	GetDescription() *string
	SetHotwordLibraryId(v string) *GetHotwordLibraryResponseBody
	GetHotwordLibraryId() *string
	SetHotwords(v []*Hotword) *GetHotwordLibraryResponseBody
	GetHotwords() []*Hotword
	SetName(v string) *GetHotwordLibraryResponseBody
	GetName() *string
	SetRequestId(v string) *GetHotwordLibraryResponseBody
	GetRequestId() *string
	SetUsageScenario(v string) *GetHotwordLibraryResponseBody
	GetUsageScenario() *string
}

type GetHotwordLibraryResponseBody struct {
	// example:
	//
	// 2020-12-23T13:33:49Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// 热词库描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// ****05512043f49f697f7425****
	HotwordLibraryId *string    `json:"HotwordLibraryId,omitempty" xml:"HotwordLibraryId,omitempty"`
	Hotwords         []*Hotword `json:"Hotwords,omitempty" xml:"Hotwords,omitempty" type:"Repeated"`
	// example:
	//
	// 热词库名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ****12e8864746a0a398****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ASR
	UsageScenario *string `json:"UsageScenario,omitempty" xml:"UsageScenario,omitempty"`
}

func (s GetHotwordLibraryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotwordLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotwordLibraryResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetHotwordLibraryResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetHotwordLibraryResponseBody) GetHotwordLibraryId() *string {
	return s.HotwordLibraryId
}

func (s *GetHotwordLibraryResponseBody) GetHotwords() []*Hotword {
	return s.Hotwords
}

func (s *GetHotwordLibraryResponseBody) GetName() *string {
	return s.Name
}

func (s *GetHotwordLibraryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotwordLibraryResponseBody) GetUsageScenario() *string {
	return s.UsageScenario
}

func (s *GetHotwordLibraryResponseBody) SetCreationTime(v string) *GetHotwordLibraryResponseBody {
	s.CreationTime = &v
	return s
}

func (s *GetHotwordLibraryResponseBody) SetDescription(v string) *GetHotwordLibraryResponseBody {
	s.Description = &v
	return s
}

func (s *GetHotwordLibraryResponseBody) SetHotwordLibraryId(v string) *GetHotwordLibraryResponseBody {
	s.HotwordLibraryId = &v
	return s
}

func (s *GetHotwordLibraryResponseBody) SetHotwords(v []*Hotword) *GetHotwordLibraryResponseBody {
	s.Hotwords = v
	return s
}

func (s *GetHotwordLibraryResponseBody) SetName(v string) *GetHotwordLibraryResponseBody {
	s.Name = &v
	return s
}

func (s *GetHotwordLibraryResponseBody) SetRequestId(v string) *GetHotwordLibraryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotwordLibraryResponseBody) SetUsageScenario(v string) *GetHotwordLibraryResponseBody {
	s.UsageScenario = &v
	return s
}

func (s *GetHotwordLibraryResponseBody) Validate() error {
	return dara.Validate(s)
}
