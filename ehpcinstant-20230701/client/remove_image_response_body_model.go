// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveImageResponseBody
	GetSuccess() *bool
}

type RemoveImageResponseBody struct {
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveImageResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveImageResponseBody) SetRequestId(v string) *RemoveImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveImageResponseBody) SetSuccess(v bool) *RemoveImageResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveImageResponseBody) Validate() error {
	return dara.Validate(s)
}
