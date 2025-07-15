// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatMediaUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateChatMediaUrlResponseBody
	GetCode() *string
	SetData(v *CreateChatMediaUrlResponseBodyData) *CreateChatMediaUrlResponseBody
	GetData() *CreateChatMediaUrlResponseBodyData
	SetHttpStatusCode(v int32) *CreateChatMediaUrlResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateChatMediaUrlResponseBody
	GetMessage() *string
	SetParams(v []*string) *CreateChatMediaUrlResponseBody
	GetParams() []*string
	SetRequestId(v string) *CreateChatMediaUrlResponseBody
	GetRequestId() *string
}

type CreateChatMediaUrlResponseBody struct {
	// example:
	//
	// OK
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateChatMediaUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 03C67DAD-EB26-41D8-949D-9B0C470FB716
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateChatMediaUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateChatMediaUrlResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChatMediaUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateChatMediaUrlResponseBody) GetData() *CreateChatMediaUrlResponseBodyData {
	return s.Data
}

func (s *CreateChatMediaUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateChatMediaUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateChatMediaUrlResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CreateChatMediaUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateChatMediaUrlResponseBody) SetCode(v string) *CreateChatMediaUrlResponseBody {
	s.Code = &v
	return s
}

func (s *CreateChatMediaUrlResponseBody) SetData(v *CreateChatMediaUrlResponseBodyData) *CreateChatMediaUrlResponseBody {
	s.Data = v
	return s
}

func (s *CreateChatMediaUrlResponseBody) SetHttpStatusCode(v int32) *CreateChatMediaUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateChatMediaUrlResponseBody) SetMessage(v string) *CreateChatMediaUrlResponseBody {
	s.Message = &v
	return s
}

func (s *CreateChatMediaUrlResponseBody) SetParams(v []*string) *CreateChatMediaUrlResponseBody {
	s.Params = v
	return s
}

func (s *CreateChatMediaUrlResponseBody) SetRequestId(v string) *CreateChatMediaUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateChatMediaUrlResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateChatMediaUrlResponseBodyData struct {
	// example:
	//
	// $iAHNCNQCo21wMwMGBAAFAAbaACOEAaQhIH6TAqogDGyb-qD2Hbj0A88AAAGRLKYVnwTOACwwYwcACM8AAAGRLRPynQ
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// https://ccc-v2-online.oss-cn-shanghai.aliyuncs.com/ccc-test/namelist.csv?Expires=1642067227&OSSAccessKeyId=****&Signature=****
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateChatMediaUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateChatMediaUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateChatMediaUrlResponseBodyData) GetMediaId() *string {
	return s.MediaId
}

func (s *CreateChatMediaUrlResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *CreateChatMediaUrlResponseBodyData) SetMediaId(v string) *CreateChatMediaUrlResponseBodyData {
	s.MediaId = &v
	return s
}

func (s *CreateChatMediaUrlResponseBodyData) SetUrl(v string) *CreateChatMediaUrlResponseBodyData {
	s.Url = &v
	return s
}

func (s *CreateChatMediaUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
