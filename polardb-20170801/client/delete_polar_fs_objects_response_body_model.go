// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarFsObjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolarFsInstanceId(v string) *DeletePolarFsObjectsResponseBody
	GetPolarFsInstanceId() *string
	SetRequestId(v string) *DeletePolarFsObjectsResponseBody
	GetRequestId() *string
}

type DeletePolarFsObjectsResponseBody struct {
	// example:
	//
	// pfs-2ze0i7*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C890995A-CF06-4F4D-8DB8-DD26C2******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePolarFsObjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarFsObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolarFsObjectsResponseBody) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DeletePolarFsObjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePolarFsObjectsResponseBody) SetPolarFsInstanceId(v string) *DeletePolarFsObjectsResponseBody {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DeletePolarFsObjectsResponseBody) SetRequestId(v string) *DeletePolarFsObjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePolarFsObjectsResponseBody) Validate() error {
	return dara.Validate(s)
}
