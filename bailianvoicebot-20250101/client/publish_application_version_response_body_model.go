// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishApplicationVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PublishApplicationVersionResponseBody
	GetCode() *string
	SetData(v string) *PublishApplicationVersionResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *PublishApplicationVersionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *PublishApplicationVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *PublishApplicationVersionResponseBody
	GetRequestId() *string
}

type PublishApplicationVersionResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 82ea16d1-425c-4c03-9be5-cc91de9779ed
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CF6D3484-19A1-5C77-863B-AC8B5754D37C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishApplicationVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishApplicationVersionResponseBody) GoString() string {
	return s.String()
}

func (s *PublishApplicationVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *PublishApplicationVersionResponseBody) GetData() *string {
	return s.Data
}

func (s *PublishApplicationVersionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PublishApplicationVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PublishApplicationVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishApplicationVersionResponseBody) SetCode(v string) *PublishApplicationVersionResponseBody {
	s.Code = &v
	return s
}

func (s *PublishApplicationVersionResponseBody) SetData(v string) *PublishApplicationVersionResponseBody {
	s.Data = &v
	return s
}

func (s *PublishApplicationVersionResponseBody) SetHttpStatusCode(v int32) *PublishApplicationVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PublishApplicationVersionResponseBody) SetMessage(v string) *PublishApplicationVersionResponseBody {
	s.Message = &v
	return s
}

func (s *PublishApplicationVersionResponseBody) SetRequestId(v string) *PublishApplicationVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishApplicationVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
