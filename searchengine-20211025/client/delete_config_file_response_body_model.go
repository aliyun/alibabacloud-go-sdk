// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteConfigFileResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *DeleteConfigFileResponseBody
	GetResult() map[string]interface{}
}

type DeleteConfigFileResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteConfigFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConfigFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConfigFileResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *DeleteConfigFileResponseBody) SetRequestId(v string) *DeleteConfigFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConfigFileResponseBody) SetResult(v map[string]interface{}) *DeleteConfigFileResponseBody {
	s.Result = v
	return s
}

func (s *DeleteConfigFileResponseBody) Validate() error {
	return dara.Validate(s)
}
