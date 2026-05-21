// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateInstanceZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *MigrateInstanceZoneResponseBody
	GetData() *bool
	SetRequestId(v string) *MigrateInstanceZoneResponseBody
	GetRequestId() *string
}

type MigrateInstanceZoneResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 819A7F0F-2951-540F-BD94-6A41ECF0281F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s MigrateInstanceZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MigrateInstanceZoneResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateInstanceZoneResponseBody) GetData() *bool {
	return s.Data
}

func (s *MigrateInstanceZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MigrateInstanceZoneResponseBody) SetData(v bool) *MigrateInstanceZoneResponseBody {
	s.Data = &v
	return s
}

func (s *MigrateInstanceZoneResponseBody) SetRequestId(v string) *MigrateInstanceZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateInstanceZoneResponseBody) Validate() error {
	return dara.Validate(s)
}
