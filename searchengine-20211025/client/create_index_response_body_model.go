// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateIndexResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *CreateIndexResponseBody
	GetResult() map[string]interface{}
}

type CreateIndexResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 407BFD91-DE7D-50BA-8F88-CDE52A3B5E46
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The results returned.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIndexResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIndexResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *CreateIndexResponseBody) SetRequestId(v string) *CreateIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIndexResponseBody) SetResult(v map[string]interface{}) *CreateIndexResponseBody {
	s.Result = v
	return s
}

func (s *CreateIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
