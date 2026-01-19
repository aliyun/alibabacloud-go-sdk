// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCorpNumberGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCorpNumberGroupResponseBody
	GetCode() *string
	SetData(v *CreateCorpNumberGroupResponseBodyData) *CreateCorpNumberGroupResponseBody
	GetData() *CreateCorpNumberGroupResponseBodyData
	SetHttpStatusCode(v int32) *CreateCorpNumberGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateCorpNumberGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCorpNumberGroupResponseBody
	GetRequestId() *string
}

type CreateCorpNumberGroupResponseBody struct {
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CreateCorpNumberGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCorpNumberGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCorpNumberGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCorpNumberGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCorpNumberGroupResponseBody) GetData() *CreateCorpNumberGroupResponseBodyData {
	return s.Data
}

func (s *CreateCorpNumberGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateCorpNumberGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCorpNumberGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCorpNumberGroupResponseBody) SetCode(v string) *CreateCorpNumberGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCorpNumberGroupResponseBody) SetData(v *CreateCorpNumberGroupResponseBodyData) *CreateCorpNumberGroupResponseBody {
	s.Data = v
	return s
}

func (s *CreateCorpNumberGroupResponseBody) SetHttpStatusCode(v int32) *CreateCorpNumberGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateCorpNumberGroupResponseBody) SetMessage(v string) *CreateCorpNumberGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCorpNumberGroupResponseBody) SetRequestId(v string) *CreateCorpNumberGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCorpNumberGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCorpNumberGroupResponseBodyData struct {
	AliyunUid       *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	NumberCount     *string `json:"NumberCount,omitempty" xml:"NumberCount,omitempty"`
	NumberGroupId   *string `json:"NumberGroupId,omitempty" xml:"NumberGroupId,omitempty"`
	NumberGroupName *string `json:"NumberGroupName,omitempty" xml:"NumberGroupName,omitempty"`
}

func (s CreateCorpNumberGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateCorpNumberGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCorpNumberGroupResponseBodyData) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *CreateCorpNumberGroupResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateCorpNumberGroupResponseBodyData) GetNumberCount() *string {
	return s.NumberCount
}

func (s *CreateCorpNumberGroupResponseBodyData) GetNumberGroupId() *string {
	return s.NumberGroupId
}

func (s *CreateCorpNumberGroupResponseBodyData) GetNumberGroupName() *string {
	return s.NumberGroupName
}

func (s *CreateCorpNumberGroupResponseBodyData) SetAliyunUid(v string) *CreateCorpNumberGroupResponseBodyData {
	s.AliyunUid = &v
	return s
}

func (s *CreateCorpNumberGroupResponseBodyData) SetDescription(v string) *CreateCorpNumberGroupResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateCorpNumberGroupResponseBodyData) SetNumberCount(v string) *CreateCorpNumberGroupResponseBodyData {
	s.NumberCount = &v
	return s
}

func (s *CreateCorpNumberGroupResponseBodyData) SetNumberGroupId(v string) *CreateCorpNumberGroupResponseBodyData {
	s.NumberGroupId = &v
	return s
}

func (s *CreateCorpNumberGroupResponseBodyData) SetNumberGroupName(v string) *CreateCorpNumberGroupResponseBodyData {
	s.NumberGroupName = &v
	return s
}

func (s *CreateCorpNumberGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
