// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileDetectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHashKey(v string) *CreateFileDetectResponseBody
	GetHashKey() *string
	SetRequestId(v string) *CreateFileDetectResponseBody
	GetRequestId() *string
}

type CreateFileDetectResponseBody struct {
	// The identifier of the file.
	//
	// example:
	//
	// 0a212417e65c26ff133cfff28f6c****
	HashKey *string `json:"HashKey,omitempty" xml:"HashKey,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFileDetectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFileDetectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileDetectResponseBody) GetHashKey() *string {
	return s.HashKey
}

func (s *CreateFileDetectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFileDetectResponseBody) SetHashKey(v string) *CreateFileDetectResponseBody {
	s.HashKey = &v
	return s
}

func (s *CreateFileDetectResponseBody) SetRequestId(v string) *CreateFileDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFileDetectResponseBody) Validate() error {
	return dara.Validate(s)
}
