// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingTalkUserOrgByAliyunTmpCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDingTalkUserOrgByAliyunTmpCodeResponseBody
	GetCode() *string
	SetData(v *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyData) *GetDingTalkUserOrgByAliyunTmpCodeResponseBody
	GetData() *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyData
	SetHttpCode(v string) *GetDingTalkUserOrgByAliyunTmpCodeResponseBody
	GetHttpCode() *string
	SetMessage(v string) *GetDingTalkUserOrgByAliyunTmpCodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDingTalkUserOrgByAliyunTmpCodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDingTalkUserOrgByAliyunTmpCodeResponseBody
	GetSuccess() *bool
}

type GetDingTalkUserOrgByAliyunTmpCodeResponseBody struct {
	Code      *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpCode  *string                                            `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDingTalkUserOrgByAliyunTmpCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDingTalkUserOrgByAliyunTmpCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBody) GetData() *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyData {
	return s.Data
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBody) SetCode(v string) *GetDingTalkUserOrgByAliyunTmpCodeResponseBody {
	s.Code = &v
	return s
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBody) SetData(v *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyData) *GetDingTalkUserOrgByAliyunTmpCodeResponseBody {
	s.Data = v
	return s
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBody) SetHttpCode(v string) *GetDingTalkUserOrgByAliyunTmpCodeResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBody) SetMessage(v string) *GetDingTalkUserOrgByAliyunTmpCodeResponseBody {
	s.Message = &v
	return s
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBody) SetRequestId(v string) *GetDingTalkUserOrgByAliyunTmpCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBody) SetSuccess(v bool) *GetDingTalkUserOrgByAliyunTmpCodeResponseBody {
	s.Success = &v
	return s
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDingTalkUserOrgByAliyunTmpCodeResponseBodyData struct {
	AssociatedUnionId *string                                                        `json:"AssociatedUnionId,omitempty" xml:"AssociatedUnionId,omitempty"`
	Nick              *string                                                        `json:"Nick,omitempty" xml:"Nick,omitempty"`
	OrgDtoList        []*GetDingTalkUserOrgByAliyunTmpCodeResponseBodyDataOrgDtoList `json:"OrgDtoList,omitempty" xml:"OrgDtoList,omitempty" type:"Repeated"`
}

func (s GetDingTalkUserOrgByAliyunTmpCodeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDingTalkUserOrgByAliyunTmpCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyData) GetAssociatedUnionId() *string {
	return s.AssociatedUnionId
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyData) GetNick() *string {
	return s.Nick
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyData) GetOrgDtoList() []*GetDingTalkUserOrgByAliyunTmpCodeResponseBodyDataOrgDtoList {
	return s.OrgDtoList
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyData) SetAssociatedUnionId(v string) *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyData {
	s.AssociatedUnionId = &v
	return s
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyData) SetNick(v string) *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyData {
	s.Nick = &v
	return s
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyData) SetOrgDtoList(v []*GetDingTalkUserOrgByAliyunTmpCodeResponseBodyDataOrgDtoList) *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyData {
	s.OrgDtoList = v
	return s
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyData) Validate() error {
	if s.OrgDtoList != nil {
		for _, item := range s.OrgDtoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDingTalkUserOrgByAliyunTmpCodeResponseBodyDataOrgDtoList struct {
	CorpId  *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	OrgId   *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
}

func (s GetDingTalkUserOrgByAliyunTmpCodeResponseBodyDataOrgDtoList) String() string {
	return dara.Prettify(s)
}

func (s GetDingTalkUserOrgByAliyunTmpCodeResponseBodyDataOrgDtoList) GoString() string {
	return s.String()
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyDataOrgDtoList) GetCorpId() *string {
	return s.CorpId
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyDataOrgDtoList) GetOrgId() *string {
	return s.OrgId
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyDataOrgDtoList) GetOrgName() *string {
	return s.OrgName
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyDataOrgDtoList) SetCorpId(v string) *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyDataOrgDtoList {
	s.CorpId = &v
	return s
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyDataOrgDtoList) SetOrgId(v string) *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyDataOrgDtoList {
	s.OrgId = &v
	return s
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyDataOrgDtoList) SetOrgName(v string) *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyDataOrgDtoList {
	s.OrgName = &v
	return s
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponseBodyDataOrgDtoList) Validate() error {
	return dara.Validate(s)
}
