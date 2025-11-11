// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateViewPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GenerateViewPointResponseBody
	GetCode() *string
	SetData(v []*GenerateViewPointResponseBodyData) *GenerateViewPointResponseBody
	GetData() []*GenerateViewPointResponseBodyData
	SetHttpStatusCode(v int32) *GenerateViewPointResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GenerateViewPointResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateViewPointResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GenerateViewPointResponseBody
	GetSuccess() *bool
}

type GenerateViewPointResponseBody struct {
	// example:
	//
	// 200
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GenerateViewPointResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 94512A33-8EC1-5452-A793-5C91F18ED2F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateViewPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateViewPointResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateViewPointResponseBody) GetCode() *string {
	return s.Code
}

func (s *GenerateViewPointResponseBody) GetData() []*GenerateViewPointResponseBodyData {
	return s.Data
}

func (s *GenerateViewPointResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GenerateViewPointResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateViewPointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateViewPointResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GenerateViewPointResponseBody) SetCode(v string) *GenerateViewPointResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateViewPointResponseBody) SetData(v []*GenerateViewPointResponseBodyData) *GenerateViewPointResponseBody {
	s.Data = v
	return s
}

func (s *GenerateViewPointResponseBody) SetHttpStatusCode(v int32) *GenerateViewPointResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateViewPointResponseBody) SetMessage(v string) *GenerateViewPointResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateViewPointResponseBody) SetRequestId(v string) *GenerateViewPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateViewPointResponseBody) SetSuccess(v bool) *GenerateViewPointResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateViewPointResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GenerateViewPointResponseBodyData struct {
	Point *string `json:"Point,omitempty" xml:"Point,omitempty"`
}

func (s GenerateViewPointResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateViewPointResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateViewPointResponseBodyData) GetPoint() *string {
	return s.Point
}

func (s *GenerateViewPointResponseBodyData) SetPoint(v string) *GenerateViewPointResponseBodyData {
	s.Point = &v
	return s
}

func (s *GenerateViewPointResponseBodyData) Validate() error {
	return dara.Validate(s)
}
