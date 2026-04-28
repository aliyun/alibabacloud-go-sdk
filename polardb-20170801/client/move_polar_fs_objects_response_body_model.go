// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMovePolarFsObjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolarFsInstanceId(v string) *MovePolarFsObjectsResponseBody
	GetPolarFsInstanceId() *string
	SetRequestId(v string) *MovePolarFsObjectsResponseBody
	GetRequestId() *string
}

type MovePolarFsObjectsResponseBody struct {
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6BD9CDE4-5E7B-4BF3-9BB8-83C73E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MovePolarFsObjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MovePolarFsObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *MovePolarFsObjectsResponseBody) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *MovePolarFsObjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MovePolarFsObjectsResponseBody) SetPolarFsInstanceId(v string) *MovePolarFsObjectsResponseBody {
	s.PolarFsInstanceId = &v
	return s
}

func (s *MovePolarFsObjectsResponseBody) SetRequestId(v string) *MovePolarFsObjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *MovePolarFsObjectsResponseBody) Validate() error {
	return dara.Validate(s)
}
