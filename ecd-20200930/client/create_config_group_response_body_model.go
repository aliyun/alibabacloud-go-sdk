// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *CreateConfigGroupResponseBody
	GetGroupId() *string
	SetMessage(v string) *CreateConfigGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateConfigGroupResponseBody
	GetRequestId() *string
}

type CreateConfigGroupResponseBody struct {
	// The ID of the configuration group.
	//
	// example:
	//
	// ccg-0ctwi5zbswtql****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The creation result of the configuration group.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EE9472BC-0B5D-5458-85CD-C52BDD******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateConfigGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConfigGroupResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateConfigGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateConfigGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConfigGroupResponseBody) SetGroupId(v string) *CreateConfigGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateConfigGroupResponseBody) SetMessage(v string) *CreateConfigGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateConfigGroupResponseBody) SetRequestId(v string) *CreateConfigGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConfigGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
