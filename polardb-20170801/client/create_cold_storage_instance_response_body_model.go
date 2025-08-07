// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateColdStorageInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetColdStorageInstanceId(v string) *CreateColdStorageInstanceResponseBody
	GetColdStorageInstanceId() *string
	SetRequestId(v string) *CreateColdStorageInstanceResponseBody
	GetRequestId() *string
}

type CreateColdStorageInstanceResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// pcs_2zeth2gf4i83e578t
	ColdStorageInstanceId *string `json:"ColdStorageInstanceId,omitempty" xml:"ColdStorageInstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F6EBB4ED-D12F-5F49-824C-9DD9C0EC4CF2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateColdStorageInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateColdStorageInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateColdStorageInstanceResponseBody) GetColdStorageInstanceId() *string {
	return s.ColdStorageInstanceId
}

func (s *CreateColdStorageInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateColdStorageInstanceResponseBody) SetColdStorageInstanceId(v string) *CreateColdStorageInstanceResponseBody {
	s.ColdStorageInstanceId = &v
	return s
}

func (s *CreateColdStorageInstanceResponseBody) SetRequestId(v string) *CreateColdStorageInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateColdStorageInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
