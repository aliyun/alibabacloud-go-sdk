// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteConfigTemplateResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteConfigTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteConfigTemplateResponseBody
	GetRequestId() *string
}

type DeleteConfigTemplateResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D16979DC-4D42-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConfigTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConfigTemplateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteConfigTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteConfigTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConfigTemplateResponseBody) SetCode(v int32) *DeleteConfigTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteConfigTemplateResponseBody) SetMessage(v string) *DeleteConfigTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteConfigTemplateResponseBody) SetRequestId(v string) *DeleteConfigTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConfigTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
