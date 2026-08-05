// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOfflineTaskErrorLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListOfflineTaskErrorLogsResponseBody
	GetRequestId() *string
	SetResult(v []*ListOfflineTaskErrorLogsResponseBodyResult) *ListOfflineTaskErrorLogsResponseBody
	GetResult() []*ListOfflineTaskErrorLogsResponseBodyResult
	SetTotalCount(v int32) *ListOfflineTaskErrorLogsResponseBody
	GetTotalCount() *int32
}

type ListOfflineTaskErrorLogsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1-2-3-4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The collection of log request bodies, log responses, retry counts, and timestamps.
	Result []*ListOfflineTaskErrorLogsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListOfflineTaskErrorLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOfflineTaskErrorLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOfflineTaskErrorLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOfflineTaskErrorLogsResponseBody) GetResult() []*ListOfflineTaskErrorLogsResponseBodyResult {
	return s.Result
}

func (s *ListOfflineTaskErrorLogsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListOfflineTaskErrorLogsResponseBody) SetRequestId(v string) *ListOfflineTaskErrorLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOfflineTaskErrorLogsResponseBody) SetResult(v []*ListOfflineTaskErrorLogsResponseBodyResult) *ListOfflineTaskErrorLogsResponseBody {
	s.Result = v
	return s
}

func (s *ListOfflineTaskErrorLogsResponseBody) SetTotalCount(v int32) *ListOfflineTaskErrorLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOfflineTaskErrorLogsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOfflineTaskErrorLogsResponseBodyResult struct {
	// **The log request body.**
	//
	// example:
	//
	// {"instance":"123","user":"xuanzhen"}
	Request *string `json:"request,omitempty" xml:"request,omitempty"`
	// **The log response.**
	//
	// example:
	//
	// [{"error":{"reason":"unable to authenticate user [elastic] for REST request [/_bulk]","header":{"WWW-Authenticate":["Basic realm=\\"security\\", charset=\\"UTF-8\\"","ApiKey"]},"type":"security_exception","root_cause":[{"reason":"unable to authenticate user [elastic] for REST request [/_bulk]","header":{"WWW-Authenticate":["Basic realm=\\"security\\", charset=\\"UTF-8\\"","ApiKey"]},"type":"security_exception"}]},"status":401}]
	Response *string `json:"response,omitempty" xml:"response,omitempty"`
	// **The number of retries.**
	//
	// example:
	//
	// 1
	Retry *string `json:"retry,omitempty" xml:"retry,omitempty"`
	// **The timestamp.**
	//
	// example:
	//
	// 1770272507085
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s ListOfflineTaskErrorLogsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListOfflineTaskErrorLogsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListOfflineTaskErrorLogsResponseBodyResult) GetRequest() *string {
	return s.Request
}

func (s *ListOfflineTaskErrorLogsResponseBodyResult) GetResponse() *string {
	return s.Response
}

func (s *ListOfflineTaskErrorLogsResponseBodyResult) GetRetry() *string {
	return s.Retry
}

func (s *ListOfflineTaskErrorLogsResponseBodyResult) GetTimestamp() *string {
	return s.Timestamp
}

func (s *ListOfflineTaskErrorLogsResponseBodyResult) SetRequest(v string) *ListOfflineTaskErrorLogsResponseBodyResult {
	s.Request = &v
	return s
}

func (s *ListOfflineTaskErrorLogsResponseBodyResult) SetResponse(v string) *ListOfflineTaskErrorLogsResponseBodyResult {
	s.Response = &v
	return s
}

func (s *ListOfflineTaskErrorLogsResponseBodyResult) SetRetry(v string) *ListOfflineTaskErrorLogsResponseBodyResult {
	s.Retry = &v
	return s
}

func (s *ListOfflineTaskErrorLogsResponseBodyResult) SetTimestamp(v string) *ListOfflineTaskErrorLogsResponseBodyResult {
	s.Timestamp = &v
	return s
}

func (s *ListOfflineTaskErrorLogsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
