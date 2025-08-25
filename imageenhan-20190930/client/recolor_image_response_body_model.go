// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecolorImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecolorImageResponseBodyData) *RecolorImageResponseBody
	GetData() *RecolorImageResponseBodyData
	SetRequestId(v string) *RecolorImageResponseBody
	GetRequestId() *string
}

type RecolorImageResponseBody struct {
	Data *RecolorImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 3A9BFC5E-3F7C-4D9A-9445-908C6D14AB5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecolorImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecolorImageResponseBody) GoString() string {
	return s.String()
}

func (s *RecolorImageResponseBody) GetData() *RecolorImageResponseBodyData {
	return s.Data
}

func (s *RecolorImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecolorImageResponseBody) SetData(v *RecolorImageResponseBodyData) *RecolorImageResponseBody {
	s.Data = v
	return s
}

func (s *RecolorImageResponseBody) SetRequestId(v string) *RecolorImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecolorImageResponseBody) Validate() error {
	return dara.Validate(s)
}

type RecolorImageResponseBodyData struct {
	// 1
	ImageList []*string `json:"ImageList,omitempty" xml:"ImageList,omitempty" type:"Repeated"`
}

func (s RecolorImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecolorImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecolorImageResponseBodyData) GetImageList() []*string {
	return s.ImageList
}

func (s *RecolorImageResponseBodyData) SetImageList(v []*string) *RecolorImageResponseBodyData {
	s.ImageList = v
	return s
}

func (s *RecolorImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
