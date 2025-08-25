// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbstractFilmVideoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AbstractFilmVideoResponseBodyData) *AbstractFilmVideoResponseBody
	GetData() *AbstractFilmVideoResponseBodyData
	SetMessage(v string) *AbstractFilmVideoResponseBody
	GetMessage() *string
	SetRequestId(v string) *AbstractFilmVideoResponseBody
	GetRequestId() *string
}

type AbstractFilmVideoResponseBody struct {
	Data    *AbstractFilmVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9A5B872B-1BF0-4D84-90DA-A2EE1F072B82
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AbstractFilmVideoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AbstractFilmVideoResponseBody) GoString() string {
	return s.String()
}

func (s *AbstractFilmVideoResponseBody) GetData() *AbstractFilmVideoResponseBodyData {
	return s.Data
}

func (s *AbstractFilmVideoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AbstractFilmVideoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AbstractFilmVideoResponseBody) SetData(v *AbstractFilmVideoResponseBodyData) *AbstractFilmVideoResponseBody {
	s.Data = v
	return s
}

func (s *AbstractFilmVideoResponseBody) SetMessage(v string) *AbstractFilmVideoResponseBody {
	s.Message = &v
	return s
}

func (s *AbstractFilmVideoResponseBody) SetRequestId(v string) *AbstractFilmVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbstractFilmVideoResponseBody) Validate() error {
	return dara.Validate(s)
}

type AbstractFilmVideoResponseBodyData struct {
	// example:
	//
	// http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/film-summary/EA61D012-5F89-4102-931E-419158BE2ADA_gb27k00.mp4?Expires=1584707613&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=weTexlBR1wmQlAhuU2JXaE7AyJ****
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s AbstractFilmVideoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AbstractFilmVideoResponseBodyData) GoString() string {
	return s.String()
}

func (s *AbstractFilmVideoResponseBodyData) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *AbstractFilmVideoResponseBodyData) SetVideoUrl(v string) *AbstractFilmVideoResponseBodyData {
	s.VideoUrl = &v
	return s
}

func (s *AbstractFilmVideoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
