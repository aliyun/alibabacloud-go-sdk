// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetClusterInterceptionConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetClusterInterceptionConfigResponseBody
	GetRequestId() *string
	SetResult(v bool) *SetClusterInterceptionConfigResponseBody
	GetResult() *bool
}

type SetClusterInterceptionConfigResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 74AB990C-539B-579B-9239-B8A2036B7337
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s SetClusterInterceptionConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetClusterInterceptionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetClusterInterceptionConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetClusterInterceptionConfigResponseBody) GetResult() *bool {
	return s.Result
}

func (s *SetClusterInterceptionConfigResponseBody) SetRequestId(v string) *SetClusterInterceptionConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetClusterInterceptionConfigResponseBody) SetResult(v bool) *SetClusterInterceptionConfigResponseBody {
	s.Result = &v
	return s
}

func (s *SetClusterInterceptionConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
