// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateGroupImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GenerateGroupImageResponseBody
	GetCode() *string
	SetData(v *GenerateGroupImageResponseBodyData) *GenerateGroupImageResponseBody
	GetData() *GenerateGroupImageResponseBodyData
	SetMessage(v string) *GenerateGroupImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateGroupImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GenerateGroupImageResponseBody
	GetSuccess() *bool
}

type GenerateGroupImageResponseBody struct {
	// The error code. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// 202
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The composite image generation result.
	Data *GenerateGroupImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateGroupImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateGroupImageResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateGroupImageResponseBody) GetCode() *string {
	return s.Code
}

func (s *GenerateGroupImageResponseBody) GetData() *GenerateGroupImageResponseBodyData {
	return s.Data
}

func (s *GenerateGroupImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateGroupImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateGroupImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GenerateGroupImageResponseBody) SetCode(v string) *GenerateGroupImageResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateGroupImageResponseBody) SetData(v *GenerateGroupImageResponseBodyData) *GenerateGroupImageResponseBody {
	s.Data = v
	return s
}

func (s *GenerateGroupImageResponseBody) SetMessage(v string) *GenerateGroupImageResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateGroupImageResponseBody) SetRequestId(v string) *GenerateGroupImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateGroupImageResponseBody) SetSuccess(v bool) *GenerateGroupImageResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateGroupImageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateGroupImageResponseBodyData struct {
	// The composite image request ID.
	//
	// example:
	//
	// a5561c14-b5cd-49a4-ab79-01a63b10d99c
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s GenerateGroupImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateGroupImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateGroupImageResponseBodyData) GetGroupId() *string {
	return s.GroupId
}

func (s *GenerateGroupImageResponseBodyData) SetGroupId(v string) *GenerateGroupImageResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *GenerateGroupImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
