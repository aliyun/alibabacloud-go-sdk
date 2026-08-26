// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveMessageAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateLiveMessageAppResponseBody
	GetAppId() *string
	SetAppKey(v string) *CreateLiveMessageAppResponseBody
	GetAppKey() *string
	SetAppSign(v string) *CreateLiveMessageAppResponseBody
	GetAppSign() *string
	SetDataCenter(v string) *CreateLiveMessageAppResponseBody
	GetDataCenter() *string
	SetRequestId(v string) *CreateLiveMessageAppResponseBody
	GetRequestId() *string
}

type CreateLiveMessageAppResponseBody struct {
	// The application ID, which is used for subsequent operations such as joining groups.
	//
	// example:
	//
	// demo
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The AppKey, which is used to generate authentication for various operations related to the AppId.
	//
	// example:
	//
	// **********************************
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The application signature. The interactive messaging service SDK requires this information.
	//
	// example:
	//
	// **************************************************************************
	AppSign *string `json:"AppSign,omitempty" xml:"AppSign,omitempty"`
	// The data center.
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 65EEDBEB-43FE-1E15-976F-3DDD753A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLiveMessageAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveMessageAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLiveMessageAppResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *CreateLiveMessageAppResponseBody) GetAppKey() *string {
	return s.AppKey
}

func (s *CreateLiveMessageAppResponseBody) GetAppSign() *string {
	return s.AppSign
}

func (s *CreateLiveMessageAppResponseBody) GetDataCenter() *string {
	return s.DataCenter
}

func (s *CreateLiveMessageAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLiveMessageAppResponseBody) SetAppId(v string) *CreateLiveMessageAppResponseBody {
	s.AppId = &v
	return s
}

func (s *CreateLiveMessageAppResponseBody) SetAppKey(v string) *CreateLiveMessageAppResponseBody {
	s.AppKey = &v
	return s
}

func (s *CreateLiveMessageAppResponseBody) SetAppSign(v string) *CreateLiveMessageAppResponseBody {
	s.AppSign = &v
	return s
}

func (s *CreateLiveMessageAppResponseBody) SetDataCenter(v string) *CreateLiveMessageAppResponseBody {
	s.DataCenter = &v
	return s
}

func (s *CreateLiveMessageAppResponseBody) SetRequestId(v string) *CreateLiveMessageAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLiveMessageAppResponseBody) Validate() error {
	return dara.Validate(s)
}
