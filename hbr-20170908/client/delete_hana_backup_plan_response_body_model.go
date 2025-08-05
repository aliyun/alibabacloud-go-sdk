// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHanaBackupPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteHanaBackupPlanResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteHanaBackupPlanResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteHanaBackupPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteHanaBackupPlanResponseBody
	GetSuccess() *bool
}

type DeleteHanaBackupPlanResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 071E4789-6256-526B-B22E-2A9CDDB9EB21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteHanaBackupPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHanaBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHanaBackupPlanResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteHanaBackupPlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteHanaBackupPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHanaBackupPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteHanaBackupPlanResponseBody) SetCode(v string) *DeleteHanaBackupPlanResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHanaBackupPlanResponseBody) SetMessage(v string) *DeleteHanaBackupPlanResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHanaBackupPlanResponseBody) SetRequestId(v string) *DeleteHanaBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHanaBackupPlanResponseBody) SetSuccess(v bool) *DeleteHanaBackupPlanResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteHanaBackupPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
