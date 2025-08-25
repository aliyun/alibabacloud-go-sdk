// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImitatePhotoStyleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ImitatePhotoStyleResponseBodyData) *ImitatePhotoStyleResponseBody
	GetData() *ImitatePhotoStyleResponseBodyData
	SetRequestId(v string) *ImitatePhotoStyleResponseBody
	GetRequestId() *string
}

type ImitatePhotoStyleResponseBody struct {
	Data *ImitatePhotoStyleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// A880432B-6D9A-4EF4-B7B7-863F38A930D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImitatePhotoStyleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImitatePhotoStyleResponseBody) GoString() string {
	return s.String()
}

func (s *ImitatePhotoStyleResponseBody) GetData() *ImitatePhotoStyleResponseBodyData {
	return s.Data
}

func (s *ImitatePhotoStyleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImitatePhotoStyleResponseBody) SetData(v *ImitatePhotoStyleResponseBodyData) *ImitatePhotoStyleResponseBody {
	s.Data = v
	return s
}

func (s *ImitatePhotoStyleResponseBody) SetRequestId(v string) *ImitatePhotoStyleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImitatePhotoStyleResponseBody) Validate() error {
	return dara.Validate(s)
}

type ImitatePhotoStyleResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-gd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/photo-style-imitation/7c4c0809-5e15-4ca7-84b3-ba16711e3255__5cb220200622-075203.jpg?Expires=1592814125&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=DNhhRFPbMBwpHCEhrLdL%2BBF%2BXs****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ImitatePhotoStyleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImitatePhotoStyleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImitatePhotoStyleResponseBodyData) GetImageURL() *string {
	return s.ImageURL
}

func (s *ImitatePhotoStyleResponseBodyData) SetImageURL(v string) *ImitatePhotoStyleResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *ImitatePhotoStyleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
