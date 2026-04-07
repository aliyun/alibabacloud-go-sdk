// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDmsKnowledgeSearchOrderInfoDO interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunAccountUid(v string) *DmsKnowledgeSearchOrderInfoDO
	GetAliyunAccountUid() *string
	SetApiKey(v string) *DmsKnowledgeSearchOrderInfoDO
	GetApiKey() *string
	SetGmtCreate(v string) *DmsKnowledgeSearchOrderInfoDO
	GetGmtCreate() *string
	SetGmtModified(v string) *DmsKnowledgeSearchOrderInfoDO
	GetGmtModified() *string
	SetId(v int64) *DmsKnowledgeSearchOrderInfoDO
	GetId() *int64
	SetOrderId(v string) *DmsKnowledgeSearchOrderInfoDO
	GetOrderId() *string
	SetWebSearchApiUrl(v string) *DmsKnowledgeSearchOrderInfoDO
	GetWebSearchApiUrl() *string
}

type DmsKnowledgeSearchOrderInfoDO struct {
	AliyunAccountUid *string `json:"AliyunAccountUid,omitempty" xml:"AliyunAccountUid,omitempty"`
	ApiKey           *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	GmtCreate        *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified      *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	OrderId          *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	WebSearchApiUrl  *string `json:"WebSearchApiUrl,omitempty" xml:"WebSearchApiUrl,omitempty"`
}

func (s DmsKnowledgeSearchOrderInfoDO) String() string {
	return dara.Prettify(s)
}

func (s DmsKnowledgeSearchOrderInfoDO) GoString() string {
	return s.String()
}

func (s *DmsKnowledgeSearchOrderInfoDO) GetAliyunAccountUid() *string {
	return s.AliyunAccountUid
}

func (s *DmsKnowledgeSearchOrderInfoDO) GetApiKey() *string {
	return s.ApiKey
}

func (s *DmsKnowledgeSearchOrderInfoDO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DmsKnowledgeSearchOrderInfoDO) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DmsKnowledgeSearchOrderInfoDO) GetId() *int64 {
	return s.Id
}

func (s *DmsKnowledgeSearchOrderInfoDO) GetOrderId() *string {
	return s.OrderId
}

func (s *DmsKnowledgeSearchOrderInfoDO) GetWebSearchApiUrl() *string {
	return s.WebSearchApiUrl
}

func (s *DmsKnowledgeSearchOrderInfoDO) SetAliyunAccountUid(v string) *DmsKnowledgeSearchOrderInfoDO {
	s.AliyunAccountUid = &v
	return s
}

func (s *DmsKnowledgeSearchOrderInfoDO) SetApiKey(v string) *DmsKnowledgeSearchOrderInfoDO {
	s.ApiKey = &v
	return s
}

func (s *DmsKnowledgeSearchOrderInfoDO) SetGmtCreate(v string) *DmsKnowledgeSearchOrderInfoDO {
	s.GmtCreate = &v
	return s
}

func (s *DmsKnowledgeSearchOrderInfoDO) SetGmtModified(v string) *DmsKnowledgeSearchOrderInfoDO {
	s.GmtModified = &v
	return s
}

func (s *DmsKnowledgeSearchOrderInfoDO) SetId(v int64) *DmsKnowledgeSearchOrderInfoDO {
	s.Id = &v
	return s
}

func (s *DmsKnowledgeSearchOrderInfoDO) SetOrderId(v string) *DmsKnowledgeSearchOrderInfoDO {
	s.OrderId = &v
	return s
}

func (s *DmsKnowledgeSearchOrderInfoDO) SetWebSearchApiUrl(v string) *DmsKnowledgeSearchOrderInfoDO {
	s.WebSearchApiUrl = &v
	return s
}

func (s *DmsKnowledgeSearchOrderInfoDO) Validate() error {
	return dara.Validate(s)
}
