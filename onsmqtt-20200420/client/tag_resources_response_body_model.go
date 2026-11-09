// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TagResourcesResponseBody
	GetCode() *string
	SetData(v bool) *TagResourcesResponseBody
	GetData() *bool
	SetMessage(v string) *TagResourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *TagResourcesResponseBody
	GetRequestId() *string
}

type TagResourcesResponseBody struct {
	// The status code of the API call.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the call.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 82B9E503-F4A1-4F30-976F-C6999FF9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *TagResourcesResponseBody) GetData() *bool {
	return s.Data
}

func (s *TagResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TagResourcesResponseBody) SetCode(v string) *TagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *TagResourcesResponseBody) SetData(v bool) *TagResourcesResponseBody {
	s.Data = &v
	return s
}

func (s *TagResourcesResponseBody) SetMessage(v string) *TagResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
