// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeSkyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ChangeSkyResponseBodyData) *ChangeSkyResponseBody
	GetData() *ChangeSkyResponseBodyData
	SetRequestId(v string) *ChangeSkyResponseBody
	GetRequestId() *string
}

type ChangeSkyResponseBody struct {
	Data *ChangeSkyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// F9D60817-EC5A-4BAC-9092-4AD42220CFA8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeSkyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeSkyResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeSkyResponseBody) GetData() *ChangeSkyResponseBodyData {
	return s.Data
}

func (s *ChangeSkyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeSkyResponseBody) SetData(v *ChangeSkyResponseBodyData) *ChangeSkyResponseBody {
	s.Data = v
	return s
}

func (s *ChangeSkyResponseBody) SetRequestId(v string) *ChangeSkyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeSkyResponseBody) Validate() error {
	return dara.Validate(s)
}

type ChangeSkyResponseBodyData struct {
	// example:
	//
	// http://viapi-cn-shanghai-dha-segmenter.oss-cn-shanghai.aliyuncs.com/upload/result_skySegmentator/2020-7-24/invi_skySegmentator_015955791588111000000_5pn2QM.jpg?Expires=1595580958&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=Sq4po8h3WAj%2BBFrCgTP3ghlXn4****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ChangeSkyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ChangeSkyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeSkyResponseBodyData) GetImageURL() *string {
	return s.ImageURL
}

func (s *ChangeSkyResponseBodyData) SetImageURL(v string) *ChangeSkyResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *ChangeSkyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
