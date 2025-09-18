// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishMmAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PublishMmAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PublishMmAppResponseBody
	GetSuccess() *bool
}

type PublishMmAppResponseBody struct {
	// example:
	//
	// xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PublishMmAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishMmAppResponseBody) GoString() string {
	return s.String()
}

func (s *PublishMmAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishMmAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PublishMmAppResponseBody) SetRequestId(v string) *PublishMmAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishMmAppResponseBody) SetSuccess(v bool) *PublishMmAppResponseBody {
	s.Success = &v
	return s
}

func (s *PublishMmAppResponseBody) Validate() error {
	return dara.Validate(s)
}
