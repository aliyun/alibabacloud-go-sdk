// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelInstanceRefreshResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelInstanceRefreshResponseBody
	GetRequestId() *string
}

type CancelInstanceRefreshResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelInstanceRefreshResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelInstanceRefreshResponseBody) GoString() string {
	return s.String()
}

func (s *CancelInstanceRefreshResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelInstanceRefreshResponseBody) SetRequestId(v string) *CancelInstanceRefreshResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelInstanceRefreshResponseBody) Validate() error {
	return dara.Validate(s)
}
