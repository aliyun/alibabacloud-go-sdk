// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishIndexVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PublishIndexVersionResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *PublishIndexVersionResponseBody
	GetResult() map[string]interface{}
}

type PublishIndexVersionResponseBody struct {
	// id of request
	//
	// example:
	//
	// E45380E8-994A-5402-9806-F114B3295FCF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the index
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s PublishIndexVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishIndexVersionResponseBody) GoString() string {
	return s.String()
}

func (s *PublishIndexVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishIndexVersionResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *PublishIndexVersionResponseBody) SetRequestId(v string) *PublishIndexVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishIndexVersionResponseBody) SetResult(v map[string]interface{}) *PublishIndexVersionResponseBody {
	s.Result = v
	return s
}

func (s *PublishIndexVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
