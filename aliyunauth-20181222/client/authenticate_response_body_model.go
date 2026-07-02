// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthenticateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AuthenticateResponseBody
	GetCode() *int32
	SetData(v *AuthenticateResponseBodyData) *AuthenticateResponseBody
	GetData() *AuthenticateResponseBodyData
	SetMessage(v string) *AuthenticateResponseBody
	GetMessage() *string
	SetRequestId(v string) *AuthenticateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AuthenticateResponseBody
	GetSuccess() *bool
}

type AuthenticateResponseBody struct {
	Code      *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AuthenticateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AuthenticateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthenticateResponseBody) GoString() string {
	return s.String()
}

func (s *AuthenticateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AuthenticateResponseBody) GetData() *AuthenticateResponseBodyData {
	return s.Data
}

func (s *AuthenticateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AuthenticateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthenticateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AuthenticateResponseBody) SetCode(v int32) *AuthenticateResponseBody {
	s.Code = &v
	return s
}

func (s *AuthenticateResponseBody) SetData(v *AuthenticateResponseBodyData) *AuthenticateResponseBody {
	s.Data = v
	return s
}

func (s *AuthenticateResponseBody) SetMessage(v string) *AuthenticateResponseBody {
	s.Message = &v
	return s
}

func (s *AuthenticateResponseBody) SetRequestId(v string) *AuthenticateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthenticateResponseBody) SetSuccess(v bool) *AuthenticateResponseBody {
	s.Success = &v
	return s
}

func (s *AuthenticateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AuthenticateResponseBodyData struct {
	Authorized    *bool                                      `json:"Authorized,omitempty" xml:"Authorized,omitempty"`
	InstanceId    *string                                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OperateCode   *string                                    `json:"OperateCode,omitempty" xml:"OperateCode,omitempty"`
	ProductCode   *string                                    `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	RecordVid     *string                                    `json:"RecordVid,omitempty" xml:"RecordVid,omitempty"`
	UserInputList *AuthenticateResponseBodyDataUserInputList `json:"UserInputList,omitempty" xml:"UserInputList,omitempty" type:"Struct"`
}

func (s AuthenticateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AuthenticateResponseBodyData) GoString() string {
	return s.String()
}

func (s *AuthenticateResponseBodyData) GetAuthorized() *bool {
	return s.Authorized
}

func (s *AuthenticateResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AuthenticateResponseBodyData) GetOperateCode() *string {
	return s.OperateCode
}

func (s *AuthenticateResponseBodyData) GetProductCode() *string {
	return s.ProductCode
}

func (s *AuthenticateResponseBodyData) GetRecordVid() *string {
	return s.RecordVid
}

func (s *AuthenticateResponseBodyData) GetUserInputList() *AuthenticateResponseBodyDataUserInputList {
	return s.UserInputList
}

func (s *AuthenticateResponseBodyData) SetAuthorized(v bool) *AuthenticateResponseBodyData {
	s.Authorized = &v
	return s
}

func (s *AuthenticateResponseBodyData) SetInstanceId(v string) *AuthenticateResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *AuthenticateResponseBodyData) SetOperateCode(v string) *AuthenticateResponseBodyData {
	s.OperateCode = &v
	return s
}

func (s *AuthenticateResponseBodyData) SetProductCode(v string) *AuthenticateResponseBodyData {
	s.ProductCode = &v
	return s
}

func (s *AuthenticateResponseBodyData) SetRecordVid(v string) *AuthenticateResponseBodyData {
	s.RecordVid = &v
	return s
}

func (s *AuthenticateResponseBodyData) SetUserInputList(v *AuthenticateResponseBodyDataUserInputList) *AuthenticateResponseBodyData {
	s.UserInputList = v
	return s
}

func (s *AuthenticateResponseBodyData) Validate() error {
	if s.UserInputList != nil {
		if err := s.UserInputList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AuthenticateResponseBodyDataUserInputList struct {
	QueryAuthRSDTO []*AuthenticateResponseBodyDataUserInputListQueryAuthRSDTO `json:"QueryAuthRSDTO,omitempty" xml:"QueryAuthRSDTO,omitempty" type:"Repeated"`
}

func (s AuthenticateResponseBodyDataUserInputList) String() string {
	return dara.Prettify(s)
}

func (s AuthenticateResponseBodyDataUserInputList) GoString() string {
	return s.String()
}

func (s *AuthenticateResponseBodyDataUserInputList) GetQueryAuthRSDTO() []*AuthenticateResponseBodyDataUserInputListQueryAuthRSDTO {
	return s.QueryAuthRSDTO
}

func (s *AuthenticateResponseBodyDataUserInputList) SetQueryAuthRSDTO(v []*AuthenticateResponseBodyDataUserInputListQueryAuthRSDTO) *AuthenticateResponseBodyDataUserInputList {
	s.QueryAuthRSDTO = v
	return s
}

func (s *AuthenticateResponseBodyDataUserInputList) Validate() error {
	if s.QueryAuthRSDTO != nil {
		for _, item := range s.QueryAuthRSDTO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AuthenticateResponseBodyDataUserInputListQueryAuthRSDTO struct {
	Content       *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ExpandContent *string `json:"ExpandContent,omitempty" xml:"ExpandContent,omitempty"`
	FieldName     *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	FieldVid      *string `json:"FieldVid,omitempty" xml:"FieldVid,omitempty"`
}

func (s AuthenticateResponseBodyDataUserInputListQueryAuthRSDTO) String() string {
	return dara.Prettify(s)
}

func (s AuthenticateResponseBodyDataUserInputListQueryAuthRSDTO) GoString() string {
	return s.String()
}

func (s *AuthenticateResponseBodyDataUserInputListQueryAuthRSDTO) GetContent() *string {
	return s.Content
}

func (s *AuthenticateResponseBodyDataUserInputListQueryAuthRSDTO) GetExpandContent() *string {
	return s.ExpandContent
}

func (s *AuthenticateResponseBodyDataUserInputListQueryAuthRSDTO) GetFieldName() *string {
	return s.FieldName
}

func (s *AuthenticateResponseBodyDataUserInputListQueryAuthRSDTO) GetFieldVid() *string {
	return s.FieldVid
}

func (s *AuthenticateResponseBodyDataUserInputListQueryAuthRSDTO) SetContent(v string) *AuthenticateResponseBodyDataUserInputListQueryAuthRSDTO {
	s.Content = &v
	return s
}

func (s *AuthenticateResponseBodyDataUserInputListQueryAuthRSDTO) SetExpandContent(v string) *AuthenticateResponseBodyDataUserInputListQueryAuthRSDTO {
	s.ExpandContent = &v
	return s
}

func (s *AuthenticateResponseBodyDataUserInputListQueryAuthRSDTO) SetFieldName(v string) *AuthenticateResponseBodyDataUserInputListQueryAuthRSDTO {
	s.FieldName = &v
	return s
}

func (s *AuthenticateResponseBodyDataUserInputListQueryAuthRSDTO) SetFieldVid(v string) *AuthenticateResponseBodyDataUserInputListQueryAuthRSDTO {
	s.FieldVid = &v
	return s
}

func (s *AuthenticateResponseBodyDataUserInputListQueryAuthRSDTO) Validate() error {
	return dara.Validate(s)
}
