// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommitServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CommitServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CommitServiceResponseBody
	GetRequestId() *string
}

type CommitServiceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CommitServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CommitServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CommitServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CommitServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CommitServiceResponseBody) SetMessage(v string) *CommitServiceResponseBody {
	s.Message = &v
	return s
}

func (s *CommitServiceResponseBody) SetRequestId(v string) *CommitServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CommitServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
