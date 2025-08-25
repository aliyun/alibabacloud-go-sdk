// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecolorHDImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecolorHDImageResponseBodyData) *RecolorHDImageResponseBody
	GetData() *RecolorHDImageResponseBodyData
	SetMessage(v string) *RecolorHDImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *RecolorHDImageResponseBody
	GetRequestId() *string
}

type RecolorHDImageResponseBody struct {
	Data    *RecolorHDImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9364CEB1-1C39-489F-B6DB-6E65577429F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecolorHDImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecolorHDImageResponseBody) GoString() string {
	return s.String()
}

func (s *RecolorHDImageResponseBody) GetData() *RecolorHDImageResponseBodyData {
	return s.Data
}

func (s *RecolorHDImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RecolorHDImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecolorHDImageResponseBody) SetData(v *RecolorHDImageResponseBodyData) *RecolorHDImageResponseBody {
	s.Data = v
	return s
}

func (s *RecolorHDImageResponseBody) SetMessage(v string) *RecolorHDImageResponseBody {
	s.Message = &v
	return s
}

func (s *RecolorHDImageResponseBody) SetRequestId(v string) *RecolorHDImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecolorHDImageResponseBody) Validate() error {
	return dara.Validate(s)
}

type RecolorHDImageResponseBodyData struct {
	// 1
	ImageList []*string `json:"ImageList,omitempty" xml:"ImageList,omitempty" type:"Repeated"`
}

func (s RecolorHDImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecolorHDImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecolorHDImageResponseBodyData) GetImageList() []*string {
	return s.ImageList
}

func (s *RecolorHDImageResponseBodyData) SetImageList(v []*string) *RecolorHDImageResponseBodyData {
	s.ImageList = v
	return s
}

func (s *RecolorHDImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
