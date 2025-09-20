// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomizedStoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetObjectId(v string) *CreateCustomizedStoryResponseBody
	GetObjectId() *string
	SetRequestId(v string) *CreateCustomizedStoryResponseBody
	GetRequestId() *string
}

type CreateCustomizedStoryResponseBody struct {
	// The ID of the story.
	//
	// example:
	//
	// 563062c0b085733f34ab****
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BC91D091-D49F-0ACD-95D5-F0621045****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomizedStoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomizedStoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomizedStoryResponseBody) GetObjectId() *string {
	return s.ObjectId
}

func (s *CreateCustomizedStoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomizedStoryResponseBody) SetObjectId(v string) *CreateCustomizedStoryResponseBody {
	s.ObjectId = &v
	return s
}

func (s *CreateCustomizedStoryResponseBody) SetRequestId(v string) *CreateCustomizedStoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomizedStoryResponseBody) Validate() error {
	return dara.Validate(s)
}
