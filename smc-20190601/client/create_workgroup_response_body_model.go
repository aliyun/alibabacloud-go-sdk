// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkgroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateWorkgroupResponseBody
	GetRequestId() *string
	SetWorkgroupId(v string) *CreateWorkgroupResponseBody
	GetWorkgroupId() *string
}

type CreateWorkgroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C8B26B44-0189-443E-9816-D951F59623A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the workgroup.
	//
	// example:
	//
	// w-***
	WorkgroupId *string `json:"WorkgroupId,omitempty" xml:"WorkgroupId,omitempty"`
}

func (s CreateWorkgroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkgroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkgroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkgroupResponseBody) GetWorkgroupId() *string {
	return s.WorkgroupId
}

func (s *CreateWorkgroupResponseBody) SetRequestId(v string) *CreateWorkgroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkgroupResponseBody) SetWorkgroupId(v string) *CreateWorkgroupResponseBody {
	s.WorkgroupId = &v
	return s
}

func (s *CreateWorkgroupResponseBody) Validate() error {
	return dara.Validate(s)
}
