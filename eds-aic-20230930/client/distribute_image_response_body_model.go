// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDistributeImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DistributeImageResponseBody
	GetRequestId() *string
}

type DistributeImageResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 440D7342-5FC2-5E7C-B2DB-D0B4EAC2BDF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DistributeImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DistributeImageResponseBody) GoString() string {
	return s.String()
}

func (s *DistributeImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DistributeImageResponseBody) SetRequestId(v string) *DistributeImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DistributeImageResponseBody) Validate() error {
	return dara.Validate(s)
}
