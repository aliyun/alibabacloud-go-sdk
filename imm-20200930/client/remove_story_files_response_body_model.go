// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveStoryFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveStoryFilesResponseBody
	GetRequestId() *string
}

type RemoveStoryFilesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6E93D6C9-5AC0-49F9-914D-E02678D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveStoryFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveStoryFilesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveStoryFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveStoryFilesResponseBody) SetRequestId(v string) *RemoveStoryFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveStoryFilesResponseBody) Validate() error {
	return dara.Validate(s)
}
