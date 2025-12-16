// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserAnalyzerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveUserAnalyzerResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *RemoveUserAnalyzerResponseBody
	GetResult() map[string]interface{}
}

type RemoveUserAnalyzerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result set. This parameter is not returned if the request is successful.
	//
	// example:
	//
	// []
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveUserAnalyzerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserAnalyzerResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserAnalyzerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveUserAnalyzerResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *RemoveUserAnalyzerResponseBody) SetRequestId(v string) *RemoveUserAnalyzerResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUserAnalyzerResponseBody) SetResult(v map[string]interface{}) *RemoveUserAnalyzerResponseBody {
	s.Result = v
	return s
}

func (s *RemoveUserAnalyzerResponseBody) Validate() error {
	return dara.Validate(s)
}
