// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAICInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetAICInstanceResponseBody
	GetRequestId() *string
}

type ResetAICInstanceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAICInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetAICInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAICInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetAICInstanceResponseBody) SetRequestId(v string) *ResetAICInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetAICInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
