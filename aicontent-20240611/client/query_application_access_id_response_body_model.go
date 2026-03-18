// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryApplicationAccessIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryApplicationAccessIdResponseBodyData) *QueryApplicationAccessIdResponseBody
	GetData() *QueryApplicationAccessIdResponseBodyData
	SetErrCode(v string) *QueryApplicationAccessIdResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QueryApplicationAccessIdResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *QueryApplicationAccessIdResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryApplicationAccessIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryApplicationAccessIdResponseBody
	GetSuccess() *bool
}

type QueryApplicationAccessIdResponseBody struct {
	// example:
	//
	// []
	Data *QueryApplicationAccessIdResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryApplicationAccessIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryApplicationAccessIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryApplicationAccessIdResponseBody) GetData() *QueryApplicationAccessIdResponseBodyData {
	return s.Data
}

func (s *QueryApplicationAccessIdResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QueryApplicationAccessIdResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QueryApplicationAccessIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryApplicationAccessIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryApplicationAccessIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryApplicationAccessIdResponseBody) SetData(v *QueryApplicationAccessIdResponseBodyData) *QueryApplicationAccessIdResponseBody {
	s.Data = v
	return s
}

func (s *QueryApplicationAccessIdResponseBody) SetErrCode(v string) *QueryApplicationAccessIdResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryApplicationAccessIdResponseBody) SetErrMessage(v string) *QueryApplicationAccessIdResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryApplicationAccessIdResponseBody) SetHttpStatusCode(v int32) *QueryApplicationAccessIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryApplicationAccessIdResponseBody) SetRequestId(v string) *QueryApplicationAccessIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryApplicationAccessIdResponseBody) SetSuccess(v bool) *QueryApplicationAccessIdResponseBody {
	s.Success = &v
	return s
}

func (s *QueryApplicationAccessIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryApplicationAccessIdResponseBodyData struct {
	// example:
	//
	// 1234567890
	ApplicationAccessId *string `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
	// example:
	//
	// MyAppSecret
	ApplicationAccessSecret *string `json:"applicationAccessSecret,omitempty" xml:"applicationAccessSecret,omitempty"`
}

func (s QueryApplicationAccessIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryApplicationAccessIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryApplicationAccessIdResponseBodyData) GetApplicationAccessId() *string {
	return s.ApplicationAccessId
}

func (s *QueryApplicationAccessIdResponseBodyData) GetApplicationAccessSecret() *string {
	return s.ApplicationAccessSecret
}

func (s *QueryApplicationAccessIdResponseBodyData) SetApplicationAccessId(v string) *QueryApplicationAccessIdResponseBodyData {
	s.ApplicationAccessId = &v
	return s
}

func (s *QueryApplicationAccessIdResponseBodyData) SetApplicationAccessSecret(v string) *QueryApplicationAccessIdResponseBodyData {
	s.ApplicationAccessSecret = &v
	return s
}

func (s *QueryApplicationAccessIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
