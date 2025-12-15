// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachNodesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AttachNodesResponseBody
	GetSuccess() *bool
}

type AttachNodesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2263XXXX-XXXX-XXXX-XXXX-XXXX2448XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AttachNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachNodesResponseBody) GoString() string {
	return s.String()
}

func (s *AttachNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachNodesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AttachNodesResponseBody) SetRequestId(v string) *AttachNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachNodesResponseBody) SetSuccess(v bool) *AttachNodesResponseBody {
	s.Success = &v
	return s
}

func (s *AttachNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
