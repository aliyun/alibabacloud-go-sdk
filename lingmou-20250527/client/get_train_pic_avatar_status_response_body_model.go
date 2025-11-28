// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrainPicAvatarStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTrainPicAvatarStatusResponseBody
	GetCode() *string
	SetData(v *GetTrainPicAvatarStatusResponseBodyData) *GetTrainPicAvatarStatusResponseBody
	GetData() *GetTrainPicAvatarStatusResponseBodyData
	SetHttpStatusCode(v int32) *GetTrainPicAvatarStatusResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTrainPicAvatarStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTrainPicAvatarStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTrainPicAvatarStatusResponseBody
	GetSuccess() *bool
}

type GetTrainPicAvatarStatusResponseBody struct {
	// example:
	//
	// 200
	Code *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetTrainPicAvatarStatusResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 7239F9E5-A4DB-55BA-B701-0CE47DBDB0A8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTrainPicAvatarStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTrainPicAvatarStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrainPicAvatarStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTrainPicAvatarStatusResponseBody) GetData() *GetTrainPicAvatarStatusResponseBodyData {
	return s.Data
}

func (s *GetTrainPicAvatarStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTrainPicAvatarStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTrainPicAvatarStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTrainPicAvatarStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTrainPicAvatarStatusResponseBody) SetCode(v string) *GetTrainPicAvatarStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetTrainPicAvatarStatusResponseBody) SetData(v *GetTrainPicAvatarStatusResponseBodyData) *GetTrainPicAvatarStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetTrainPicAvatarStatusResponseBody) SetHttpStatusCode(v int32) *GetTrainPicAvatarStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTrainPicAvatarStatusResponseBody) SetMessage(v string) *GetTrainPicAvatarStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetTrainPicAvatarStatusResponseBody) SetRequestId(v string) *GetTrainPicAvatarStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrainPicAvatarStatusResponseBody) SetSuccess(v bool) *GetTrainPicAvatarStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetTrainPicAvatarStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTrainPicAvatarStatusResponseBodyData struct {
	// example:
	//
	// M1YJTNTH2yoLmLnzKdYHeGBg
	AvatarId *string `json:"avatarId,omitempty" xml:"avatarId,omitempty"`
	// example:
	//
	// //daily-avatar-property.oss-cn-beijing.aliyuncs.com/avatar-share-property/AVATAR_2D_PIC/Mt.CMTMRYX4TNIU2/retalking_output.mp4?Expires=1764327167&OSSAccessKeyId=LTAI5tBRPnF5JkRCidYA8kw9&Signature=%2BH%2BSBpNDPiMQtPyl8vraEHMv9X8%3D
	PreviewURL *string `json:"previewURL,omitempty" xml:"previewURL,omitempty"`
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetTrainPicAvatarStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTrainPicAvatarStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTrainPicAvatarStatusResponseBodyData) GetAvatarId() *string {
	return s.AvatarId
}

func (s *GetTrainPicAvatarStatusResponseBodyData) GetPreviewURL() *string {
	return s.PreviewURL
}

func (s *GetTrainPicAvatarStatusResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetTrainPicAvatarStatusResponseBodyData) SetAvatarId(v string) *GetTrainPicAvatarStatusResponseBodyData {
	s.AvatarId = &v
	return s
}

func (s *GetTrainPicAvatarStatusResponseBodyData) SetPreviewURL(v string) *GetTrainPicAvatarStatusResponseBodyData {
	s.PreviewURL = &v
	return s
}

func (s *GetTrainPicAvatarStatusResponseBodyData) SetStatus(v string) *GetTrainPicAvatarStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTrainPicAvatarStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
