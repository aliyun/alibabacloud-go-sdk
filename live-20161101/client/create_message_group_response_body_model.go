// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMessageGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMessageGroupResponseBody
	GetRequestId() *string
	SetResult(v *CreateMessageGroupResponseBodyResult) *CreateMessageGroupResponseBody
	GetResult() *CreateMessageGroupResponseBodyResult
}

type CreateMessageGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *CreateMessageGroupResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateMessageGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMessageGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMessageGroupResponseBody) GetResult() *CreateMessageGroupResponseBodyResult {
	return s.Result
}

func (s *CreateMessageGroupResponseBody) SetRequestId(v string) *CreateMessageGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMessageGroupResponseBody) SetResult(v *CreateMessageGroupResponseBodyResult) *CreateMessageGroupResponseBody {
	s.Result = v
	return s
}

func (s *CreateMessageGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateMessageGroupResponseBodyResult struct {
	// The extended field.
	//
	// example:
	//
	// test001
	Extension map[string]interface{} `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The ID of the message group.
	//
	// example:
	//
	// AE35-****-T95F
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s CreateMessageGroupResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateMessageGroupResponseBodyResult) GetExtension() map[string]interface{} {
	return s.Extension
}

func (s *CreateMessageGroupResponseBodyResult) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateMessageGroupResponseBodyResult) SetExtension(v map[string]interface{}) *CreateMessageGroupResponseBodyResult {
	s.Extension = v
	return s
}

func (s *CreateMessageGroupResponseBodyResult) SetGroupId(v string) *CreateMessageGroupResponseBodyResult {
	s.GroupId = &v
	return s
}

func (s *CreateMessageGroupResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
