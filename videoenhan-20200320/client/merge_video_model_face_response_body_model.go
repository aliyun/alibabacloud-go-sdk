// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMergeVideoModelFaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *MergeVideoModelFaceResponseBodyData) *MergeVideoModelFaceResponseBody
	GetData() *MergeVideoModelFaceResponseBodyData
	SetMessage(v string) *MergeVideoModelFaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *MergeVideoModelFaceResponseBody
	GetRequestId() *string
}

type MergeVideoModelFaceResponseBody struct {
	Data    *MergeVideoModelFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 334F180F-3B50-51CB-B4CB-9A86A542D3BC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MergeVideoModelFaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MergeVideoModelFaceResponseBody) GoString() string {
	return s.String()
}

func (s *MergeVideoModelFaceResponseBody) GetData() *MergeVideoModelFaceResponseBodyData {
	return s.Data
}

func (s *MergeVideoModelFaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MergeVideoModelFaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MergeVideoModelFaceResponseBody) SetData(v *MergeVideoModelFaceResponseBodyData) *MergeVideoModelFaceResponseBody {
	s.Data = v
	return s
}

func (s *MergeVideoModelFaceResponseBody) SetMessage(v string) *MergeVideoModelFaceResponseBody {
	s.Message = &v
	return s
}

func (s *MergeVideoModelFaceResponseBody) SetRequestId(v string) *MergeVideoModelFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *MergeVideoModelFaceResponseBody) Validate() error {
	return dara.Validate(s)
}

type MergeVideoModelFaceResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-gd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/video-face-fusion/334F180F-3B50-51CB-B4CB-9A86A542D3BC-5716-20210906-074905.mp4?Expires=1630916420&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=OEtNqVYxXRwkyO4BrsYVJ8q5bx****
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s MergeVideoModelFaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MergeVideoModelFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *MergeVideoModelFaceResponseBodyData) GetVideoURL() *string {
	return s.VideoURL
}

func (s *MergeVideoModelFaceResponseBodyData) SetVideoURL(v string) *MergeVideoModelFaceResponseBodyData {
	s.VideoURL = &v
	return s
}

func (s *MergeVideoModelFaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
