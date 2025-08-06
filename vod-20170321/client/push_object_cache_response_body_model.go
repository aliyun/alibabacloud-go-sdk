// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushObjectCacheResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPushTaskId(v string) *PushObjectCacheResponseBody
	GetPushTaskId() *string
	SetRequestId(v string) *PushObjectCacheResponseBody
	GetRequestId() *string
}

type PushObjectCacheResponseBody struct {
	PushTaskId *string `json:"PushTaskId,omitempty" xml:"PushTaskId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushObjectCacheResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushObjectCacheResponseBody) GoString() string {
	return s.String()
}

func (s *PushObjectCacheResponseBody) GetPushTaskId() *string {
	return s.PushTaskId
}

func (s *PushObjectCacheResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushObjectCacheResponseBody) SetPushTaskId(v string) *PushObjectCacheResponseBody {
	s.PushTaskId = &v
	return s
}

func (s *PushObjectCacheResponseBody) SetRequestId(v string) *PushObjectCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushObjectCacheResponseBody) Validate() error {
	return dara.Validate(s)
}
