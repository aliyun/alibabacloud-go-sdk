// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSearchStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveSearchStrategyResponseBody
	GetRequestId() *string
}

type RemoveSearchStrategyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F5099063-6B86-F398-D843-905F9EFB683A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RemoveSearchStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveSearchStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSearchStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveSearchStrategyResponseBody) SetRequestId(v string) *RemoveSearchStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveSearchStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
