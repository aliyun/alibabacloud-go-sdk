// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveImageSubtitlesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RemoveImageSubtitlesResponseBodyData) *RemoveImageSubtitlesResponseBody
	GetData() *RemoveImageSubtitlesResponseBodyData
	SetRequestId(v string) *RemoveImageSubtitlesResponseBody
	GetRequestId() *string
}

type RemoveImageSubtitlesResponseBody struct {
	Data *RemoveImageSubtitlesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 939B2103-EE28-4F2D-9773-9A37AD00E5B7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveImageSubtitlesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveImageSubtitlesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveImageSubtitlesResponseBody) GetData() *RemoveImageSubtitlesResponseBodyData {
	return s.Data
}

func (s *RemoveImageSubtitlesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveImageSubtitlesResponseBody) SetData(v *RemoveImageSubtitlesResponseBodyData) *RemoveImageSubtitlesResponseBody {
	s.Data = v
	return s
}

func (s *RemoveImageSubtitlesResponseBody) SetRequestId(v string) *RemoveImageSubtitlesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveImageSubtitlesResponseBody) Validate() error {
	return dara.Validate(s)
}

type RemoveImageSubtitlesResponseBodyData struct {
	// example:
	//
	// http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/image-desubtitle/2020-03-23-08/02%3A50-e8af2ea3-bddc-4ec8-b21c-493ee687465e.jpg?Expires=1584952370&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=qVnfWZJ2QtI9NRWQ410FsEFioq****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RemoveImageSubtitlesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RemoveImageSubtitlesResponseBodyData) GoString() string {
	return s.String()
}

func (s *RemoveImageSubtitlesResponseBodyData) GetImageURL() *string {
	return s.ImageURL
}

func (s *RemoveImageSubtitlesResponseBodyData) SetImageURL(v string) *RemoveImageSubtitlesResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *RemoveImageSubtitlesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
