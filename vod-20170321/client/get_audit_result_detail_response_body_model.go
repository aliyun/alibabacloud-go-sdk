// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditResultDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIAuditResultDetail(v *GetAuditResultDetailResponseBodyAIAuditResultDetail) *GetAuditResultDetailResponseBody
	GetAIAuditResultDetail() *GetAuditResultDetailResponseBodyAIAuditResultDetail
	SetRequestId(v string) *GetAuditResultDetailResponseBody
	GetRequestId() *string
}

type GetAuditResultDetailResponseBody struct {
	AIAuditResultDetail *GetAuditResultDetailResponseBodyAIAuditResultDetail `json:"AIAuditResultDetail,omitempty" xml:"AIAuditResultDetail,omitempty" type:"Struct"`
	RequestId           *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAuditResultDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAuditResultDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuditResultDetailResponseBody) GetAIAuditResultDetail() *GetAuditResultDetailResponseBodyAIAuditResultDetail {
	return s.AIAuditResultDetail
}

func (s *GetAuditResultDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAuditResultDetailResponseBody) SetAIAuditResultDetail(v *GetAuditResultDetailResponseBodyAIAuditResultDetail) *GetAuditResultDetailResponseBody {
	s.AIAuditResultDetail = v
	return s
}

func (s *GetAuditResultDetailResponseBody) SetRequestId(v string) *GetAuditResultDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuditResultDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAuditResultDetailResponseBodyAIAuditResultDetail struct {
	List  []*GetAuditResultDetailResponseBodyAIAuditResultDetailList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Total *int32                                                     `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetAuditResultDetailResponseBodyAIAuditResultDetail) String() string {
	return dara.Prettify(s)
}

func (s GetAuditResultDetailResponseBodyAIAuditResultDetail) GoString() string {
	return s.String()
}

func (s *GetAuditResultDetailResponseBodyAIAuditResultDetail) GetList() []*GetAuditResultDetailResponseBodyAIAuditResultDetailList {
	return s.List
}

func (s *GetAuditResultDetailResponseBodyAIAuditResultDetail) GetTotal() *int32 {
	return s.Total
}

func (s *GetAuditResultDetailResponseBodyAIAuditResultDetail) SetList(v []*GetAuditResultDetailResponseBodyAIAuditResultDetailList) *GetAuditResultDetailResponseBodyAIAuditResultDetail {
	s.List = v
	return s
}

func (s *GetAuditResultDetailResponseBodyAIAuditResultDetail) SetTotal(v int32) *GetAuditResultDetailResponseBodyAIAuditResultDetail {
	s.Total = &v
	return s
}

func (s *GetAuditResultDetailResponseBodyAIAuditResultDetail) Validate() error {
	return dara.Validate(s)
}

type GetAuditResultDetailResponseBodyAIAuditResultDetailList struct {
	Index          *string `json:"Index,omitempty" xml:"Index,omitempty"`
	Object         *string `json:"Object,omitempty" xml:"Object,omitempty"`
	PornLabel      *string `json:"PornLabel,omitempty" xml:"PornLabel,omitempty"`
	PornScore      *string `json:"PornScore,omitempty" xml:"PornScore,omitempty"`
	TerrorismLabel *string `json:"TerrorismLabel,omitempty" xml:"TerrorismLabel,omitempty"`
	TerrorismScore *string `json:"TerrorismScore,omitempty" xml:"TerrorismScore,omitempty"`
	Timestamp      *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetAuditResultDetailResponseBodyAIAuditResultDetailList) String() string {
	return dara.Prettify(s)
}

func (s GetAuditResultDetailResponseBodyAIAuditResultDetailList) GoString() string {
	return s.String()
}

func (s *GetAuditResultDetailResponseBodyAIAuditResultDetailList) GetIndex() *string {
	return s.Index
}

func (s *GetAuditResultDetailResponseBodyAIAuditResultDetailList) GetObject() *string {
	return s.Object
}

func (s *GetAuditResultDetailResponseBodyAIAuditResultDetailList) GetPornLabel() *string {
	return s.PornLabel
}

func (s *GetAuditResultDetailResponseBodyAIAuditResultDetailList) GetPornScore() *string {
	return s.PornScore
}

func (s *GetAuditResultDetailResponseBodyAIAuditResultDetailList) GetTerrorismLabel() *string {
	return s.TerrorismLabel
}

func (s *GetAuditResultDetailResponseBodyAIAuditResultDetailList) GetTerrorismScore() *string {
	return s.TerrorismScore
}

func (s *GetAuditResultDetailResponseBodyAIAuditResultDetailList) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetAuditResultDetailResponseBodyAIAuditResultDetailList) SetIndex(v string) *GetAuditResultDetailResponseBodyAIAuditResultDetailList {
	s.Index = &v
	return s
}

func (s *GetAuditResultDetailResponseBodyAIAuditResultDetailList) SetObject(v string) *GetAuditResultDetailResponseBodyAIAuditResultDetailList {
	s.Object = &v
	return s
}

func (s *GetAuditResultDetailResponseBodyAIAuditResultDetailList) SetPornLabel(v string) *GetAuditResultDetailResponseBodyAIAuditResultDetailList {
	s.PornLabel = &v
	return s
}

func (s *GetAuditResultDetailResponseBodyAIAuditResultDetailList) SetPornScore(v string) *GetAuditResultDetailResponseBodyAIAuditResultDetailList {
	s.PornScore = &v
	return s
}

func (s *GetAuditResultDetailResponseBodyAIAuditResultDetailList) SetTerrorismLabel(v string) *GetAuditResultDetailResponseBodyAIAuditResultDetailList {
	s.TerrorismLabel = &v
	return s
}

func (s *GetAuditResultDetailResponseBodyAIAuditResultDetailList) SetTerrorismScore(v string) *GetAuditResultDetailResponseBodyAIAuditResultDetailList {
	s.TerrorismScore = &v
	return s
}

func (s *GetAuditResultDetailResponseBodyAIAuditResultDetailList) SetTimestamp(v string) *GetAuditResultDetailResponseBodyAIAuditResultDetailList {
	s.Timestamp = &v
	return s
}

func (s *GetAuditResultDetailResponseBodyAIAuditResultDetailList) Validate() error {
	return dara.Validate(s)
}
