// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDayuMigrateStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryDayuMigrateStatusResponseBody
	GetRequestId() *string
	SetCode(v string) *QueryDayuMigrateStatusResponseBody
	GetCode() *string
	SetData(v string) *QueryDayuMigrateStatusResponseBody
	GetData() *string
}

type QueryDayuMigrateStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryDayuMigrateStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDayuMigrateStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDayuMigrateStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDayuMigrateStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryDayuMigrateStatusResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryDayuMigrateStatusResponseBody) SetRequestId(v string) *QueryDayuMigrateStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDayuMigrateStatusResponseBody) SetCode(v string) *QueryDayuMigrateStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDayuMigrateStatusResponseBody) SetData(v string) *QueryDayuMigrateStatusResponseBody {
	s.Data = &v
	return s
}

func (s *QueryDayuMigrateStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
