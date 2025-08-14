// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDropSearchLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DropSearchLibResponseBody
	GetCode() *string
	SetRequestId(v string) *DropSearchLibResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DropSearchLibResponseBody
	GetSuccess() *string
}

type DropSearchLibResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
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
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DropSearchLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DropSearchLibResponseBody) GoString() string {
	return s.String()
}

func (s *DropSearchLibResponseBody) GetCode() *string {
	return s.Code
}

func (s *DropSearchLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DropSearchLibResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DropSearchLibResponseBody) SetCode(v string) *DropSearchLibResponseBody {
	s.Code = &v
	return s
}

func (s *DropSearchLibResponseBody) SetRequestId(v string) *DropSearchLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *DropSearchLibResponseBody) SetSuccess(v string) *DropSearchLibResponseBody {
	s.Success = &v
	return s
}

func (s *DropSearchLibResponseBody) Validate() error {
	return dara.Validate(s)
}
