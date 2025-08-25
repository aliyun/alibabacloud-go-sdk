// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMergeVideoFaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *MergeVideoFaceResponseBodyData) *MergeVideoFaceResponseBody
	GetData() *MergeVideoFaceResponseBodyData
	SetMessage(v string) *MergeVideoFaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *MergeVideoFaceResponseBody
	GetRequestId() *string
}

type MergeVideoFaceResponseBody struct {
	Data    *MergeVideoFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DEF90E76-B62D-45EF-8835-CA3C83842B18
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MergeVideoFaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MergeVideoFaceResponseBody) GoString() string {
	return s.String()
}

func (s *MergeVideoFaceResponseBody) GetData() *MergeVideoFaceResponseBodyData {
	return s.Data
}

func (s *MergeVideoFaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MergeVideoFaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MergeVideoFaceResponseBody) SetData(v *MergeVideoFaceResponseBodyData) *MergeVideoFaceResponseBody {
	s.Data = v
	return s
}

func (s *MergeVideoFaceResponseBody) SetMessage(v string) *MergeVideoFaceResponseBody {
	s.Message = &v
	return s
}

func (s *MergeVideoFaceResponseBody) SetRequestId(v string) *MergeVideoFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *MergeVideoFaceResponseBody) Validate() error {
	return dara.Validate(s)
}

type MergeVideoFaceResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-gd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/video-face-fusion/A657011C-82B4-4705-A5DB-69B18B7CE89D.mp4?Expires=1606378308&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=Hl3cq5XedTGCscOSr0OGVxAS2o****
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s MergeVideoFaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MergeVideoFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *MergeVideoFaceResponseBodyData) GetVideoURL() *string {
	return s.VideoURL
}

func (s *MergeVideoFaceResponseBodyData) SetVideoURL(v string) *MergeVideoFaceResponseBodyData {
	s.VideoURL = &v
	return s
}

func (s *MergeVideoFaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
