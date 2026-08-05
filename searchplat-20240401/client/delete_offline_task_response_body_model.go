// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOfflineTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteOfflineTaskResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteOfflineTaskResponseBody
	GetResult() *bool
}

type DeleteOfflineTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 2423C841-91C4-5E51-B296-590D367967FC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	//
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteOfflineTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOfflineTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOfflineTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOfflineTaskResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteOfflineTaskResponseBody) SetRequestId(v string) *DeleteOfflineTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOfflineTaskResponseBody) SetResult(v bool) *DeleteOfflineTaskResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteOfflineTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
