// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTitleIntelligenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTitleIntelligenceResponseBody
	GetCode() *int32
	SetData(v *GetTitleIntelligenceResponseBodyData) *GetTitleIntelligenceResponseBody
	GetData() *GetTitleIntelligenceResponseBodyData
	SetMessage(v string) *GetTitleIntelligenceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTitleIntelligenceResponseBody
	GetRequestId() *string
}

type GetTitleIntelligenceResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTitleIntelligenceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D70487B4-8891-4D24-BB64-8788E53671FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTitleIntelligenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTitleIntelligenceResponseBody) GoString() string {
	return s.String()
}

func (s *GetTitleIntelligenceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTitleIntelligenceResponseBody) GetData() *GetTitleIntelligenceResponseBodyData {
	return s.Data
}

func (s *GetTitleIntelligenceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTitleIntelligenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTitleIntelligenceResponseBody) SetCode(v int32) *GetTitleIntelligenceResponseBody {
	s.Code = &v
	return s
}

func (s *GetTitleIntelligenceResponseBody) SetData(v *GetTitleIntelligenceResponseBodyData) *GetTitleIntelligenceResponseBody {
	s.Data = v
	return s
}

func (s *GetTitleIntelligenceResponseBody) SetMessage(v string) *GetTitleIntelligenceResponseBody {
	s.Message = &v
	return s
}

func (s *GetTitleIntelligenceResponseBody) SetRequestId(v string) *GetTitleIntelligenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTitleIntelligenceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTitleIntelligenceResponseBodyData struct {
	// example:
	//
	// Custom Hello Kitty PVC Cartoon Apple for Home Garden Complete Apple Bath Shower
	Titles *string `json:"Titles,omitempty" xml:"Titles,omitempty"`
}

func (s GetTitleIntelligenceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTitleIntelligenceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTitleIntelligenceResponseBodyData) GetTitles() *string {
	return s.Titles
}

func (s *GetTitleIntelligenceResponseBodyData) SetTitles(v string) *GetTitleIntelligenceResponseBodyData {
	s.Titles = &v
	return s
}

func (s *GetTitleIntelligenceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
