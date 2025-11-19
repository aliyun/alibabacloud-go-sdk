// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNonExistImageIds(v *UpdateImageInfosResponseBodyNonExistImageIds) *UpdateImageInfosResponseBody
	GetNonExistImageIds() *UpdateImageInfosResponseBodyNonExistImageIds
	SetRequestId(v string) *UpdateImageInfosResponseBody
	GetRequestId() *string
}

type UpdateImageInfosResponseBody struct {
	// The IDs of the images that do not exist.
	NonExistImageIds *UpdateImageInfosResponseBodyNonExistImageIds `json:"NonExistImageIds,omitempty" xml:"NonExistImageIds,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateImageInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageInfosResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageInfosResponseBody) GetNonExistImageIds() *UpdateImageInfosResponseBodyNonExistImageIds {
	return s.NonExistImageIds
}

func (s *UpdateImageInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateImageInfosResponseBody) SetNonExistImageIds(v *UpdateImageInfosResponseBodyNonExistImageIds) *UpdateImageInfosResponseBody {
	s.NonExistImageIds = v
	return s
}

func (s *UpdateImageInfosResponseBody) SetRequestId(v string) *UpdateImageInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateImageInfosResponseBody) Validate() error {
	if s.NonExistImageIds != nil {
		if err := s.NonExistImageIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateImageInfosResponseBodyNonExistImageIds struct {
	ImageId []*string `json:"ImageId,omitempty" xml:"ImageId,omitempty" type:"Repeated"`
}

func (s UpdateImageInfosResponseBodyNonExistImageIds) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageInfosResponseBodyNonExistImageIds) GoString() string {
	return s.String()
}

func (s *UpdateImageInfosResponseBodyNonExistImageIds) GetImageId() []*string {
	return s.ImageId
}

func (s *UpdateImageInfosResponseBodyNonExistImageIds) SetImageId(v []*string) *UpdateImageInfosResponseBodyNonExistImageIds {
	s.ImageId = v
	return s
}

func (s *UpdateImageInfosResponseBodyNonExistImageIds) Validate() error {
	return dara.Validate(s)
}
