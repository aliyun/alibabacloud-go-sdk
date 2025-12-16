// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDataCollectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveDataCollectionResponseBody
	GetRequestId() *string
	SetResult(v string) *RemoveDataCollectionResponseBody
	GetResult() *string
}

type RemoveDataCollectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// -
	//
	// example:
	//
	// {}
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveDataCollectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveDataCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDataCollectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveDataCollectionResponseBody) GetResult() *string {
	return s.Result
}

func (s *RemoveDataCollectionResponseBody) SetRequestId(v string) *RemoveDataCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDataCollectionResponseBody) SetResult(v string) *RemoveDataCollectionResponseBody {
	s.Result = &v
	return s
}

func (s *RemoveDataCollectionResponseBody) Validate() error {
	return dara.Validate(s)
}
