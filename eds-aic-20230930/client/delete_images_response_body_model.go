// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteImagesResponseBodyData) *DeleteImagesResponseBody
	GetData() *DeleteImagesResponseBodyData
	SetRequestId(v string) *DeleteImagesResponseBody
	GetRequestId() *string
}

type DeleteImagesResponseBody struct {
	// The images.
	Data *DeleteImagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 4610632D-D661-5982-B3D7-5D3FD183F595
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteImagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteImagesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImagesResponseBody) GetData() *DeleteImagesResponseBodyData {
	return s.Data
}

func (s *DeleteImagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteImagesResponseBody) SetData(v *DeleteImagesResponseBodyData) *DeleteImagesResponseBody {
	s.Data = v
	return s
}

func (s *DeleteImagesResponseBody) SetRequestId(v string) *DeleteImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteImagesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteImagesResponseBodyData struct {
	// The IDs of the images that failed to be deleted.
	FailDeleteImageIds []*string `json:"FailDeleteImageIds,omitempty" xml:"FailDeleteImageIds,omitempty" type:"Repeated"`
	// The IDs of the images that are successfully deleted.
	SuccessDeleteImageIds []*string `json:"SuccessDeleteImageIds,omitempty" xml:"SuccessDeleteImageIds,omitempty" type:"Repeated"`
}

func (s DeleteImagesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteImagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteImagesResponseBodyData) GetFailDeleteImageIds() []*string {
	return s.FailDeleteImageIds
}

func (s *DeleteImagesResponseBodyData) GetSuccessDeleteImageIds() []*string {
	return s.SuccessDeleteImageIds
}

func (s *DeleteImagesResponseBodyData) SetFailDeleteImageIds(v []*string) *DeleteImagesResponseBodyData {
	s.FailDeleteImageIds = v
	return s
}

func (s *DeleteImagesResponseBodyData) SetSuccessDeleteImageIds(v []*string) *DeleteImagesResponseBodyData {
	s.SuccessDeleteImageIds = v
	return s
}

func (s *DeleteImagesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
