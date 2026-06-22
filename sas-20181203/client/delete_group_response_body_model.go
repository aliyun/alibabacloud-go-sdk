// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteGroupResponseBody
	GetCode() *string
	SetRequestId(v string) *DeleteGroupResponseBody
	GetRequestId() *string
}

type DeleteGroupResponseBody struct {
	// The result code. A value of **200*	- indicates success. Other values indicate failure. You can use this field to determine the cause of the failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The request ID. Alibaba Cloud generates a unique ID for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGroupResponseBody) SetCode(v string) *DeleteGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGroupResponseBody) SetRequestId(v string) *DeleteGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
