// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRayLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRayLogResponseBody
	GetCode() *string
	SetData(v *GetRayLogResponseBodyData) *GetRayLogResponseBody
	GetData() *GetRayLogResponseBodyData
	SetMessage(v string) *GetRayLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRayLogResponseBody
	GetRequestId() *string
}

type GetRayLogResponseBody struct {
	// The response code. A value of 1000000 indicates a successful request. Other values indicate a failed request. You can view the specific error description in the message parameter.
	//
	// example:
	//
	// 1000000
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetRayLogResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B06AADC1-1627-5B1C-B62F-A2226C122897
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRayLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRayLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetRayLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRayLogResponseBody) GetData() *GetRayLogResponseBodyData {
	return s.Data
}

func (s *GetRayLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRayLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRayLogResponseBody) SetCode(v string) *GetRayLogResponseBody {
	s.Code = &v
	return s
}

func (s *GetRayLogResponseBody) SetData(v *GetRayLogResponseBodyData) *GetRayLogResponseBody {
	s.Data = v
	return s
}

func (s *GetRayLogResponseBody) SetMessage(v string) *GetRayLogResponseBody {
	s.Message = &v
	return s
}

func (s *GetRayLogResponseBody) SetRequestId(v string) *GetRayLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRayLogResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRayLogResponseBodyData struct {
	// The file download URL.
	//
	// example:
	//
	// https://mybucket.cn-hangzhou.com/xxxxxx
	AccessUrl *string `json:"accessUrl,omitempty" xml:"accessUrl,omitempty"`
}

func (s GetRayLogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRayLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRayLogResponseBodyData) GetAccessUrl() *string {
	return s.AccessUrl
}

func (s *GetRayLogResponseBodyData) SetAccessUrl(v string) *GetRayLogResponseBodyData {
	s.AccessUrl = &v
	return s
}

func (s *GetRayLogResponseBodyData) Validate() error {
	return dara.Validate(s)
}
