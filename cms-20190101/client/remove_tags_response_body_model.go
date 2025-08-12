// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveTagsResponseBody
	GetCode() *string
	SetMessage(v string) *RemoveTagsResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveTagsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveTagsResponseBody
	GetSuccess() *bool
	SetTag(v *RemoveTagsResponseBodyTag) *RemoveTagsResponseBody
	GetTag() *RemoveTagsResponseBodyTag
}

type RemoveTagsResponseBody struct {
	// The status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// Illegal parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 195390D2-69D0-4D9E-81AA-A7F5BC1B91EB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The deleted tags.
	Tag *RemoveTagsResponseBodyTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Struct"`
}

func (s RemoveTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveTagsResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTagsResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveTagsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveTagsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveTagsResponseBody) GetTag() *RemoveTagsResponseBodyTag {
	return s.Tag
}

func (s *RemoveTagsResponseBody) SetCode(v string) *RemoveTagsResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveTagsResponseBody) SetMessage(v string) *RemoveTagsResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveTagsResponseBody) SetRequestId(v string) *RemoveTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveTagsResponseBody) SetSuccess(v bool) *RemoveTagsResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveTagsResponseBody) SetTag(v *RemoveTagsResponseBodyTag) *RemoveTagsResponseBody {
	s.Tag = v
	return s
}

func (s *RemoveTagsResponseBody) Validate() error {
	return dara.Validate(s)
}

type RemoveTagsResponseBodyTag struct {
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s RemoveTagsResponseBodyTag) String() string {
	return dara.Prettify(s)
}

func (s RemoveTagsResponseBodyTag) GoString() string {
	return s.String()
}

func (s *RemoveTagsResponseBodyTag) GetTags() []*string {
	return s.Tags
}

func (s *RemoveTagsResponseBodyTag) SetTags(v []*string) *RemoveTagsResponseBodyTag {
	s.Tags = v
	return s
}

func (s *RemoveTagsResponseBodyTag) Validate() error {
	return dara.Validate(s)
}
