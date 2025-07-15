// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessChannelOfStagingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAccessChannelOfStagingResponseBody
	GetCode() *string
	SetData(v *GetAccessChannelOfStagingResponseBodyData) *GetAccessChannelOfStagingResponseBody
	GetData() *GetAccessChannelOfStagingResponseBodyData
	SetHttpStatusCode(v int32) *GetAccessChannelOfStagingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAccessChannelOfStagingResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAccessChannelOfStagingResponseBody
	GetRequestId() *string
}

type GetAccessChannelOfStagingResponseBody struct {
	Code           *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetAccessChannelOfStagingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccessChannelOfStagingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccessChannelOfStagingResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessChannelOfStagingResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAccessChannelOfStagingResponseBody) GetData() *GetAccessChannelOfStagingResponseBodyData {
	return s.Data
}

func (s *GetAccessChannelOfStagingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAccessChannelOfStagingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAccessChannelOfStagingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccessChannelOfStagingResponseBody) SetCode(v string) *GetAccessChannelOfStagingResponseBody {
	s.Code = &v
	return s
}

func (s *GetAccessChannelOfStagingResponseBody) SetData(v *GetAccessChannelOfStagingResponseBodyData) *GetAccessChannelOfStagingResponseBody {
	s.Data = v
	return s
}

func (s *GetAccessChannelOfStagingResponseBody) SetHttpStatusCode(v int32) *GetAccessChannelOfStagingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAccessChannelOfStagingResponseBody) SetMessage(v string) *GetAccessChannelOfStagingResponseBody {
	s.Message = &v
	return s
}

func (s *GetAccessChannelOfStagingResponseBody) SetRequestId(v string) *GetAccessChannelOfStagingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccessChannelOfStagingResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAccessChannelOfStagingResponseBodyData struct {
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetAccessChannelOfStagingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAccessChannelOfStagingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAccessChannelOfStagingResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *GetAccessChannelOfStagingResponseBodyData) SetToken(v string) *GetAccessChannelOfStagingResponseBodyData {
	s.Token = &v
	return s
}

func (s *GetAccessChannelOfStagingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
