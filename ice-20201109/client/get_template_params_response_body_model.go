// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateParamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParamList(v []*GetTemplateParamsResponseBodyParamList) *GetTemplateParamsResponseBody
	GetParamList() []*GetTemplateParamsResponseBodyParamList
	SetRequestId(v string) *GetTemplateParamsResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *GetTemplateParamsResponseBody
	GetTemplateId() *string
}

type GetTemplateParamsResponseBody struct {
	// The queried parameters.
	ParamList []*GetTemplateParamsResponseBodyParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ****2876-6263-4B75-8F2C-CD0F7FCF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The template ID.
	//
	// example:
	//
	// ******419c8741c1b4325f035b******
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetTemplateParamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateParamsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateParamsResponseBody) GetParamList() []*GetTemplateParamsResponseBodyParamList {
	return s.ParamList
}

func (s *GetTemplateParamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemplateParamsResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTemplateParamsResponseBody) SetParamList(v []*GetTemplateParamsResponseBodyParamList) *GetTemplateParamsResponseBody {
	s.ParamList = v
	return s
}

func (s *GetTemplateParamsResponseBody) SetRequestId(v string) *GetTemplateParamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateParamsResponseBody) SetTemplateId(v string) *GetTemplateParamsResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateParamsResponseBody) Validate() error {
	if s.ParamList != nil {
		for _, item := range s.ParamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTemplateParamsResponseBodyParamList struct {
	// The original subtitle content.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The thumbnail URL of the original material.
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// video1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The URL of the original material.
	MediaUrl *string `json:"MediaUrl,omitempty" xml:"MediaUrl,omitempty"`
	// The material type.
	//
	// Valid values:
	//
	// 	- Video
	//
	// 	- Text
	//
	// 	- Image
	//
	// example:
	//
	// Image
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTemplateParamsResponseBodyParamList) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateParamsResponseBodyParamList) GoString() string {
	return s.String()
}

func (s *GetTemplateParamsResponseBodyParamList) GetContent() *string {
	return s.Content
}

func (s *GetTemplateParamsResponseBodyParamList) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *GetTemplateParamsResponseBodyParamList) GetKey() *string {
	return s.Key
}

func (s *GetTemplateParamsResponseBodyParamList) GetMediaUrl() *string {
	return s.MediaUrl
}

func (s *GetTemplateParamsResponseBodyParamList) GetType() *string {
	return s.Type
}

func (s *GetTemplateParamsResponseBodyParamList) SetContent(v string) *GetTemplateParamsResponseBodyParamList {
	s.Content = &v
	return s
}

func (s *GetTemplateParamsResponseBodyParamList) SetCoverUrl(v string) *GetTemplateParamsResponseBodyParamList {
	s.CoverUrl = &v
	return s
}

func (s *GetTemplateParamsResponseBodyParamList) SetKey(v string) *GetTemplateParamsResponseBodyParamList {
	s.Key = &v
	return s
}

func (s *GetTemplateParamsResponseBodyParamList) SetMediaUrl(v string) *GetTemplateParamsResponseBodyParamList {
	s.MediaUrl = &v
	return s
}

func (s *GetTemplateParamsResponseBodyParamList) SetType(v string) *GetTemplateParamsResponseBodyParamList {
	s.Type = &v
	return s
}

func (s *GetTemplateParamsResponseBodyParamList) Validate() error {
	return dara.Validate(s)
}
