// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyLensServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ApplyLensServiceResponseBody
	GetRequestId() *string
}

type ApplyLensServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyLensServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyLensServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyLensServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyLensServiceResponseBody) SetRequestId(v string) *ApplyLensServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyLensServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
