// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGlobalDistributeCacheResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalInstanceId(v string) *CreateGlobalDistributeCacheResponseBody
	GetGlobalInstanceId() *string
	SetInstanceId(v string) *CreateGlobalDistributeCacheResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *CreateGlobalDistributeCacheResponseBody
	GetRequestId() *string
}

type CreateGlobalDistributeCacheResponseBody struct {
	GlobalInstanceId *string `json:"GlobalInstanceId,omitempty" xml:"GlobalInstanceId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E681E498-5A0D-44F2-B1A7-912DC3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGlobalDistributeCacheResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGlobalDistributeCacheResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGlobalDistributeCacheResponseBody) GetGlobalInstanceId() *string {
	return s.GlobalInstanceId
}

func (s *CreateGlobalDistributeCacheResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateGlobalDistributeCacheResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGlobalDistributeCacheResponseBody) SetGlobalInstanceId(v string) *CreateGlobalDistributeCacheResponseBody {
	s.GlobalInstanceId = &v
	return s
}

func (s *CreateGlobalDistributeCacheResponseBody) SetInstanceId(v string) *CreateGlobalDistributeCacheResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateGlobalDistributeCacheResponseBody) SetRequestId(v string) *CreateGlobalDistributeCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGlobalDistributeCacheResponseBody) Validate() error {
	return dara.Validate(s)
}
