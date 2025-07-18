// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitVectorDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *InitVectorDatabaseResponseBody
	GetMessage() *string
	SetRequestId(v string) *InitVectorDatabaseResponseBody
	GetRequestId() *string
	SetStatus(v string) *InitVectorDatabaseResponseBody
	GetStatus() *string
}

type InitVectorDatabaseResponseBody struct {
	// The error message returned if the request fails.
	//
	// example:
	//
	// failed to connect database, detailMsg: getConnection fail::SQL State: 28P01, Error Code: 0, Error Message: FATAL: password
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **Success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s InitVectorDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitVectorDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *InitVectorDatabaseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InitVectorDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitVectorDatabaseResponseBody) GetStatus() *string {
	return s.Status
}

func (s *InitVectorDatabaseResponseBody) SetMessage(v string) *InitVectorDatabaseResponseBody {
	s.Message = &v
	return s
}

func (s *InitVectorDatabaseResponseBody) SetRequestId(v string) *InitVectorDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitVectorDatabaseResponseBody) SetStatus(v string) *InitVectorDatabaseResponseBody {
	s.Status = &v
	return s
}

func (s *InitVectorDatabaseResponseBody) Validate() error {
	return dara.Validate(s)
}
