// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCartoonizedImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GenerateCartoonizedImageResponseBodyData) *GenerateCartoonizedImageResponseBody
	GetData() *GenerateCartoonizedImageResponseBodyData
	SetMessage(v string) *GenerateCartoonizedImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateCartoonizedImageResponseBody
	GetRequestId() *string
}

type GenerateCartoonizedImageResponseBody struct {
	Data    *GenerateCartoonizedImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 48f38719-f0c2-4784-a9cc-30df95e393a9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateCartoonizedImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateCartoonizedImageResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateCartoonizedImageResponseBody) GetData() *GenerateCartoonizedImageResponseBodyData {
	return s.Data
}

func (s *GenerateCartoonizedImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateCartoonizedImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateCartoonizedImageResponseBody) SetData(v *GenerateCartoonizedImageResponseBodyData) *GenerateCartoonizedImageResponseBody {
	s.Data = v
	return s
}

func (s *GenerateCartoonizedImageResponseBody) SetMessage(v string) *GenerateCartoonizedImageResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateCartoonizedImageResponseBody) SetRequestId(v string) *GenerateCartoonizedImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateCartoonizedImageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateCartoonizedImageResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-damo-eas-cn-shanghai.oss-cn-shanghai.aliyuncs.com/generative-cartoon/2023-02-02/5a3e5760-ff27-4321-8976-d05656fb716a/20230202_154009511910_pclb0gomva.jpg?Expires=1675325422&amp;OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&amp;Signature=UmAa7HxeumVkDfrdoL02dtztwS****
	ResultUrl *string `json:"ResultUrl,omitempty" xml:"ResultUrl,omitempty"`
}

func (s GenerateCartoonizedImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateCartoonizedImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateCartoonizedImageResponseBodyData) GetResultUrl() *string {
	return s.ResultUrl
}

func (s *GenerateCartoonizedImageResponseBodyData) SetResultUrl(v string) *GenerateCartoonizedImageResponseBodyData {
	s.ResultUrl = &v
	return s
}

func (s *GenerateCartoonizedImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
