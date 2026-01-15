// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteImageResponseBody
	GetCode() *int32
	SetData(v *DeleteImageResponseBodyData) *DeleteImageResponseBody
	GetData() *DeleteImageResponseBodyData
	SetMessage(v string) *DeleteImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteImageResponseBody
	GetSuccess() *bool
}

type DeleteImageResponseBody struct {
	// The error code returned.
	//
	// 	- A value of 0 indicates that the operation is successful.
	//
	// 	- Values other than 0 indicate errors.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the deleted images.
	Data *DeleteImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0703956F-9BCC-48FA-99F7-96C0BF449C69
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteImageResponseBody) GetData() *DeleteImageResponseBodyData {
	return s.Data
}

func (s *DeleteImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteImageResponseBody) SetCode(v int32) *DeleteImageResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteImageResponseBody) SetData(v *DeleteImageResponseBodyData) *DeleteImageResponseBody {
	s.Data = v
	return s
}

func (s *DeleteImageResponseBody) SetMessage(v string) *DeleteImageResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteImageResponseBody) SetRequestId(v string) *DeleteImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteImageResponseBody) SetSuccess(v bool) *DeleteImageResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteImageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteImageResponseBodyData struct {
	// The name (PicName) of the deleted image.
	//
	// example:
	//
	// 5555.jpg
	PicNames []*string `json:"PicNames,omitempty" xml:"PicNames,omitempty" type:"Repeated"`
}

func (s DeleteImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteImageResponseBodyData) GetPicNames() []*string {
	return s.PicNames
}

func (s *DeleteImageResponseBodyData) SetPicNames(v []*string) *DeleteImageResponseBodyData {
	s.PicNames = v
	return s
}

func (s *DeleteImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
