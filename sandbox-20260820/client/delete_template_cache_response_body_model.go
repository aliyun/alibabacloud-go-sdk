// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateCacheResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteTemplateCacheResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteTemplateCacheResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteTemplateCacheResponseBody
	GetRequestId() *string
}

type DeleteTemplateCacheResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// B5AD8B54-4358-5F5B-ACAA-52F2016459C6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteTemplateCacheResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateCacheResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateCacheResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteTemplateCacheResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTemplateCacheResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTemplateCacheResponseBody) SetCode(v string) *DeleteTemplateCacheResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTemplateCacheResponseBody) SetMessage(v string) *DeleteTemplateCacheResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTemplateCacheResponseBody) SetRequestId(v string) *DeleteTemplateCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTemplateCacheResponseBody) Validate() error {
	return dara.Validate(s)
}
