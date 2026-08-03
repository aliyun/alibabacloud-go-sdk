// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectImageTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMessage(v string) *SelectImageTaskResponseBody
	GetErrorMessage() *string
	SetFailed(v string) *SelectImageTaskResponseBody
	GetFailed() *string
	SetGenerationSource(v string) *SelectImageTaskResponseBody
	GetGenerationSource() *string
	SetGmtCreate(v string) *SelectImageTaskResponseBody
	GetGmtCreate() *string
	SetImageInfos(v []*SelectImageTaskResponseBodyImageInfos) *SelectImageTaskResponseBody
	GetImageInfos() []*SelectImageTaskResponseBodyImageInfos
	SetRequestId(v string) *SelectImageTaskResponseBody
	GetRequestId() *string
	SetScene(v string) *SelectImageTaskResponseBody
	GetScene() *string
	SetStatus(v string) *SelectImageTaskResponseBody
	GetStatus() *string
	SetSubtaskProcessing(v string) *SelectImageTaskResponseBody
	GetSubtaskProcessing() *string
	SetSuccess(v string) *SelectImageTaskResponseBody
	GetSuccess() *string
	SetTotal(v string) *SelectImageTaskResponseBody
	GetTotal() *string
}

type SelectImageTaskResponseBody struct {
	ErrorMessage      *string                                  `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Failed            *string                                  `json:"failed,omitempty" xml:"failed,omitempty"`
	GenerationSource  *string                                  `json:"generationSource,omitempty" xml:"generationSource,omitempty"`
	GmtCreate         *string                                  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	ImageInfos        []*SelectImageTaskResponseBodyImageInfos `json:"imageInfos,omitempty" xml:"imageInfos,omitempty" type:"Repeated"`
	RequestId         *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Scene             *string                                  `json:"scene,omitempty" xml:"scene,omitempty"`
	Status            *string                                  `json:"status,omitempty" xml:"status,omitempty"`
	SubtaskProcessing *string                                  `json:"subtaskProcessing,omitempty" xml:"subtaskProcessing,omitempty"`
	Success           *string                                  `json:"success,omitempty" xml:"success,omitempty"`
	Total             *string                                  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s SelectImageTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SelectImageTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SelectImageTaskResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SelectImageTaskResponseBody) GetFailed() *string {
	return s.Failed
}

func (s *SelectImageTaskResponseBody) GetGenerationSource() *string {
	return s.GenerationSource
}

func (s *SelectImageTaskResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *SelectImageTaskResponseBody) GetImageInfos() []*SelectImageTaskResponseBodyImageInfos {
	return s.ImageInfos
}

func (s *SelectImageTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SelectImageTaskResponseBody) GetScene() *string {
	return s.Scene
}

func (s *SelectImageTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *SelectImageTaskResponseBody) GetSubtaskProcessing() *string {
	return s.SubtaskProcessing
}

func (s *SelectImageTaskResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *SelectImageTaskResponseBody) GetTotal() *string {
	return s.Total
}

func (s *SelectImageTaskResponseBody) SetErrorMessage(v string) *SelectImageTaskResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SelectImageTaskResponseBody) SetFailed(v string) *SelectImageTaskResponseBody {
	s.Failed = &v
	return s
}

func (s *SelectImageTaskResponseBody) SetGenerationSource(v string) *SelectImageTaskResponseBody {
	s.GenerationSource = &v
	return s
}

func (s *SelectImageTaskResponseBody) SetGmtCreate(v string) *SelectImageTaskResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *SelectImageTaskResponseBody) SetImageInfos(v []*SelectImageTaskResponseBodyImageInfos) *SelectImageTaskResponseBody {
	s.ImageInfos = v
	return s
}

func (s *SelectImageTaskResponseBody) SetRequestId(v string) *SelectImageTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SelectImageTaskResponseBody) SetScene(v string) *SelectImageTaskResponseBody {
	s.Scene = &v
	return s
}

func (s *SelectImageTaskResponseBody) SetStatus(v string) *SelectImageTaskResponseBody {
	s.Status = &v
	return s
}

func (s *SelectImageTaskResponseBody) SetSubtaskProcessing(v string) *SelectImageTaskResponseBody {
	s.SubtaskProcessing = &v
	return s
}

func (s *SelectImageTaskResponseBody) SetSuccess(v string) *SelectImageTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SelectImageTaskResponseBody) SetTotal(v string) *SelectImageTaskResponseBody {
	s.Total = &v
	return s
}

func (s *SelectImageTaskResponseBody) Validate() error {
	if s.ImageInfos != nil {
		for _, item := range s.ImageInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SelectImageTaskResponseBodyImageInfos struct {
	CustomImageUrl *string `json:"customImageUrl,omitempty" xml:"customImageUrl,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	ImageH         *string `json:"imageH,omitempty" xml:"imageH,omitempty"`
	ImageW         *string `json:"imageW,omitempty" xml:"imageW,omitempty"`
}

func (s SelectImageTaskResponseBodyImageInfos) String() string {
	return dara.Prettify(s)
}

func (s SelectImageTaskResponseBodyImageInfos) GoString() string {
	return s.String()
}

func (s *SelectImageTaskResponseBodyImageInfos) GetCustomImageUrl() *string {
	return s.CustomImageUrl
}

func (s *SelectImageTaskResponseBodyImageInfos) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *SelectImageTaskResponseBodyImageInfos) GetImageH() *string {
	return s.ImageH
}

func (s *SelectImageTaskResponseBodyImageInfos) GetImageW() *string {
	return s.ImageW
}

func (s *SelectImageTaskResponseBodyImageInfos) SetCustomImageUrl(v string) *SelectImageTaskResponseBodyImageInfos {
	s.CustomImageUrl = &v
	return s
}

func (s *SelectImageTaskResponseBodyImageInfos) SetGmtCreate(v string) *SelectImageTaskResponseBodyImageInfos {
	s.GmtCreate = &v
	return s
}

func (s *SelectImageTaskResponseBodyImageInfos) SetImageH(v string) *SelectImageTaskResponseBodyImageInfos {
	s.ImageH = &v
	return s
}

func (s *SelectImageTaskResponseBodyImageInfos) SetImageW(v string) *SelectImageTaskResponseBodyImageInfos {
	s.ImageW = &v
	return s
}

func (s *SelectImageTaskResponseBodyImageInfos) Validate() error {
	return dara.Validate(s)
}
