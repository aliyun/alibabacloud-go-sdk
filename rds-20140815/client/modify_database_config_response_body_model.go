// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatabaseConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyDatabaseConfigResponseBody
	GetCode() *string
	SetMessage(v string) *ModifyDatabaseConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyDatabaseConfigResponseBody
	GetRequestId() *string
}

type ModifyDatabaseConfigResponseBody struct {
	// The code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8B993DA9-5272-5414-94E3-4CA8BA0146C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDatabaseConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatabaseConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyDatabaseConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyDatabaseConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDatabaseConfigResponseBody) SetCode(v string) *ModifyDatabaseConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDatabaseConfigResponseBody) SetMessage(v string) *ModifyDatabaseConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyDatabaseConfigResponseBody) SetRequestId(v string) *ModifyDatabaseConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDatabaseConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
