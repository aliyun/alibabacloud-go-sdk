// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyBroadcastSceneFromTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CopyBroadcastSceneFromTemplateResponseBody
	GetCode() *string
	SetData(v *BroadcastScene) *CopyBroadcastSceneFromTemplateResponseBody
	GetData() *BroadcastScene
	SetHttpStatusCode(v int32) *CopyBroadcastSceneFromTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CopyBroadcastSceneFromTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *CopyBroadcastSceneFromTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CopyBroadcastSceneFromTemplateResponseBody
	GetSuccess() *bool
}

type CopyBroadcastSceneFromTemplateResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string         `json:"code,omitempty" xml:"code,omitempty"`
	Data *BroadcastScene `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 435B4F80-8DEB-5CF6-AC86-395CB6CF28C9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CopyBroadcastSceneFromTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyBroadcastSceneFromTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CopyBroadcastSceneFromTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *CopyBroadcastSceneFromTemplateResponseBody) GetData() *BroadcastScene {
	return s.Data
}

func (s *CopyBroadcastSceneFromTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CopyBroadcastSceneFromTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CopyBroadcastSceneFromTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyBroadcastSceneFromTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CopyBroadcastSceneFromTemplateResponseBody) SetCode(v string) *CopyBroadcastSceneFromTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CopyBroadcastSceneFromTemplateResponseBody) SetData(v *BroadcastScene) *CopyBroadcastSceneFromTemplateResponseBody {
	s.Data = v
	return s
}

func (s *CopyBroadcastSceneFromTemplateResponseBody) SetHttpStatusCode(v int32) *CopyBroadcastSceneFromTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CopyBroadcastSceneFromTemplateResponseBody) SetMessage(v string) *CopyBroadcastSceneFromTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CopyBroadcastSceneFromTemplateResponseBody) SetRequestId(v string) *CopyBroadcastSceneFromTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyBroadcastSceneFromTemplateResponseBody) SetSuccess(v bool) *CopyBroadcastSceneFromTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *CopyBroadcastSceneFromTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
