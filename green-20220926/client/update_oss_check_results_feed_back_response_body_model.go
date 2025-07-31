// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOssCheckResultsFeedBackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateOssCheckResultsFeedBackResponseBody
	GetData() *bool
	SetRequestId(v string) *UpdateOssCheckResultsFeedBackResponseBody
	GetRequestId() *string
}

type UpdateOssCheckResultsFeedBackResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOssCheckResultsFeedBackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOssCheckResultsFeedBackResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOssCheckResultsFeedBackResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateOssCheckResultsFeedBackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOssCheckResultsFeedBackResponseBody) SetData(v bool) *UpdateOssCheckResultsFeedBackResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateOssCheckResultsFeedBackResponseBody) SetRequestId(v string) *UpdateOssCheckResultsFeedBackResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOssCheckResultsFeedBackResponseBody) Validate() error {
	return dara.Validate(s)
}
