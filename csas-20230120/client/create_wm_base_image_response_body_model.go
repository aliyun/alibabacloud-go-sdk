// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWmBaseImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateWmBaseImageResponseBodyData) *CreateWmBaseImageResponseBody
	GetData() *CreateWmBaseImageResponseBodyData
	SetRequestId(v string) *CreateWmBaseImageResponseBody
	GetRequestId() *string
}

type CreateWmBaseImageResponseBody struct {
	Data *CreateWmBaseImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWmBaseImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWmBaseImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWmBaseImageResponseBody) GetData() *CreateWmBaseImageResponseBodyData {
	return s.Data
}

func (s *CreateWmBaseImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWmBaseImageResponseBody) SetData(v *CreateWmBaseImageResponseBodyData) *CreateWmBaseImageResponseBody {
	s.Data = v
	return s
}

func (s *CreateWmBaseImageResponseBody) SetRequestId(v string) *CreateWmBaseImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWmBaseImageResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateWmBaseImageResponseBodyData struct {
	// example:
	//
	// fafb432cdede9b20640e12105845386e-496883833-8242409229217337*****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// https://example.com/test-*****.png
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// 17185*****
	ImageUrlExp *int64 `json:"ImageUrlExp,omitempty" xml:"ImageUrlExp,omitempty"`
}

func (s CreateWmBaseImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateWmBaseImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateWmBaseImageResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *CreateWmBaseImageResponseBodyData) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CreateWmBaseImageResponseBodyData) GetImageUrlExp() *int64 {
	return s.ImageUrlExp
}

func (s *CreateWmBaseImageResponseBodyData) SetImageId(v string) *CreateWmBaseImageResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *CreateWmBaseImageResponseBodyData) SetImageUrl(v string) *CreateWmBaseImageResponseBodyData {
	s.ImageUrl = &v
	return s
}

func (s *CreateWmBaseImageResponseBodyData) SetImageUrlExp(v int64) *CreateWmBaseImageResponseBodyData {
	s.ImageUrlExp = &v
	return s
}

func (s *CreateWmBaseImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
