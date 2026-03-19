// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPptTemplateSelectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPptTemplateSelectorResponseBody
	GetCode() *string
	SetData(v *GetPptTemplateSelectorResponseBodyData) *GetPptTemplateSelectorResponseBody
	GetData() *GetPptTemplateSelectorResponseBodyData
	SetHttpStatusCode(v int32) *GetPptTemplateSelectorResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPptTemplateSelectorResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPptTemplateSelectorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPptTemplateSelectorResponseBody
	GetSuccess() *bool
}

type GetPptTemplateSelectorResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetPptTemplateSelectorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPptTemplateSelectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPptTemplateSelectorResponseBody) GoString() string {
	return s.String()
}

func (s *GetPptTemplateSelectorResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPptTemplateSelectorResponseBody) GetData() *GetPptTemplateSelectorResponseBodyData {
	return s.Data
}

func (s *GetPptTemplateSelectorResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPptTemplateSelectorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPptTemplateSelectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPptTemplateSelectorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPptTemplateSelectorResponseBody) SetCode(v string) *GetPptTemplateSelectorResponseBody {
	s.Code = &v
	return s
}

func (s *GetPptTemplateSelectorResponseBody) SetData(v *GetPptTemplateSelectorResponseBodyData) *GetPptTemplateSelectorResponseBody {
	s.Data = v
	return s
}

func (s *GetPptTemplateSelectorResponseBody) SetHttpStatusCode(v int32) *GetPptTemplateSelectorResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPptTemplateSelectorResponseBody) SetMessage(v string) *GetPptTemplateSelectorResponseBody {
	s.Message = &v
	return s
}

func (s *GetPptTemplateSelectorResponseBody) SetRequestId(v string) *GetPptTemplateSelectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPptTemplateSelectorResponseBody) SetSuccess(v bool) *GetPptTemplateSelectorResponseBody {
	s.Success = &v
	return s
}

func (s *GetPptTemplateSelectorResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPptTemplateSelectorResponseBodyData struct {
	Career    []*GetPptTemplateSelectorResponseBodyDataCareer    `json:"Career,omitempty" xml:"Career,omitempty" type:"Repeated"`
	Colour    []*GetPptTemplateSelectorResponseBodyDataColour    `json:"Colour,omitempty" xml:"Colour,omitempty" type:"Repeated"`
	SuitScene []*GetPptTemplateSelectorResponseBodyDataSuitScene `json:"SuitScene,omitempty" xml:"SuitScene,omitempty" type:"Repeated"`
	SuitStyle []*GetPptTemplateSelectorResponseBodyDataSuitStyle `json:"SuitStyle,omitempty" xml:"SuitStyle,omitempty" type:"Repeated"`
}

func (s GetPptTemplateSelectorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPptTemplateSelectorResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPptTemplateSelectorResponseBodyData) GetCareer() []*GetPptTemplateSelectorResponseBodyDataCareer {
	return s.Career
}

func (s *GetPptTemplateSelectorResponseBodyData) GetColour() []*GetPptTemplateSelectorResponseBodyDataColour {
	return s.Colour
}

func (s *GetPptTemplateSelectorResponseBodyData) GetSuitScene() []*GetPptTemplateSelectorResponseBodyDataSuitScene {
	return s.SuitScene
}

func (s *GetPptTemplateSelectorResponseBodyData) GetSuitStyle() []*GetPptTemplateSelectorResponseBodyDataSuitStyle {
	return s.SuitStyle
}

func (s *GetPptTemplateSelectorResponseBodyData) SetCareer(v []*GetPptTemplateSelectorResponseBodyDataCareer) *GetPptTemplateSelectorResponseBodyData {
	s.Career = v
	return s
}

func (s *GetPptTemplateSelectorResponseBodyData) SetColour(v []*GetPptTemplateSelectorResponseBodyDataColour) *GetPptTemplateSelectorResponseBodyData {
	s.Colour = v
	return s
}

func (s *GetPptTemplateSelectorResponseBodyData) SetSuitScene(v []*GetPptTemplateSelectorResponseBodyDataSuitScene) *GetPptTemplateSelectorResponseBodyData {
	s.SuitScene = v
	return s
}

func (s *GetPptTemplateSelectorResponseBodyData) SetSuitStyle(v []*GetPptTemplateSelectorResponseBodyDataSuitStyle) *GetPptTemplateSelectorResponseBodyData {
	s.SuitStyle = v
	return s
}

func (s *GetPptTemplateSelectorResponseBodyData) Validate() error {
	if s.Career != nil {
		for _, item := range s.Career {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Colour != nil {
		for _, item := range s.Colour {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SuitScene != nil {
		for _, item := range s.SuitScene {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SuitStyle != nil {
		for _, item := range s.SuitStyle {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPptTemplateSelectorResponseBodyDataCareer struct {
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0
	IsHot *int64 `json:"IsHot,omitempty" xml:"IsHot,omitempty"`
	// example:
	//
	// 教育培训
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPptTemplateSelectorResponseBodyDataCareer) String() string {
	return dara.Prettify(s)
}

func (s GetPptTemplateSelectorResponseBodyDataCareer) GoString() string {
	return s.String()
}

func (s *GetPptTemplateSelectorResponseBodyDataCareer) GetId() *int64 {
	return s.Id
}

func (s *GetPptTemplateSelectorResponseBodyDataCareer) GetIsHot() *int64 {
	return s.IsHot
}

func (s *GetPptTemplateSelectorResponseBodyDataCareer) GetName() *string {
	return s.Name
}

func (s *GetPptTemplateSelectorResponseBodyDataCareer) SetId(v int64) *GetPptTemplateSelectorResponseBodyDataCareer {
	s.Id = &v
	return s
}

func (s *GetPptTemplateSelectorResponseBodyDataCareer) SetIsHot(v int64) *GetPptTemplateSelectorResponseBodyDataCareer {
	s.IsHot = &v
	return s
}

func (s *GetPptTemplateSelectorResponseBodyDataCareer) SetName(v string) *GetPptTemplateSelectorResponseBodyDataCareer {
	s.Name = &v
	return s
}

func (s *GetPptTemplateSelectorResponseBodyDataCareer) Validate() error {
	return dara.Validate(s)
}

type GetPptTemplateSelectorResponseBodyDataColour struct {
	// example:
	//
	// #FCC462
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 橙色
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPptTemplateSelectorResponseBodyDataColour) String() string {
	return dara.Prettify(s)
}

func (s GetPptTemplateSelectorResponseBodyDataColour) GoString() string {
	return s.String()
}

func (s *GetPptTemplateSelectorResponseBodyDataColour) GetCode() *string {
	return s.Code
}

func (s *GetPptTemplateSelectorResponseBodyDataColour) GetId() *int64 {
	return s.Id
}

func (s *GetPptTemplateSelectorResponseBodyDataColour) GetName() *string {
	return s.Name
}

func (s *GetPptTemplateSelectorResponseBodyDataColour) SetCode(v string) *GetPptTemplateSelectorResponseBodyDataColour {
	s.Code = &v
	return s
}

func (s *GetPptTemplateSelectorResponseBodyDataColour) SetId(v int64) *GetPptTemplateSelectorResponseBodyDataColour {
	s.Id = &v
	return s
}

func (s *GetPptTemplateSelectorResponseBodyDataColour) SetName(v string) *GetPptTemplateSelectorResponseBodyDataColour {
	s.Name = &v
	return s
}

func (s *GetPptTemplateSelectorResponseBodyDataColour) Validate() error {
	return dara.Validate(s)
}

type GetPptTemplateSelectorResponseBodyDataSuitScene struct {
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 教育培训
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetPptTemplateSelectorResponseBodyDataSuitScene) String() string {
	return dara.Prettify(s)
}

func (s GetPptTemplateSelectorResponseBodyDataSuitScene) GoString() string {
	return s.String()
}

func (s *GetPptTemplateSelectorResponseBodyDataSuitScene) GetId() *int64 {
	return s.Id
}

func (s *GetPptTemplateSelectorResponseBodyDataSuitScene) GetTitle() *string {
	return s.Title
}

func (s *GetPptTemplateSelectorResponseBodyDataSuitScene) SetId(v int64) *GetPptTemplateSelectorResponseBodyDataSuitScene {
	s.Id = &v
	return s
}

func (s *GetPptTemplateSelectorResponseBodyDataSuitScene) SetTitle(v string) *GetPptTemplateSelectorResponseBodyDataSuitScene {
	s.Title = &v
	return s
}

func (s *GetPptTemplateSelectorResponseBodyDataSuitScene) Validate() error {
	return dara.Validate(s)
}

type GetPptTemplateSelectorResponseBodyDataSuitStyle struct {
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 扁平简约
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetPptTemplateSelectorResponseBodyDataSuitStyle) String() string {
	return dara.Prettify(s)
}

func (s GetPptTemplateSelectorResponseBodyDataSuitStyle) GoString() string {
	return s.String()
}

func (s *GetPptTemplateSelectorResponseBodyDataSuitStyle) GetId() *int64 {
	return s.Id
}

func (s *GetPptTemplateSelectorResponseBodyDataSuitStyle) GetTitle() *string {
	return s.Title
}

func (s *GetPptTemplateSelectorResponseBodyDataSuitStyle) SetId(v int64) *GetPptTemplateSelectorResponseBodyDataSuitStyle {
	s.Id = &v
	return s
}

func (s *GetPptTemplateSelectorResponseBodyDataSuitStyle) SetTitle(v string) *GetPptTemplateSelectorResponseBodyDataSuitStyle {
	s.Title = &v
	return s
}

func (s *GetPptTemplateSelectorResponseBodyDataSuitStyle) Validate() error {
	return dara.Validate(s)
}
