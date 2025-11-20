// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateListByUserIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextCursor(v int64) *GetTemplateListByUserIdResponseBody
	GetNextCursor() *int64
	SetRequestId(v string) *GetTemplateListByUserIdResponseBody
	GetRequestId() *string
	SetTemplateList(v []*GetTemplateListByUserIdResponseBodyTemplateList) *GetTemplateListByUserIdResponseBody
	GetTemplateList() []*GetTemplateListByUserIdResponseBodyTemplateList
}

type GetTemplateListByUserIdResponseBody struct {
	// example:
	//
	// 12312131231
	NextCursor *int64 `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId    *string                                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TemplateList []*GetTemplateListByUserIdResponseBodyTemplateList `json:"templateList,omitempty" xml:"templateList,omitempty" type:"Repeated"`
}

func (s GetTemplateListByUserIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateListByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateListByUserIdResponseBody) GetNextCursor() *int64 {
	return s.NextCursor
}

func (s *GetTemplateListByUserIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemplateListByUserIdResponseBody) GetTemplateList() []*GetTemplateListByUserIdResponseBodyTemplateList {
	return s.TemplateList
}

func (s *GetTemplateListByUserIdResponseBody) SetNextCursor(v int64) *GetTemplateListByUserIdResponseBody {
	s.NextCursor = &v
	return s
}

func (s *GetTemplateListByUserIdResponseBody) SetRequestId(v string) *GetTemplateListByUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateListByUserIdResponseBody) SetTemplateList(v []*GetTemplateListByUserIdResponseBodyTemplateList) *GetTemplateListByUserIdResponseBody {
	s.TemplateList = v
	return s
}

func (s *GetTemplateListByUserIdResponseBody) Validate() error {
	if s.TemplateList != nil {
		for _, item := range s.TemplateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTemplateListByUserIdResponseBodyTemplateList struct {
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 11111
	ReportCode *string `json:"ReportCode,omitempty" xml:"ReportCode,omitempty"`
	// example:
	//
	// https://scsss/sss
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetTemplateListByUserIdResponseBodyTemplateList) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateListByUserIdResponseBodyTemplateList) GoString() string {
	return s.String()
}

func (s *GetTemplateListByUserIdResponseBodyTemplateList) GetIconUrl() *string {
	return s.IconUrl
}

func (s *GetTemplateListByUserIdResponseBodyTemplateList) GetName() *string {
	return s.Name
}

func (s *GetTemplateListByUserIdResponseBodyTemplateList) GetReportCode() *string {
	return s.ReportCode
}

func (s *GetTemplateListByUserIdResponseBodyTemplateList) GetUrl() *string {
	return s.Url
}

func (s *GetTemplateListByUserIdResponseBodyTemplateList) SetIconUrl(v string) *GetTemplateListByUserIdResponseBodyTemplateList {
	s.IconUrl = &v
	return s
}

func (s *GetTemplateListByUserIdResponseBodyTemplateList) SetName(v string) *GetTemplateListByUserIdResponseBodyTemplateList {
	s.Name = &v
	return s
}

func (s *GetTemplateListByUserIdResponseBodyTemplateList) SetReportCode(v string) *GetTemplateListByUserIdResponseBodyTemplateList {
	s.ReportCode = &v
	return s
}

func (s *GetTemplateListByUserIdResponseBodyTemplateList) SetUrl(v string) *GetTemplateListByUserIdResponseBodyTemplateList {
	s.Url = &v
	return s
}

func (s *GetTemplateListByUserIdResponseBodyTemplateList) Validate() error {
	return dara.Validate(s)
}
