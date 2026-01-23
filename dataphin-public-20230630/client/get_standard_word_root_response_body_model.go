// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardWordRootResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetStandardWordRootResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetStandardWordRootResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetStandardWordRootResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetStandardWordRootResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetStandardWordRootResponseBody
	GetSuccess() *bool
	SetWordRootInfo(v *GetStandardWordRootResponseBodyWordRootInfo) *GetStandardWordRootResponseBody
	GetWordRootInfo() *GetStandardWordRootResponseBodyWordRootInfo
}

type GetStandardWordRootResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	WordRootInfo *GetStandardWordRootResponseBodyWordRootInfo `json:"WordRootInfo,omitempty" xml:"WordRootInfo,omitempty" type:"Struct"`
}

func (s GetStandardWordRootResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStandardWordRootResponseBody) GoString() string {
	return s.String()
}

func (s *GetStandardWordRootResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetStandardWordRootResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetStandardWordRootResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetStandardWordRootResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStandardWordRootResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetStandardWordRootResponseBody) GetWordRootInfo() *GetStandardWordRootResponseBodyWordRootInfo {
	return s.WordRootInfo
}

func (s *GetStandardWordRootResponseBody) SetCode(v string) *GetStandardWordRootResponseBody {
	s.Code = &v
	return s
}

func (s *GetStandardWordRootResponseBody) SetHttpStatusCode(v int32) *GetStandardWordRootResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetStandardWordRootResponseBody) SetMessage(v string) *GetStandardWordRootResponseBody {
	s.Message = &v
	return s
}

func (s *GetStandardWordRootResponseBody) SetRequestId(v string) *GetStandardWordRootResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStandardWordRootResponseBody) SetSuccess(v bool) *GetStandardWordRootResponseBody {
	s.Success = &v
	return s
}

func (s *GetStandardWordRootResponseBody) SetWordRootInfo(v *GetStandardWordRootResponseBodyWordRootInfo) *GetStandardWordRootResponseBody {
	s.WordRootInfo = v
	return s
}

func (s *GetStandardWordRootResponseBody) Validate() error {
	if s.WordRootInfo != nil {
		if err := s.WordRootInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStandardWordRootResponseBodyWordRootInfo struct {
	// example:
	//
	// avg
	Abbreviation *string `json:"Abbreviation,omitempty" xml:"Abbreviation,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 30012021
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// test
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// average
	FullName *string `json:"FullName,omitempty" xml:"FullName,omitempty"`
	// example:
	//
	// 30012021
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// example:
	//
	// test
	LastModifierName *string `json:"LastModifierName,omitempty" xml:"LastModifierName,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 平均值
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetStandardWordRootResponseBodyWordRootInfo) String() string {
	return dara.Prettify(s)
}

func (s GetStandardWordRootResponseBodyWordRootInfo) GoString() string {
	return s.String()
}

func (s *GetStandardWordRootResponseBodyWordRootInfo) GetAbbreviation() *string {
	return s.Abbreviation
}

func (s *GetStandardWordRootResponseBodyWordRootInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetStandardWordRootResponseBodyWordRootInfo) GetCreator() *string {
	return s.Creator
}

func (s *GetStandardWordRootResponseBodyWordRootInfo) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetStandardWordRootResponseBodyWordRootInfo) GetDescription() *string {
	return s.Description
}

func (s *GetStandardWordRootResponseBodyWordRootInfo) GetFullName() *string {
	return s.FullName
}

func (s *GetStandardWordRootResponseBodyWordRootInfo) GetLastModifier() *string {
	return s.LastModifier
}

func (s *GetStandardWordRootResponseBodyWordRootInfo) GetLastModifierName() *string {
	return s.LastModifierName
}

func (s *GetStandardWordRootResponseBodyWordRootInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetStandardWordRootResponseBodyWordRootInfo) GetName() *string {
	return s.Name
}

func (s *GetStandardWordRootResponseBodyWordRootInfo) SetAbbreviation(v string) *GetStandardWordRootResponseBodyWordRootInfo {
	s.Abbreviation = &v
	return s
}

func (s *GetStandardWordRootResponseBodyWordRootInfo) SetCreateTime(v string) *GetStandardWordRootResponseBodyWordRootInfo {
	s.CreateTime = &v
	return s
}

func (s *GetStandardWordRootResponseBodyWordRootInfo) SetCreator(v string) *GetStandardWordRootResponseBodyWordRootInfo {
	s.Creator = &v
	return s
}

func (s *GetStandardWordRootResponseBodyWordRootInfo) SetCreatorName(v string) *GetStandardWordRootResponseBodyWordRootInfo {
	s.CreatorName = &v
	return s
}

func (s *GetStandardWordRootResponseBodyWordRootInfo) SetDescription(v string) *GetStandardWordRootResponseBodyWordRootInfo {
	s.Description = &v
	return s
}

func (s *GetStandardWordRootResponseBodyWordRootInfo) SetFullName(v string) *GetStandardWordRootResponseBodyWordRootInfo {
	s.FullName = &v
	return s
}

func (s *GetStandardWordRootResponseBodyWordRootInfo) SetLastModifier(v string) *GetStandardWordRootResponseBodyWordRootInfo {
	s.LastModifier = &v
	return s
}

func (s *GetStandardWordRootResponseBodyWordRootInfo) SetLastModifierName(v string) *GetStandardWordRootResponseBodyWordRootInfo {
	s.LastModifierName = &v
	return s
}

func (s *GetStandardWordRootResponseBodyWordRootInfo) SetModifyTime(v string) *GetStandardWordRootResponseBodyWordRootInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetStandardWordRootResponseBodyWordRootInfo) SetName(v string) *GetStandardWordRootResponseBodyWordRootInfo {
	s.Name = &v
	return s
}

func (s *GetStandardWordRootResponseBodyWordRootInfo) Validate() error {
	return dara.Validate(s)
}
