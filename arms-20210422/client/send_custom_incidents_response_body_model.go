// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendCustomIncidentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendCustomIncidentsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendCustomIncidentsResponseBody
	GetSuccess() *bool
}

type SendCustomIncidentsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendCustomIncidentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendCustomIncidentsResponseBody) GoString() string {
	return s.String()
}

func (s *SendCustomIncidentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendCustomIncidentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendCustomIncidentsResponseBody) SetRequestId(v string) *SendCustomIncidentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendCustomIncidentsResponseBody) SetSuccess(v bool) *SendCustomIncidentsResponseBody {
	s.Success = &v
	return s
}

func (s *SendCustomIncidentsResponseBody) Validate() error {
	return dara.Validate(s)
}
