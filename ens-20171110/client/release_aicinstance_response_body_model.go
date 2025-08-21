// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseAICInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseAICInstanceResponseBody
	GetRequestId() *string
}

type ReleaseAICInstanceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseAICInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseAICInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseAICInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseAICInstanceResponseBody) SetRequestId(v string) *ReleaseAICInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseAICInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
