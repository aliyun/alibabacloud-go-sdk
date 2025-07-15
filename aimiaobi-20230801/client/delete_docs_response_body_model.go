// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDocsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteDocsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteDocsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDocsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDocsResponseBody
	GetSuccess() *bool
}

type DeleteDocsResponseBody struct {
	// example:
	//
	// successful
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2C565EDD-E624-5FED-8565-0A9CB0C8CC46
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDocsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDocsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDocsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteDocsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDocsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDocsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDocsResponseBody) SetCode(v string) *DeleteDocsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDocsResponseBody) SetHttpStatusCode(v int32) *DeleteDocsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDocsResponseBody) SetMessage(v string) *DeleteDocsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDocsResponseBody) SetRequestId(v string) *DeleteDocsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDocsResponseBody) SetSuccess(v bool) *DeleteDocsResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDocsResponseBody) Validate() error {
	return dara.Validate(s)
}
