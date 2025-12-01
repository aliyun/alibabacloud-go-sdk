// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlarmDetailByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetAlarmDetailByIdRequest
	GetId() *int64
}

type GetAlarmDetailByIdRequest struct {
	// Primary key ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20077761
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetAlarmDetailByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmDetailByIdRequest) GoString() string {
	return s.String()
}

func (s *GetAlarmDetailByIdRequest) GetId() *int64 {
	return s.Id
}

func (s *GetAlarmDetailByIdRequest) SetId(v int64) *GetAlarmDetailByIdRequest {
	s.Id = &v
	return s
}

func (s *GetAlarmDetailByIdRequest) Validate() error {
	return dara.Validate(s)
}
