// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PublishScriptResponseBody
	GetCode() *string
	SetData(v string) *PublishScriptResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *PublishScriptResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *PublishScriptResponseBody
	GetMessage() *string
	SetParams(v []*string) *PublishScriptResponseBody
	GetParams() []*string
	SetRequestId(v string) *PublishScriptResponseBody
	GetRequestId() *string
}

type PublishScriptResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 64241e64-190c-45d1-af66-06f51c07b091
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance af81a389-91f0-4157-8d82-720edd02b66a
	//
	//  does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishScriptResponseBody) GoString() string {
	return s.String()
}

func (s *PublishScriptResponseBody) GetCode() *string {
	return s.Code
}

func (s *PublishScriptResponseBody) GetData() *string {
	return s.Data
}

func (s *PublishScriptResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PublishScriptResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PublishScriptResponseBody) GetParams() []*string {
	return s.Params
}

func (s *PublishScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishScriptResponseBody) SetCode(v string) *PublishScriptResponseBody {
	s.Code = &v
	return s
}

func (s *PublishScriptResponseBody) SetData(v string) *PublishScriptResponseBody {
	s.Data = &v
	return s
}

func (s *PublishScriptResponseBody) SetHttpStatusCode(v int32) *PublishScriptResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PublishScriptResponseBody) SetMessage(v string) *PublishScriptResponseBody {
	s.Message = &v
	return s
}

func (s *PublishScriptResponseBody) SetParams(v []*string) *PublishScriptResponseBody {
	s.Params = v
	return s
}

func (s *PublishScriptResponseBody) SetRequestId(v string) *PublishScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishScriptResponseBody) Validate() error {
	return dara.Validate(s)
}
