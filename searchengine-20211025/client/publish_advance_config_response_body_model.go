// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishAdvanceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PublishAdvanceConfigResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *PublishAdvanceConfigResponseBody
	GetResult() map[string]interface{}
}

type PublishAdvanceConfigResponseBody struct {
	// The ID of the request
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s PublishAdvanceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishAdvanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *PublishAdvanceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishAdvanceConfigResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *PublishAdvanceConfigResponseBody) SetRequestId(v string) *PublishAdvanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishAdvanceConfigResponseBody) SetResult(v map[string]interface{}) *PublishAdvanceConfigResponseBody {
	s.Result = v
	return s
}

func (s *PublishAdvanceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
