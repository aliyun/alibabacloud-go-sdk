// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelLensServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelLensServiceResponseBody
	GetRequestId() *string
}

type CancelLensServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelLensServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelLensServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CancelLensServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelLensServiceResponseBody) SetRequestId(v string) *CancelLensServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelLensServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
