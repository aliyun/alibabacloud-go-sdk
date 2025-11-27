// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseReadWriteSplittingConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseReadWriteSplittingConnectionResponseBody
	GetRequestId() *string
}

type ReleaseReadWriteSplittingConnectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5A77D650-27A1-4E08-AD9E-59008EDB6927
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseReadWriteSplittingConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseReadWriteSplittingConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseReadWriteSplittingConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseReadWriteSplittingConnectionResponseBody) SetRequestId(v string) *ReleaseReadWriteSplittingConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseReadWriteSplittingConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
