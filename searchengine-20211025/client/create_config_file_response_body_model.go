// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateConfigFileResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *CreateConfigFileResponseBody
	GetResult() map[string]interface{}
}

type CreateConfigFileResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FE03180A-0E29-5474-8A86-33F0683294A4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateConfigFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConfigFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConfigFileResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *CreateConfigFileResponseBody) SetRequestId(v string) *CreateConfigFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConfigFileResponseBody) SetResult(v map[string]interface{}) *CreateConfigFileResponseBody {
	s.Result = v
	return s
}

func (s *CreateConfigFileResponseBody) Validate() error {
	return dara.Validate(s)
}
